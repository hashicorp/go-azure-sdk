package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingServiceDefaultPriceType string

const (
	BookingServiceDefaultPriceType_CallUs      BookingServiceDefaultPriceType = "callUs"
	BookingServiceDefaultPriceType_FixedPrice  BookingServiceDefaultPriceType = "fixedPrice"
	BookingServiceDefaultPriceType_Free        BookingServiceDefaultPriceType = "free"
	BookingServiceDefaultPriceType_Hourly      BookingServiceDefaultPriceType = "hourly"
	BookingServiceDefaultPriceType_NotSet      BookingServiceDefaultPriceType = "notSet"
	BookingServiceDefaultPriceType_PriceVaries BookingServiceDefaultPriceType = "priceVaries"
	BookingServiceDefaultPriceType_StartingAt  BookingServiceDefaultPriceType = "startingAt"
	BookingServiceDefaultPriceType_Undefined   BookingServiceDefaultPriceType = "undefined"
)

func PossibleValuesForBookingServiceDefaultPriceType() []string {
	return []string{
		string(BookingServiceDefaultPriceType_CallUs),
		string(BookingServiceDefaultPriceType_FixedPrice),
		string(BookingServiceDefaultPriceType_Free),
		string(BookingServiceDefaultPriceType_Hourly),
		string(BookingServiceDefaultPriceType_NotSet),
		string(BookingServiceDefaultPriceType_PriceVaries),
		string(BookingServiceDefaultPriceType_StartingAt),
		string(BookingServiceDefaultPriceType_Undefined),
	}
}

func (s *BookingServiceDefaultPriceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingServiceDefaultPriceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingServiceDefaultPriceType(input string) (*BookingServiceDefaultPriceType, error) {
	vals := map[string]BookingServiceDefaultPriceType{
		"callus":      BookingServiceDefaultPriceType_CallUs,
		"fixedprice":  BookingServiceDefaultPriceType_FixedPrice,
		"free":        BookingServiceDefaultPriceType_Free,
		"hourly":      BookingServiceDefaultPriceType_Hourly,
		"notset":      BookingServiceDefaultPriceType_NotSet,
		"pricevaries": BookingServiceDefaultPriceType_PriceVaries,
		"startingat":  BookingServiceDefaultPriceType_StartingAt,
		"undefined":   BookingServiceDefaultPriceType_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingServiceDefaultPriceType(input)
	return &out, nil
}
