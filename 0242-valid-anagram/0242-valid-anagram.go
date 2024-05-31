func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    alps := make([]int, 26)
    for i:=0; i<len(s); i++ {
        alps[s[i]-'a']++
        alps[t[i]-'a']--
    }
    for _, alp := range alps {
        if alp != 0 {
            return false
        }
    }

    return true
}