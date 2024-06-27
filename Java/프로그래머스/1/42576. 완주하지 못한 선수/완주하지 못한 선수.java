import java.util.*;

class Solution {
    public String solution(String[] participant, String[] completion) {
        String answer = "";
        Map<String, Integer> cm = new HashMap<>();
        
        for (String p : participant) {
            cm.put(p, cm.getOrDefault(p, 0) + 1);
        }
        
        for (String c : completion) {
            cm.put(c, cm.get(c) - 1);
        }
        
        for (String k : cm.keySet()){
            if (cm.get(k) != 0) {
                answer = k;
                break;
            }
        }
        return answer;
    }
}