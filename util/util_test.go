package util

import (
	"github.com/almonteb/buildmaid/fileman"
	"reflect"
	"testing"
)

func TestRemovalCandidatesNone(t *testing.T) {
	maxBuilds := 0
	builds := []fileman.Build{getBuild("2"), getBuild("3"), getBuild("4")}
	expCandidates := []fileman.Build{getBuild("2"), getBuild("3"), getBuild("4")}
	assertCandidates(t, GetRemovalCandidates(builds, maxBuilds), expCandidates)
}

func TestRemovalCandidates(t *testing.T) {
	maxBuilds := 2
	builds := []fileman.Build{getBuild("2"), getBuild("3"), getBuild("4")}
	expCandidates := []fileman.Build{getBuild("2")}
	assertCandidates(t, GetRemovalCandidates(builds, maxBuilds), expCandidates)
}

func TestRemovalCandidatesZero(t *testing.T) {
	maxBuilds := 3
	builds := []fileman.Build{getBuild("2"), getBuild("3"), getBuild("4")}
	expCandidates := []fileman.Build{}
	assertCandidates(t, GetRemovalCandidates(builds, maxBuilds), expCandidates)
}

func TestFilterIgnore(t *testing.T) {
	builds := []fileman.Build{getBuild("latest"), getBuild("2"), getBuild("3"), getBuild("test")}
	ignores := []string{"latest", "3"}
	exp := []fileman.Build{getBuild("2"), getBuild("test")}
	filtered := FilterIgnores(builds, ignores)
	assertCandidates(t, filtered, exp)
}

func getBuild(name string) fileman.Build {
	return fileman.Build{Name: name, Path: "dummyPath"}
}

func assertCandidates(t *testing.T, candidates, expCandidates []fileman.Build) {
	if len(candidates) != len(expCandidates) {
		t.Errorf("Expected %d but outcome: %+v", len(expCandidates), candidates)
	}
	if !reflect.DeepEqual(candidates, expCandidates) {
		t.Errorf("Expected %+v but outcome: %+v", expCandidates, candidates)
	}
}
