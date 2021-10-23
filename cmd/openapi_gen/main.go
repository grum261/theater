package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/grum261/theater/internal/rest"
	"gopkg.in/yaml.v2"
)

func main() {
	var output string

	flag.StringVar(&output, "path", "", "")
	flag.Parse()

	if output == "" {
		panic(errors.New("не передан путь"))
	}

	swagger := rest.NewOpenAPI()

	// data, err := json.Marshal(&swagger)
	// if err != nil {
	// 	panic(err)
	// }

	// if err := os.WriteFile(path.Join(output, "openapi3.json"), data, 0664); err != nil {
	// 	panic(err)
	// }

	data, err := yaml.Marshal(&swagger)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path.Join(output, "openapi3.yaml"), data, 0664); err != nil {
		panic(err)
	}

	fmt.Println("Все сгенерировано")
}
