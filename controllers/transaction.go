package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ersa97/paper-test/helpers"
	"github.com/ersa97/paper-test/models"
	"github.com/gorilla/mux"
)

func (m *PaperService) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}

	var body models.Transaction
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	userid := helpers.GetAuthorizationTokenValue(r, "userid")

	body.UserId = int(userid.(float64))

	result, err := models.AddTransaction(body, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "add transaction success",
		Data: map[string]interface{}{
			"trxid":      result.TrxId,
			"code_from":  result.CodeFrom,
			"code_to":    result.CodeTo,
			"user_id":    result.UserId,
			"amount":     result.Amount,
			"status":     helpers.GetStatus(result.Status),
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt,
		},
	})
}

func (m *PaperService) GetDetailTransaction(w http.ResponseWriter, r *http.Request) {
	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}
	trxid, _ := strconv.Atoi(mux.Vars(r)["trxid"])
	result, err := models.GetDetailTransaction(trxid, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get detail transaction success",
		Data: map[string]interface{}{
			"trxid":      result.TrxId,
			"code_from":  result.CodeFrom,
			"code_to":    result.CodeTo,
			"user_id":    result.UserId,
			"amount":     result.Amount,
			"status":     helpers.GetStatus(result.Status),
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt,
		},
	})

}

func (m *PaperService) GetListTransaction(w http.ResponseWriter, r *http.Request) {
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
	code, err := strconv.Atoi(r.URL.Query()["code"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required page limit",
		})
		return
	}

	result, err := models.GetListTransaction(limit, page, code, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get transaction list success",
		Data:    result,
	})
}

func (m *PaperService) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var trx models.Transaction

	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}

	trxid, _ := strconv.Atoi(mux.Vars(r)["trxid"])
	err := json.NewDecoder(r.Body).Decode(&trx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}
	trx.TrxId = trxid
	userid := helpers.GetAuthorizationTokenValue(r, "userid")
	trx.UserId = int(userid.(float64))

	_, err = models.UpdateTransaction(trx, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	result, err := models.GetDetailTransaction(trx.TrxId, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "update transaction success",
		Data: map[string]interface{}{
			"trxid":      result.TrxId,
			"code_from":  result.CodeFrom,
			"code_to":    result.CodeTo,
			"user_id":    result.UserId,
			"amount":     result.Amount,
			"status":     helpers.GetStatus(result.Status),
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt,
		},
	})
}
func (m *PaperService) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	verif := helpers.VerifyUuid(r, m.DB)
	if !verif {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "you are not logged in",
		})
		return
	}

	trxid, _ := strconv.Atoi(mux.Vars(r)["trxid"])

	trx := models.Transaction{
		TrxId: trxid,
	}

	err := models.DeleteTransaction(trx, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "delete transaction success",
	})
}
