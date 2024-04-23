package backtracking

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	res := readBinaryWatchByEnum(9)
	fmt.Println(res)
}

func readBinaryWatch(turnedOn int) []string {
	return nil
}

// 枚举法
// 枚举小时的所有可能值 [0,11]，以及分钟的所有可能值 [0,59]，
// 并计算二者的二进制中 1 的个数之和，若为 turnedOn ，则将其加入到答案中。
func readBinaryWatchByEnum(turnedOn int) []string {
	res := make([]string, 0)
	for h := 0; h < 12; h++ {
		for m := 0; m < 59; m++ {
			// bits.OnesCount8返回入参中比特位为1的个数
			if bits.OnesCount8(uint8(h))+bits.OnesCount8(uint8(m)) == turnedOn {
				res = append(res, fmt.Sprintf("%v:%02d", h, m))
			}
		}
	}
	return res
}
