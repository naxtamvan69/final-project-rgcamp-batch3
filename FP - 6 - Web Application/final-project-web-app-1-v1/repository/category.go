package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	// get category data by user id using gorm with context
	var categories []entity.Category
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Select("*").Where("user_id = ?", id).Scan(&categories).Error
	if err != nil {
		return nil, err
	} else if len(categories) == 0 {
		return []entity.Category{}, nil
	}
	return categories, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	// store category data using gorm with context
	err = r.db.WithContext(ctx).Create(&category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	// store many category data using gorm
	err := r.db.WithContext(ctx).Create(&categories).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	// get category data by id using gorm with context
	var category entity.Category
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Select("*").Where("id = ?", id).Scan(&category).Error
	if err != nil {
		return entity.Category{}, err
	} else if category.ID == 0 {
		return entity.Category{}, nil
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	// update category data using gorm with context
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	// delete category data by id using gorm with context
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
