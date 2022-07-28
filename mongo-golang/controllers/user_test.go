package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Olderest/mongo-golang/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

//mongoDB mock instance

type MockUserController struct{
	UserOperations
}

func (m MockUserController) Create(user models.User, session *mgo.Session){
	
}

func TestCreate(t *testing.T){
	mock := MockUserController{}
	uc := UserController{session:nil, UserOperations: mock}
	values := map[string]string{"name": "Charan",}
	jsonData, err := json.Marshal(values)
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
    	if err != nil {
        	t.Fatal(err)
    	}

    	rr := httptest.NewRecorder()
    	handler :=   http.HandlerFunc(uc.CreateUser)

    	handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, "This is my result", rr.Body.String())
}