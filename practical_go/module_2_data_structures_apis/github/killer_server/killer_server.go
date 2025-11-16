package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	err := KillerServer("server.pid")
	if err != nil {
		fmt.Println("error", err)

		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file not found")
		}

		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func KillerServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	/*
		- defer aconbtece quando a func termina, nao importa o que (mesmo com panic)
		- funciona no nivel da FUNCAO - se usar dentro de um for, provavelmente vai dar bug
		- defer sao executados em ordem REVERSA (stack, LIFO)
		Idiom: tenta pegar um recursa, checa por erro, e fazer um defer para ´soltar' o recurso
	*/
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("close", "file", pidFile, "error", err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", pidFile, err)
	}

	slog.Info("killing", "pid", pid)
	if err := os.Remove(pidFile); err != nil {
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}
