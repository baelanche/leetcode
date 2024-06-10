class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        prev_dict = {}
        
        for i, value in enumerate(nums):
            if value in prev_dict:
                if (abs(i - prev_dict[value]) <= k):
                    return True
            prev_dict[value] = i

        return False