package main

import (
	"fmt"
	"os"
	"os/exec"
)

func underlyingError(err error) error {
	switch errtype := err.(type) {
	case *os.PathError:
		return errtype.Err
	case *os.LinkError:
		return errtype.Err
	case *os.SyscallError:
		return errtype.Err
	case *exec.Error:
		return errtype.Err
	default:
		return err
	}
}

func knownError(err error) {
	switch err {
	case os.ErrClosed:
		fmt.Println("errClosed")
	case os.ErrInvalid:
		fmt.Println("errInvalid")
	case os.ErrPermission:
		fmt.Println("errPermission")
	}
}
