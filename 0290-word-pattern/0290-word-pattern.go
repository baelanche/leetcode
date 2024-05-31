import (
    "strings"
)

func wordPattern(pattern string, s string) bool {
    keys := make(map[string]string)
    vals := make(map[string]string)
    snippets := strings.Split(s, " ")
    if len(pattern) != len(snippets) {
        return false
    }

    for i, snippet := range snippets {
        key := string(pattern[i])
        if keys[key] == "" && vals[snippet] == "" {
            keys[key] = snippet
            vals[snippet] = key
        } else {
            if !(keys[key] == snippet && vals[snippet] == key) {
                return false
            }
        }
    }
    return true
}