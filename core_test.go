package textsbox

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tb := New()
	assert.NotNil(t, tb)
}

func TestLoad(t *testing.T) {
	file, err := os.Open("./fixtures/fr.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tb := New()
	err = tb.Load(file)
	assert.Nil(t, err)
}

func TestFind(t *testing.T) {
	tb := New()

	filenameList := []string{"fr.yml", "en.yml"}
	for _, lang := range filenameList {
		filename := fmt.Sprintf("./fixtures/%s", lang)
		err := tb.LoadFile(filename)
		assert.Nil(t, err)
	}

	tb.AddKeyAlias("fr", "fr-FR", "FR-FR")

	data := map[string]map[string]interface{}{
		"fr": {
			"hello_msg": "Bonjour tout le monde",
		},
		"FR-FR": {
			"hello_msg": "Bonjour tout le monde",
		},
		"fr-FR": {
			"hello_msg": "Bonjour tout le monde",
		},
		"en": {
			"hello_msg":    "Hello World !",
			"nested.value": "value",
		},
	}

	for key, values := range data {
		for pattern, expectedValue := range values {
			value, err := tb.Find(key, pattern)
			assert.Nil(t, err)
			assert.Equal(t, expectedValue, value)
		}
	}

	_, err := tb.Find("fr", "foo.bar")
	assert.Error(t, ErrNotFound{}, err)
}
