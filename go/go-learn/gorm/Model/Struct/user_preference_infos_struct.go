package Struct

import (
	"fmt"
	"hbtest/Model"
)

type UserPreferenceInfos struct {
	Id       int                 `gorm:"primary_key" json:"id"`
	UserId   int                 `json:"user_id,omitempty"`
	Nickname string              `json:"nickname,omitempty"`
	Birth    string              `json:"-"`
	Sex      string              `json:"-"`
	Height   int                 `json:"-"`
	Hobby    string              `json:"-"`
	Remark   []map[string]string `json:"-"`
}

type Test struct {
	Test *UserPreferenceInfos
	Name string
}

//添加数据
func (User *UserPreferenceInfos) Add() {
	err := Model.Db.Create(User).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}
