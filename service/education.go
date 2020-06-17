package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"time"
)

var Education education

type education int

func (education) Create(form *model.Education) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Table("project_exprience").Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (education) Updata(form *model.Education) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("project_exprience").Update(form); err != nil {
		return err
	}
	return nil
}
func (education) List(page *model.Page, education *model.Education, users *[]model.Education) error {
	if err := cs.Sql.Table("project_exprience").Where("dtd=false").Limit(page.Limit(), page.Skip()).Find(users, education); err != nil {
		return err
	}
	return nil
}
func (education) Delete(education *model.Education) error {
	education.Dtd = true
	if _, err := cs.Sql.Table("project_exprience").Update(education); err != nil {
		return err
	}
	return nil
}
func (education) Get(education *model.Education) error {

	if _, err := cs.Sql.Table("project_exprience").ID(education.Id).Update(education); err != nil {
		return err
	}
	return nil
}
