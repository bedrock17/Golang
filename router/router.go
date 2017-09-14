package router

import "net/http"

import "fmt"

type Handler interface {
	ServeHTTP(http.ResponseWriter,*http.Request)
}

type Router struct  {
	Handlers map[string]map[string]http.HandlerFunc "핸들러 모음"
}


func (r *Router) HandleFunc(method, pattern string, h http.HandlerFunc) {

	//http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.Handlers[method]
	if !ok{
		//등록된 맵이 없다면 새로운 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.Handlers[method] = m
	}
	//http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}


func (r *Router) ServeHTTP(res http.ResponseWriter,  req *http.Request) {

	//debug
	fmt.Println(req.Method, req.URL.Path)

	if m, ok := r.Handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			//요청 URL 에 해당하는 핸들러 수행
			h(res, req)
			return
		}
	}

	http.NotFound(res, req)
	return 
}


