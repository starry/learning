package queue

// MovingAverage  数据流中的移动平均值
type MovingAverage struct {
	data []int
	size int
}

// CreateMovingAverage Initialize your data structure here.
func CreateMovingAverage(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

// Next add num in the queue
func (m *MovingAverage) Next(val int) float64 {

	m.data = append(m.data, val)

	if m.size < len(m.data) {
		m.data = m.data[1:]
	}

	var result int

	for i := 0; i < len(m.data); i++ {
		result += m.data[i]
	}

	return float64(result) / float64(len(m.data))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
