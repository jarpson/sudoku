// 作者：姜嵩
// 日期：2018-02-01

// 数独基本结构及函数定义
package sudoku

import "fmt"

// 数独基本结构定义
// 1-9 表示固定的值，0 表示未知值
type Sudoku struct {
	M_number [9][9]int
}

// 对比两个数独是否相等
func (m *Sudoku) Equal(sd *Sudoku) bool {
	return m.M_number == sd.M_number
}

// 生成副本
func (m *Sudoku) Dump() *Sudoku {
	return &Sudoku{m.M_number}
}

// 获取指定序号元素值
func (m *Sudoku) GetIdxVal(idx, mod int) int {
	x, y := GetIdxPos(idx, mod)
	if x < 0 || y < 0 {
		return -1
	}
	return m.M_number[y][x]
}

// 设置指定序号元素值
func (m *Sudoku) SetIdxVal(idx, mod, val int) {
	x, y := GetIdxPos(idx, mod)
	if x < 0 || y < 0 || val < 0 || val > 9 {
		return
	}
	m.M_number[y][x] = val
}

func (m *Sudoku) GetNullNum() int {
	num := 0
	for idx := 0; idx < 81; idx++ {
		if m.GetIdxVal(idx, BY_LINE) == 0 {
			num++
		}
	}
	return num
}

// 输出内容
func (m *Sudoku) Print() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Printf("%d  ", m.M_number[y][x])
			if x%3 == 2 {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
		if y%3 == 2 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// 返回字串形式，如 123456789456789123……
func (m *Sudoku) String() string {
	var content [81]byte
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			content[9*y+x] = 0x30 + byte(m.M_number[y][x])
		}
	}
	return string(content[:])
}

// 输入字符串形式的数独，如果格式不合法返回false
func (m *Sudoku) SetString(str string) bool {
	if len(str) != 81 {
		return false // 长度不合法
	}
	for i, v := range str {
		c := int(v) - 0x30
		if c < 0 || c > 9 {
			return false // 值不正确
		}
		m.M_number[i/9][i%9] = c
	}
	return true
}

// 输入数独 ,如果输入不合法返回false
func (m *Sudoku) Scan() bool {
	fmt.Printf("input sudoku:\n")
	var val int
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			for {
				_, err := fmt.Scanf("%d", &val)
				if err == nil {
					break
				}
			}
			if val < 0 || val > 9 {
				return false // 输入不合法
			}
			m.M_number[y][x] = val
		}
	}
	fmt.Println("scan over")
	return true
}

// 遍历方式
const (
	BY_LINE = iota // 按行
	BY_COW         // 按列
	BY_AREA        // 按宫格
)

// 获取遍历下一个节点
// 输入：节点排序号、遍历方式
// 输出：x, y 坐标(小于0 表示idx异常）
func GetIdxPos(idx, mod int) (x int, y int) {
	if idx >= 81 {
		return -1, -1
	}
	switch mod {
	case BY_LINE:
		x = idx % 9
		y = idx / 9
	case BY_COW:
		y = idx % 9
		x = idx / 9
	case BY_AREA:
		area := idx / 9
		x1 := (area % 3) * 3
		y1 := (area / 3) * 3

		pos := idx % 9
		x = x1 + pos%3
		y = y1 + pos/3
	}
	return
}

// 获取节点序号
// 输入： x,y坐标 （编号从0-8）
// 输出：序号 0-80， 异常点返回 -1
func GetPosIdx(x, y, mod int) int {
	if x >= 9 || y >= 9 {
		return -1
	}
	switch mod {
	case BY_LINE:
		return y*9 + x
	case BY_COW:
		return x*9 + y
	case BY_AREA:
		x1 := x / 3
		y1 := y / 3
		id := y1*9*3 + x1*9

		x2 := x % 3
		y2 := y % 3
		return id + y2*3 + x2
	}
	return -1
}
