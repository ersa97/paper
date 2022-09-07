package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ersa97/paper-test/helpers"
	"github.com/ersa97/paper-test/models"
	"github.com/pborman/uuid"
)

func (m *PaperService) Register(w http.ResponseWriter, r *http.Request) {

	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	exist, err := models.CheckUser(body.Email, m.DB)
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	passwordbyte := []byte(body.Password)
	passwordmd5 := md5.Sum(passwordbyte)
	password := hex.EncodeToString(passwordmd5[:])

	body.Password = string(password)

	result, err := models.AddUser(body, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	err = models.AddUserToken(result.Id, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "register successful",
		Data: map[string]interface{}{
			"Id":    result.Id,
			"Name":  result.Name,
			"Email": result.Email,
		},
	})
}

func (m *PaperService) Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	result, err := models.GetUser(body.Email, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	if fmt.Sprintf("%x", md5.Sum([]byte(body.Password))) != result.Password {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "wrong password",
		})
		return
	}
	uid := uuid.New()
	err = models.UpdateUserToken(result.Id, uid, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	token := helpers.GenerateToken(result.Id, result.Name, body.Email, uid)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "login success",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func (m *PaperService) Logout(w http.ResponseWriter, r *http.Request) {

	id := int(helpers.GetAuthorizationTokenValue(r, "userid").(float64))

	err := models.UpdateUserToken(id, "", m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "logout successful",
	})
}
