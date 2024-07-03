package main


// https://leetcode.cn/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    len := m+n
    end1 := m-1
    end2 := n-1



    for i := len-1; i >= 0; i-- {
    	if end1 >= 0 && (end2 < 0 || nums1[end1] > nums2[end2]){
    		nums1[i] = nums1[end1]
    		end1 -- 
    	}else{
    		nums1[i] = nums2[end2]
    		end2--
    	}
    }
}