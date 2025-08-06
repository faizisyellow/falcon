package generate

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/faizisyellow/falcon/internal/template/api"
)

type ProjectStructure struct {
	Parent   string
	Cmd      []string
	Internal []string
	Root     []string
}

type Content interface {
	GetContent() (string, []byte)
}

func (ps *ProjectStructure) Mkdir() error {

	for _, cmd := range ps.Cmd {

		err := os.MkdirAll(fmt.Sprintf("%s/cmd/%s", ps.Parent, cmd), 0755)
		if err != nil {
			return err
		}

	}

	for _, inter := range ps.Internal {

		err := os.MkdirAll(fmt.Sprintf("%s/internal/%s", ps.Parent, inter), 0755)
		if err != nil {
			return err
		}

	}

	for _, root := range ps.Root {

		err := os.WriteFile(fmt.Sprintf("%s/%s", ps.Parent, root), []byte(""), 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

func (ps *ProjectStructure) CreateContentCmd(fail chan<- error, pckg string, filename string, data []byte) {

	f, err := os.Create(fmt.Sprintf("%s/cmd/%s/%s", ps.Parent, pckg, filename))
	if err != nil {
		fail <- err
		return
	}

	_, err = f.Write(data)
	if err != nil {
		fail <- err
		return
	}

	fail <- nil
}

var cmds = []string{
	"api",
	"migrate/migrations",
}

var internals = []string{
	"auth",
	"db",
	"keys",
	"repository",
	"service",
	"uploader",
	"utils",
}

var roots = []string{".env", ".air.toml"}

func GenerateNewProject(dst string, opts []string) error {

	project := ProjectStructure{
		Parent:   dst,
		Cmd:      cmds,
		Internal: internals,
		Root:     roots,
	}

	err := project.Mkdir()
	if err != nil {
		return err
	}

	for _, cmd := range project.Cmd {
		if cmd == cmds[0] {

			fail := make(chan error)
			pckg := cmds[0]

			go project.CreateContentCmd(fail, pckg, "main.go", api.Main)
			go project.CreateContentCmd(fail, pckg, "api.go", api.Api)
			go project.CreateContentCmd(fail, pckg, "auth.go", api.Auth)
			go project.CreateContentCmd(fail, pckg, "errors.go", api.Errors)
			go project.CreateContentCmd(fail, pckg, "json.go", api.Json)
			go project.CreateContentCmd(fail, pckg, "middlewares.go", api.Middlewares)
			go project.CreateContentCmd(fail, pckg, "users.go", api.Users)

			err := <-fail
			if err != nil {
				return err
			}
		}
	}

	return nil

}
