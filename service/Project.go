package service

import (
	"io"
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"lncios.cn/resume/util"
	"mime/multipart"
	"os"
	"time"
)

var ProjectExprience projectExprience

type projectExprience int

func (projectExprience) Create(form *model.ProjectExprience) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (projectExprience) Updata(form *model.ProjectExprience) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Id(form.Id).Update(form); err != nil {
		return err
	}
	return nil
}
func (projectExprience) List(page *model.Page, projectExprience *model.ProjectExprience, users *[]model.ProjectExprience) error {
	ss := cs.Sql.NewSession()
	defer ss.Close()
	if page.K != "" {
		ss.Where("kind=?", page.K)
	}
	if cnt, err := ss.Where("dtd=false").Limit(page.Limit(), page.Skip()).FindAndCount(users, projectExprience); err != nil {
		return err
	} else {
		page.Cnt = cnt
		return nil
	}
}
func (projectExprience) Delete(projectExprience *model.ProjectExprience) error {
	projectExprience.Dtd = true
	if _, err := cs.Sql.Update(projectExprience); err != nil {
		return err
	}
	return nil
}
func (projectExprience) Get(projectExprience *model.ProjectExprience) error {

	if _, err := cs.Sql.ID(projectExprience.Id).Get(projectExprience); err != nil {
		return err
	}
	return nil
}
func (projectExprience) UploadFile(form *model.File, file multipart.File) error {
	name := util.GeneralNumber(10)
	storepath := "./images" // 存储位置
	reldoc := "/" + time.Now().Format("2006-01-02")

	if err := os.MkdirAll(storepath+"/"+reldoc, 0755); err != nil {
		return err
	}
	relpath := storepath + "/" + reldoc + "/" + name + form.Suffix
	out, err := os.Create(relpath)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err = io.Copy(out, file); err != nil {
		return err
	}

	form.RelativePath = reldoc + "/" + name + form.Suffix
	form.PrefixUri = "http://www.lncios.cn"
	form.UniqueName = name
	return nil
}
