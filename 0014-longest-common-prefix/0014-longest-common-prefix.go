func longestCommonPrefix(strs []string) string {
    common := strs[0]
    for _, str := range strs {
        if len(common) != len(str) {
            if len(common) < len(str) {
                common = common[:len(common)]
            } else {
                common = common[:len(str)]
            }
        }
        for j:=0; j<len(common) && j<len(str); j++ {
            if common[j] != str[j] {
                common = common[:j]
            }
        }
    }
    return common
}