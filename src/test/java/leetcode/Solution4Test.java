package leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootApplication
public class Solution4Test{
    @Autowired
    private Solution4 solution;
    @Test
    public void Usage1(){
        assertEquals(2, solution.findMedianSortedArrays(new int[]{1, 3}, new int[]{2}));
    }
}