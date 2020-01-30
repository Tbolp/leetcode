package leetcode;

import org.springframework.stereotype.Component;

@Component
public class Solution45{
    private int jumpcache(int len, int[] nums, int[] cache){
        if(cache[len - 1] != -1)
            return cache[len - 1];
        int res = len - 1;
        for(int i = len-2; i >= 0; --i){
            if(cache[i] == -1) continue;
            if(nums[i] + i >= len - 1){
                int step = jumpcache(i + 1, nums, cache) + 1;
                if(res > step) res = step;
            }
        }
        cache[len - 1] = res;
        return res;
    }
    public int jump(int[] nums) {
        int[] cache = nums.clone();
        for(int i = 0; i < cache.length; ++i){
            cache[i] = -1;
        }
        cache[0] = 0;
        int step = 0;
        for(int i = 0; i < cache.length; ++i){
            if(nums[i] + i > step ||
                nums[i] + i >= cache.length-1){
                jumpcache(i + 1, nums, cache);
                step = nums[i] + i;
            }
        }
        return jumpcache(nums.length, nums, cache);
    }
}