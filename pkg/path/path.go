package path

import (
	"os"
	"path/filepath"
)

// GetWorkdirFromManifestPath sets the path
func GetWorkdirFromManifestPath(manifestPath string) string {
	dir := filepath.Dir(manifestPath)
	if filepath.Base(dir) == ".okteto" {
		dir = filepath.Dir(dir)
	}
	return dir
}

func GetRelativePathFromCWD(cwd, path string) (string, error) {
	if path == "" || !filepath.IsAbs(path) {
		return path, nil
	}

	relativePath, err := filepath.Rel(cwd, path)
	if err != nil {
		return "", err
	}
	return relativePath, nil
}

// GetManifestPathFromWorkdir returns the path from a workdir
func GetManifestPathFromWorkdir(manifestPath, workdir string) string {
	mPath, err := filepath.Rel(workdir, manifestPath)
	if err != nil {
		return ""
	}
	return mPath
}

func UpdateCWDtoManifestPath(manifestPath string) (string, error) {
	workdir := GetWorkdirFromManifestPath(manifestPath)
	if err := os.Chdir(workdir); err != nil {
		return "", err
	}
	updatedManifestPath := GetManifestPathFromWorkdir(manifestPath, workdir)
	return updatedManifestPath, nil
}
