package leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootApplication
public class Solution45Test{
    @Autowired
    private Solution45 solution;
    @Test
    public void Usage1(){
        assertEquals(2, solution.jump(new int[]{2, 3, 1, 1, 4}));
    }
    @Test
    public void Usage2(){
        assertEquals(1, solution.jump(new int[]{6, 1, 1, 1, 4}));
    }
    @Test
    public void Usage3(){
        assertEquals(5, solution.jump(new int[]{5,6,4,4,6,9,4,4,7,4,4,8,2,6,8,1,5,9,6,5,2,7,9,7,9,6,9,4,1,6,8,8,4,4,2,0,3,8,5}));
    }
    @Test
    public void Usage4(){
        assertEquals(5, solution.jump(new int[]{1, 1, 1, 1, 1 ,1}));
    }
    @Test
    public void Usage5(){
        int[] nums= new int[24999];
        for(int i = 0; i < nums.length; ++i)
            nums[i] = 1;
        assertEquals(24998, solution.jump(nums));
    }
    @Test
    public void Usage6(){
        assertEquals(2, solution.jump(new int[]{8,7,6,5,4,3,2,1,1,0}));
    }
}