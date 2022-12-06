package types

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

//go:embed testdata/cloud.json
var CLOUD_JSON string

func Test(t *testing.T) { TestingT(t) }

type JsonSuite struct{}

var _ = Suite(&JsonSuite{})

func (s *JsonSuite) TestLoadJson(c *C) {
	json.NewDecoder(strings.NewReader(CLOUD_JSON))
}
