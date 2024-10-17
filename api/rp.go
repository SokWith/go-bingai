package api

import (
	"adams549659584/go-proxy-bingai/common"
	"adams549659584/go-proxy-bingai/common/helper"
	"net/http"
	"strings"
)

func Rpcore(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	r.URL.Path = strings.ReplaceAll(r.URL.Path, "/rp/", "/web/rp/")
	common.NewSingleHostReverseProxy(common.BING_RPURL).ServeHTTP(w, r)
}


func CreateMU(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	
	common.NewSingleHostReverseProxy(common.BING_CREATEURL).ServeHTTP(w, r)
}