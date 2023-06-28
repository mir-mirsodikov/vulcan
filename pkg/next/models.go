package next

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
	GetDirPath() string
	GetResourcePath() string
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

func (p *Page) GetPagePath(ext string) string {
	return p.Path + "page" + ext
}

func (c *Component) GetComponentPath(ext string) string {
	return c.Path + c.Name + ext
}
