package types

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

//go:embed testdata/cloud.json
var CLOUD_JSON string

func Test(t *testing.T) { TestingT(t) }

type JsonSuite struct{}

var _ = Suite(&JsonSuite{})

// Load static json API description
func (s *JsonSuite) TestLoadJson(c *C) {
	var result ApiDescriptionDocument
	if err := json.NewDecoder(strings.NewReader(CLOUD_JSON)).Decode(&result); err != nil {
		c.Errorf("Error deserialising cloud.json")
	}
}

func (s *JsonSuite) TestLoadJsonFindUnknownFields(c *C) {
	var result ApiDescriptionDocument
	json.NewDecoder(strings.NewReader(CLOUD_JSON)).Decode(&result)
	a := func(node *JsonNode) {
		v := reflect.ValueOf(node)
		if v.Kind() != reflect.Ptr && v.Kind() != reflect.Interface {
			return
		}
		if f := reflect.ValueOf(node).Elem().FieldByName("X"); f.IsValid() {
			if len(f.MapKeys()) != 0 {
				c.Errorf("Not empty value in ...")
			}
		}
		fmt.Printf("Visited %s\n", node.Path)
	}
	(&JsonNode{Node: result, Path: "#"}).Accept(a)
}

type JsonVisitor interface {
	Visit(node *JsonNode)
}

type JsonNode struct {
	Node any
	Path string
}

func (self *JsonNode) Accept(visitor func(*JsonNode)) {
	v := reflect.ValueOf(self.Node)
	if v.IsZero() {
		return
	}
	t := v.Type()
	if t.Kind() == reflect.Struct {
		numField := v.NumField()
		for idx := 0; idx < numField; idx++ {
			if t.Field(idx).Name == "X" {
				continue
			}
			if v.Field(idx).Kind() == reflect.Interface {
				(&JsonNode{Node: v.Field(idx).Interface(), Path: self.Path + "." + t.Field(idx).Name}).Accept(visitor)
			}
			if v.Field(idx).Kind() == reflect.Slice {
				for idx2 := 0; idx2 < v.Field(idx).Len(); idx2++ {
					var value any
					if v.Field(idx).Index(idx2).CanInterface() {
						value = v.Field(idx).Index(idx2).Interface()
						(&JsonNode{Node: value, Path: fmt.Sprintf("%s.%s[%d]", self.Path, t.Field(idx).Name, idx2)}).Accept(visitor)
					}
				}
			}
		}
	}
	visitor(self)
}
