package handlerservicerepository

import (
	"testing"

	filestorage "github.com/faizisyellow/falcon/internal/fileStorage"
	"github.com/stretchr/testify/require"
)

func TestCreateDirectories(t *testing.T) {

	testCase := struct {
		expected filestorage.FileStorage
	}{
		expected: filestorage.FileStorage{
			Name: "",
			Children: []*filestorage.FileStorage{
				{
					Name:    "cmd",
					Content: nil,
					Children: []*filestorage.FileStorage{
						{
							Name:     "api",
							Children: nil,
							Content:  nil,
						},
						{
							Name:    "migrate",
							Content: nil,
							Children: []*filestorage.FileStorage{
								{
									Name:     "migrations",
									Children: nil,
									Content:  nil,
								},
							},
						},
					},
				},
				{
					Name:    "internal",
					Content: nil,
					Children: []*filestorage.FileStorage{
						{
							Name:     "db",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "auth",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "service",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "repository",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "keys",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "uploader",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "utils",
							Children: nil,
							Content:  nil,
						},
						{
							Name:     "logger",
							Children: nil,
							Content:  nil,
						},
					},
				},
				{
					Name:     "docs",
					Content:  nil,
					Children: nil,
				},
				{
					Name:     "bin",
					Content:  nil,
					Children: nil,
				},
				{
					Name:     "log",
					Content:  nil,
					Children: nil,
				},
			},
		},
	}

	t.Run("should success create directories", func(t *testing.T) {

		dirs := New()

		require.EqualValues(t, testCase.expected, dirs)
	})
}
