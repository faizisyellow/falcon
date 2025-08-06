package generate

import (
	"os"
	"testing"
)

func TestMakeProject(t *testing.T) {

	t.Run("should success create project directories", func(t *testing.T) {

		tempDir, err := os.MkdirTemp(".", "")
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

		tempDir, err := os.MkdirTemp(".", "")
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

		fail := make(chan error)

		dat, err := os.ReadFile("./api.go")
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
		}

		go pr.CreateContentCmd(fail, "api", f.Name(), dat)

		err = <-fail
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
			return
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
