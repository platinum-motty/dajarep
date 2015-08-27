// sharep_test.go
package main

import (
	"strings"
	"testing"
)

func TestSharep(t *testing.T) {

	input := `人民の人民による人民のための政治
アルミ缶の上にあるミカン
トンネルを抜けるとそこは雪国であった
智代子のチョコ
布団が吹っ飛んだ
我輩は猫である
猫が寝転んだ
その意見にはついていけん
靴を靴箱に入れる
傘を貸さない
イカは如何なものか
親譲りの無鉄砲で子供の時から損ばかりしている`
	ans := `アルミ缶の上にあるミカン
智代子のチョコ
布団が吹っ飛んだ
猫が寝転んだ
その意見にはついていけん
傘を貸さない
イカは如何なものか`
	res := strings.Join(Sharep(input), "\n")
	if res != ans {
		t.Errorf("Share(x) =\n%s\n, want \n%s", res, ans)
		return
	}
}
