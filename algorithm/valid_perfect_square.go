package main

import "fmt"

//有效的完全平方数
//leetcode: https://leetcode.com/problems/valid-perfect-square/

//使用二分查找
//Time complexity O(log n)
//Space complexity O(1)
func isPerfectSquare(num int) bool {
    if num < 1 {
        return false
    }
    left, right := 1, num

    for left <= right {
        mid := left + (right-left)/2
        //fmt.Println(left, mid, right)
        //不要使用mid*mid，因为有可能会越界
        if mid == num/mid && num%mid == 0{
            return true
        } else if mid < num/mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}

func main(){
    n := 100
    ret := isPerfectSquare(n)
    fmt.Println(ret)
}

