package solve0661

/**
 * @index 661
 * @title 图片平滑器
 * @difficulty 简单
 * @tags array,matrix
 * @draft false
 * @link https://leetcode.cn/problems/image-smoother/
 * @frontendId 661
 */

func imageSmoother(img [][]int) [][]int {
	mod := 1 << 8
	m, n := len(img), len(img[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for ii := i - 1; ii <= i+1; ii++ {
				for jj := j - 1; jj <= j+1; jj++ {
					if ii >= 0 && ii < m && jj >= 0 && jj < n {
						sum += img[ii][jj] % mod
						count++
					}
				}
			}
			img[i][j] += sum / count * mod
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			img[i][j] /= mod
		}
	}
	return img
}
