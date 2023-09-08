package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingServiceDefaultPriceType string

const (
	BookingServiceDefaultPriceTypecallUs      BookingServiceDefaultPriceType = "CallUs"
	BookingServiceDefaultPriceTypefixedPrice  BookingServiceDefaultPriceType = "FixedPrice"
	BookingServiceDefaultPriceTypefree        BookingServiceDefaultPriceType = "Free"
	BookingServiceDefaultPriceTypehourly      BookingServiceDefaultPriceType = "Hourly"
	BookingServiceDefaultPriceTypenotSet      BookingServiceDefaultPriceType = "NotSet"
	BookingServiceDefaultPriceTypepriceVaries BookingServiceDefaultPriceType = "PriceVaries"
	BookingServiceDefaultPriceTypestartingAt  BookingServiceDefaultPriceType = "StartingAt"
	BookingServiceDefaultPriceTypeundefined   BookingServiceDefaultPriceType = "Undefined"
)

func PossibleValuesForBookingServiceDefaultPriceType() []string {
	return []string{
		string(BookingServiceDefaultPriceTypecallUs),
		string(BookingServiceDefaultPriceTypefixedPrice),
		string(BookingServiceDefaultPriceTypefree),
		string(BookingServiceDefaultPriceTypehourly),
		string(BookingServiceDefaultPriceTypenotSet),
		string(BookingServiceDefaultPriceTypepriceVaries),
		string(BookingServiceDefaultPriceTypestartingAt),
		string(BookingServiceDefaultPriceTypeundefined),
	}
}

func parseBookingServiceDefaultPriceType(input string) (*BookingServiceDefaultPriceType, error) {
	vals := map[string]BookingServiceDefaultPriceType{
		"callus":      BookingServiceDefaultPriceTypecallUs,
		"fixedprice":  BookingServiceDefaultPriceTypefixedPrice,
		"free":        BookingServiceDefaultPriceTypefree,
		"hourly":      BookingServiceDefaultPriceTypehourly,
		"notset":      BookingServiceDefaultPriceTypenotSet,
		"pricevaries": BookingServiceDefaultPriceTypepriceVaries,
		"startingat":  BookingServiceDefaultPriceTypestartingAt,
		"undefined":   BookingServiceDefaultPriceTypeundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingServiceDefaultPriceType(input)
	return &out, nil
}
