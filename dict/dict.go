package dict

// 微信接口
const (
	WxLogin  = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WxUnid   = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
	WXreqUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
)

const (

	// Ebl小程序
	LibrarayId    = "wx98bbdb4a7d2c92f3"
	LibrarySecret = "bcb29d9aceddd5e9304222627d517769"
)
