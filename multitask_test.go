package multitask

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("a")
			time.Sleep(time.Second)
			return errors.New("sdf")
		},
		func() error {
			fmt.Println("b")
			time.Sleep(2 * time.Second)
			return nil
		},
		func() error {
			fmt.Println("c")
			time.Sleep(2 * time.Second)
			return errors.New("sdf")
		},
		func() error {
			fmt.Println("d")
			time.Sleep(2 * time.Second)
			return nil
		},
		func() error {
			fmt.Println("e")
			time.Sleep(4 * time.Second)
			return nil
		},
		func() error {
			fmt.Println("f")
			time.Sleep(4 * time.Second)
			return nil
		},
	}
	Run(tasks, 3, 2)
}
