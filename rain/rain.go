package rain

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/gookit/color"
)

func ExitIf(err any, args ...any) {
	switch e := err.(type) {
	case nil:
		return
	case error:
		if e != nil {
			fmt.Println(e.Error())
			os.Exit(1)
		}
	case string:
		if e != "" {
			fmt.Printf(e, args...)
			fmt.Println()
			os.Exit(1)
		}
	default:
		fmt.Printf("%+v\n", e)
		os.Exit(1)
	}
}

func WaitCtrlC() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func NeedHelp() bool {
	return len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help")
}

// IsInteractive Return true if os.Stdin appears to be interactive
func IsInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&os.ModeCharDevice != 0
}

func OpenBrower(uri string) error {
	// 不同平台启动指令不同
	var commands = map[string]string{
		"windows": "explorer",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("invalid platform: %s", runtime.GOOS)
	}
	cmd := exec.Command(run, uri)
	return cmd.Run()
}

// Shell 执行 shell 脚本
func Shell(ctx context.Context, input string) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", input)
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = buf

	if err := cmd.Run(); err != nil {
		errf := strings.TrimSpace(buf.String())
		return "", fmt.Errorf("exec: %w (%q)", err, errf)
	}
	return buf.String(), nil
}

func DebugCmd(params string) {
	if len(params) == 0 {
		return
	}
	os.Args = append([]string{os.Args[0]}, strings.Fields(params)...)
}

func DebugArgs(strs ...string) {
	if len(strs) == 0 {
		return
	}
	os.Args = append([]string{os.Args[0]}, strs...)
}

func DebugEnvs(envs map[string]string) {
	for k, v := range envs {
		if v == "" {
			os.Unsetenv(k)
		} else {
			os.Setenv(k, v)
		}
	}
}

type Clog int

func (c Clog) Str(s any) string {
	return color.HEX(fmt.Sprintf("%x", c)).Sprint(s)
}
