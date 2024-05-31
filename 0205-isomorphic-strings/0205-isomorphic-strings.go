func isIsomorphic(s string, t string) bool {
    keys := make(map[byte]byte)
    vals := make(map[byte]byte)
    for i:=0; i<len(s); i++ {
        key := s[i] - 'a' + 1
        val := t[i] - 'a' + 1
        if keys[key] == 0 && vals[val] == 0 {
            keys[key] = val
            vals[val] = key
        } else {
            if !(keys[key] == val && vals[val] == key) {
                return false
            }
        }
    }
    return true
}