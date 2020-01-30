package leetcode;

import org.springframework.stereotype.Component;

@Component
public class Solution3 {

    boolean isrepeat(String s){
        for(int i = 0; i < s.length()-1; ++i){
            if(s.charAt(s.length() - 1) == s.charAt(i)){
                return true;
            }
        }
        return false;
    }

    public int lengthOfLongestSubstring(String s) {
        if(s.length() == 0)
            return 0;
        int size = 1;
        int maxsize = size;
        for(int i = 0; i < s.length(); ++i){
            while(i+size <= s.length() && 
                isrepeat(s.substring(i, i+size)) == false){
                if(size > maxsize) 
                    maxsize = size;
                size += 1;
            }
            size -= 1;
        }
        return maxsize;
    }
}