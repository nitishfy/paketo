package read

import (
	"fmt"
	"github.com/goreleaser/goreleaser/pkg/context"
	"github.com/nitishfy/paketo/types"
	"os"
	"sigs.k8s.io/release-sdk/obs"
	"sigs.k8s.io/yaml"
)

func ReadYAML(filepath string) (*types.Project, error) {
	o := obs.OBS{}
	ctx := context.Context{}
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the content of the file: %v", err)
	}
	fmt.Println("Printing the file content:")
	os.Stdout.Write(content)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall read: %v", err)
	}
	var prjs types.Project
	err = yaml.Unmarshal(content, &prjs)
	if err != nil {
		return nil, fmt.Errorf("error is: %v", err)
	}

	remotePrj, _ := o.GetProjectMetaFile(ctx, prjs.Name)
	if remotePrj.Name != prjs.Name {
		fmt.Println("Project not found")
		return nil, nil
	}

	return &prjs, nil
}
