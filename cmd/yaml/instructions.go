package yaml

import (
	"fmt"
	"io"
	"os"

	"github.com/IbraheemHaseeb7/apee-i/cmd"
	y "gopkg.in/yaml.v3"
)

type YAMLReader struct {}

func (r *YAMLReader) ReadInstructions(filepath string) (*cmd.Structure, error) {
	
	file, err := os.Open(filepath)
	if err != nil { return &cmd.Structure{}, fmt.Errorf("Could not open file") }

	defer file.Close()

	fileRawContents, err := io.ReadAll(file)
	if err != nil { return &cmd.Structure{}, fmt.Errorf("Could not read file contents") }
	
	fileContents := new(cmd.Structure)
	err = y.Unmarshal(fileRawContents, &fileContents)
	if err != nil { return &cmd.Structure{}, fmt.Errorf("Could not map elements in yaml") }

	return fileContents, nil
}
