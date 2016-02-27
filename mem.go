// Package mem easily gets current memory usage
package mem

import (
	"fmt"
	"runtime"
)

// Get returns a string of current memory usage such as "8KB" or "16MB"
func Get() string {
	s := new(runtime.MemStats)
	runtime.ReadMemStats(s)
	m := s.TotalAlloc

	var i int
	for {
		if m > 1024 {
			m = m / 1024
		} else {
			break
		}
		i++
	}

	b := make(map[int]string, 6)
	b[0] = "B"
	b[1] = "KB"
	b[2] = "MB"
	b[3] = "GB"
	b[4] = "TB"
	b[5] = "PB"

	if i > 5 {
		panic("Mem: We don't deal in anything larger than Petabytes")
	}

	return fmt.Sprintf("%d"+b[i], m)
}
