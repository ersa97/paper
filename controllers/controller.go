package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ersa97/paper-test/models"
	"github.com/jinzhu/gorm"
)

type PaperService struct {
	DB *gorm.DB
}

func (m *PaperService) TestAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "API is working",
		Data:    nil,
	})
}
