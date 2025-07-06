package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// ProductStatusType defines the allowed status values
type ProductStatusType string

const (
	Publish ProductStatusType = "Publish"
	Draft   ProductStatusType = "Draft"
)

type Product struct {
	BaseModel
	Name          string            `gorm:"type:string;size:100;not null"`
	Description   string            `gorm:"type:string;size:255;null"`
	Price         float64           `gorm:"type:float;not null"`
	Stock         int               `gorm:"type:int;not null"`
	CategoryId    int               `gorm:"type:int;not null"`
	Status        ProductStatusType `gorm:"type:string;size:10;not null"`
	Slug          string            `gorm:"type:string;size:100;not null"`
	Category      Category          `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	AverageRating float64           `gorm:"type:float;default:0"`
	CountViews    int               `gorm:"type:int;default:0"`
	Images        []ProductImage    `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Reviews       []ProductReview   `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
type Category struct {
	BaseModel
	Name        string    `gorm:"type:string;size:100;not null"`
	Description string    `gorm:"type:string;size:255;null"`
	Products    []Product `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
type ProductImage struct {
	BaseModel
	ProductId int     `gorm:"type:int;not null"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Image     File    `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ImageId   int     `gorm:"uniqueIndex:idx_ProductId_ProductImageId;not null"`
	IsMain    bool    `gorm:"type:boolean;default:false;uniqueIndex:idx_ProductId_IsMain"`
}

type ProductReview struct {
	BaseModel
	ProductId int     `gorm:"type:int;not null"`
	Rating    int     `gorm:"type:int;not null"`
	Comment   string  `gorm:"type:string;size:255;null"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId    int     `gorm:"type:int;not null"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

type File struct {
	BaseModel
	Name        string `gorm:"size:100;type:string;not null"`
	Directory   string `gorm:"size:100;type:string;not null"`
	Description string `gorm:"size:500;type:string;not null"`
	MimeType    string `gorm:"size:20;type:string;not null"`
}

// BeforeSave is a GORM hook that runs before saving a product
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	if strings.Contains(p.Slug, " ") {
		return errors.New("slug cannot contain spaces")
	}
	return nil
}

// BeforeSave is a GORM hook that ensures only one main image exists per product
func (p *ProductImage) BeforeSave(tx *gorm.DB) (err error) {
	// if the image is marked as main, ensure no other images are marked as main for the same product
	if p.IsMain {
		err := tx.Model(&ProductImage{}).
			Where("product_id = ? AND is_main = ? ", p.ProductId, true).
			Update("is_main", false).Error
		if err != nil {
			return err
		}
	}

	return nil
}
