package models

import (
	"github.com/mir-mirsodikov/vulcan/utils"
)

type Component struct {
	Name            string
	ServerComponent bool
	Path            string
}

type Page struct {
	Name            string
	ServerComponent bool
	Path            string
}

type ResourceInterface interface {
	SetPath(path string)
}

func (p *Page) SetPath(desiredPath string) {
	if desiredPath != "" {
		p.Path = "app/" + desiredPath + "/" + p.Name + "/"
	} else {
		p.Path = "app/" + p.Name + "/"
	}
}

func (c *Component) SetPath(desiredPath string) {
	if desiredPath != "" {
		c.Path = "component/" + desiredPath + "/"
	} else {
		c.Path = "component/"
	}
}

func (p *Page) GetPagePath() string {
	ext := utils.GetFileExtension()
	return p.Path + "page" + ext
}

func (c *Component) GetComponentPath() string {
	ext := utils.GetFileExtension()
	return c.Path + c.Name + ext
}
