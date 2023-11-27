package read

import (
	"fmt"
	"github.com/nitishfy/paketo/types"
	"os"
	"sigs.k8s.io/yaml"
)

func ReadYAML(filepath string) (*types.Projects, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the content of the file: %v", err)
	}

	var prjs types.Projects
	err = yaml.Unmarshal(content, &prjs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall read: %v", err)
	}

	return &prjs, nil
}
