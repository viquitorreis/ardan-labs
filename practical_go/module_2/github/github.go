package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// dado um github user login, retornar o name e o numero de repos publicos
// api.github.com/users/<username>

func main() {
	fmt.Println(UserInfo("viquitorreis"))
}

func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err making get request: %+v", err)
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("err bad status: %s\n", resp.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	// io.Copy(os.Stdout, resp.Body)

	// struct anonima para fazer o unmarshal
	var reply struct {
		Name        string
		PublicRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		fmt.Printf("error: %+v\n", err)
		return "", 0, err
	}

	return reply.Name, reply.PublicRepos, nil
}

/* JSON <-> Go (json.org)

Types - Conversao dos tipos json -> go
string <-> string
true/false <-> bool
number <-> float64, float32, int, int8, ..., int64, uint, uint8, ... (se nao especificar nada, Go converte para float64)

Tipos compostos (compound types)
array <-> []T, []any
object <-> map[string]any, struct

encoding/json API

	Como fazer o encoding do JSON dependendo do tipo dos dados recebidos

	Fonte -> Destino

	- JSON -> []byte -> Go: Usa o Unmarshal
	- Go -> []byte -> JSON: Marshal
	- JSON -> io.Reader -> Go: Decoder
	- Go -> io.Writer -> JSON: Encoder

*/
