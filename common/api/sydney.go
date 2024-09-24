package api

import (
	"adams549659584/go-proxy-bingai/common"
	"adams549659584/go-proxy-bingai/common/helper"
	"net/http"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	// 检查并删除 USRLOC cookie
//	if _, err := r.Cookie("USRLOC"); err == nil {
//		c := &http.Cookie{
//			Name:   "USRLOC",
//			MaxAge: -1,
//		}
//		http.SetCookie(w, c)
//	}
	// 设置 USRLOC cookie 的值为 JA=1
	c := &http.Cookie{
		Name:  "USRLOC",
		Value: "JA=1",
		Path:  "/",
	}
	http.SetCookie(w, c)

	// 检查用户是否已认证
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}

	// 代理请求到 BING_SYDNEY_URL
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}

