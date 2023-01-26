package secretvar

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

type TestStruct struct {
	Var1 string
	Var2 SecretString
	Var3 string
}

func TestSecretString(t *testing.T) {
	ss := SecretString("secret-secret-secret")
	ts := TestStruct{
		Var1: "string-string-string",
		Var2: "secret-secret-secret",
		Var3: "string-string-string",
	}
	// string
	fmt.Println(ss)
	fmt.Println(ss.Get())
	fmt.Println(ts)
	// json
	vrJson, _ := json.MarshalIndent(ts, "", "  ")
	fmt.Println(string(vrJson))
	// yaml
	vrYaml, _ := yaml.Marshal(ts)
	fmt.Println(string(vrYaml))
}
