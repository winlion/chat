package model

import "time"

type Community struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//名称
	Name   string 		`xorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid       int64	`xorm:"bigint(20)" form:"ownerid" json:"ownerid"`   // 什么角色
	//群logo
	Icon	   string 		`xorm:"varchar(250)" form:"icon" json:"icon"`
	//como
	Cate      int	`xorm:"int(11)" form:"cate" json:"cate"`   // 什么角色
	//描述
	Memo    string	`xorm:"varchar(120)" form:"memo" json:"memo"`   // 什么角色
	//
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
}
const (
		COMMUNITY_CATE_COM = 0x01
	)