package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"time"
)

var ProjectExprience projectExprience

type projectExprience int

func (projectExprience) Create(form *model.ProjectExprience) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Table("project_exprience").Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (projectExprience) Updata(form *model.ProjectExprience) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("project_exprience").Update(form); err != nil {
		return err
	}
	return nil
}
func (projectExprience) List(page *model.Page, projectExprience *model.ProjectExprience, users *[]model.ProjectExprience) error {
	if err := cs.Sql.Table("project_exprience").Where("dtd=false").Limit(page.Limit(), page.Skip()).Find(users, projectExprience); err != nil {
		return err
	}
	return nil
}
func (projectExprience) Delete(projectExprience *model.ProjectExprience) error {
	projectExprience.Dtd = true
	if _, err := cs.Sql.Table("project_exprience").Update(projectExprience); err != nil {
		return err
	}
	return nil
}
func (projectExprience) Get(projectExprience *model.ProjectExprience) error {

	if _, err := cs.Sql.Table("project_exprience").ID(projectExprience.Id).Update(projectExprience); err != nil {
		return err
	}
	return nil
}
