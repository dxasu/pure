package storage

import (
	"os"
	"path/filepath"

	"github.com/dxasu/pure/rain"
)

type DirType int

const (
	DirTemp DirType = iota
	DirCurrent
	DirTarget // full path to target directory
)

type DataPath struct {
	dir string
}

// NewDataPath creates a new DataPath with the specified directory type and optional subdirectory.
// If the directory does not exist, it will be created.
// If subDir is empty, it uses the current working directory for DirCurrent or the system temp directory for DirTemp.
func NewDataPath(t DirType, subDir string) *DataPath {
	dp := &DataPath{}
	if t < DirTemp || t > DirTarget {
		rain.ExitIf(os.ErrInvalid)
	}

	if dp.dir != "" {
		rain.ExitIf(os.ErrInvalid)
	}
	switch t {
	case DirCurrent, DirTemp:
		var path string
		if t == DirCurrent {
			path, _ = os.Getwd()
		} else {
			path = os.TempDir()
		}
		if subDir == "" {
			dp.dir = path
		} else {
			dp.dir = filepath.Join(path, subDir)
		}
	case DirTarget:
		if len(subDir) == 0 {
			rain.ExitIf("subDir cannot be empty for DirTarget")
		}
		dp.dir = subDir
	default:
		rain.ExitIf("invalid DirType: %d", t)
	}

	if _, err := os.Stat(dp.dir); os.IsNotExist(err) {
		err := os.MkdirAll(dp.dir, 0755)
		rain.ExitIf(err)
	}
	return dp
}

func (dp *DataPath) GetFilePath(file string) string {
	return filepath.Join(dp.dir, file)
}

func (dp *DataPath) GetDir() string {
	return dp.dir
}
