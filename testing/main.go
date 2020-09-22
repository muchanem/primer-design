package main


import (
	"sort"
	"fmt"
)
func sortByVal(m map[string]int) []string{
	s := make([]string, 0)	
    type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range m {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value < ss[j].Value
    })
	for _, kv := range ss {
		s = append(s, kv.Key)
	}
	return s
} 
func main() {
		amap := map[string]int{
			"something": 10,
			"yo": 20,
			"blah": 17,
		}
		fmt.Println(sortByVal(amap))
}