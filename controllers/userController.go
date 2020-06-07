package controllers

import (
	"encoding/json"
	"gorillatest/model/frombackend"

	//"gorillatest/dal"
	//"gorillatest/models"
	bzuser "gorillatest/bizrules/user"
	"gorillatest/model/tobackend"
	"io/ioutil"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var userRegister tobackend.UserRegister

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(body, &userRegister)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := bzuser.Register(userRegister)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin tobackend.UserLogin

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := bzuser.Login(userLogin)
	if err != nil {
		var errors frombackend.Errors
		errors.Message = append(errors.Message, err.Error())
		output, _ := json.Marshal(errors)
		http.Error(w, string(output), http.StatusUnprocessableEntity)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		var errors frombackend.Errors
		errors.Message = append(errors.Message, err.Error())
		output, _ = json.Marshal(errors)
		http.Error(w, string(output), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	stoken := r.Header.Get("authorization")
	if len(stoken) == 0 {
		http.Error(w, "Не передан токен authorization", http.StatusUnauthorized)
		return
	}

	user, err := bzuser.GetUser(stoken)
	if err != nil {
		var errors frombackend.Errors
		errors.Message = append(errors.Message, err.Error())
		output, _ := json.Marshal(errors)
		http.Error(w, string(output), http.StatusUnprocessableEntity)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		var errors frombackend.Errors
		errors.Message = append(errors.Message, err.Error())
		output, _ = json.Marshal(errors)
		http.Error(w, string(output), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
