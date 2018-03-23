// 作者：姜嵩
// 日期：2018-02-01

// 获取难度等级
package sudoku


// 个数需要与 LEVEL_GRADE_NUM 一致
// 0: 0 解
// 1: 0-700 简单
// 2: 700-800 中等
// 3: 800-900 困难
// 4: 900-970 骨灰级
// 5: 970+ 生存模式
const (
	LEVEL_SOLUTION	= iota
	LEVEL_EASY
	LEVEL_NORMAL
	LEVEL_HARD
	LEVEL_VERYHARD
	LEVEL_SURVIVAL
	LEVEL_GRADE_NUM // 难度等级档次个数
)
// 获取难度等级
// 难度值= 每个空格方格所在的行列宫格中空格数之和, 参考了：http://www.cnblogs.com/candyhuang/archive/2011/12/17/2153668.html
func (m *Sudoku) GetLevel() int {
	sum, num := 0, 0
	// 遍历行、列、宫格
	for mod := 0; mod < 3; mod++ {
		line := 0
		for idx := 0; idx < 81; idx++ {
			if m.GetIdxVal(idx, mod) == 0 {
				line++
			}
			if (idx+1)%9 == 0 {
				sum += line * line
				if mod == 0 { // 只计算一次
					num += line
				}
				line = 0
			}
		}
	}
	return sum - 2*num // 行、列、宫格各被计算一次，多计算了两次
}

// 计算难度等级
func (m *Sudoku) GetLevelGrade() int {
	level := m.GetLevel()
	if level > 970 {
		return LEVEL_SURVIVAL
	} else if level > 900 {
		return LEVEL_VERYHARD
	} else if level > 800 {
		return LEVEL_HARD
	} else if level > 700 {
		return LEVEL_NORMAL
	} else if level > 0 {
		return LEVEL_EASY
	}
	return LEVEL_SOLUTION
}
// 统计 x,y 所在 行/列/宫格 中空方格的个数
// 输入：x, y 坐标 及 统计方式
// 输出：个数
func (m *Sudoku) getNullNumber(x, y, mod int) int {
	num := 0 // 空方格的个数
	idx := GetPosIdx(x, y, mod)
	start := idx - idx%9 //
	for i := 0; i < 9; i++ {
		if m.GetIdxVal(start+i, mod) == 0 {
			num++
		}
	}
	return num
}

// 将一个点置空后，难度等级增加值
func (m *Sudoku) GetLevelDiffAfter(x, y int) int {
	if m.M_number[y][x] == 0 {
		return 0
	}

	num := 0
	for mod := 0; mod < 3; mod++ {
		num += m.getNullNumber(x, y, mod)
	}
	return num*2 + 1
}
