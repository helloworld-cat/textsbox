package textsbox

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type TextsBox struct {
	Files    []string
	Data     []yaml.MapSlice
	Cache    map[string]interface{}
	KeyAlias map[string]string
}

func New() *TextsBox {
	return &TextsBox{
		Data:     make([]yaml.MapSlice, 0),
		Cache:    make(map[string]interface{}),
		KeyAlias: make(map[string]string),
	}
}

func (tb *TextsBox) AddKeyAlias(key string, aliases ...string) {
	for _, alias := range aliases {
		tb.KeyAlias[alias] = key
	}
}

func (tb *TextsBox) LoadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tb.Load(file); err != nil {
		return err
	}

	return nil
}

func (tb *TextsBox) Load(r io.Reader) error {
	fc, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	ms := yaml.MapSlice{}
	if err := yaml.Unmarshal(fc, &ms); err != nil {
		return err
	}
	tb.Data = append(tb.Data, ms)

	return nil
}

func (tb *TextsBox) ResetCache() {
	tb.Cache = make(map[string]interface{})
}

func (tb *TextsBox) Find(key, pattern string) (interface{}, error) {
	alias, exists := tb.KeyAlias[key]
	if exists {
		key = alias
	}

	pattern = fmt.Sprintf("%s.%s", key, pattern)

	if value, exists := tb.Cache[pattern]; exists {
		return value, nil
	}

	for _, ms := range tb.Data {
		if value, err := find(pattern, ms); err == nil {
			tb.Cache[pattern] = value
			return value, nil
		}
	}
	return "", ErrNotFound{Pattern: pattern}
}

type ErrNotFound struct {
	Pattern string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("Pattern `%s` not found", e.Pattern)
}

func find(pattern string, ms yaml.MapSlice) (interface{}, error) {
	fItems := strings.Split(pattern, ".")
	for _, item := range ms {
		if item.Key.(string) == fItems[0] {
			if v, ok := item.Value.(yaml.MapSlice); ok {
				return find(strings.Join(fItems[1:], "."), v)
			}
			return item.Value, nil
		}
	}
	return "", ErrNotFound{Pattern: pattern}
}
