package utils

import "testing"

func TestGetModule(t *testing.T) {

	t.Run("should get module name", func(t *testing.T) {

		exist, err := IsModuleExist()
		if err != nil {
			t.Error(err)
			return
		}

		var result string
		expected := "github.com/faizisyellow/falcon"

		if exist {
			result = GetModuleName()
		}

		if result != expected {
			t.Errorf("expected %v but got %v", expected, result)
		}

	})
}
