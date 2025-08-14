package filestorage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFile(t *testing.T) {

	testCases := []struct {
		name     string
		input    Data
		file     FileStorage
		expected FileStorage
	}{
		{
			name: "Add sub-directory to empty root",
			input: Data{
				Parent: nil,
				File:   FileStorage{Name: "foo"},
			},
			file: FileStorage{
				Name:     "root",
				Children: nil,
				Content:  nil,
			},
			expected: FileStorage{
				Name:    "root",
				Content: nil,
				Children: []*FileStorage{
					{Name: "foo", Children: nil, Content: nil},
				},
			},
		},
		{
			name: "Add sub-directory to root that have already subs",
			input: Data{
				Parent: nil,
				File: FileStorage{
					Name:     "bar",
					Children: nil,
					Content:  nil,
				},
			},
			file: FileStorage{
				Name:    "root",
				Content: nil,
				Children: []*FileStorage{
					{Name: "foo", Children: nil, Content: nil},
				},
			},
			expected: FileStorage{
				Name:    "root",
				Content: nil,
				Children: []*FileStorage{
					{Name: "foo", Children: nil, Content: nil},
					{Name: "bar", Children: nil, Content: nil},
				},
			},
		},
		{
			name: "Add a file to root directory",
			input: Data{
				Parent: nil,
				File: FileStorage{
					Name:     "foo.txt",
					Children: nil,
					Content:  []byte("hello world"),
				},
			},
			file: FileStorage{
				Name:     "root",
				Children: nil,
				Content:  nil,
			},
			expected: FileStorage{
				Name:    "root",
				Content: nil,
				Children: []*FileStorage{
					{Name: "foo.txt", Children: nil, Content: []byte("hello world")},
				},
			},
		},
		{
			name: "Add a file to sub-directory",
			input: Data{
				Parent: stringToP("foo"),
				File: FileStorage{
					Name:     "bar.txt",
					Content:  []byte("lizzy mcalpine"),
					Children: nil,
				},
			},
			file: FileStorage{
				Name:    "lizzy-mcalpine",
				Content: nil,
				Children: []*FileStorage{
					{
						Name:     "foo",
						Children: nil,
						Content:  nil,
					},
					{
						Name:     "xoxo",
						Children: nil,
						Content:  nil,
					},
				},
			},
			expected: FileStorage{
				Name:    "lizzy-mcalpine",
				Content: nil,
				Children: []*FileStorage{
					{
						Name:    "foo",
						Content: nil,
						Children: []*FileStorage{
							{
								Name:     "bar.txt",
								Content:  []byte("lizzy mcalpine"),
								Children: nil,
							},
						},
					},
					{
						Name:     "xoxo",
						Children: nil,
						Content:  nil,
					},
				},
			},
		},
		{
			name: "Add a file to sub of sub-directory",
			input: Data{
				Parent: stringToP("db"),
				File: FileStorage{
					Name:     "db.go",
					Children: nil,
					Content:  []byte("lizzy mcalpine"),
				},
			},
			file: FileStorage{
				Name:    "falcon",
				Content: nil,
				Children: []*FileStorage{
					{
						Name:    "internal",
						Content: nil,
						Children: []*FileStorage{
							{
								Name:     "db",
								Children: nil,
								Content:  nil,
							},
							{
								Name:     "services",
								Children: nil,
								Content:  nil,
							},
						},
					},
				},
			},
			expected: FileStorage{
				Name:    "falcon",
				Content: nil,
				Children: []*FileStorage{
					{
						Name:    "internal",
						Content: nil,
						Children: []*FileStorage{
							{
								Name:    "db",
								Content: nil,
								Children: []*FileStorage{
									{
										Name:     "db.go",
										Children: nil,
										Content:  []byte("lizzy mcalpine"),
									},
								},
							},
							{
								Name:     "services",
								Children: nil,
								Content:  nil,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tc.file.Add(tc.input)

			require.EqualValues(t, tc.expected, tc.file, "should match")

		})
	}

}

func stringToP(text string) *string {

	return &text
}
