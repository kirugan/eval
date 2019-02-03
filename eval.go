package eval

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"plugin"
	"strings"
)

const entryFunctionName = "Main"

func Eval(code string) error {
	fd, err := ioutil.TempFile("", "go-eval*.go")
	if err != nil {
		return err
	}
	defer cleanupTmpFile(fd)

	nW, err := fd.Write([]byte(code))
	if err != nil {
		return err
	}
	if nW != len(code) {
		return fmt.Errorf("partial write %d of %d bytes written", nW, len(code))
	}
	fd.Sync()

	basename := strings.TrimSuffix(fd.Name(), path.Ext(fd.Name()))
	pluginFile := basename + ".so"

	var out bytes.Buffer
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginFile, fd.Name())
	cmd.Stderr = &out

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("%v (stderr='%s')", err, out.String())
	}

	pl, err := plugin.Open(pluginFile)
	if err != nil {
		return err
	}

	symb, err := pl.Lookup(entryFunctionName)
	if err != nil {
		return err
	}

	if main, ok := symb.(func()); ok {
		main()
		return nil
	}

	// the reason why we here is a signature check failure
	return fmt.Errorf("entry function '%v' has wrong signature", entryFunctionName)
}

func cleanupTmpFile(fd *os.File) {
	fd.Close()
	go os.Remove(fd.Name())
}