func isSubsequence(s string, t string) bool {
    for i, j := 0, 0;; j++ {
        if len(s) <= i {
            return true
        }
        if len(t) <= j {
            return false
        }

        if s[i] == t[j] {
            i++
        }
    }
}