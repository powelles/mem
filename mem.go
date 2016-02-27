// Package mem easily gets current memory usage
package mem

import (
	"fmt"
	"runtime"
)

func get() *MemStats{
	s := new(runtime.MemStats)
	runtime.ReadMemStats(s)
	return s
}

// func Allocated returns a string of current memory usage such as "8KB" or "16MB"
func Allocated() string {
	m := get().Alloc
	
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

	if i > 3 {
		panic("github.com/powelles/mem: We don't deal in anything larger than Gigabytes")
	}

	return fmt.Sprintf("%d"+b[i], m)
}
