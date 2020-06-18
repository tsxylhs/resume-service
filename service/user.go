package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"lncios.cn/resume/password"
	"time"
)

var User user

type user int

func (user) Login(form *model.User) (model.User, bool) {
	selectUser := &model.User{}
	if _, err := cs.Sql.Table("user").Where("username=?", form.Username).Get(selectUser); err != nil {
		return *selectUser, false
	}
	if selectUser.Id > 0 {
		if password.Validate(form.Password, selectUser.Password) {
			return *selectUser, true
		} else {
			return *selectUser, false
		}
	}
	return *selectUser, false
}
func (user) Create(form *model.User) error {
	form.BeforeInsert()
	form.Password, _ = password.Encrypt(form.Password)
	if _, err := cs.Sql.Table("user").Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (user) Updata(form *model.User) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("user").Update(form); err != nil {
		return err
	}
	return nil
}
func (user) List(page *model.Page, user *model.User, users *[]model.User) error {
	if err := cs.Sql.Table("user").Where("dtd=false").Limit(page.Limit(), page.Skip()).Find(users, user); err != nil {
		return err
	}
	return nil
}
func (user) Delete(user *model.User) error {
	user.Dtd = true
	if _, err := cs.Sql.Table("user").Update(user); err != nil {
		return err
	}
	return nil
}
func (user) Get(user *model.User) error {

	if _, err := cs.Sql.Table("user").ID(user.Id).Update(user); err != nil {
		return err
	}
	return nil
}
