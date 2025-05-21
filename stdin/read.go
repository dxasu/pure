package stdin

import (
	"errors"
	"io"
	"os"
)

var ErrNoStdin = errors.New("no stdin")

func GetStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	if hasStdin {
		out, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return out[:len(out)-1], nil
	}
	return nil, ErrNoStdin
}
