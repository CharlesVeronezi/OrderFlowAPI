package orders

import "time"

type Order struct {
	OrOrderid         string    `json:"or_orderID"`
	OrUserid          string    `json:"or_userID"`
	OrProducts        []Product `json:"or_products"`
	OrTotalamount     float64   `json:"or_total_amount"`
	OrOrderstatus     string    `json:"or_order_status"`
	OrPaymentmethod   string    `json:"or_payment_method"`
	OrShippingaddress Address   `json:"or_shipping_address"`
	OrCreatedat       time.Time `json:"or_created_at"`
	OrUpdatedat       time.Time `json:"or_updated_at"`
}

type Address struct {
	AdAddresid string `json:"ad_addressID"`
	AdStreet   string `json:"ad_street"`
	AdCity     string `json:"ad_city"`
	AdState    string `json:"ad_state"`
	AdZip      string `json:"ad_zip"`
	AdCountry  string `json:"ad_country"`
}

type Product struct {
	PrProductID   string  `json:"pr_productID"`
	PrDescription string  `json:"pr_description"`
	PrPrice       float64 `json:"pr_price"`
	PrVbActive    bool    `json:"pr_vb_active"`
}

type User struct {
	UsUserId    string `json:"us_userID"`
	UsFirstname string `json:"us_firstname"`
	UsLastname  string `json:"us_lastname"`
	UsEmail     string `json:"us_email"`
	UsVbActive  bool   `json:"us_vb_active"`
}
