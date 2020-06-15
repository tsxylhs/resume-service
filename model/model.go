package model

import (
	"github.com/bwmarrin/snowflake"

	//"lncios.cn/resume/cs"

	//"lncios.cn/resume/cs"
	"time"
)

var idgen *snowflake.Node

type Base struct {
	Id  int64     `xorm:"pk  'id'" json:"id,string" form:"id"`
	Crt time.Time `xorm:"crt" json:"crt"`
	Lut time.Time `xorm:"lut" json:"lut"`
	Dtd bool      `xorm:"dtd" json:"-"`
}

type User struct {
	Base     `xorm:"extends"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Mcc      string `json:"-" form:"mcc"`
	Mobile   string `json:"mobile" form:"mobile"`
	AvatarId string `json:"-"`
	Language int    `json:"-"`
	ImageOne string `json:"imageOne"`
	ImageTwo string `json:"imageTwo"`
	OpenId   string
	RoleIds  []string
	Groups   []string
	OrgId    int64 `json:",string"`
	Status   int
	Content  string `json:"content"`
	IdNo     string `json:"idNo" form:"idNo"` // 身份证号码
	Code     string ` xorm:"-" json:"code"`
}
type Role struct {
	Id          string
	Name        string `json:"name" form:"name"`
	Code        string `form:"code"`
	Permissions []string
	Description string    `json:"description" form:"description"`
	Crt         time.Time `json:"crt"`
	Lut         time.Time `json:"-"`
	Dtd         bool
	OwnerId     int64
}
type Message struct {
	Base     `xorm:"extends"`
	Commpany string `json:"commpany"`
	Email    string `json:"email"`
	Content  string `json:"content"`
}
type ProjectExprience struct {
	Base       `xorm:"extends"`
	Title      string `json:"title" form:"title"`
	SubTitle   string `json:"subTitle" form:"subTitle"`
	CoverImage string `json:"coverImage" from:"coverImage"`
	Content    string `json:"content" from:"content"`
}
type WorkExprience struct {
	Base      `xorm:"extends"`
	Commpany  string    `json:"commpany"`
	Jobs      string    `json:"jobs"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Harvest   string    `json:"harvest"`
}
type Education struct {
	Base    `xorm:"extends"`
	Name    string `json:"name"`
	Harvest string `json:"harvest"`
}
type Page struct {
	P   int    `json:"p" form:"p"`
	Ps  int    `json:"ps" form:"ps"`
	Cnt int64  `json:"cnt"`
	K   string `josn:"k"`
	Pc  int    `json:"pc"`
	Od  string `json:"od,omitempty"`
}

func (page *Page) GetPage() *Page {
	return page
}

func (page *Page) GetPager(count int64) *Page {
	page.Cnt = count
	if page.P < 1 {
		page.P = 1
	}
	if page.Ps < 1 {
		page.Ps = 10
	}
	page.Pc = int(page.Cnt)/page.Ps + 1
	return page
}

func (page *Page) Skip() int {
	if page.Ps > 0 {
		return (page.P - 1) * page.Ps
	}

	return (page.P - 1) * 10
}

func (page *Page) Limit() int {
	if page.Ps > 0 {
		return page.Ps
	}

	return 10
}

func (b *Base) BeforeInsert() {
	b.Id, _ = Next()
	now := time.Now()
	b.Crt = now
	b.Lut = now
}

var node *snowflake.Node

func Next() (int64, error) {
	return int64(node.Generate()), nil
}
func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}
