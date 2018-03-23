// 作者：姜嵩
// 日期：2018-02-01

// 随机生成合法数独解
package sudoku

//import "fmt"

// 随机生成数独解
func RandomSudoku() Sudoku {
	var sd Sudoku

	// 设置对角线随机, 对角线的三个宫格互不相干，因此先填充完提高效率
	start := [3]int{0, 36, 72}
	for _, s := range start {
		rnd := RandOrd()
		for i := 0; i < 9; i++ {
			sd.SetIdxVal(s+i, BY_AREA, rnd[i])
		}
	}
	// 随机填充剩下的宫格
	ans := sd.NewAnswer()
	ans.GetRandSolution()
	return ans.Sudoku
}

// 检查是否有唯一解
// 输入： sd 有唯一解的数独
//		idx 准备置空的格式
// 输出：
//		true:置空后仍有唯一解
func checkOnly(sd *Sudoku, idx int) bool {
	ans := sd.NewAnswer()
	src := ans.GetIdxVal(idx, BY_LINE)
	for i := 1; i <= 9; i++ {
		if i == src {
			continue
		}
		ans.SetIdxVal(idx, BY_LINE, i)
		if ans.GetNextSolution() == ANS_OK {
			return false
		}
	}
	return true
}

// 生成数独
// 生成不同难度等级
func RandLevelSudoku() (result [LEVEL_GRADE_NUM + 1]*Sudoku) {
	sd := RandomSudoku()
	result[LEVEL_SOLUTION] = &sd

	ret, rnd := randSudoku(&sd)
	
	level := ret.GetLevelGrade()
	result[level] = ret.Dump()

	// 填充，生成低难度数独
	if level > LEVEL_EASY {
		ret2 := ret.Dump() // copy
		for i := len(rnd) - 1; i >= 0; i-- {
			if ret2.GetIdxVal(rnd[i], BY_LINE) == 0 {
				ret2.SetIdxVal(rnd[i], BY_LINE, sd.GetIdxVal(rnd[i], BY_LINE))
				level = ret2.GetLevelGrade()
				if result[level] == nil {
					result[level] = ret2.Dump()
					if level == LEVEL_EASY {
						break; // 已生成最低档的数独
					}
				}
			}
		}
	}

	// 尝试去掉部分结点
	for _, idx := range rnd {
		if ret.GetIdxVal(idx, BY_LINE) != 0 && checkOnly(ret, idx) {
			ret.SetIdxVal(idx, BY_LINE, 0)
			level = ret.GetLevelGrade()
			if result[level] == nil || level == LEVEL_SURVIVAL { // 最后一级用最难的
				result[level] = ret.Dump()
			}
		}
	}
	return
}

// 生成合法数独
// 先生成随机解样板，从样板中随机同步17个格子（17 是可唯一解下限）
// 然后一个一个随机同步，同时检查解是否唯一，解唯一即停止
func randSudoku(psd *Sudoku) (*Sudoku, [81]int) {
	var ret Sudoku	
	maxnum := 37	// 测试 level 1 通常是34 这里多设几
redo:
	rnd := RandCreateOrd()
	for i := 0; i < 17; i++ { // 非空格数大于17才可能解唯一
		idx := rnd[i]
		ret.SetIdxVal(idx, BY_LINE, psd.GetIdxVal(idx, BY_LINE))
	}

	count := 17
	for {
		// 判断解是否唯一
		ans := ret.NewAnswer()
		ans.GetRandSolution()
		if ans.Equal(psd) { // 随机解与预设解相等，判断是否是唯一解
			ans.Clear()
			ans.GetNextSolution()
			if ans.GetNextSolution() != ANS_OK {
				// 得到唯一解
				break
			}
			// 非唯一
			if ans.Equal(psd) {
				continue
			}
		}
		num := 0
		// 解不唯一时，赋差异值(差异赋值方式可有效降低平均查找次数）
		for i := 17; i < 81; i++ {
			idx := rnd[i]
			val := psd.GetIdxVal(idx, BY_LINE)
			if ans.GetIdxVal(idx, BY_LINE) != val {
				ret.SetIdxVal(idx, BY_LINE, val)
				num++
				count++
				if num >= 3 { // 为提高效率，一次赋多个
					break
				}
			}
		}
		if count > maxnum { // ，如果填充的值过大会增加去点的计算量，及时止损重新生成随机序列
			ret = Sudoku{}
			maxnum += 5 // 阀值加5
			goto redo
		}
	}
	return &ret, rnd
}
