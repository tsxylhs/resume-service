package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"time"
)

var WorkExprience workExprience

type workExprience int

func (workExprience) Create(form *model.WorkExprience) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Table("project_exprience").Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (workExprience) Updata(form *model.WorkExprience) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("project_exprience").Update(form); err != nil {
		return err
	}
	return nil
}
func (workExprience) List(page *model.Page, workExprience *model.WorkExprience, users *[]model.WorkExprience) error {
	if err := cs.Sql.Table("project_exprience").Where("dtd=false").Limit(page.Limit(), page.Skip()).Find(users, workExprience); err != nil {
		return err
	}
	return nil
}
func (workExprience) Delete(workExprience *model.WorkExprience) error {
	workExprience.Dtd = true
	if _, err := cs.Sql.Table("project_exprience").Update(workExprience); err != nil {
		return err
	}
	return nil
}
func (workExprience) Get(workExprience *model.WorkExprience) error {

	if _, err := cs.Sql.Table("project_exprience").ID(workExprience.Id).Update(workExprience); err != nil {
		return err
	}
	return nil
}
