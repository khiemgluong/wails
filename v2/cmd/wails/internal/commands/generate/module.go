package generate

import (
	"github.com/leaanthony/clir"
	"github.com/wailsapp/wails/v2/internal/shell"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// AddModuleCommand adds the `module` subcommand for the `generate` command
func AddModuleCommand(app *clir.Cli, parent *clir.Command, w io.Writer) error {

	command := parent.NewSubCommand("module", "Generate wailsjs modules")

	command.Action(func() error {

		filename := "wailsbindings"
		if runtime.GOOS == "windows" {
			filename += ".exe"
		}
		// go build -tags bindings -o bindings.exe
		tempDir := os.TempDir()
		filename = filepath.Join(tempDir, filename)

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		_, _, err = shell.RunCommand(cwd, "go", "build", "-tags", "bindings", "-o", filename)
		if err != nil {
			return err
		}

		_, _, err = shell.RunCommand(cwd, filename)
		if err != nil {
			return err
		}

		err = os.Remove(filename)
		if err != nil {
			return err
		}

		return nil
	})
	return nil
}
