package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"csv-storage/db"
	"csv-storage/models"
	"csv-storage/utils"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the CSV storage service")
}

// Create a Promotion from the request body
func CreatePromotion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := r.Context()

	var promotion models.Promotion

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	if err := r.Body.Close(); err != nil {
		utils.HandleError(err)
	}

	// Save JSON to the promotion struct
	if err := json.Unmarshal(body, &promotion); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		err = json.NewEncoder(w).Encode(err)
		utils.HandleError(err)

	}

	// storing at the database
	err = db.CreatePromotion(ctx, promotion)
	utils.HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// Retriving the promotion with id
func FindPromotionById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	id := ps.ByName("id")
	promotion, err := db.FindPromotionById(ctx, id)
	utils.HandleError(err)
	if promotion.Id != "" {
		err = json.NewEncoder(w).Encode(promotion)
		utils.HandleError(err)
	} else {
		fmt.Fprintf(w, "Not found Object")
	}
}
