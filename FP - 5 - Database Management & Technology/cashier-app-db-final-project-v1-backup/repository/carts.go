package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	var joinCart []model.JoinCart
	err := c.db.Table("carts").Select("carts.id, carts.product_id, products.name, carts.quantity, carts.total_price").Joins("INNER JOIN products ON carts.product_id = products.id").Scan(&joinCart).Error
	if err != nil {
		return nil, err
	}
	return joinCart, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		var searchCart model.Cart
		err := tx.Model(&model.Cart{}).Where("product_id = ?", product.ID).Take(&searchCart).Error
		if err != nil {
			cart := model.Cart{
				ProductID:  product.ID,
				Quantity:   1,
				TotalPrice: product.Price - (product.Price * (product.Discount / 100)),
			}
			err = tx.Create(&cart).Error
			if err != nil {
				return err
			}

			err = tx.Model(&model.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
			if err != nil {
				return err
			}
		}

		err = tx.Model(&model.Cart{}).Where("id = ?", searchCart.ID).Update("quantity", searchCart.Quantity+1).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Cart{}).Where("id = ?", searchCart.ID).Update("total_price", searchCart.TotalPrice+searchCart.TotalPrice).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
		if err != nil {
			return err
		}
		return nil
	})
	// TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {

	return c.db.Transaction(func(tx *gorm.DB) error {
		var cart model.Cart
		var product model.Product

		if err := tx.Model(&model.Cart{}).Where("id = ?", id).Take(&cart).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Product{}).Where("id = ?", productID).Take(&product).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Delete(&model.Cart{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Product{}).Where("id = ?", productID).Update("stock", int(cart.Quantity)+product.Stock).Error; err != nil {
			return err
		}

		return nil
	}) // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	err := c.db.Model(&model.Cart{}).Where("id = ?", id).Updates(cart).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
