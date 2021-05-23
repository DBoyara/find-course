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
	Rate     float64   `json:"rate" bson:"rate"`           // Процентная ставка
	DataFrom time.Time `json:"data_from" bson:"data_from"` // Дата, с которой действует ставка
}

type EarlyRepayment struct {
	Type          string    `json:"type" bson:"type"`                   // Тип: разовый, раз в месяц, раз в квартал...
	DataFrom      time.Time `json:"data_from" bson:"data_from"`         // Дата платежа
	Amount        uint      `json:"new_amount" bson:"new_amount"`       // Сумма платежа(ей)
	Recalculation string    `json:"recalculation" bson:"recalculation"` // уменьшить платеж-срок
}

type UserCalculation struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClientName         string             `json:"client_name" bson:"client_name"`
	CreditAmount       uint               `json:"amount" bson:"amount"`                                                // Сумма кредита
	CreditTerm         uint16             `json:"term" bson:"term"`                                                    // Срок кредита
	CreditTermType     string             `json:"term_type" bson:"term_type"`                                          // год-месяц
	PaymentType        string             `json:"payment_type" bson:"payment_type"`                                    // Вид платежа: аннуитетный, дифферинцированный
	CostRealEstate     uint               `json:"cost" bson:"cost"`                                                    // Стоимость недвижимости
	InitialPayment     uint               `json:"iniatal_payment" bson:"iniatal_payment"`                              // Первоначальный платеж
	InitialPaymentType string             `json:"iniatal_payment_type" bson:"iniatal_payment_type"`                    // руб-%
	PercentageRates    PercentageRate     `json:"percentage_rate,omitempty" bson:"percentage_rate,embedded,omitempty"` // Процентная ставка
	EarlyRepayments    EarlyRepayment     `json:"early_repayment,omitempty" bson:"early_repayment,embedded,omitempty"` // Досрочные погашения
}
