package rotarylog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var _ io.WriteCloser = (*RotaryLog)(nil)

// RotaryLog struct
type RotaryLog struct {
	sync.Mutex
	Age   time.Duration
	Num   int
	Size  int
	file  *os.File
	sTime time.Time
	size  int64
}

// New return instance of RotaryLog
// defaults
// age  86400 rotate every 24h0m0s
// num  7     files
// size 0     no limit size
func NewRotaryLog(logfile string, age, num, size int) (*RotaryLog, error) {
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	Age := time.Duration(0)
	if age > 0 {
		Age = time.Duration(age) * time.Second
	}
	num--
	if num < 0 {
		num = 7
	}
	Size := 0
	if size > 0 {
		Size = size * 1048576
	}
	rl := &RotaryLog{
		Age:   Age,
		Num:   num,
		Size:  Size,
		file:  f,
		sTime: time.Now(),
	}
	// rotate if needed
	if i, err := rl.file.Stat(); err == nil {
		if rl.Age > 0 && time.Since(i.ModTime()) >= rl.Age {
			if err := rl.rotate(); err != nil {
				return nil, err
			}
		} else if rl.Size > 0 && i.Size() > int64(rl.Size) {
			if err := rl.rotate(); err != nil {
				return nil, err
			}
		}
	}
	return rl, nil
}

// Write implements io.Writer
func (rl *RotaryLog) Write(p []byte) (n int, err error) {
	rl.Lock()
	defer rl.Unlock()

	writeLen := int64(len(p))

	// rotate based on Age and size
	if rl.Age > 0 && time.Since(rl.sTime) >= rl.Age {
		rl.sTime = time.Now()
		if err := rl.rotate(); err != nil {
			return 0, err
		}
	} else if rl.Size > 0 && rl.size+writeLen > int64(rl.Size) {
		if err := rl.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = rl.file.Write(p)
	rl.size += int64(n)
	return n, err
}

// Close implements io.Closer, and closes the current logfile
func (rl *RotaryLog) Close() error {
	rl.Lock()
	defer rl.Unlock()
	return rl.close()
}

// close closes the file if it is open
func (rl *RotaryLog) close() error {
	if rl.file == nil {
		return nil
	}
	err := rl.file.Close()
	rl.file = nil
	return err
}

// Rotate helper function for rotate
func (rl *RotaryLog) Rotate() error {
	rl.Lock()
	defer rl.Unlock()
	return rl.rotate()
}

// rotate close existing log file and create a new one
func (rl *RotaryLog) rotate() error {
	format := "%s-%d%s"
	path := rl.file.Name()
	ext := filepath.Ext(path)
	rl.close()
	// rotate logs
	for i := rl.Num; i >= 0; i-- {
		logfile := fmt.Sprintf(format, strings.TrimSuffix(path, ext), i, ext)
		if _, err := os.Stat(logfile); err == nil {
			// delete old file
			if i == rl.Num {
				os.Remove(logfile)
			} else if err := os.Rename(logfile, fmt.Sprintf(format, strings.TrimSuffix(path, ext), i+1, ext)); err != nil {
				return err
			}
		}
	}
	// create logfile 0
	if err := os.Rename(path, fmt.Sprintf(format, strings.TrimSuffix(path, ext), 0, ext)); err != nil {
		return err
	}
	// create new log file
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	rl.file = f
	rl.size = 0
	return nil
}
