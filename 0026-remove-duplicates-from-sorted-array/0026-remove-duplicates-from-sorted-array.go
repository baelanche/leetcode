func removeDuplicates(nums []int) int {
    var i, j = 0, 0
    for j < len(nums) {
        if i == j {
            j++
        } else {
            if nums[i] == nums[j] {
                j++
            } else {
                i++
                nums[i] = nums[j]
            }
        }
    }
    return i+1
}