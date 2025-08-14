package handlerservicerepository

import (
	filestorage "github.com/faizisyellow/falcon/internal/fileStorage"
	"github.com/faizisyellow/falcon/internal/utils"
)

func New() filestorage.FileStorage {

	var root filestorage.FileStorage
	var directories []*filestorage.Data

	cmd := filestorage.NewDir("cmd", nil)
	api := filestorage.NewDir("api", utils.StringToPoint("cmd"))
	migrate := filestorage.NewDir("migrate", utils.StringToPoint("cmd"))
	migrations := filestorage.NewDir("migrations", utils.StringToPoint("migrate"))

	internal := filestorage.NewDir("internal", nil)
	db := filestorage.NewDir("db", utils.StringToPoint("internal"))
	auth := filestorage.NewDir("auth", utils.StringToPoint("internal"))
	service := filestorage.NewDir("service", utils.StringToPoint("internal"))
	repository := filestorage.NewDir("repository", utils.StringToPoint("internal"))
	keys := filestorage.NewDir("keys", utils.StringToPoint("internal"))
	uploader := filestorage.NewDir("uploader", utils.StringToPoint("internal"))
	utilsDir := filestorage.NewDir("utils", utils.StringToPoint("internal"))
	logger := filestorage.NewDir("logger", utils.StringToPoint("internal"))

	docs := filestorage.NewDir("docs", nil)
	bin := filestorage.NewDir("bin", nil)
	log := filestorage.NewDir("log", nil)

	directories = append(
		directories,
		cmd,
		api,
		migrate,
		internal,
		db,
		auth,
		service,
		repository,
		keys,
		uploader,
		docs,
		bin,
		migrations,
		utilsDir,
		log,
		logger,
	)

	for _, dir := range directories {
		root.Add(*dir)
	}

	return root
}
