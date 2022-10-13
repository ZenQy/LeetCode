package solve0151

/**
 * @index 151
 * @title 反转字符串中的单词
 * @difficulty 中等
 * @tags two-pointers,string
 * @draft false
 * @link https://leetcode.cn/problems/reverse-words-in-a-string/
 * @frontendId 151
 */

func reverseWords(s string) string {
	bs, count := []byte(s), len(s)
	lold, rold := 0, 0
	lnew, rnew := 0, -2
	reverseBytes(bs)

	for i := 0; i < count; i++ {
		if bs[i] != ' ' && (i == 0 || bs[i-1] == ' ') {
			lold = i
		}
		if bs[i] != ' ' && (i+1 == count || bs[i+1] == ' ') {
			rold = i
			lnew = rnew + 2
			rnew = rold - lold + lnew
			reverseBytes(bs[lnew : rold+1])
		}

	}
	return string(bs[:rnew+1])
}

func reverseBytes(bs []byte) {
	count := len(bs)
	for i := 0; i < count/2; i++ {
		bs[i], bs[count-1-i] = bs[count-1-i], bs[i]
	}
}
