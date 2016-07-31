package util

import (
	"github.com/almonteb/buildmaid/fileman"
	"sort"
	"strconv"
)

func GetRemovalCandidates(candidates []fileman.Build, maxBuilds int) []fileman.Build {
	num := len(candidates) - maxBuilds
	if num <= 0 {
		return []fileman.Build{}
	}

	sort.Sort(byName{candidates})
	return candidates[0:num]
}

func FilterIgnores(candidates []fileman.Build, ignores []string) []fileman.Build {
	if len(ignores) == 0 {
		return candidates
	}

	for i := len(candidates) - 1; i >= 0; i-- {
		for _, ignore := range ignores {
			if ignore == candidates[i].Name {
				candidates = append(candidates[:i], candidates[i+1:]...)
			}
		}
	}
	return candidates
}

type names []fileman.Build

func (s names) Len() int {
	return len(s)
}
func (s names) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type byName struct{ names }

func (s byName) Less(i, j int) bool {
	i, _ = strconv.Atoi(s.names[i].Name)
	j, _ = strconv.Atoi(s.names[j].Name)
	return i < j
}
