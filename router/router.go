package router

import "net/http"
import "strings"
import "fmt"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Router struct {
	Handlers map[string]map[string]http.HandlerFunc "핸들러 모음"
}

func (router *Router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	//http 메서드로 등록된 맵이 있는지 확인
	m, ok := router.Handlers[method]
	if !ok {
		//등록된 맵이 없다면 새로운 맵을 생성
		m = make(map[string]http.HandlerFunc)
		router.Handlers[method] = m
	}
	//http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

func match(pattern, path string) (bool, map[string]string) {

	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		fmt.Println("1", len(patterns), len(paths))
		fmt.Println("1", patterns, paths)
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			fmt.Println("param!", patterns[i], paths[i])
			params[patterns[i][1:]] = paths[i]
		default:
			fmt.Println("2", patterns, paths)
			return false, nil
		}
	}

	return true, params

}

func (router *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//debug
	fmt.Println(req.Method, req.URL.Path)

	//	if method, ok := router.Handlers[req.Method]; ok {

	for pattern, handler := range router.Handlers[req.Method] {
		if ok, params := match(pattern, req.URL.Path); ok {
			//요청 URL 에 해당하는 핸들러 수행
			fmt.Println(pattern, params)
			handler(res, req)
			return
		}
	}

	http.NotFound(res, req)
	return
}
