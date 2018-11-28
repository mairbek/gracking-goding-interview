package leetcode

// The idea is to keep a max heap with elements that are less than the median.
// And the min heap with elements that are greater than the median.

type SlidingMedianFinder struct {
	minHeap *Heap
	maxHeap *Heap
}

func NewSlidingMedianFinder() SlidingMedianFinder {
	return SlidingMedianFinder{
		minHeap: NewMinHeap(),
		maxHeap: NewMaxHeap(),
	}
}

func (this *SlidingMedianFinder) AddNum(num int) {
	insertMax := this.minHeap.Len == 0
	if !insertMax {
		_, val := this.minHeap.Min()
		insertMax = num > val
	}
	if insertMax {
		this.minHeap.Insert(num)
		if this.minHeap.Len > this.maxHeap.Len+1 {
			_, val := this.minHeap.ExtractMin()
			this.maxHeap.Insert(val)
		}
	} else {
		this.maxHeap.Insert(num)
		if this.maxHeap.Len > this.minHeap.Len+1 {
			_, val := this.maxHeap.ExtractMin()
			this.minHeap.Insert(val)
		}
	}
}

func (this *SlidingMedianFinder) FindMedian() float64 {
	len := this.minHeap.Len + this.maxHeap.Len
	if len == 0 {
		return 0
	}
	if this.maxHeap.Len > this.minHeap.Len {
		_, result := this.maxHeap.Min()
		return float64(result)
	}
	if this.minHeap.Len > this.maxHeap.Len {
		_, result := this.minHeap.Min()
		return float64(result)
	}
	_, a := this.minHeap.Min()
	_, b := this.maxHeap.Min()
	return (float64(a) + float64(b)) / 2
}

func medianSlidingWindow(nums []int, k int) []float64 {
	// Brute force
	result := make([]float64, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		mf := NewSlidingMedianFinder()
		for j := 0; j < k; j++ {
			if j >= len(nums) {
				break
			}
			mf.AddNum(nums[i+j])
		}
		result = append(result, mf.FindMedian())
	}
	return result
}
