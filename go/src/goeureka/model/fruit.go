package model

type Fruit struct {
        Name                  string    `json:"name"`
        CaloriesPerServing    int       `json:"caloriesPerServing"`
        CalorieSources        string    `json:"calorieSources"`
}
