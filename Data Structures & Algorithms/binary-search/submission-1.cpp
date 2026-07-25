class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0, rigth = nums.size() - 1;
        while(left <= rigth){
            int middle = left + ((rigth - left) / 2);
            if(target < nums[middle]){
                rigth = middle - 1;
            }else if(target > nums[middle]){
                left = middle + 1;
            }else{
                return middle;
            }
        }
        return -1;
    }
};
