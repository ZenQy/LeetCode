package solve0495

/**
 * @index 495
 * @title 提莫攻击
 * @difficulty 简单
 * @tags array,simulation
 * @draft false
 * @link https://leetcode-cn.com/problems/teemo-attacking/
 * @frontendId 495
 */

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := duration
	for i := 0; i < len(timeSeries)-1; i++ {
		sub := timeSeries[i+1] - timeSeries[i]
		if sub > duration {
			total += duration
		} else {
			total += sub
		}
	}

	return total
}
