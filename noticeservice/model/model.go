package model

import (
	_ "github.com/lib/pq"
)

type Model struct {
	CreateUser STRING `db:"createuser" json:"createUser,omitempty"`
	CreateTime INT64  `db:"createtime" json:"createTime,omitempty"`
	UpdateUser STRING `db:"updateuser" json:"updateUser,omitempty"`
	UpdateTime INT64  `db:"updatetime" json:"updateTime,omitempty"`
}
