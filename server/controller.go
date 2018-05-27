package server

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/builder/remotecontext/git"
	"github.com/docker/docker/pkg/stringid"
	"github.com/jhoonb/archivex"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"time"
)

func EmptyBuildWithTags(tags ...string) *Build {
	return &Build{
		Status: PENDING,
		Tags:   &tags,
	}
}

func BuildImage(project Project, build *Build) error {

	buildId := stringid.GenerateRandomID()
	tempDir := os.TempDir()

	build.ProjectId = project.Id
	build.Id = buildId
	build.Status = PENDING
	build.Date = time.Now().Unix()

	path, err := git.Clone(project.GitRepo)
	if err != nil {
		build.Status = FAILED
		return err
	}

	tar := new(archivex.TarFile)
	tar.Create(filepath.Join(tempDir, buildId))
	tar.AddAll(path, false)
	tar.Close()

	dockerBuildContext, err := os.Open(filepath.Join(tempDir, buildId+".tar"))
	defer dockerBuildContext.Close()

	var tags []string
	for _, tag := range *build.Tags {
		tags = append(tags, filepath.Join(project.Name, tag))
	}

	_, err = docker.ImageBuild(context.Background(), dockerBuildContext, types.ImageBuildOptions{
		Tags:           tags,
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
	})

	if err != nil {
		build.Status = FAILED
		return err
	}

	build.Status = FINISHED
	return nil
}
