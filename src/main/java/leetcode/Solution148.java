package leetcode;

import org.springframework.stereotype.Component;

@Component
public class Solution148 {
    static void swap(ListNode start, ListNode n1p, ListNode n2p) {
        if (n1p != null && n2p != null) {
            ListNode n1 = n1p.next;
            ListNode n2 = n2p.next;
            ListNode n1n = n1.next;
            ListNode n2n = n2.next;
            n1p.next = n2;
            n2.next = n1n;
            n2p.next = n1;
            n1.next = n2n;
        } else if (n1p == null) {
            ListNode n1 = start;
            ListNode n2 = n2p.next;
            ListNode n1n = n1.next;
            ListNode n2n = n2.next;
            n2.next = n1n;
            n2p.next = n1;
            n1.next = n2n;
        } else if (n2p == null) {
            ListNode n1 = n1p.next;
            ListNode n2 = start;
            ListNode n1n = n1.next;
            ListNode n2n = n2.next;
            n1p.next = n2;
            n2.next = n1n;
            n1.next = n2n;
        }
    }

    public ListNode sortList(ListNode head) {
        ListNode n1 = head;
        ListNode n1p = null;
        while (n1 != null) {
            ListNode n2 = n1.next;
            ListNode n2p = n1;
            while (n2 != null) {
                if (n1.val > n2.val) {
                    swap(head, n1p, n2p);
                    n2 = n1p.next;
                }
                n2p = n2;
                n2 = n2.next;
            }
            n1p = n1;
            n1 = n1.next;
        }
        return head;
    }
    
}
