package utils

import "testing"

func TestRemovePunctuation(t *testing.T) {
	str := "iosadfio902184][\2131[21]4[12][4/[1]\\24[]\\-*["
	str = RemovePunctuation(str)
	if str != "iosadfio902184\2131214124124" {
		t.Fatal()
	}
}
