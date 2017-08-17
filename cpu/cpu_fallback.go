// +build !darwin,!linux,!freebsd,!openbsd,!solaris,!windows

package cpu

import (
	"time"

	"github.com/WigWagCo/gopsutil/internal/common"
)

func Times(percpu bool) ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func Info() ([]InfoStat, error) {
	return []InfoStat{}, common.ErrNotImplementedError
}
