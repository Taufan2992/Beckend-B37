package handlers

import (
	"encoding/json"
	productdto "golang/dto/product"
	dto "golang/dto/result"
	topingdto "golang/dto/toping"
	"golang/models"
	"golang/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerToping struct {
	TopingRepository repositories.TopingRepository
}

func HandlerToping(TopingRepository repositories.TopingRepository) *handlerToping {
	return &handlerToping{TopingRepository}
}

func (h *handlerToping) FindTopings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	topings, err := h.TopingRepository.FindTopings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// looping untuk melakukan tampilan gambar pada postman
	for i, p := range topings {
		topings[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: topings}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) GetToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var toping models.Toping
	toping, err := h.TopingRepository.GetToping(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// menampilkan gambar
	toping.Image = path_file + toping.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: toping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) CreateToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := productdto.ProductRequest{
		Title: r.FormValue("title"),
		Price: price,
	}

	// request := new(topingdto.TopingRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toping := models.Toping{
		Title: request.Title,
		Price: request.Price,
		Image: filename,
	}

	// err := mysql.DB.Create(&toping).Error
	toping, err = h.TopingRepository.CreateToping(toping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toping, _ = h.TopingRepository.GetToping(toping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: toping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) UpdateToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(topingdto.UpdateTopingRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	toping, err := h.TopingRepository.GetToping(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		toping.Title = request.Title
	}

	if request.Price != 0 {
		toping.Price = request.Price
	}

	if request.Image != "" {
		toping.Image = request.Image
	}

	data, err := h.TopingRepository.UpdateToping(toping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) DeleteToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	toping, err := h.TopingRepository.GetToping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TopingRepository.DeleteToping(toping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// func convertResponseToping(u models.Toping) models.TopingResponse {
// 	return models.TopingResponse{
// 		ID:    u.ID,
// 		Name:  u.Name,
// 		Desc:  u.Desc,
// 		Price: u.Price,
// 	}
// }
