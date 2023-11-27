package cmd

import (
	"errors"
	"fmt"
	path "github.com/nitishfy/paketo/pkg/path"
	read "github.com/nitishfy/paketo/pkg/read"
	"github.com/nitishfy/paketo/types"
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/release-sdk/obs"
)

type Options struct {
	ManifestPath string
	OBSClient    *obs.OBS
}

func Reconcile() *cobra.Command {
	opts := &Options{}
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

			prjs, err := read.ReadYAML(opts.ManifestPath)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			fmt.Println(prjs)
		},
	}
	cmd.Flags().StringP("manifest", "m", "", "path to read manifest")
	return cmd
}

func Compare(manifest *types.Projects, opts *Options) {
	// yet to be implemented
}

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
