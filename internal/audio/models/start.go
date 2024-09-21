package models

import (
	"context"
	"fmt"
	"io"
	"os"
	"syscall"
)

func DownloadAndExtract(url string) error {

	// Create context which quits on SIGINT or SIGQUIT
	ctx := ContextForSignal(os.Interrupt, syscall.SIGQUIT)

	// Progress filehandle
	progress := os.Stdout
	if *flagQuiet {
		progress, err := os.Open(os.DevNull)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(-1)
		}
		defer progress.Close()
	}

	// Get output path
	out, err := GetOut()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(-1)
	}

	// Download file from url - exit on error or interrupt
	if path, err := Download(ctx, progress, url, out); err == nil || err == io.EOF {
		fmt.Fprintln(progress, "downloading successful")
		return nil
	} else if err == context.Canceled {
		os.Remove(path)
		fmt.Fprintln(progress, "\nInterrupted")
		return err
	} else if err == context.DeadlineExceeded {
		os.Remove(path)
		fmt.Fprintln(progress, "Timeout downloading model")
		return err
	} else {
		os.Remove(path)
		fmt.Fprintln(os.Stderr, "Error:", err)
		return err
	}

}
