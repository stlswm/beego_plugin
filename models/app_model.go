package model

import "github.com/stlswm/beego_plugin/orm_r"

// AppModel
type AppModel struct {
	orm *orm_r.Ormer
}

// 设置orm
func (m *AppModel) SetOrm(o *orm_r.Ormer) *AppModel {
	m.orm = o
	return m
}

// 获取orm对象
func (m *AppModel) GetOrm() *orm_r.Ormer {
	if m.orm != nil {
		return m.orm
	}
	m.orm = orm_r.NewOrm()
	return m.orm
}
