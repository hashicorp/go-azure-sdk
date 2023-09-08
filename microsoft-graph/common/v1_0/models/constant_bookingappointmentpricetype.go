package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointmentPriceType string

const (
	BookingAppointmentPriceTypecallUs      BookingAppointmentPriceType = "CallUs"
	BookingAppointmentPriceTypefixedPrice  BookingAppointmentPriceType = "FixedPrice"
	BookingAppointmentPriceTypefree        BookingAppointmentPriceType = "Free"
	BookingAppointmentPriceTypehourly      BookingAppointmentPriceType = "Hourly"
	BookingAppointmentPriceTypenotSet      BookingAppointmentPriceType = "NotSet"
	BookingAppointmentPriceTypepriceVaries BookingAppointmentPriceType = "PriceVaries"
	BookingAppointmentPriceTypestartingAt  BookingAppointmentPriceType = "StartingAt"
	BookingAppointmentPriceTypeundefined   BookingAppointmentPriceType = "Undefined"
)

func PossibleValuesForBookingAppointmentPriceType() []string {
	return []string{
		string(BookingAppointmentPriceTypecallUs),
		string(BookingAppointmentPriceTypefixedPrice),
		string(BookingAppointmentPriceTypefree),
		string(BookingAppointmentPriceTypehourly),
		string(BookingAppointmentPriceTypenotSet),
		string(BookingAppointmentPriceTypepriceVaries),
		string(BookingAppointmentPriceTypestartingAt),
		string(BookingAppointmentPriceTypeundefined),
	}
}

func parseBookingAppointmentPriceType(input string) (*BookingAppointmentPriceType, error) {
	vals := map[string]BookingAppointmentPriceType{
		"callus":      BookingAppointmentPriceTypecallUs,
		"fixedprice":  BookingAppointmentPriceTypefixedPrice,
		"free":        BookingAppointmentPriceTypefree,
		"hourly":      BookingAppointmentPriceTypehourly,
		"notset":      BookingAppointmentPriceTypenotSet,
		"pricevaries": BookingAppointmentPriceTypepriceVaries,
		"startingat":  BookingAppointmentPriceTypestartingAt,
		"undefined":   BookingAppointmentPriceTypeundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingAppointmentPriceType(input)
	return &out, nil
}
