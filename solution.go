// 作者：姜嵩
// 日期：2018-02-01

// 随机生成合法数独解
package sudoku

type Answer struct {
	Sudoku
	Travel []*int
}

func (m *Sudoku) NewAnswer() *Answer {
	result := &Answer{Sudoku: *m}
	result.Travel = make([]*int, result.GetNullNum())
	num := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if result.M_number[y][x] == 0 {
				result.Travel[num] = &result.M_number[y][x]
				num++
			}
		}
	}
	return result
}

// 检查当前値是否非法
func (m *Answer) CheckErr() bool {
	for mod := 0; mod < 3; mod++ {
		for i := 0; i < 9; i++ {
			var flag uint = 0
			for j := 0; j < 9; j++ {
				v := uint(m.GetIdxVal(9*i+j, mod))
				if v != 0 {
					// 排查重复值
					if (flag & (1 << v)) != 0 {
						return true
					}
					flag |= 1 << v
				}
			}
		}
	}
	return false
}

const (
	ANS_OK = iota
	ANS_NO
	ANS_ERR
)

// 随机求一个解
func (m *Answer) ranAnswer(deep int) int {
	if deep < 0 {
		return ANS_ERR
	} else if deep >= len(m.Travel) {
		return ANS_OK
	}

	rnd := RandOrd()

	p := m.Travel[deep]
	for _, val := range rnd {
		*p = val
		// 当前预估值不合理，剪枝
		if m.CheckErr() {
			continue
		}
		// 当前预估值正常，深搜
		ret := m.ranAnswer(deep + 1)
		if ret == ANS_OK {
			// 找到最终解
			return ANS_OK
		}
		// 未找到继续迭代
	}
	*p = 0
	return ANS_NO
}

// 将猜测结果清空
func (m *Answer) Clear() {
	for _, p := range m.Travel {
		*p = 0
	}
}

// 解数独
// 输入：数独基础结构
// 输出：单个解数组
func (m *Answer) GetRandSolution() int {
	m.Clear()
	if m.CheckErr() {
		return ANS_NO // 输入无解
	}
	return m.ranAnswer(0)
}

// 顺序迭代下一个解
func (m *Answer) nextAnswer(deep int) int {
	if deep < 0 {
		return ANS_ERR
	} else if deep >= len(m.Travel) {
		return ANS_OK
	}

	p := m.Travel[deep]
	if *p != 0 && deep < len(m.Travel)-1 { // 从上次结果继续迭代
		if ANS_OK == m.nextAnswer(deep+1) {
			// 找到最终解
			return ANS_OK
		}
	}
	for *p < 9 {
		*p++
		// 当前预估值不合理，剪枝
		if m.CheckErr() {
			continue
		}
		// 当前预估值正常，深搜
		ret := m.nextAnswer(deep + 1)
		if ret == ANS_OK {
			// 找到最终解
			return ANS_OK
		}
		// 未找到继续迭代
	}
	*p = 0
	return ANS_NO
}

// 解数独
// 输入：数独基础结构
// 输出：单个解数组
func (m *Answer) GetNextSolution() int {
	if m.CheckErr() {
		return ANS_NO // 输入无解
	}
	return m.nextAnswer(0)
}
