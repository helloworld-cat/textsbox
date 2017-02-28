package textsbox

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
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

	patternPath := strings.Split(pattern, ".")
	for _, ms := range tb.Data {
		if value, err := find(patternPath, ms); err == nil {
			tb.Cache[pattern] = value
			return value, nil
		}
	}
	return "", ErrNotFound{PatternPath: patternPath}
}

type ErrNotFound struct {
	PatternPath []string
}

func (e ErrNotFound) Error() string {
	pattern := strings.Join(e.PatternPath, ".")
	return fmt.Sprintf("Pattern `%s` not found", pattern)
}

func find(patternPath []string, ms yaml.MapSlice) (interface{}, error) {
	for _, item := range ms {
		if item.Key.(string) == patternPath[0] {
			if v, ok := item.Value.(yaml.MapSlice); ok {
				return find(patternPath[1:], v)
			}
			return item.Value, nil
		}
	}
	return "", ErrNotFound{PatternPath: patternPath}
}
