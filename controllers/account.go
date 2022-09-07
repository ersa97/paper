package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ersa97/paper-test/helpers"
	"github.com/ersa97/paper-test/models"
	"github.com/gorilla/mux"
)

func (m *PaperService) AddAccount(w http.ResponseWriter, r *http.Request) {
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	var body models.Account
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	result, err := models.AddAccount(body, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "add account success",
		Data: map[string]interface{}{
			"code":       result.Code,
			"name":       result.Name,
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt,
		},
	})

}
func (m *PaperService) GetDetailAccount(w http.ResponseWriter, r *http.Request) {

	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	code, _ := strconv.Atoi(mux.Vars(r)["code"])

	result, err := models.GetAccountDetail(code, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get detail account success",
		Data: map[string]interface{}{
			"code":       result.Code,
			"name":       result.Name,
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt,
		},
	})
}
func (m *PaperService) GetAccountList(w http.ResponseWriter, r *http.Request) {
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required field limit",
		})
		return
	}
	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required page limit",
		})
		return
	}

	result, err := models.GetAccountList(limit, page, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get account list success",
		Data:    result,
	})

}
func (m *PaperService) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	code, _ := strconv.Atoi(mux.Vars(r)["code"])
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}
	acc.Code = code

	result, err := models.UpdateAccount(acc, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "update account success",
		Data:    result,
	})

}
func (m *PaperService) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	code, _ := strconv.Atoi(mux.Vars(r)["code"])

	acc := models.Account{
		Code: code,
	}

	err := models.DeleteAccount(acc, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "delete account success",
	})

}
