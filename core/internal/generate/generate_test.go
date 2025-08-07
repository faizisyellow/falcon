package generate

import (
	"os"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestMakeProject(t *testing.T) {

	t.Run("should success create project directories", func(t *testing.T) {

		tempDir, err := os.MkdirTemp("", "")
		if err != nil {
			t.Error(err)
			return
		}

		defer os.RemoveAll(tempDir)

		pr := setupData(tempDir)

		err = pr.Mkdir()
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}
	})

	t.Run("should success create content", func(t *testing.T) {

		tempDir, err := os.MkdirTemp("", "")
		if err != nil {
			t.Error(err)
			return
		}

		defer os.RemoveAll(tempDir)

		pr := setupData(tempDir)
		err = pr.Mkdir()
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}

		f, err := os.Create("./api.go")
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}

		defer f.Close()
		defer os.Remove(f.Name())

		f.Write([]byte("lizzy mcalpine"))

		var errg errgroup.Group
		dat, err := os.ReadFile("./api.go")
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}

		var wg sync.WaitGroup

		wg.Add(1)
		errg.Go(func() error {
			return pr.CreateContentCmd(&wg, "api", f.Name(), dat)
		})

		err = errg.Wait()
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}
	})

	t.Run("should success generate new project", func(t *testing.T) {

		tempDir, err := os.MkdirTemp(".", "")
		if err != nil {
			t.Error(err)
			return
		}

		defer os.RemoveAll(tempDir)

		err = GenerateNewProject(tempDir, []string{"chi", "mysql"})
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
		}
	})
}

func setupData(parent string) ProjectStructure {
	return ProjectStructure{
		Parent: parent,
		Cmd: []string{
			"api",
			"migrate/migrations",
		},
		Internal: []string{
			"auth",
			"db",
			"keys",
			"repository",
			"service",
			"uploader",
			"utils",
		},

		Root: []string{
			".env", ".air.toml",
		},
	}
}
