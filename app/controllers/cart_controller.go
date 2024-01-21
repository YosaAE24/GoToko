package controllers

import (
	"fmt"
	"gotoko-postgres/app/models"
	"math"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetShoppingCartID(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionShoppingCart)
	if session.Values["cart-id"] == nil {
		session.Values["cart-id"] = uuid.New().String()
		_ = session.Save(r, w)
	}

	return fmt.Sprintf("%v", session.Values["cart-id"])
}

func GetShoppingCart(db *gorm.DB, cartID string) (*models.Cart, error) {
	var cart models.Cart

	existCart, err := cart.GetCart(db, cartID)
	if err != nil {
		existCart, _ = cart.CreateCart(db, cartID)
	}

	_, _ = existCart.CalculateCart(db, cartID)

	updatedCart, _ := cart.GetCart(db, cartID)

	totalWeight := 0
	productModel := models.Product{}
	for _, cartItem := range updatedCart.CartItems {
		product, _ := productModel.FindByID(db, cartItem.ProductID)

		productWeight, _ := product.Weight.Float64()
		ceilWeight := math.Ceil(productWeight)

		itemWeight := cartItem.Qty * int(ceilWeight)

		totalWeight += itemWeight
	}

	updatedCart.TotalWeight = totalWeight

	return updatedCart, nil
}

func (server *Server) GetCart(w http.ResponseWriter, r *http.Request){
	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)

	fmt.Println("cart id ====>", cart.ID)
	fmt.Println("cart id ====>", cart.CartItems)
}

func (server *Server) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	productID := r.FormValue("product_id")
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, productID)

	if err != nil {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	if qty > product.Stock {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)
	_, err = cart.AddItem(server.DB, models.CartItem{
		ProductID: productID,
		Qty:       qty,
	})
	if err != nil {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	// SetFlash(w, r, "success", "Item berhasil ditambahkan")
	http.Redirect(w, r, "/carts", http.StatusSeeOther)
}