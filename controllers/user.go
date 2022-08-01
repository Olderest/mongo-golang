package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Olderest/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
	UserOperations
}

type UserOperations interface{
	Create(user models.User,session *mgo.Session)
	DeleteUser()
	GetUser()
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{session:s}
}

func (uc UserController) Create(user models.User, session *mgo.Session){

	session.DB("User_collection").C("Users").Insert(user)
	
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("User_collection").C("Users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.UserOperations.Create(u, uc.session)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

//func (uc UserController) UpdateUser(w http.ResponseWriter, r http.Request, p httprouter.Params){
//	id := p.ByName("id")
//
//	if !bson.IsObjectIdHex(id){
//		w.WriteHeader(http.StatusNotFound)
//	} 
//
//	oid := bson.ObjectIdHex(id)
//
//	if err := uc.session.DB("User_collection").C("Users").Remove(oid); err != nil{
//		w.WriteHeader(404)
//	}
//
//	u := models.User{}
//
//	json.NewDecoder(r.Body).Decode(&u)
//
//	u.Id = bson.ObjectId(id)
//
//	uc.session.DB("User_collection").C("Users").Insert(u)
//
//	uj, err := json.Marshal(u)
//	if err != nil {
//		fmt.Println(err)
//	}
//	w.Header().Set("Content-type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	fmt.Fprintf(w, "%s\n", uj)
//
//}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("User_collection").C("Users").Remove(oid); err != nil{
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s Deleted User", oid)
}
