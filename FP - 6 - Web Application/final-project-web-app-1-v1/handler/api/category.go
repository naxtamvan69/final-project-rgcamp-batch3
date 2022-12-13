package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	/*
		Get user id from context
		if user id is empty return 400 code with message invalid user id
	*/
	userId := fmt.Sprintf("%s", r.Context().Value("id"))
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	/*
		Call categoryService.GetCategory with parameter context and user id
		if in this service return error, then return 500 code with message error internal server
	*/
	userIdInt, _ := strconv.Atoi(userId)
	categories, err := c.categoryService.GetCategories(r.Context(), userIdInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 200 code with category data by user id
	*/
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	/*
		Check if category request is empty then return 400 code with message invalid category request
	*/
	if (entity.CategoryRequest{} == category) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	/*
		Get user id from context
		if user id is empty return 400 code with message invalid user id
	*/
	userId := fmt.Sprintf("%s", r.Context().Value("id"))
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	/*
		Call categoryService.StoreCategory with parameter context and category
		if in this service return error, then return 500 code with message error internal server
	*/
	userIdInt, _ := strconv.Atoi(userId)
	newCategory := entity.Category{
		Type:   category.Type,
		UserID: userIdInt,
	}
	newCategory, err = c.categoryService.StoreCategory(r.Context(), &newCategory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 201 code with user id, category id and message success create new category
	*/
	resp := map[string]interface{}{
		"user_id":     newCategory.UserID,
		"category_id": newCategory.ID,
		"message":     "success create new category",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	/*
		Get user id from context
		if user id is empty return 400 code with message invalid user id
	*/
	userId := fmt.Sprintf("%s", r.Context().Value("id"))
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	/*
		Get category id from query parameter url
		if category id is empty return 400 code with message invalid category id
	*/
	categoryId := r.URL.Query().Get("category_id")
	if categoryId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid parameter id"))
		return
	}

	/*
		Call categoryService.DeleteCategory with parameter context and category id
		if this service return error then return 500 code with message error internal server
	*/
	categoryIdInt, _ := strconv.Atoi(categoryId)
	err := c.categoryService.DeleteCategory(r.Context(), categoryIdInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 200 code with user id, category id and message success delete category
	*/
	resp := map[string]interface{}{
		"user_id":     userId,
		"category_id": categoryIdInt,
		"message":     "success delete category",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
