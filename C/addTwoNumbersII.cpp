#include <iostream>
#include <stack>

using namespace std;

/**
 * Definition for singly-linked list.
 */
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
 
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        if (NULL == l1 && NULL == l2) {
            return NULL;
        }
        if (NULL == l1 && NULL != l2) {
            return l2;
        }
        if (NULL != l1 && NULL == l2) {
            return l1;
        }

        stack<int> stack1;
        stack<int> stack2;
        stack<int> stack3;

        while (l1) {
            stack1.push(l1->val);
            l1 = l1->next;
        }
        while (l2) {
            stack2.push(l2->val);
            l2 = l2->next;
        }

        int flag = 0;
        while (!stack1.empty() || !stack2.empty()) {
            int x = (stack1.empty() ? 0 : stack1.top());
            int y = (stack2.empty() ? 0 : stack2.top());

            int sum = x + y + flag;
            stack3.push(sum % 10);
            flag = sum / 10;

            stack1.pop();
            stack2.pop();
        }

        if (flag > 0) {
            stack3.push(flag);
        }

        ListNode *lret = new ListNode(0);
        ListNode *tmp = lret;

        while (!stack3.empty()) {
            tmp->next = new ListNode(stack3.top());
            stack3.pop();
            tmp = tmp->next;
        }

        return lret->next;
    }
};

void print(ListNode *l)
{
    if (NULL == l) {
        return;
    }
    while (l) {
        cout << l->val << " ";
        l = l->next;
    }
    cout << endl;
}

int main(int argc, char const *argv[])
{
    Solution s;
    int a1[] = {7,2,4,3};
    int a2[] = {5,6,4};

    ListNode *l1 = new ListNode(0);
    ListNode *l2 = new ListNode(0);
    ListNode *tmp = l1;
    for (int i = 0; i < 4; i ++) {
        tmp->next = new ListNode(a1[i]);
        tmp = tmp->next;
    }
    tmp = l2;
    for (int i = 0; i < 3; i ++) {
        tmp->next = new ListNode(a2[i]);
        tmp = tmp->next;
    }
    // print(l1->next);
    // print(l2->next);

    ListNode *ret = s.addTwoNumbers(l1->next, l2->next);
    print(ret);
    return 0;
}
