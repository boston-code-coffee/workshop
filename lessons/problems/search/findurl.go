package search

import (
	"fmt"
	"sort"
)

// findUrl returns the string target if its present in the sorted list.
func findUrlSet(logUrls []string, target string) string {
	urlSet := make(map[string]struct{})
	for _, s := range logUrls {
		urlSet[s] = struct{}{}
	}
	if _, ok := urlSet[target]; ok {
		return target
	}
	return ""

}
func findUrlBinary(logUrls []string, target string) string {
	first := 0
	last := len(logUrls) - 1
	var mid int
	for first <= last {
		mid = first + (last-first)/2
		s := logUrls[mid]
		switch {
		case target < s:
			last = mid - 1
		case s < target:
			first = mid + 1
		default:
			return target
		}
	}

	return ""
}

func findUrlsSort(logUrls []string, targets []string) []string {
	sort.Strings(targets)
	n := len(targets)
	m := len(logUrls)
	found := make([]string, 0, m)
	i := 0
	j := 0
	for i < n && j < m {
		target := targets[i]
		url := logUrls[j]

		switch {
		case url < target:
			j++
		case target < url:
			i++
		default:
			found = append(found, target)
			i++
			j++
		}
	}
	return found
}

func findUrlsMap(logUrls []string, targets []string) []string {
	urlSet := make(map[string]struct{})
	for _, s := range logUrls {
		urlSet[s] = struct{}{}
	}
	found := make([]string, 0, len(targets))
	for _, target := range targets {
		if _, ok := urlSet[target]; ok {
			found = append(found, target)
		}
	}
	return found
}

func main() {
	fmt.Println("vim-go")
}
