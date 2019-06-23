package medium

import "container/heap"

// Given a non-empty array of integers, return the k most frequent elements.
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func topKFrequent(nums []int, k int) []int {
	numCnt := make(map[int]int)
	for _, v := range nums {
		numCnt[v] += 1
	}

	h := new(minHeap)
	var i int
	for val, freq := range numCnt {
		i++
		if i <= k {
			*h = append(*h, item{freq: freq, val: val})
			if i == k {
				heap.Init(h)
			}
		} else {
			if freq > (*h)[0].freq {
				heap.Pop(h)
				heap.Push(h, item{freq: freq, val: val})
			}
		}
	}

	var ans []int
	for h.Len() != 0 {
		ans = append(ans, heap.Pop(h).(item).val)
	}

	return ans
}

type item struct {
	freq, val int
}

type minHeap []item

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].freq < (*h)[j].freq
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(item))
}

// The official soloution:
//
// Approach 1: Heap
//
// class Solution {
// 	public List<Integer> topKFrequent(int[] nums, int k) {
// 	  // build hash map : character and how often it appears
// 	  HashMap<Integer, Integer> count = new HashMap();
// 	  for (int n: nums) {
// 		count.put(n, count.getOrDefault(n, 0) + 1);
// 	  }
//
// 	  // init heap 'the less frequent element first'
// 	  PriorityQueue<Integer> heap =
// 			  new PriorityQueue<Integer>((n1, n2) -> count.get(n1) - count.get(n2));
//
// 	  // keep k top frequent elements in the heap
// 	  for (int n: count.keySet()) {
// 		heap.add(n);
// 		if (heap.size() > k)
// 		  heap.poll();
// 	  }
//
// 	  // build output list
// 	  List<Integer> top_k = new LinkedList();
// 	  while (!heap.isEmpty())
// 		top_k.add(heap.poll());
// 	  Collections.reverse(top_k);
// 	  return top_k;
// 	}
// }
