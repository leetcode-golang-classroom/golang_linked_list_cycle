# golang_linked_list_cycle

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` *if there is a cycle in the linked list*. Otherwise, return `false`.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

```

**Example 2:**

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

```

**Example 3:**

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

```

**Constraints:**

- The number of the nodes in the list is in the range `[0, 104]`.
- `105 <= Node.val <= 105`
- `pos` is `1` or a **valid index** in the linked-list.

## 解析

題目要求給定一個單向鏈結串列，實作一個演算法判斷是否有 cycle

一個單向鏈結串列會出現 cycle ，當透過指標遇到過去走訪過的結點

從這個判斷代表可以透過紀錄走訪過的結點

每次走訪到下一個結點時，檢查是否已經走訪過即可

這樣只要實作一個 hashMap 儲存走過的結點指標

當走完整個鏈結串列都沒遇到重複的結點，代表沒有 cycle

如果有遇到代表有 cycle

以這個方式的時間複雜與空間複雜度都是 O(n)

如果要把空間複雜度要降低到 O(1)

就必須要理解一下[Floyd's algorithm](https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare)裏面所提到的特性

當一個鏈結串列具有 cycle ，從一個點出發兩個人以不同速度前進，一定會在某一個點走到同一個點

所以可以透過使用兩個 pointer，一個每次往前一格，一個每次往前兩格

當兩個 pointer都非空 檢查是否是相同的

假如有遇到相同的，代表有 cycle

假如直到其中一個是空都沒有相同， 代表沒有 cycle

## 程式碼

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
  slow := head
  if slow == nil {
      return false
  }
  fast := head.Next
  if fast == nil {
      return false
  }
  for slow != nil && fast != nil {
      if slow == fast {
          return true
      }
      slow = slow.Next
      fast = fast.Next
      if fast != nil {
          fast = fast.Next
      }
  }
  
  return false
}
```

## 困難點

1. 要理解何時會發生 cycle

## Solve Point

- [x]  Understand what problem the question would like to solve
- [x]  Analysis Complexity