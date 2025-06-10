package main

import (
	"fmt"
	"go-breeders/models"
	"go-breeders/pets"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	fmt.Println(page)
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfmonth(w http.ResponseWriter, r *http.Request) {
	// get the breed
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	// get the dog of the month from the database
	dom, _ := app.App.Models.Dog.GetDogOfMonthById(1)

	layout := "2006-02-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	// create the dog and decorate it
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			Id:               1,
			DogName:          "Sam",
			BreedId:          breed.Id,
			Color:            "Black & Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: 0,
			Description:      "Sam is a very good boy",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}
	// serve the webpage
	data := make(map[string]any)
	data["dog"] = dog
	app.render(w, "dog-of-month.page.gohtml", &templateData{
		Data: data,
	})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("dog")
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("cat")
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().SetSpecies("dog").SetBreed("mixed breed").
		SetWeight(15).SetDescription("a mixed breed of unknown origin. probably has some german shepherd heritage").
		SetColor("black and white").SetAge(3).SetAgeEstimated(true).Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	// create a cat using the builder pattern
	p, err := pets.NewPetBuilder().SetSpecies("cat").SetBreed("felis silverstris catus").
		SetWeight(4).SetDescription("a beautiful house cat").
		SetColor("black and white").SetAge(1).SetAgeEstimated(true).Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// setup toolbox
	var t toolbox.Tools
	// get species from url itself
	species := chi.URLParam(r, "species")
	// get breed from the URL
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)
	fmt.Println("species", species)
	fmt.Println("breed", breed)
	// create a pet from abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// write the result as JSON
	_ = t.WriteJSON(w, http.StatusOK, pet)
}
