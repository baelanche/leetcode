func majorityElement(nums []int) int {
    m := make(map[int]int)

    for _, n := range nums {
        m[n] += 1
    }

    var ans, cnt = 0, 0
    for i, val := range m {
        if cnt < val {
            cnt = val
            ans = i
        }
    }

    return ans
}