package JZ_Offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	ans := make([]int, 0)

	var dfs func(head *ListNode)
	dfs = func(head *ListNode) {
		if head == nil {
			return
		}
		dfs(head.Next)
		ans = append(ans, head.Val)
	}
	dfs(head)
	return ans
}
