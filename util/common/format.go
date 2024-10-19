package common

import (
	"fmt"
)

func FormatTraffic(trafficBytes int64) string {
  units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
  unitIndex := 0
  size := float64(trafficBytes)

  for size >= 1024 && unitIndex < len(units)-1 {
    if unitIndex == 4 {
       size /= 1000
    } else {
       size /= 1024
    }
    unitIndex++
  }
  return fmt.Sprintf("%.2f%s", size, units[unitIndex])
}
