package project

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestParse(t *testing.T) {
	const yml = `
# reviewdog.yml

runner:
  golint:
    cmd: golint ./...
    level: info
    errorformat:
      - "%f:%l:%c: %m"
  govet:
    cmd: go tool vet -all -shadowstrict .
    format: govet
    level: warning
  namekey:
    cmd: echo 'name'
    name: nameoverwritten
    format: checkstyle
    level: error
`

	want := &Config{
		Runner: map[string]*Runner{
			"golint": {
				Cmd:         "golint ./...",
				Errorformat:` ▋
