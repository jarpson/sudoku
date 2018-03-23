// 作者：姜嵩
// 日期：2018-02-02

// 数独库测试程序
package main

import (
	"../../sudoku"
	"flag"
	"fmt"
	"time"
)

// 生成及求解数独
// 例： 090001400200005700000200086040300500510090024002004010670003000005600009004900070
func main() {
	gen := flag.Int("gen", 0, "gen sudoku mod, 0: gen randrom, 1:scanf, 2: input sudoku string(use -str)")
	sudostr := flag.String("str", "", "sudoku string")
	flag.Parse()

	// 生成数独
	var sd sudoku.Sudoku
	switch *gen {
		case 1: // 直接输入
			if !sd.Scan() {
				fmt.Println("input wrong")
				return
			}
		case 2: // 通过命令行参数输入字符串型的数独
			if !sd.SetString(*sudostr) {
				fmt.Println("sudostr err, please check -str")
				return
			}
		default: // 随机生成
			b1 := time.Now().UnixNano()
			arr := sudoku.RandLevelSudoku()
			b2 := time.Now().UnixNano()
			for i, s := range arr {
				if s != nil {
					fmt.Printf("gen gread:%d,level:%d, np:%d\n", i, s.GetLevel(), 81-s.GetNullNum())
					s.Print()
				}
			}
			fmt.Printf("gen time:%v\n", b2-b1)
			return
	}
	fmt.Printf("gen1 level:%d, np:%d\n", sd.GetLevel(), 81-sd.GetNullNum())
	sd.Print()

	// 求解
	ans := sd.NewAnswer()
	b1 := time.Now().UnixNano()
	ret := ans.GetNextSolution()
	b2 := time.Now().UnixNano()
	fmt.Printf("result:%d, time:%v\n", ret, b2 - b1)
	if ret == 0 { // 有解
		ans.Print()
		fmt.Println(ans.String())

		// 检查解是否唯一 
		ret = ans.GetNextSolution()
		if ret == 0 {
			fmt.Printf("answer not unique:\n")
			ans.Print()
		}
	}
}
