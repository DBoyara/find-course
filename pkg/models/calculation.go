package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calculation struct {
	Base
	UserID          uint   `json:"user_id"`
	UserCalculation string `json:"user_calc"`
}

type PercentageRate struct {
	Rate     float64   `json:"rate" bson:"double"`         // Процентная ставка
	DataFrom time.Time `json:"data_from" bson:"timestamp"` // Дата, с которой действует ставка
}

type EarlyRepayment struct {
	Type          string    `json:"type" bson:"string"`          // Тип: разовый, раз в месяц, раз в квартал...
	DataFrom      time.Time `json:"data_from" bson:"timestamp"`  // Дата платежа
	Amount        uint      `json:"new_amount" bson:"int64"`     // Сумма платежа(ей)
	Recalculation string    `json:"recalculation" bson:"string"` // уменьшить платеж-срок
}

type UserCalculation struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ClientName         string             `json:"client_name" bson:"string"`
	CreditAmount       uint               `json:"amount" bson:"int64"`                // Сумма кредита
	CreditTerm         uint16             `json:"term" bson:"int64"`                  // Срок кредита
	CreditTermType     string             `json:"term_type" bson:"string"`            // год-месяц
	PaymentType        string             `json:"payment_type" bson:"string"`         // Вид платежа: аннуитетный, дифферинцированный
	CostRealEstate     uint               `json:"cost" bson:"int64"`                  // Стоимость недвижимости
	InitialPayment     uint               `json:"iniatal_payment" bson:"int64"`       // Первоначальный платеж
	InitialPaymentType string             `json:"iniatal_payment_type" bson:"string"` // руб-%
	PercentageRates    PercentageRate     `json:"percentage_rate" gorm:"embedded"`    // Процентная ставка
	EarlyRepayments    EarlyRepayment     `json:"early_repayment" gorm:"embedded"`    // Досрочные погашения
}
