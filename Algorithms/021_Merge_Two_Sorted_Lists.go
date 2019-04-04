// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

package main

func test021() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    l3 := new(ListNode)
    l4 := l3
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            l4.Val = l1.Val
            l4.Next = new(ListNode)
            l4 = l4.Next
            l1 = l1.Next
        } else {
            l4.Val = l2.Val
            l4.Next = new(ListNode)
            l4 = l4.Next
            l2 = l2.Next
        }
    }
    for l1 != nil {
        l4.Val = l1.Val
        if l1.Next == nil {
            break
        }
        l4.Next = new(ListNode)
        l4 = l4.Next
        l1 = l1.Next
    }
    for l2 != nil {
        l4.Val = l2.Val
        if l2.Next == nil {
            break
        }
        l4.Next = new(ListNode)
        l4 = l4.Next
        l2 = l2.Next
    }
    return l3
}