package main

import (
	"fmt"
	"sort"
)

// CanAttendAll reports whether one person can attend every meeting:
// after sorting by start, no meeting may begin before the previous ends.
// A meeting starting exactly when the previous ends is fine.
func CanAttendAll(intervals [][2]int) bool {
	sorted := append([][2]int{}, intervals...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][0] < sorted[j][0] })
	for i := 1; i < len(sorted); i++ {
		if sorted[i][0] < sorted[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendAll([][2]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(CanAttendAll([][2]int{{7, 10}, {2, 4}}))
}
