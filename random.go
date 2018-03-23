// 作者：姜嵩
// 日期：2018-02-02

// 生成随机9宫格序列
package sudoku

import (
	"math/big"
	"math/rand"
	"time"
	//"fmt"
)

// 1-9 排例，用于求第N个排列
var g_gird [9]int

//  用于生成数独挖洞
var g_create [81]big.Int

var g_rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	// 初始化
	g_gird[0] = 1
	for i := 1; i < len(g_gird); i++ {
		g_gird[i] = (i + 1) * g_gird[i-1]
	}
	// 初始化
	g_create[0].SetInt64(1)
	for i := 1; i < len(g_create); i++ {
		g_create[i].Mul(&g_create[i-1], big.NewInt(int64(i+1)))
		//fmt.Printf("%d\t%v\n", i, g_create[i].String())
	}
}

// 1-9随机排列
func RandOrd() (ret [9]int) {
	l := len(ret)
	n := g_rnd.Intn(g_gird[l-1])
	for i := 0; i < l; i++ {
		ret[i] = i + 1
	}
	for i := 0; i < l-1; i++ {
		mod := g_gird[l-2-i]
		ord := n / mod
		n = n % mod
		if ord > 0 { // 需要移位
			tmp := ret[i+ord]
			for j := i + ord; j > i; j-- {
				ret[j] = ret[j-1]
			}
			ret[i] = tmp
		}
	}
	return
}

// 1-9随机排列
func RandCreateOrd() (ret [81]int) {
	l := len(ret)
	var n, mod big.Int
	n.Rand(g_rnd, &g_create[l-1])
	for i := 0; i < l; i++ {
		ret[i] = i + 1
	}
	for i := 0; i < l-1; i++ {
		mod.QuoRem(&n, &g_create[l-2-i], &n)
		if mod.Sign() > 0 { // 需要移位
			ord := int(mod.Uint64())
			tmp := ret[i+ord]
			for j := i + ord; j > i; j-- {
				ret[j] = ret[j-1]
			}
			ret[i] = tmp
		}
	}
	return
}
