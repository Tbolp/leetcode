package leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootApplication
public class Solution3Test{
    @Autowired
    private Solution3 solution;
    @Test
    public void Usage1(){
        assertEquals(3, solution.lengthOfLongestSubstring("abcabcbb"));
    }
    @Test
    public void Usage2(){
        assertEquals(3, solution.lengthOfLongestSubstring("pwwkew"));
    }
    @Test
    public void Usage3(){
        assertEquals(1, solution.lengthOfLongestSubstring("bbbbb"));
    }
    @Test
    public void Usage4(){
        assertEquals(0, solution.lengthOfLongestSubstring(""));
    }
    @Test
    public void Usage5(){
        assertEquals(2, solution.lengthOfLongestSubstring("au"));
    }
}