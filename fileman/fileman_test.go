package fileman

import (
	"github.com/almonteb/buildmaid/config"
	"reflect"
	"testing"
)

func TestFactoryFs(t *testing.T) {
	c := config.Project{FileMan: "fs"}
	fm, err := NewFileMan(c)
	if err != nil {
		t.Error(err)
	}
	fmType := reflect.TypeOf(fm)
	if fmType.String() != reflect.TypeOf(&FileManFs{}).String() {
		t.Errorf("Incorrect filemanager: %s", fmType.String())
	}
}

func TestFactoryNeg(t *testing.T) {
	c := config.Project{FileMan: "notreal"}
	fm, err := NewFileMan(c)
	if fm != nil {
		t.Errorf("Expected no filemanager for %+v", c)
	}
	if err == nil {
		t.Error("Should have received an error")
	}
}
