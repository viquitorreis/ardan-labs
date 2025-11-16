package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
}

// returns SHA1 signature of uncompressed file
// cat http.log.gz | gunzip | sha1sum
// descomprimir apenas se o nome do arquivo termina com ".gz "
func SHA1Sig(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// | gunzip
	r, err := gzip.NewReader(f)
	if err != nil {
		return "", fmt.Errorf("%q - gzip: %w", fileName, err)
	}
	defer r.Close()

	if strings.HasSuffix(f.Name(), ".gz") {
		outFile, err := os.Create("out.log")
		if err != nil {
			return "", fmt.Errorf("error creating file: %v", err)
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, r); err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)

		}
	}

	// | sha1sum
	w := sha1.New() // writer
	// conseguimos mover dados entre um io.Reader e um io.Writer usando io.Copy
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	// %x -> formata em hexadecimal
	// se passar bytes para ele vai formatar em 2 digitos hexadecimais por byte
	return fmt.Sprintf("%x", sig), nil
}

/*
type Reader interface {
	Read(p []byte) (n  int, err error)
}
*/
