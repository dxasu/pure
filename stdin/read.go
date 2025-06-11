package stdin

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/dxasu/pure/rain"
)

var ErrStdin = errors.New("invalid stdin")

// GetStdin reads from standard input and returns the data as a byte slice.
// If there is no data in stdin, it returns an error.
// If the input ends with a newline, it is removed from the output.
// If stdin is not available, it returns an error.
// It is useful for reading piped input or when the program is run without arguments.
// If the input is empty, it returns ErrStdin.
// It is recommended to use this function when you expect input from the user or another program.
// It reads all data from stdin until EOF and returns it as a byte slice.
// If there is no data, it returns an error.
// This function is useful for command-line tools that need to process input from the user or other programs.
// It can be used in conjunction with other functions that process or manipulate the input data.
// Example usage:
//
//	data, err := GetStdin()
//	if err != nil {
//	    fmt.Println("Error reading stdin:", err)
//	    return
//	}
//	fmt.Println("Received input:", string(data))
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
		if len(out) > 0 && out[len(out)-1] == '\n' {
			out = out[:len(out)-1]
		}
		if len(out) == 0 {
			return nil, ErrStdin
		}
		return out, nil
	}
	return nil, ErrStdin
}

// GetInput retrieves input data from command line arguments or standard input.
// If command line arguments are provided, it joins them into a single string.
// If no command line arguments are provided, it attempts to read from standard input.
// If standard input is not available, it tries to read from the clipboard.
// It returns the input data as a string.
// If no input is available from any source, it returns an empty string.
// This function is useful for applications that need to process input from various sources,
// such as command line arguments, standard input, or clipboard.
func GetInput() string {
	var data string
	if len(os.Args) > 1 {
		data = strings.Join(os.Args[1:], " ")
	} else if d, err := GetStdin(); err != nil {
		content, err := clipboard.ReadAll()
		rain.ExitIf(err)
		data = content
	} else {
		data = string(d)
	}
	return data
}
