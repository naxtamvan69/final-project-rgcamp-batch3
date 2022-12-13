package repository

import (
	"a21hc3NpZ25tZW50/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	err := p.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	var result []model.Product
	err := p.db.Raw("SELECT * FROM products WHERE deleted_at IS NULL").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	err := p.db.Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	err := p.db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
