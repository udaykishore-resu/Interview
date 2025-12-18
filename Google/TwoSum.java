package Google;

import java.util.HashMap;
import java.util.Map;
public class TwoSum {

    private static Map<Integer, Integer>  hasTwoSum(Integer[] nums, int k) {
       Map<Integer, Integer> resMap = new HashMap<>();
        for (int num : nums) {
            Integer comp = k - num;
            if (resMap.containsKey(comp)) {
                return Map.of(num, comp);
            }
            resMap.put(num,comp);
        }
       return Map.of();
    }

    public static  void main(String[] args){
        Integer[] nums = {10, 15, 3, 7,8, 9, 14};
        int k = 17;
        Map<Integer, Integer>  result = hasTwoSum(nums, k); 
        System.out.println(result);
    }  
}