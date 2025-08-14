package generate

import (
	"os"
	"path/filepath"

	filestorage "github.com/faizisyellow/falcon/internal/fileStorage"
	handlerservicerepository "github.com/faizisyellow/falcon/internal/pattern/handlerServiceRepository"
	"github.com/faizisyellow/falcon/internal/template/air"
	"github.com/faizisyellow/falcon/internal/template/api"
	"github.com/faizisyellow/falcon/internal/template/auth"
	"github.com/faizisyellow/falcon/internal/template/db"
	"github.com/faizisyellow/falcon/internal/template/docs"
	"github.com/faizisyellow/falcon/internal/template/env"
	"github.com/faizisyellow/falcon/internal/template/gitignore"
	"github.com/faizisyellow/falcon/internal/template/jwt"
	"github.com/faizisyellow/falcon/internal/template/keys"
	loggertemplate "github.com/faizisyellow/falcon/internal/template/loggerTemplate"
	"github.com/faizisyellow/falcon/internal/template/makefile"
	"github.com/faizisyellow/falcon/internal/template/migrate/migrations"
	"github.com/faizisyellow/falcon/internal/template/repository"
	"github.com/faizisyellow/falcon/internal/template/service"
	"github.com/faizisyellow/falcon/internal/template/uploader"
	utilsTemplate "github.com/faizisyellow/falcon/internal/template/utilsTemplate"
)

type Options struct {
	Db     string
	Router string
}

type Filedata struct {
	Dir  string
	Name string
	Data []byte
}

func GenerateNewProject(opts Options) error {

	var files []Filedata

	apiByte, err := api.ApiData(nil)
	if err != nil {
		return err
	}

	authByte, err := api.AuthData(nil)
	if err != nil {
		return err
	}

	responseByte, err := api.ResponseData(nil)
	if err != nil {
		return err
	}

	mainByte, err := api.MainData(nil)
	if err != nil {
		return err
	}

	middlewaresByte, err := api.MiddlewaresData(nil)
	if err != nil {
		return err
	}

	middlewaresTestByte, err := api.MiddlewaresTestData(nil)
	if err != nil {
		return err
	}

	usersByte, err := api.UsersData(nil)
	if err != nil {
		return err
	}

	jsonByte, err := api.JsonData(nil)
	if err != nil {
		return err
	}

	routerOpt := api.RouterOpt{
		Router: opts.Router,
	}
	muxByte, err := api.MuxData(routerOpt)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "api", Name: "api.go", Data: apiByte})
	files = append(files, Filedata{Dir: "api", Name: "auth.go", Data: authByte})
	files = append(files, Filedata{Dir: "api", Name: "response.go", Data: responseByte})
	files = append(files, Filedata{Dir: "api", Name: "main.go", Data: mainByte})
	files = append(files, Filedata{Dir: "api", Name: "middlewares.go", Data: middlewaresByte})
	files = append(files, Filedata{Dir: "api", Name: "middlewares_test.go", Data: middlewaresTestByte})
	files = append(files, Filedata{Dir: "api", Name: "users.go", Data: usersByte})
	files = append(files, Filedata{Dir: "api", Name: "json.go", Data: jsonByte})
	files = append(files, Filedata{Dir: "api", Name: "mux.go", Data: muxByte})

	migrationsUserDownByte, err := migrations.CreateUserMigrateDownData(nil)
	if err != nil {
		return err
	}

	migrationsUserUpByte, err := migrations.CreateUserMigrateUpData(nil)
	if err != nil {
		return err
	}

	migrationsInvitationDownByte, err := migrations.CreateInvitationsMigrateDownData(nil)
	if err != nil {
		return err
	}

	migrationsInvitationUpByte, err := migrations.CreateInvitationsMigrateUpData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "migrations", Name: "000001_create_users_table.down.sql", Data: migrationsUserDownByte})
	files = append(files, Filedata{Dir: "migrations", Name: "000001_create_users_table.up.sql", Data: migrationsUserUpByte})

	files = append(files, Filedata{Dir: "migrations", Name: "000002_create_invitations_table.up.sql", Data: migrationsInvitationUpByte})
	files = append(files, Filedata{Dir: "migrations", Name: "000002_create_invitations_table.down.sql", Data: migrationsInvitationDownByte})

	authPckgByte, err := auth.AuthData(nil)
	if err != nil {
		return err
	}

	jwtByte, err := jwt.JwtData(nil)
	if err != nil {
		return err
	}

	docsByte, err := docs.DocsData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "docs", Name: "docs.go", Data: docsByte})

	files = append(files, Filedata{Dir: "auth", Name: "jwt.go", Data: jwtByte})
	files = append(files, Filedata{Dir: "auth", Name: "auth.go", Data: authPckgByte})

	dbOpt := db.DBOption{
		DB: opts.Db,
	}
	dbByte, err := db.DBData(dbOpt)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "db", Name: "db.go", Data: dbByte})
	envByte, err := env.EnvData(nil)
	if err != nil {
		return err
	}

	airByte, err := air.AirData(nil)
	if err != nil {
		return err
	}

	gitignoreByte, err := gitignore.GitignoreData(nil)
	if err != nil {
		return err
	}

	makefileByte, err := makefile.MakefileData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "", Name: "makefile", Data: makefileByte})
	files = append(files, Filedata{Dir: "", Name: ".air.toml", Data: airByte})
	files = append(files, Filedata{Dir: "", Name: ".gitignore", Data: gitignoreByte})
	files = append(files, Filedata{Dir: "", Name: ".env.sample", Data: envByte})

	keysByte, err := keys.KeysData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "keys", Name: "keys.go", Data: keysByte})

	usersRepoByte, err := repository.UsersData(nil)
	if err != nil {
		return err
	}

	invitationByte, err := repository.InvitationData(nil)
	if err != nil {
		return err
	}

	repositoryByte, err := repository.RepositoryData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "repository", Name: "users.go", Data: usersRepoByte})
	files = append(files, Filedata{Dir: "repository", Name: "invitation.go", Data: invitationByte})
	files = append(files, Filedata{Dir: "repository", Name: "repository.go", Data: repositoryByte})

	loggerByte, err := loggertemplate.LoggerData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "logger", Name: "logger.go", Data: loggerByte})

	serviceByte, err := service.ServiceData(nil)
	if err != nil {
		return err
	}

	usersServiceByte, err := service.UsersData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "service", Name: "service.go", Data: serviceByte})
	files = append(files, Filedata{Dir: "service", Name: "users.go", Data: usersServiceByte})

	contentContextRepoByte, err := utilsTemplate.ContentContextData(nil)
	if err != nil {
		return err
	}

	passwordByte, err := utilsTemplate.PasswordData(nil)
	if err != nil {
		return err
	}

	tokenByte, err := utilsTemplate.TokenData(nil)
	if err != nil {
		return err
	}

	pointerByte, err := utilsTemplate.PointerData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "utils", Name: "contentContext.go", Data: contentContextRepoByte})
	files = append(files, Filedata{Dir: "utils", Name: "password.go", Data: passwordByte})
	files = append(files, Filedata{Dir: "utils", Name: "token.go", Data: tokenByte})
	files = append(files, Filedata{Dir: "utils", Name: "pointer.go", Data: pointerByte})

	uploaderByte, err := uploader.UplaoderData(nil)
	if err != nil {
		return err
	}

	files = append(files, Filedata{Dir: "uploader", Name: "uploader.go", Data: uploaderByte})

	projects := handlerservicerepository.New()

	for _, file := range files {

		dir := &file.Dir

		if file.Dir == "" {
			dir = nil
		}

		projects.Add(filestorage.Data{
			Parent: dir,
			File: filestorage.FileStorage{
				Name:     file.Name,
				Content:  file.Data,
				Children: nil,
				IsFile:   true,
			},
		})
	}

	err = CreateProject(projects)
	if err != nil {
		return err
	}

	return nil
}

func CreateProject(project filestorage.FileStorage) error {
	for _, fs := range project.Children {
		if err := createNode(project.Name, fs); err != nil {
			return err
		}
	}
	return nil
}

func createNode(basePath string, fs *filestorage.FileStorage) error {
	path := filepath.Join(basePath, fs.Name)

	if fs.IsFile {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		if len(fs.Content) > 0 {
			if _, err := f.Write(fs.Content); err != nil {
				f.Close()
				return err
			}
		}
		return f.Close()
	}

	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	for _, child := range fs.Children {
		if err := createNode(path, child); err != nil {
			return err
		}
	}
	return nil
}
