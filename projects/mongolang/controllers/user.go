package controllers

import {
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController(s)
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjextIdHex(id){
		w.WriterHeader(http.StatusNotFound)
	}

	oid := bson.ObjectHex(id)

	u := modules.User{}

	uc.Session.DB("mongolang").C("users").FindId(oid).One
}

// GetUser
// CreateUser
// DeleteUser