class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cnt = 0
        output = 0
        for num in nums:
            if num == output:
                cnt += 1
            else:
                if cnt < 1:
                    output = num
                else:
                    cnt -= 1
        return output
        