package eval

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
)

func Eval(code string) error {
	fd, err := ioutil.TempFile("", "go-eval")
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

	var out bytes.Buffer
	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Stderr = &out

	if err = cmd.Run(); err != nil {
		// for debug
		fmt.Println("STDOUT:" + out.String())
		return errors.New("err: " + err.Error())
	}

	// add plugin logic

	return nil
}

func cleanupTmpFile(fd *os.File) {
	fd.Close()
	go os.Remove(fd.Name())
}