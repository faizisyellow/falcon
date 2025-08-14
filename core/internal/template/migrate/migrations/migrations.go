package migrations

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed 000001_create_users_table.up.sql.tmpl
var createUserMigrateUp embed.FS

func CreateUserMigrateUpData(data any) ([]byte, error) {

	temp, err := template.ParseFS(createUserMigrateUp, "000001_create_users_table.up.sql.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

//go:embed 000001_create_users_table.down.sql.tmpl
var createUserMigrateDown embed.FS

func CreateUserMigrateDownData(data any) ([]byte, error) {

	temp, err := template.ParseFS(createUserMigrateDown, "000001_create_users_table.down.sql.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

//go:embed 000002_create_invitations_table.down.sql.tmpl
var createInvitationsMigrateDown embed.FS

func CreateInvitationsMigrateDownData(data any) ([]byte, error) {

	temp, err := template.ParseFS(createInvitationsMigrateDown, "000002_create_invitations_table.down.sql.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

//go:embed 000002_create_invitations_table.up.sql.tmpl
var createInvitationsMigrateUp embed.FS

func CreateInvitationsMigrateUpData(data any) ([]byte, error) {

	temp, err := template.ParseFS(createInvitationsMigrateUp, "000002_create_invitations_table.up.sql.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
