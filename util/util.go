package util

import (
	"sort"
	"strconv"
)

func GetRemovalCandidates(candidates []string, maxBuilds int) []string {
	num := len(candidates) - maxBuilds
	if num <= 0 {
		return []string{}
	}

	sort.Sort(byName{candidates})
	return candidates[0:num]
}

type names []string

func (s names) Len() int {
	return len(s)
}
func (s names) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type byName struct{ names }

func (s byName) Less(i, j int) bool {
	i, _ = strconv.Atoi(s.names[i])
	j, _ = strconv.Atoi(s.names[j])
	return i < j
}
