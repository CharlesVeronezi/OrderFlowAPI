package api

import "time"

type DefaultResponseID struct {
	ID      string `json:"id"`
	Message string `json:"message"`
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
	PrStock       int     `json:"pr_stock"`
	PrPrice       float64 `json:"pr_price"`
	PrVbActive    bool    `json:"pr_vbactive"`
}

type Order_Product struct {
	OrpOrderid    string  `json:"orp_orderID"`
	OrpProductid  string  `json:"orp_productID"`
	OrpQuantidad  int     `json:"orp_quantidad"`
	OrpTotalprice float64 `json:"orp_totalprice"`
}

type User struct {
	UsUserId    string `json:"us_userID"`
	UsFirstname string `json:"us_firstname"`
	UsLastname  string `json:"us_lastname"`
	UsEmail     string `json:"us_email"`
	UsVbActive  bool   `json:"us_vbactive"`
}

type Order struct {
	OrOrderid         string          `json:"or_orderid"`
	OrUserid          string          `json:"or_userid"`
	OrProducts        []Order_Product `json:"or_products"`
	OrTotalamount     float64         `json:"or_total_amount"`
	OrOrderstatus     string          `json:"or_order_status"`
	OrPaymentmethod   string          `json:"or_payment_method"`
	OrShippingaddress string          `json:"or_shipping_address"`
	OrCreatedat       time.Time       `json:"or_created_at"`
	OrUpdatedat       time.Time       `json:"or_updated_at"`
}
