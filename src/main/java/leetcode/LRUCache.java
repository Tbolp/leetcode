package leetcode;

import java.util.HashMap;
import java.util.LinkedList;

public class LRUCache {

    private HashMap<Integer, Integer> _val;
    private LinkedList<Integer> _order;
    int _capacity;

    public LRUCache(int capacity) {
        _val = new HashMap<Integer, Integer>(capacity);
        _order = new LinkedList<Integer>();
        _capacity = capacity;
    }

    public int get(int key) {
        if (_val.containsKey(key) == false) {
            return -1;
        } else {
            _order.removeLastOccurrence(key);
            _order.add(key);
            return _val.get(key);
        }
    }

    public void put(int key, int value) {
        if(_val.containsKey(key)){
            _val.put(key, value);
            _order.removeLastOccurrence(key);
            _order.add(key);
            return;
        }
        if(_val.size() == _capacity){
            Integer val = _order.remove();
            _val.remove(val);
        }
        _val.put(key, value);
        _order.add(key);
    }
}