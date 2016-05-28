package Golang

import "net/http"

type router struct  {
	handlers map[string]map[string]http.HandlerFunc "핸들러 모음"
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc){
	//http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok{
		//등록된 맵이 없다면 새로운 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	//http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

type Handler interface {
	ServerHTTP(http.ResponseWriter,*http.Request)
}

func (r* router) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			//요청 URL 에 해당하는 핸들러 수행
			h(w, req)
		}
	}
	http.NotFound(w,req)
}


