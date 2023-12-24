package cmd

import (
	"context"
	"errors"
	"fmt"
	path "github.com/nitishfy/paketo/pkg/path"
	"github.com/nitishfy/paketo/types"
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/release-sdk/obs"
	"sigs.k8s.io/yaml"
)

type Options struct {
	ManifestPath string
	OBSClient    *obs.OBS
}

func Reconcile() *cobra.Command {
	ctx := context.Background()
	opts := &Options{}
	o := &obs.OBS{}
	cmd := &cobra.Command{
		Use:   "reconcile",
		Short: "reconcile command for Paketo",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			manifestPath, _ := cmd.Flags().GetString("manifest")
			opts.ManifestPath = manifestPath

			err := GetManifestPath(opts)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			prjs, err := LoadManifest(manifestPath)
			if err != nil {
				fmt.Errorf("unable to load the manifest file: %v", err)
				return
			}

			for _, prj := range prjs.Projects {
				remotePrj, _ := o.GetProjectMetaFile(ctx, prj.Name)
				if remotePrj != nil && remotePrj.Name != prj.Name {
					fmt.Printf("Project %s doesn't exit!", prj.Name)
				}
			}

			fmt.Println("Everything working well so far!")
		},
	}
	cmd.Flags().StringP("manifest", "m", "", "path to read manifest")
	return cmd
}

func LoadManifest(filepath string) (*types.Projects, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read the file content: %v", err)
	}

	var prjs types.Projects
	err = yaml.Unmarshal(content, &prjs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml: %v", err)
	}

	return &prjs, nil
}

//func Compare(local, remote *types.Project) (bool, error) {
//
//}

func GetManifestPath(opts *Options) error {
	if opts.ManifestPath != "" {
		// if path is absolute, it is transformed from root path to a rel path
		initialCWD, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get the current working directory: %w", err)
		}

		manifestPathFlag, err := path.GetRelativePathFromCWD(initialCWD, opts.ManifestPath)
		if err != nil {
			return err
		}
		opts.ManifestPath = manifestPathFlag

		// when the manifest path is set by the cmd flag, we are moving cwd so the cmd is executed from that dir
		uptManifestPath, err := path.UpdateCWDtoManifestPath(opts.ManifestPath)
		if err != nil {
			return err
		}
		opts.ManifestPath = uptManifestPath

		if _, err := os.Stat(opts.ManifestPath); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s file doesn't exist", opts.ManifestPath)
		}
	}

	return nil
}
