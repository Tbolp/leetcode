package leetcode;

import org.junit.Test;

public class Solution146Test{
    @Test
    public void Usage1(){
        LRUCache cache = new LRUCache(2);
        cache.get(2);       
        cache.put(2, 6);    
        cache.get(1);       
        cache.put(1, 5);    
        cache.put(1, 2);       
        cache.get(1);       
        cache.get(2);       
    }
}