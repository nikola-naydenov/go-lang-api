package service

import (
	"encoding/json"
	"fmt"
	"goeureka/model"
	"io/ioutil"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	status := make(map[string]interface{})
	status["status"] = "OK"
	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	health := make(map[string]interface{})
	health["health"] = "OK"
	if err := json.NewEncoder(w).Encode(health); err != nil {
		panic(err)
	}
}

func GetSwaggerJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dir, _ := os.Getwd()
	data, _ := ioutil.ReadFile(dir + "/templates/swagger.json")
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}

func Fruits(w http.ResponseWriter, r *http.Request) {
	fruits := make([]model.Fruit, 0, 4)
	v1 := model.Fruit{ Name: "Guava", CalorieSources: "Protein 13%, Carbohydrates 75%, Fat 11%", CaloriesPerServing: 112}
	v2 := model.Fruit{ Name: "Passion Fruit", CalorieSources: "Protein 8%, Carbohydrates 86%, Fat 6%", CaloriesPerServing: 229}
	v3 := model.Fruit{ Name: "Raspberries", CalorieSources: "Protein 8%, Carbohydrates 88%, Fat 4%", CaloriesPerServing: 25}
	v4 := model.Fruit{ Name: "Orange", CalorieSources: "Protein 7%, Carbohydrates 91%, Fat 2%", CaloriesPerServing: 69}
	fruits = append(fruits, v1, v2, v3, v4)
	if len(fruits) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(fruits); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(model.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}