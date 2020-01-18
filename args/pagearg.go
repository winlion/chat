package args

import (
	"fmt"
	"time"
)

type PageArg struct {
	//从哪页开始
	Pagefrom int  `json:"pagefrom" form:"pagefrom"`
	//每页大小
	Pagesize int  `json:"pagesize" form:"pagesize"`
	//关键词
	Kword string  `json:"kword" form:"kword"`
	//asc：“id”  id asc
	Asc string 	`json:"asc" form:"asc"`
	Desc string 	`json:"desc" form:"desc"`
	//
	Name string 	`json:"name" form:"name"`
	//
	Userid int64 	`json:"userid" form:"userid"`
	//dstid
	Dstid int64 	`json:"dstid" form:"dstid"`
	//时间点1
	Datefrom time.Time 	`json:"datafrom" form:"datafrom"`
	//时间点2
	Dateto time.Time 	`json:"dateto" form:"dateto"`
	//
	Total  int64  		`json:"total" form:"total"`
}

func (p*PageArg) GetPageSize() int{
	if p.Pagesize==0{
		return 100
	}else{
		return p.Pagesize
	}

}
func (p*PageArg) GetPageFrom() int{
	if p.Pagefrom<0{
		return 0
	}else{
		return p.Pagefrom
	}
}

func (p*PageArg) GetOrderBy() string{
	if len(p.Asc)>0{
		return fmt.Sprintf(" %s asc",p.Asc)
	} else if len(p.Desc)>0{
		return fmt.Sprintf(" %s desc",p.Desc)
	}else{
		return ""
	}
}

