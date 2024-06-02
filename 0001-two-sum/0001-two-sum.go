func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    ans := make([]int, 2)

    for i, num := range nums {
        m[num] = i
    }

    for i := 0; i<len(nums); i++ {
        sub := target - nums[i]
        _, exists := m[sub]
        if exists && m[sub] != i {
            ans[0] = i
            ans[1] = m[sub]
            return ans
        }
    }
    return ans
}