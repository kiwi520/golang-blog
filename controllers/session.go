package controllers

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"net/url"
	_ "reflect"
	"sync"
	"time"
)

type SessionManager struct {
	Lock       sync.Mutex
	Smap       map[string]*list.Element
	SL         *list.List //gc
	Cookiename string
	Expires    int
}
type Session struct {
	Key     string
	Sid     string
	Expires int
	Value   interface{}
}

/*
说明：只在start中用，并已经加锁，因此这里不需要,否则会引起死锁
*/
func (sm *SessionManager) get(sid string) (Session, error) {
	if element, ok := sm.Smap[sid]; ok {
		sm.updateList(sid)
		return *(element.Value.(*Session)), nil
	}
	fmt.Printf("can't find session by sid:%s\n", sid)
	return *sm.NewSession("", sid, ""), errors.New("can't find session by sid")
}

func (sm *SessionManager) NewSession(key string, sid string, value interface{}) *Session {
	return &Session{
		Key:     key,
		Sid:     sid,
		Expires: sm.Expires,
		Value:   value,
	}
}

func (sm *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	cookie, err := r.Cookie(sm.Cookiename)
	if err != nil || cookie.Value == "" {
		sid := sm.sessionId()
		session = *sm.NewSession("", sid, "")
		cookie := http.Cookie{Name: sm.Cookiename, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: sm.Expires}
		http.SetCookie(w, &cookie)
		//fmt.Printf("\ncookie.value:%s,sid:%s\n", cookie.Value, sid)
		//log.Printf("cookiename:%s,err:%s", cookie, err.Error())
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		//fmt.Printf("\nget(sid):cookie.value:%s,sid:%s\n", cookie.Value, sid)
		session, _ = sm.get(sid)
		//log.Printf("key:%s,cookie.value:%s,sid:%s\n", session.Key, cookie.Value, sid)
	}
	return
}

//垃圾回收
func (sm *SessionManager) Listen() {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	var n *list.Element
	for e := sm.SL.Front(); e != nil; e = n {
		n = e.Next()
		if e.Value.(*Session).Expires == 0 {
			delete(sm.Smap, e.Value.(*Session).Sid)
			sm.SL.Remove(e)
		} else {
			e.Value.(*Session).Expires--
		}
	}

	time.AfterFunc(time.Duration(sm.Expires)*time.Second, func() { sm.Listen() })
}

func (sm *SessionManager) GetByKey(key string) (Session, error) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	for _, element := range sm.Smap {
		if element.Value.(*Session).Key == key {
			sm.updateList(element.Value.(*Session).Sid)
			return *(element.Value.(*Session)), nil
		}
	}
	fmt.Println("can't find session by key")
	return *sm.NewSession("", "", ""), errors.New("can't find session by key")
}

func (sm *SessionManager) Set(key string, sid string) (isExits bool) {
	isExits = false
	if _, ok := sm.Smap[sid]; ok {
		sm.SL.Remove(sm.Smap[sid])
		isExits = true
	}

	element := sm.SL.PushBack(sm.NewSession(key, sid, ""))
	sm.Smap[sid] = element
	return
}

func (sm *SessionManager) Del(key string) (bl bool) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	bl = false
	for _, element := range sm.Smap {
		if element.Value.(*Session).Key == key {
			delete(sm.Smap, element.Value.(*Session).Sid)
			sm.SL.Remove(element)
			bl = true
		}
	}
	return
}

//private函数，Lock之后使用，不用也不能上锁，否则死锁
func (sm *SessionManager) updateList(sid string) {

	if element, ok := sm.Smap[sid]; ok {
		element.Value.(*Session).Expires = sm.Expires
		sm.SL.MoveToBack(element)
	}
}

func (sm *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}