class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        unordered_set <int> map;
        for(int n : nums){
            if (map.count(n)){
                return true;
            }
            map.insert(n);
        }
        return false;
    }
};
