package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointmentPriceType string

const (
	BookingAppointmentPriceType_CallUs      BookingAppointmentPriceType = "callUs"
	BookingAppointmentPriceType_FixedPrice  BookingAppointmentPriceType = "fixedPrice"
	BookingAppointmentPriceType_Free        BookingAppointmentPriceType = "free"
	BookingAppointmentPriceType_Hourly      BookingAppointmentPriceType = "hourly"
	BookingAppointmentPriceType_NotSet      BookingAppointmentPriceType = "notSet"
	BookingAppointmentPriceType_PriceVaries BookingAppointmentPriceType = "priceVaries"
	BookingAppointmentPriceType_StartingAt  BookingAppointmentPriceType = "startingAt"
	BookingAppointmentPriceType_Undefined   BookingAppointmentPriceType = "undefined"
)

func PossibleValuesForBookingAppointmentPriceType() []string {
	return []string{
		string(BookingAppointmentPriceType_CallUs),
		string(BookingAppointmentPriceType_FixedPrice),
		string(BookingAppointmentPriceType_Free),
		string(BookingAppointmentPriceType_Hourly),
		string(BookingAppointmentPriceType_NotSet),
		string(BookingAppointmentPriceType_PriceVaries),
		string(BookingAppointmentPriceType_StartingAt),
		string(BookingAppointmentPriceType_Undefined),
	}
}

func (s *BookingAppointmentPriceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingAppointmentPriceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingAppointmentPriceType(input string) (*BookingAppointmentPriceType, error) {
	vals := map[string]BookingAppointmentPriceType{
		"callus":      BookingAppointmentPriceType_CallUs,
		"fixedprice":  BookingAppointmentPriceType_FixedPrice,
		"free":        BookingAppointmentPriceType_Free,
		"hourly":      BookingAppointmentPriceType_Hourly,
		"notset":      BookingAppointmentPriceType_NotSet,
		"pricevaries": BookingAppointmentPriceType_PriceVaries,
		"startingat":  BookingAppointmentPriceType_StartingAt,
		"undefined":   BookingAppointmentPriceType_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingAppointmentPriceType(input)
	return &out, nil
}
