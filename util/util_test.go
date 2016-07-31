package util

import (
	"testing"
	"reflect"
)

func TestRemovalCandidatesNone(t *testing.T) {
	maxBuilds := 0
	builds := []string{"2", "3", "4"}
	expCandidates := []string{"2", "3", "4"}
	removalHelper(t, builds, expCandidates, maxBuilds)
}

func TestRemovalCandidates(t *testing.T) {
	maxBuilds := 2
	builds := []string{"2", "3", "4"}
	expCandidates := []string {"2"}
	removalHelper(t, builds, expCandidates, maxBuilds)
}

func TestRemovalCandidatesZero(t *testing.T) {
	maxBuilds := 3
	builds := []string{"2", "3", "4"}
	expCandidates := []string{}
	removalHelper(t, builds, expCandidates, maxBuilds)
}

func removalHelper(t *testing.T, builds, expCandidates []string, maxBuilds int) {
	candidates := GetRemovalCandidates(builds, maxBuilds)
	if len(candidates) != len(expCandidates) {
		t.Errorf("Expected %d but outcome: %+v", len(expCandidates), candidates)
	}
	if !reflect.DeepEqual(candidates, expCandidates) {
		t.Errorf("Expected %+v but outcome: %+v", expCandidates, candidates)
	}
}
