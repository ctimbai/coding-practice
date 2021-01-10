#include <iostream>
#include <vector>
#include <climits>
#include <cmath>

using namespace std;

class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        int size = nums.size();
        if (size < 3) {
            return INT_MIN;
        }
        
        int min = INT_MAX, ret = 0;
        int i = 0;
        while (i < size) {
            int j = i + 1, k = size - 1;
            while (j < k) {
                int sum = nums[i] + nums[j] + nums[k];
                int val = abs(sum - target);
                if (val < min) {
                    min = val;
                    ret = sum;
                }
                k --;
            }
            i ++;
        }
        return ret;
    }
};

int main(int argc, char const *argv[])
{
    Solution s;
    int a[] = {1,2,4,8,16,32,64,128};
    vector<int> vec(a, a+8);
    cout << s.threeSumClosest(vec, 82);
    return 0;
}
