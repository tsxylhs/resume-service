package cs

import (
	"github.com/xormplus/xorm"
	"lncios.cn/resume/newSession"
)

var Sql *xorm.Engine
var SessionMgr *newSession.SessionMgr = nil
