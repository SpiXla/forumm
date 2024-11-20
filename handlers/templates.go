package handlers

import (
	"text/template"
)

var (
	IndTp    *template.Template
	LogTp    *template.Template
	RegTp    *template.Template
	PostTp   *template.Template
	LoggedTp *template.Template
	CatTp *template.Template
	LogCatTp *template.Template
	MyPostTp *template.Template
	ProfileTp *template.Template
)

func ParseFiles() error {
	var err error

	IndTp, err = template.ParseFiles("html/index.html")
	if err != nil {
		return err
	}

	LogTp, err = template.ParseFiles("html/login.html")
	if err != nil {
		return err
	}

	RegTp, err = template.ParseFiles("html/register.html")
	if err != nil {
		return err
	}
	LoggedTp, err = template.ParseFiles("html/logged.html")
	if err != nil {
		return err
	}
	PostTp, err = template.ParseFiles("html/post.html")
	if err != nil {
		return err
	}
	CatTp, err = template.ParseFiles("html/category.html")
	if err != nil {
		return err
	}
	LogCatTp, err = template.ParseFiles("html/category-log.html")
	if err != nil {
		return err
	}
	MyPostTp, err = template.ParseFiles("html/mypost.html")
	if err != nil {
		return err
	}
	ProfileTp, err = template.ParseFiles("html/profile.html")
	if err != nil {
		return err
	}

	return nil
}
