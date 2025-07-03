package models

// Cart represents a shopping cart in the system.
type Cart struct {
	BaseModel
	UserId     int        `gorm:"type:int;not null;uniqueIndex"`
	User       User       `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CartItems  []CartItem `gorm:"foreignKey:CartId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	TotalPrice float64    `gorm:"type:float;default:0"`
	TotalItems int        `gorm:"type:int;default:0"`
}

// CartItem represents an item in the shopping cart.
type CartItem struct {
	BaseModel
	ProductId int     `gorm:"type:int;not null"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CartId    int     `gorm:"type:int;not null"`
	Cart      Cart    `gorm:"foreignKey:CartId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Quantity  int     `gorm:"type:int;default:1"`
	UnitPrice float64 `gorm:"type:float;default:0"`
}

// Payment represents a payment in the system.
type Payment struct {
	BaseModel
	Amount      float64 `gorm:"type:float;not null"`
	Status      string  `gorm:"type:string;size:50;not null"`
	AuthorityId string  `gorm:"type:string;size:100;not null,uniqueIndex"`
	RefId       int     `gorm:"type:int;not null"`
	UserId      int     `gorm:"type:int;not null"`
	User        User    `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	OrderId     int     `gorm:"type:int;not null"`
	Order       Order   `gorm:"foreignKey:OrderId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

// Order represents an order in the system.
type Order struct {
	BaseModel
	UserId     int         `gorm:"type:int;not null"`
	User       User        `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	TotalPrice float64     `gorm:"type:float;default:0"`
	TotalItems int         `gorm:"type:int;default:0"`
	Status     string      `gorm:"type:string;size:50;not null"`
	Address    string      `gorm:"type:string;size:255;not null"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	BaseModel
	ProductId int     `gorm:"type:int;not null"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	OrderId   int     `gorm:"type:int;not null"`
	Order     Order   `gorm:"foreignKey:OrderId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Quantity  int     `gorm:"type:int;default:1"`
	UnitPrice float64 `gorm:"type:float;default:0"`
}
