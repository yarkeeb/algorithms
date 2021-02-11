package linked_list

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	if res := MergeTwoLists(l1, l2); res.ToString() != l3.ToString() {
		t.Errorf("MergeTwoLists result incorrect, got: %s expected: %s", res.ToString(), l3.ToString())
	}
}
