// Package mem easily gets current memory usage
package mem

import (
	"runtime"
	"strconv"
)

// func Allocated returns a string of current memory usage such as "8KB" or "16MB"
func Allocated() string {
	r := new(runtime.MemStats)
	runtime.ReadMemStats(r)
	b, i := r.HeapAlloc, 0

	for {
		if b < 1024 {
			break
		}

		b /= 1024
		i++
	}

	s := strconv.FormatUint(b, 10)

	switch i {
	case 0:
		// Byte
		return s + "B"
	case 1:
		// Kilobyte
		return s + "KiB"
	case 2:
		// Megabyte
		return s + "MiB"
	case 3:
		// Gigabyte
		return s + "GiB"
	case 4:
		// Terabyte
		return s + "TiB"
	case 5:
		// Petabyte
		return s + "PiB"
	case 6:
		// Exabyte
		return s + "EiB"
	case 7:
		// Zettabyte
		return s + "ZiB"
	case 8:
		// Yottabyte
		return s + "YiB"
	case 9:
		// Ronnabyte (I was a bit bummed I didn't get to use Brontobytes)
		return s + "RiB"
	case 10:
		// Quettabyte
		return s + "QiB"
	default:
		// Anything larger than the unreasonabbly large Quettabyte gets represented by it's power
		return s + " ** " + strconv.Itoa(i) + " Bytes"
	}
}
