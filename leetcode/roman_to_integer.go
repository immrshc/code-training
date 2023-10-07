package leetcode

var SymbolValue = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt1(s string) int {
	var num, memo int
	for _, char := range s {
		v := SymbolValue[string(char)]
		if memo != 0 && memo < v {
			// 余分に追加してしまった値を一度引く
			num -= memo
			num = num + v - memo
		} else {
			num += v
		}
		memo = v
	}
	return num
}

func RomanToInt2(s string) int {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return SymbolValue[s]
	}
	var total int
	for i := 0; i < n; i++ {
		former := SymbolValue[string(runes[i])]
		// 前後の文字のペアの前が計算の対象になる
		// つまり一番最後の文字は対象にならないので、さらに後ろに仮想的な文字（ゼロ）を設ける
		var latter int
		if i+1 < n {
			latter = SymbolValue[string(runes[i+1])]
		}
		if former >= latter {
			total += former
		} else {
			total -= former
		}
	}
	return total
}
