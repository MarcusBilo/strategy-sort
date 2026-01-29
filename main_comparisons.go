package main

import (
	"cmp"
	"fmt"
	"math/bits"
	"math/rand"
	"sort"
)

var comparisons uint64
var swaps uint64
var moves uint64

func StrategySort[T cmp.Ordered](arr []T) {
	n := len(arr)

	switch n {
	case 0, 1:
		return
	case 2:
		compareAndSwap(arr, 0, 1)
		return
	case 3:
		compareAndSwap(arr, 0, 2)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 1, 2)
		return
	case 4:
		compareAndSwap(arr, 0, 3)
		compareAndSwap(arr, 1, 2)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 1, 2)
		return
	case 5:
		compareAndSwap(arr, 0, 4)
		compareAndSwap(arr, 0, 2)
		compareAndSwap(arr, 1, 4)
		compareAndSwap(arr, 1, 3)
		compareAndSwap(arr, 2, 4)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 1, 2)
		compareAndSwap(arr, 3, 4)
		return
	case 6:
		compareAndSwap(arr, 0, 5)
		compareAndSwap(arr, 1, 3)
		compareAndSwap(arr, 2, 4)
		compareAndSwap(arr, 0, 2)
		compareAndSwap(arr, 1, 4)
		compareAndSwap(arr, 3, 5)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 4, 5)
		compareAndSwap(arr, 1, 2)
		compareAndSwap(arr, 3, 4)
		compareAndSwap(arr, 2, 4)
		return
	case 7:
		compareAndSwap(arr, 0, 6)
		compareAndSwap(arr, 1, 5)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 0, 2)
		compareAndSwap(arr, 1, 4)
		compareAndSwap(arr, 3, 6)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 3, 5)
		compareAndSwap(arr, 4, 6)
		compareAndSwap(arr, 1, 3)
		compareAndSwap(arr, 2, 4)
		compareAndSwap(arr, 5, 6)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 4, 5)
		compareAndSwap(arr, 1, 2)
		compareAndSwap(arr, 3, 4)
		return
	case 8:
		compareAndSwap(arr, 0, 7)
		compareAndSwap(arr, 1, 6)
		compareAndSwap(arr, 2, 5)
		compareAndSwap(arr, 3, 4)
		compareAndSwap(arr, 0, 2)
		compareAndSwap(arr, 1, 3)
		compareAndSwap(arr, 4, 6)
		compareAndSwap(arr, 5, 7)
		compareAndSwap(arr, 0, 1)
		compareAndSwap(arr, 2, 4)
		compareAndSwap(arr, 3, 5)
		compareAndSwap(arr, 6, 7)
		compareAndSwap(arr, 1, 3)
		compareAndSwap(arr, 4, 6)
		compareAndSwap(arr, 2, 3)
		compareAndSwap(arr, 4, 5)
		compareAndSwap(arr, 1, 2)
		compareAndSwap(arr, 3, 4)
		compareAndSwap(arr, 5, 6)
		return
	}

	if n >= 9 && n <= 100 {
		insertionSortQuad(arr)
		return
	}

	ascending := true
	for i := 1; i < n; i++ {
		comparisons++
		if arr[i] < arr[i-1] {
			ascending = false
			break
		}
	}
	if ascending {
		return
	}

	descending := true
	for i := 1; i < n; i++ {
		comparisons++
		if arr[i] > arr[i-1] {
			descending = false
			break
		}
	}
	if descending {
		for i := 0; i < n/2; i++ {
			swaps++
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		return
	}

	posDisorder, negDisorder := 0, 0
	threshold := n / 10

	for i := 6; i < n; i += 2 {
		comparisons++
		if arr[i] < arr[i-2] {
			comparisons++
			if arr[i] < arr[i-4] {
				comparisons++
				if arr[i] < arr[i-6] {
					posDisorder++
				}
			}
		}

		comparisons++
		if arr[i] > arr[i-2] {
			comparisons++
			if arr[i] > arr[i-4] {
				comparisons++
				if arr[i] > arr[i-6] {
					negDisorder++
				}
			}
		}

		if posDisorder > threshold && negDisorder > threshold {
			dualPivotQS(arr)
			return
		}
	}

	if posDisorder*100 < negDisorder {
		mergeIterative(arr)
		return
	}
	if negDisorder*100 < posDisorder {
		for k := 0; k < n/2; k++ {
			swaps++
			arr[k], arr[n-1-k] = arr[n-1-k], arr[k]
		}
		mergeIterative(arr)
		return
	}

	for i := 7; i < n; i += 2 {
		comparisons++
		if arr[i] < arr[i-2] {
			comparisons++
			if arr[i] < arr[i-4] {
				comparisons++
				if arr[i] < arr[i-6] {
					posDisorder++
				}
			}
		}

		comparisons++
		if arr[i] > arr[i-2] {
			comparisons++
			if arr[i] > arr[i-4] {
				comparisons++
				if arr[i] > arr[i-6] {
					negDisorder++
				}
			}
		}
	}

	if posDisorder*50 <= negDisorder {
		mergeIterative(arr)
		return
	}
	if negDisorder*50 <= posDisorder {
		for i := 0; i < n/2; i++ {
			swaps++
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		mergeIterative(arr)
		return
	}

	dualPivotQS(arr)
}

func compareAndSwap[T cmp.Ordered](arr []T, i, j int) {
	comparisons++
	if arr[i] > arr[j] {
		swaps++
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func insertionSortQuadRange[T cmp.Ordered](a []T, low, high int) {
	if high-low < 1 {
		return
	}

	i := low + 1

	for ; i+3 <= high; i += 4 {
		x0 := a[i]
		x1 := a[i+1]
		x2 := a[i+2]
		x3 := a[i+3]

		// 4-element sorting network (5 comparisons max)
		comparisons++
		if x0 > x3 {
			swaps++
			x0, x3 = x3, x0
		}
		comparisons++
		if x1 > x2 {
			swaps++
			x1, x2 = x2, x1
		}
		comparisons++
		if x0 > x1 {
			swaps++
			x0, x1 = x1, x0
		}
		comparisons++
		if x2 > x3 {
			swaps++
			x2, x3 = x3, x2
		}
		comparisons++
		if x1 > x2 {
			swaps++
			x1, x2 = x2, x1
		}

		j := i - 1

		for j >= low {
			comparisons++
			if a[j] <= x3 {
				break
			}
			moves++
			a[j+4] = a[j]
			j--
		}
		a[j+4] = x3

		for j >= low {
			comparisons++
			if a[j] <= x2 {
				break
			}
			moves++
			a[j+3] = a[j]
			j--
		}
		a[j+3] = x2

		for j >= low {
			comparisons++
			if a[j] <= x1 {
				break
			}
			moves++
			a[j+2] = a[j]
			j--
		}
		a[j+2] = x1

		for j >= low {
			comparisons++
			if a[j] <= x0 {
				break
			}
			moves++
			a[j+1] = a[j]
			j--
		}
		a[j+1] = x0
	}

	for ; i <= high; i++ {
		x := a[i]
		j := i - 1
		for j >= low {
			comparisons++
			if a[j] <= x {
				break
			}
			moves++
			a[j+1] = a[j]
			j--
		}
		a[j+1] = x
	}
}

func insertionSortQuad[T cmp.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}
	insertionSortQuadRange(a, 0, len(a)-1)
}

func mergeIterative[T cmp.Ordered](a []T) {
	const bufferSize = 320
	const minRun = 40
	n := len(a)
	if n <= 1 {
		return
	}
	var buf [bufferSize]T

	for i := 0; i < n; i += minRun {
		end := i + minRun
		if end > n {
			end = n
		}
		insertionSortQuadRange(a, i, end-1)
	}

	width := minRun
	for width < n {
		i := 0
		for i+width < n {
			from := i
			mid := i + width
			to := i + 2*width
			if to > n {
				to = n
			}

			comparisons++
			if a[mid-1] <= a[mid] {
				i += 2 * width
				continue
			}

			len1 := mid - from
			len2 := to - mid

			if len1 <= bufferSize {
				mergeLeft(a, from, mid, to, &buf)
			} else if len2 <= bufferSize {
				mergeRight(a, from, mid, to, &buf)
			} else {
				mergeIterativeFallback(a, from, mid, to, &buf)
			}

			i += 2 * width
		}
		width *= 2
	}
}

type task struct {
	from, pivot, to int
}

func mergeIterativeFallback[T cmp.Ordered](a []T, from, pivot, to int, buf *[320]T) {
	const bufferSize = 320
	var stack [64]task
	sp := 0
	stack[sp] = task{from, pivot, to}
	sp++

	for sp > 0 {
		sp--
		t := stack[sp]

		from, pivot, to = t.from, t.pivot, t.to
		len1 := pivot - from
		len2 := to - pivot

		if len1 == 0 || len2 == 0 {
			continue
		}

		comparisons++
		if a[pivot-1] <= a[pivot] {
			continue
		}

		for len1 > 0 {
			comparisons++
			if a[from+len1-1] <= a[pivot] {
				len1--
			} else {
				break
			}
		}

		for len2 > 0 {
			comparisons++
			if a[pivot-1] <= a[pivot] {
				pivot++
				len2--
			} else {
				break
			}
		}

		comparisons++
		if a[to-1] < a[from] {
			rotateContrevRange(a, from, pivot, to)
			continue
		}

		if len1 <= bufferSize {
			mergeLeft(a, from, pivot, to, buf)
			continue
		}
		if len2 <= bufferSize {
			mergeRight(a, from, pivot, to, buf)
			continue
		}

		var firstCut, secondCut int
		var len22 int

		if len1 > len2 {
			firstCut = from + len1/2
			secondCut = lower2(a, pivot, to, firstCut)
			len22 = secondCut - pivot
		} else {
			len22 = len2 / 2
			secondCut = pivot + len22
			firstCut = upper2(a, from, pivot, secondCut)
		}

		for firstCut < pivot && pivot < secondCut {
			comparisons++
			if a[pivot-1] <= a[pivot] {
				pivot++
			} else {
				break
			}
		}

		for firstCut < pivot && pivot < secondCut {
			comparisons++
			if a[firstCut] <= a[pivot] {
				firstCut++
			} else {
				break
			}
		}

		if firstCut < pivot && pivot < secondCut {
			rotL := pivot - firstCut
			rotR := secondCut - pivot
			if min(rotL, rotR) <= bufferSize {
				rotateBuf(a, firstCut, pivot, secondCut, buf)
			} else {
				rotateContrevRange(a, firstCut, pivot, secondCut)
			}
		}

		newMid := firstCut + len22

		stack[sp] = task{
			from:  newMid,
			pivot: secondCut,
			to:    to,
		}
		sp++

		stack[sp] = task{
			from:  from,
			pivot: firstCut,
			to:    newMid,
		}
		sp++
	}
}

func lower2[T cmp.Ordered](a []T, from, to, val int) int {
	length := to - from
	for length > 0 {
		half := length / 2
		mid := from + half
		comparisons++
		if a[mid] < a[val] {
			from = mid + 1
			length -= half + 1
		} else {
			length = half
		}
	}
	return from
}

func upper2[T cmp.Ordered](a []T, from, to, val int) int {
	length := to - from
	for length > 0 {
		half := length / 2
		mid := from + half
		comparisons++
		if a[val] < a[mid] {
			length = half
		} else {
			from = mid + 1
			length -= half + 1
		}
	}
	return from
}

func rotateContrevRange[T any](a []T, from, mid, to int) {
	rotateContrev(a[from:to], mid-from, to-mid)
}

// Conjoined Triple Reversal rotation (Igor van den Hoven)
func rotateContrev[T any](a []T, left, right int) {
	if left == 0 || right == 0 {
		return
	}

	pta := 0
	ptb := left
	ptc := left
	ptd := left + right

	var loop int
	var swap T

	if left > right {
		loop = right / 2

		for loop > 0 {
			loop--
			ptb--
			ptd--

			moves += 4
			swap = a[ptb]
			a[ptb] = a[pta]
			a[pta] = a[ptc]
			a[ptc] = a[ptd]
			a[ptd] = swap

			pta++
			ptc++
		}

		loop = (ptb - pta) / 2

		for loop > 0 {
			loop--
			ptb--
			ptd--

			moves += 3
			swap = a[ptb]
			a[ptb] = a[pta]
			a[pta] = a[ptd]
			a[ptd] = swap

			pta++
		}

		loop = (ptd - pta) / 2

		for loop > 0 {
			loop--
			ptd--

			moves += 2
			swap = a[pta]
			a[pta] = a[ptd]
			a[ptd] = swap

			pta++
		}

	} else if left < right {
		loop = left / 2

		for loop > 0 {
			loop--
			ptb--
			ptd--

			moves += 4
			swap = a[ptb]
			a[ptb] = a[pta]
			a[pta] = a[ptc]
			a[ptc] = a[ptd]
			a[ptd] = swap

			pta++
			ptc++
		}

		loop = (ptd - ptc) / 2

		for loop > 0 {
			loop--
			ptd--

			moves += 3
			swap = a[ptc]
			a[ptc] = a[ptd]
			a[ptd] = a[pta]
			a[pta] = swap

			pta++
			ptc++
		}

		loop = (ptd - pta) / 2

		for loop > 0 {
			loop--
			ptd--

			moves += 2
			swap = a[pta]
			a[pta] = a[ptd]
			a[ptd] = swap

			pta++
		}

	} else {
		loop = left

		for loop > 0 {
			loop--

			moves += 2
			swap = a[pta]
			a[pta] = a[ptb]
			a[ptb] = swap

			pta++
			ptb++
		}
	}
}

func rotateBuf[T cmp.Ordered](a []T, from, mid, to int, buf *[320]T) {
	len1 := mid - from
	len2 := to - mid

	if len1 <= len2 {
		b := buf[:len1]
		moves += uint64(len1)
		copy(b, a[from:mid])
		moves += uint64(len2)
		copy(a[from:], a[mid:to])
		moves += uint64(len1)
		copy(a[from+len2:], b)
	} else {
		b := buf[:len2]
		moves += uint64(len2)
		copy(b, a[mid:to])
		moves += uint64(len1)
		copy(a[from+len2:], a[from:mid])
		moves += uint64(len2)
		copy(a[from:], b)
	}
}

func mergeRight[T cmp.Ordered](a []T, from, mid, to int, buf *[320]T) {
	len2 := to - mid
	b := buf[:len2]
	moves += uint64(len2)
	copy(b, a[mid:to])

	i := mid - 1
	j := len2 - 1
	k := to - 1

	for j >= 1 && i >= from+1 {
		comparisons++
		if a[i] > buf[j] {
			moves++
			a[k] = a[i]
			i--
		} else {
			moves++
			a[k] = buf[j]
			j--
		}
		k--

		comparisons++
		if a[i] > buf[j] {
			moves++
			a[k] = a[i]
			i--
		} else {
			moves++
			a[k] = buf[j]
			j--
		}
		k--
	}

	for j >= 0 && i >= from {
		comparisons++
		if a[i] > buf[j] {
			moves++
			a[k] = a[i]
			i--
		} else {
			moves++
			a[k] = buf[j]
			j--
		}
		k--
	}

	if j >= 0 {
		copy(a[k-j:k+1], buf[:j+1])
	}
}

func mergeLeft[T cmp.Ordered](a []T, from, mid, to int, buf *[320]T) {
	len1 := mid - from
	b := buf[:len1]
	moves += uint64(len1)
	copy(b, a[from:mid])

	i := 0
	j := mid
	k := from

	for i < len1-1 && j < to-1 {
		comparisons++
		if buf[i] <= a[j] {
			moves++
			a[k] = buf[i]
			i++
		} else {
			moves++
			a[k] = a[j]
			j++
		}
		k++

		comparisons++
		if buf[i] <= a[j] {
			moves++
			a[k] = buf[i]
			i++
		} else {
			moves++
			a[k] = a[j]
			j++
		}
		k++
	}

	for i < len1 && j < to {
		comparisons++
		if buf[i] <= a[j] {
			moves++
			a[k] = buf[i]
			i++
		} else {
			moves++
			a[k] = a[j]
			j++
		}
		k++
	}

	if i < len1 {
		copy(a[k:k+(len1-i)], buf[i:len1])
	}
}

func dualPivotQS[T cmp.Ordered](a []T) {
	n := len(a)

	cutoff := 2*bits.Len(uint(n)) - 1 // approximately 2*log2(n)
	const insertionThreshold = 100

	var stack [512]int
	top := 0
	stack[top], stack[top+1] = 0, n-1
	top += 2

	for top > 0 {
		top -= 2
		low, high := stack[top], stack[top+1]
		if low >= high {
			continue
		}

		if top >= cutoff {
			shellSortPardons09Range(a, low, high)
			continue
		}

		if high-low+1 <= insertionThreshold {
			insertionSortQuadRange(a, low, high)
			continue
		}

		if allEqualRange(a, low, high) {
			continue
		}

		dualPivotMedianOfMedians(a, low, high)
		p1, p2 := a[low], a[high]

		comparisons++
		if p1 != p2 && highDuplicateDensity(a, low, high, p1, p2) {
			threeWayPartitionPush(a, low, high, p1, &stack, &top)
			continue
		}

		lt := low + 1
		gt := high - 1
		i := lt

		for i <= gt {
			x := a[i]

			comparisons++
			if x < p1 {
				if i != lt {
					swaps++
					a[i], a[lt] = a[lt], x
				}
				lt++
				i++
				continue
			}

			comparisons++
			if x > p2 {
				if i != gt {
					swaps++
					a[i], a[gt] = a[gt], x
				}
				gt--
				continue
			}

			i++
		}

		lt--
		gt++
		swaps++
		a[low], a[lt] = a[lt], a[low]

		swaps++
		a[high], a[gt] = a[gt], a[high]

		pushDualPivotStack(low, lt-1, lt+1, gt-1, gt+1, high, &stack, &top)
	}
}

func shellSortPardons09Range[T cmp.Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}
	n := high - low + 1

	gaps := []int{1031612713, 217378076, 45806244, 9651787, 2034035, 428481, 90358, 19001, 4025, 836, 182, 34, 9, 1}

	for _, gap := range gaps {
		if gap >= n {
			continue
		}
		for i := low + gap; i <= high; i++ {
			temp := arr[i]
			j := i
			for j-gap >= low {
				comparisons++
				if arr[j-gap] <= temp {
					break
				}
				moves++
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

func dualPivotMedianOfMedians[T cmp.Ordered](a []T, low, high int) {
	n := high - low + 1
	if n < 9 {
		mid := low + (high-low)/2
		medianOf3(a, low, mid, high)
		return
	}

	step := n / 8

	i1 := low
	i2 := low + step
	i3 := low + 2*step
	i4 := low + 3*step
	i5 := low + 4*step
	i6 := low + 5*step
	i7 := low + 6*step
	i8 := low + 7*step
	i9 := high

	medianOf3(a, i1, i2, i3)
	medianOf3(a, i4, i5, i6)
	medianOf3(a, i7, i8, i9)

	medianOf3(a, i2, i5, i8)

	a[low], a[i2] = a[i2], a[low]
	a[high], a[i8] = a[i8], a[high]

	comparisons++
	if a[high] < a[low] {
		swaps++
		a[low], a[high] = a[high], a[low]
	}
}

func medianOf3[T cmp.Ordered](a []T, low, mid, high int) {
	comparisons++
	if a[mid] < a[low] {
		swaps++
		a[mid], a[low] = a[low], a[mid]
	}
	comparisons++
	if a[high] < a[mid] {
		swaps++
		a[high], a[mid] = a[mid], a[high]
	}
	comparisons++
	if a[mid] < a[low] {
		swaps++
		a[mid], a[low] = a[low], a[mid]
	}
}

func allEqualRange[T cmp.Ordered](a []T, low, high int) bool {
	x := a[low]

	limit := low + 32
	if limit > high {
		limit = high
	}

	for i := low + 1; i <= limit; i++ {
		comparisons++
		if a[i] != x {
			return false
		}
	}
	for i := limit + 1; i <= high; i++ {
		comparisons++
		if a[i] != x {
			return false
		}
	}
	return true
}

func highDuplicateDensity[T cmp.Ordered](a []T, low, high int, p1, p2 T) bool {
	eqCount := 0
	sampleSize := high - low + 1
	if sampleSize > 16 {
		sampleSize = 16
	}
	for i := low; i < low+sampleSize; i++ {
		comparisons++
		if a[i] == p1 || a[i] == p2 {
			eqCount++
		}
	}

	return float64(eqCount)/float64(sampleSize) >= 0.9 || p1 == p2
}

func threeWayPartitionPush[T cmp.Ordered](a []T, low, high int, pivot T, stack *[512]int, top *int) {
	lt, i, gt := low, low, high
	for i <= gt {
		comparisons++
		if a[i] < pivot {
			a[i], a[lt] = a[lt], a[i]
			lt++
			i++
		} else {
			comparisons++
			if a[i] > pivot {
				a[i], a[gt] = a[gt], a[i]
				gt--
			} else {
				i++
			}
		}
	}

	leftSize := lt - low
	rightSize := high - gt

	if leftSize < rightSize {
		if gt+1 <= high {
			stack[*top], stack[*top+1] = gt+1, high
			*top += 2
		}
		if low < lt-1 {
			stack[*top], stack[*top+1] = low, lt-1
			*top += 2
		}
	} else {
		if low < lt-1 {
			stack[*top], stack[*top+1] = low, lt-1
			*top += 2
		}
		if gt+1 <= high {
			stack[*top], stack[*top+1] = gt+1, high
			*top += 2
		}
	}
}

func pushDualPivotStack(loLeft, hiLeft, loMid, hiMid, loRight, hiRight int, stack *[512]int, top *int) {
	leftSize := hiLeft - loLeft + 1
	midSize := hiMid - loMid + 1
	rightSize := hiRight - loRight + 1

	if leftSize < rightSize {
		if rightSize > 0 {
			stack[*top], stack[*top+1] = loRight, hiRight
			*top += 2
		}
		if midSize > 0 {
			stack[*top], stack[*top+1] = loMid, hiMid
			*top += 2
		}
		if leftSize > 0 {
			stack[*top], stack[*top+1] = loLeft, hiLeft
			*top += 2
		}
	} else {
		if leftSize > 0 {
			stack[*top], stack[*top+1] = loLeft, hiLeft
			*top += 2
		}
		if midSize > 0 {
			stack[*top], stack[*top+1] = loMid, hiMid
			*top += 2
		}
		if rightSize > 0 {
			stack[*top], stack[*top+1] = loRight, hiRight
			*top += 2
		}
	}
}

func checkSort(name string, sortFunc func([]int), arr []int) {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	sortFunc(arrCopy)

	stdSorted := make([]int, len(arr))
	copy(stdSorted, arr)
	sort.Ints(stdSorted)

	correct := true
	for i := range arr {
		if arrCopy[i] != stdSorted[i] {
			correct = false
			break
		}
	}

	if correct {
		fmt.Printf("%s: Sorting correct!\n", name)
	} else {
		fmt.Printf("%s: Sorting incorrect!\n", name)
	}
}

func main() {
	size := 1 * 1000 // * 1000
	arr := make([]int, size)
	for i := range arr {
		// arr[i] = rand.Intn(size) - 100000 // 1M
		// arr[i] = rand.Intn(size) - 100 // 1K
		// arr[i] = i
		// arr[i] = size - i
		// arr[i] = rand.Intn(10000) - 5000 // 1M
		arr[i] = rand.Intn(10) - 5 // 1K
	}

	/*
		for i := 0; i < size/10; i++ {
			pos1 := rand.Intn(size-6) + 3   // pos1 ∈ [3, size-3)
			pos2 := pos1 - 3 + rand.Intn(7) // 7 values: pos1-3 .. pos1+3
			arr[pos1], arr[pos2] = arr[pos2], arr[pos1]
		}

		for i := 0; i < size/10; i++ {
			a, b := rand.Intn(size), rand.Intn(size)
			arr[a], arr[b] = arr[b], arr[a]
		}
	*/

	sorts := []struct {
		name string
		fn   func([]int)
	}{
		// {"dualPivotQS", dualPivotQS[int]},
		// {"mergeIterative", mergeIterative[int]},
		{"StrategySort", StrategySort[int]},
	}

	for _, s := range sorts {
		checkSort(s.name, s.fn, arr)
	}

	fmt.Println(comparisons, "comparisons")
}
