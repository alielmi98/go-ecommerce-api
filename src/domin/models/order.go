package models

// Cart represents a shopping cart in the system.
type Cart struct {
	BaseModel
	UserId     int       `gorm:"type:int;not null"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Products   []Product `gorm:"many2many:cart_products;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	TotalPrice float64   `gorm:"type:float;default:0"`
	TotalItems int       `gorm:"type:int;default:0"`
}

// Payment represents a payment in the system.
type Payment struct {
	BaseModel
	Amount    float64 `gorm:"type:float;not null"`
	Status    string  `gorm:"type:string;size:50;not null"`
	PaymentId string  `gorm:"type:string;size:100;not null"`
	UserId    int     `gorm:"type:int;not null"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	OrderId   int     `gorm:"type:int;not null"`
	Order     Cart    `gorm:"foreignKey:OrderId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

// Order represents an order in the system.
type Order struct {
	BaseModel
	UserId     int       `gorm:"type:int;not null"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Products   []Product `gorm:"many2many:order_products;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	TotalPrice float64   `gorm:"type:float;default:0"`
	TotalItems int       `gorm:"type:int;default:0"`
	Status     string    `gorm:"type:string;size:50;not null"`
	Address    string    `gorm:"type:string;size:255;not null"`
	PaymentId  int       `gorm:"type:int;not null"`
	Payment    Payment   `gorm:"foreignKey:PaymentId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
