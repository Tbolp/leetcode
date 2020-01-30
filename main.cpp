#include <unordered_map>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> m;
        int i = 0;
        for(auto& num : nums){
            m.emplace(num, i);
            ++i;
        }
        i = 0;
        for(auto& num : nums){
            auto it = m.find(target - num);
            if(it != m.end() &&
                i != it->second){
                return {i, it->second};
            }
            ++i;
        }
    }
};