package main

import (
	"net/http"

	"github.com/AutomationMK/ecommerce/internal/models"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", &templateData{}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}

// PaymentSucceded parses all cardholder and payment post data and renders a succeeded
// template page
func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read posted data
	cardHolder := r.Form.Get("cardholder_name")
	cardholderEmail := r.Form.Get("cardholder_email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]any)
	data["cardholder"] = cardHolder
	data["cardholder_email"] = cardholderEmail
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err = app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// ChargeOnce renders the buy-once template page
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	widget := models.Widget{
		ID:             1,
		Name:           "Custom Widget",
		Description:    "A very nice widget",
		InventoryLevel: 10,
		Price:          1000,
	}

	data := make(map[string]any)
	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}
