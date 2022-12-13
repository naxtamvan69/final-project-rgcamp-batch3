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

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
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
		Get task id from parameter query url task_id
		if task id is empty then call taskService.GetTask
		if task id not empty call taskService.GetTaskById
		if category id not empty then call taskService.GetTaskByCategoryId
	*/
	taskId := r.URL.Query().Get("task_id")
	categoryId := r.URL.Query().Get("category_id")
	if categoryId != "" {
		/*
			Call taskService.GetTaskByCategoryID with parameter context and user id
			if in this service return error, then return 500 code with message error internal server
		*/
		categoryIdInt, _ := strconv.Atoi(categoryId)
		tasks, err := t.taskService.GetTaskByCategoryID(r.Context(), categoryIdInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		/*
			return 200 code with task data by user id
		*/
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	} else if taskId == "" {
		/*
			Call taskService.GetTask with parameter context and user id
			if in this service return error, then return 500 code with message error internal server
		*/
		userIdInt, _ := strconv.Atoi(taskId)
		tasks, err := t.taskService.GetTasks(r.Context(), userIdInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}

		/*
			return 200 code with task data by user id
		*/
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	} else if taskId != "" {
		/*
			Call taskService.GetTaskById with parameter context and task id
			if in this service return error, then return 500 code with message error internal server
		*/
		taskIdInt, _ := strconv.Atoi(taskId)
		task, err := t.taskService.GetTaskByID(r.Context(), taskIdInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}

		/*
			return 200 code with task data by user id
		*/
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	/*
		Check if task request is empty then return 400 code with message invalid task request
	*/
	if (entity.TaskRequest{} == task) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
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
		Call taskService.StoreTask with parameter context and task
		if this service return empty then return 500 code with message error internal server
	*/
	userIdInt, _ := strconv.Atoi(userId)
	newTask := entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      userIdInt,
	}

	newTask, err = t.taskService.StoreTask(r.Context(), &newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 201 code with user id, task id and message success create new task
	*/
	resp := map[string]interface{}{
		"user_id": newTask.UserID,
		"task_id": newTask.ID,
		"message": "success create new task",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
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
		Get task id from parameter query url task_id
		if task id is empty then return 400 code with message invalid task id
	*/
	taskId := r.URL.Query().Get("task_id")
	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task id"))
		return
	}

	/*
		Call serviceTask.deleteTask with parameter context and task id
		if this service return error then return 500 code with message error internal server
	*/
	taskIdInt, _ := strconv.Atoi(taskId)
	err := t.taskService.DeleteTask(r.Context(), taskIdInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 200 code with user id, task id and message success delete task
	*/
	resp := map[string]interface{}{
		"user_id": userId,
		"task_id": taskIdInt,
		"message": "success delete task",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
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
		Get task id from url query task_id
		if task id is empty return 400 code with message invalid task id
	*/
	taskId := r.URL.Query().Get("task_id")
	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task id"))
		return
	}

	/*
		Call taskService.UpdateTask with parameter context and task
		if this service return empty then return 500 code with message error internal server
	*/
	userIdInt, _ := strconv.Atoi(userId)
	newTask := entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		UserID:      userIdInt,
	}

	_, err = t.taskService.UpdateTask(r.Context(), &newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	/*
		return 200 code with user id, task id and message success update task
	*/
	resp := map[string]interface{}{
		"user_id": userId,
		"task_id": newTask.ID,
		"message": "success update task",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
