// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package pgstore

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Address struct {
	AdAddressid uuid.UUID
	AdStreet    string
	AdCity      string
	AdState     string
	AdZip       string
	AdCountry   string
}

type Order struct {
	OrOrderid         uuid.UUID
	OrUserid          pgtype.UUID
	OrTotalamount     pgtype.Numeric
	OrOrderstatus     string
	OrPaymentmethod   string
	OrShippingaddress pgtype.UUID
	OrCreatedat       pgtype.Timestamp
	OrUpdatedat       pgtype.Timestamp
}

type OrdersProduct struct {
	OrpOrderid    pgtype.UUID
	OrpProductid  pgtype.UUID
	OrpQuantidad  int32
	OrpTotalprice pgtype.Numeric
}

type Product struct {
	PrProductid   uuid.UUID
	PrDescription string
	PrStock       int32
	PrPrice       pgtype.Numeric
	PrVbactive    bool
}

type User struct {
	UsUserid    uuid.UUID
	UsFirstname string
	UsLastname  pgtype.Text
	UsEmail     pgtype.Text
	UsVbactive  bool
}
