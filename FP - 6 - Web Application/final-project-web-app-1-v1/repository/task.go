package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	// get task data by user id using gorm with context
	var tasks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Select("*").Where("user_id = ?", id).Scan(&tasks).Error
	if err != nil {
		return nil, err
	} else if len(tasks) == 0 {
		return []entity.Task{}, nil
	}
	return tasks, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	// store task using gorm with context
	err = r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return 0, err
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	// get task data by id using gorm with context
	var task entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Select("*").Where("id = ?", id).Scan(&task).Error
	if err != nil {
		return entity.Task{}, err
	} else if task.ID == 0 {
		return entity.Task{}, nil
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	// get task data by category id using gorm with context
	var tasks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Select("*").Where("category_id = ?", catId).Scan(&tasks).Error
	if err != nil {
		return nil, err
	} else if len(tasks) == 0 {
		return []entity.Task{}, nil
	}
	return tasks, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	// update task using gorm with context
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	// delete task using gorm with context
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Task{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
