package generate

import (
	_ "embed"
	"fmt"
	"os"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/faizisyellow/falcon/internal/template/air"
	"github.com/faizisyellow/falcon/internal/template/api"
	"github.com/faizisyellow/falcon/internal/template/auth"
	"github.com/faizisyellow/falcon/internal/template/db"
	"github.com/faizisyellow/falcon/internal/template/env"
	"github.com/faizisyellow/falcon/internal/template/keys"
	"github.com/faizisyellow/falcon/internal/template/repository"
	"github.com/faizisyellow/falcon/internal/template/service"
	"github.com/faizisyellow/falcon/internal/template/uploader"
	"github.com/faizisyellow/falcon/internal/template/utils"
	"golang.org/x/sync/errgroup"
)

type ProjectStructure struct {
	Parent   string
	Bin      string
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

	err := os.MkdirAll(fmt.Sprintf("%s/%s", ps.Parent, ps.Bin), 0755)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProjectStructure) CreateContentCmd(wg *sync.WaitGroup, pckg string, filename string, data []byte) error {

	defer wg.Done()

	var sub = "cmd"

	return ps.WriteData(pckg, data, filename, sub)

}

func (ps *ProjectStructure) CreateContentInternal(wg *sync.WaitGroup, pckg string, filename string, data []byte) error {

	defer wg.Done()
	var sub = "internal"

	return ps.WriteData(pckg, data, filename, sub)

}

func (ps *ProjectStructure) CreateContentRoot(wg *sync.WaitGroup, filename string, data []byte) error {

	defer wg.Done()

	f, err := os.Create(fmt.Sprintf("%s/%s", ps.Parent, filename))
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	var info = lipgloss.NewStyle().
		Faint(true).
		Foreground(lipgloss.Color("12")).
		Render(fmt.Sprintf("creating %s/%s", ps.Parent, filename))

	fmt.Println(info)

	return nil
}

func (ps *ProjectStructure) WriteData(pckg string, data []byte, filename string, sub string) error {

	f, err := os.Create(fmt.Sprintf("%s/%s/%s/%s", ps.Parent, sub, pckg, filename))
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	var info = lipgloss.NewStyle().
		Faint(true).
		Foreground(lipgloss.Color("12")).
		Render(fmt.Sprintf("creating %s/%s/%s/%s", ps.Parent, sub, pckg, filename))

	fmt.Println(info)

	return nil
}

func GenerateNewProject(dst string, opts []string) error {

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

	var wg sync.WaitGroup
	var errg errgroup.Group

	project := ProjectStructure{
		Parent:   dst,
		Cmd:      cmds,
		Internal: internals,
		Root:     roots,
		Bin:      "bin",
	}

	err := project.Mkdir()
	if err != nil {
		return err
	}

	for _, cmd := range project.Cmd {
		if cmd == cmds[0] {

			pckg := cmds[0]

			contentApi := make(map[string][]byte, 0)
			contentApi["main.go"] = api.Main
			contentApi["api.go"] = api.Api
			contentApi["auth.go"] = api.Auth
			contentApi["errors.go"] = api.Errors
			contentApi["json.go"] = api.Json
			contentApi["middlewares.go"] = api.Middlewares
			contentApi["users.go"] = api.Users

			for k, v := range contentApi {
				wg.Add(1)
				errg.Go(func() error {

					return project.CreateContentCmd(&wg, pckg, k, v)
				})
			}

		}
	}

	for _, internal := range project.Internal {

		switch internal {
		case internals[0]:

			wg.Add(1)
			errg.Go(func() error {
				return project.CreateContentInternal(&wg, internals[0], "auth.go", auth.Auth)
			})

		case internals[1]:

			if len(opts) <= 0 {
				return fmt.Errorf("options empty")
			}

			if opts[1] == "mysql" {

				wg.Add(1)
				errg.Go(func() error {
					return project.CreateContentInternal(&wg, internals[1], "db.go", db.Mysql)
				})

			}

		case internals[2]:

			wg.Add(1)
			errg.Go(func() error {
				return project.CreateContentInternal(&wg, internals[2], "keys.go", keys.Keys)
			})

		case internals[3]:

			contentRepository := make(map[string][]byte)
			contentRepository["invitation.go"] = repository.Invitation
			contentRepository["users.go"] = repository.Users
			contentRepository["repository.go"] = repository.Repository

			for k, v := range contentRepository {
				wg.Add(1)
				errg.Go(func() error {
					return project.CreateContentInternal(&wg, internals[3], k, v)
				})
			}

		case internals[4]:

			contentService := make(map[string][]byte)
			contentService["service.go"] = service.Service
			contentService["users.go"] = service.Users

			for k, v := range contentService {
				wg.Add(1)
				errg.Go(func() error {
					return project.CreateContentInternal(&wg, internals[4], k, v)
				})
			}

		case internals[5]:
			wg.Add(1)
			errg.Go(func() error {
				return project.CreateContentInternal(&wg, internals[5], "uploader.go", uploader.Uploader)
			})

		case internals[6]:

			contentUtils := make(map[string][]byte)
			contentUtils["password.go"] = utils.Password
			contentUtils["token.go"] = utils.Token
			contentUtils["contentContext.go"] = utils.ContentContext

			for k, v := range contentUtils {
				wg.Add(1)
				errg.Go(func() error {

					return project.CreateContentInternal(&wg, internals[6], k, v)
				})
			}

		}
	}

	for _, root := range project.Root {

		switch root {

		case roots[0]:

			wg.Add(1)
			errg.Go(func() error {
				return project.CreateContentRoot(&wg, roots[0], env.Env)
			})

		case roots[1]:

			wg.Add(1)
			errg.Go(func() error {
				return project.CreateContentRoot(&wg, roots[1], air.Air)
			})

		}
	}

	err = errg.Wait()
	if err != nil {
		return err
	}

	return nil

}
