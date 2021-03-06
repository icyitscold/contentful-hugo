package translate

import (
	"strings"

	"github.com/icyitscold/contentful-hugo/mapper"
)

func Dir(item mapper.Item) string {
	dir := "./content/"
	contentType := item.ContentType()
	if contentType != "homepage" {
		dir += strings.ToLower(contentType) + "/"
	}
	return dir
}

func Filename(item mapper.Item) string {
	dir := Dir(item)
	if dir == "./content/" {
		return dir + "_index.md"
	}

	return dir + item.Sys.ID + ".md"
}
