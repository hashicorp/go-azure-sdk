package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointmentInvoiceStatus string

const (
	BookingAppointmentInvoiceStatus_Canceled   BookingAppointmentInvoiceStatus = "canceled"
	BookingAppointmentInvoiceStatus_Corrective BookingAppointmentInvoiceStatus = "corrective"
	BookingAppointmentInvoiceStatus_Draft      BookingAppointmentInvoiceStatus = "draft"
	BookingAppointmentInvoiceStatus_Open       BookingAppointmentInvoiceStatus = "open"
	BookingAppointmentInvoiceStatus_Paid       BookingAppointmentInvoiceStatus = "paid"
	BookingAppointmentInvoiceStatus_Reviewing  BookingAppointmentInvoiceStatus = "reviewing"
)

func PossibleValuesForBookingAppointmentInvoiceStatus() []string {
	return []string{
		string(BookingAppointmentInvoiceStatus_Canceled),
		string(BookingAppointmentInvoiceStatus_Corrective),
		string(BookingAppointmentInvoiceStatus_Draft),
		string(BookingAppointmentInvoiceStatus_Open),
		string(BookingAppointmentInvoiceStatus_Paid),
		string(BookingAppointmentInvoiceStatus_Reviewing),
	}
}

func (s *BookingAppointmentInvoiceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingAppointmentInvoiceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingAppointmentInvoiceStatus(input string) (*BookingAppointmentInvoiceStatus, error) {
	vals := map[string]BookingAppointmentInvoiceStatus{
		"canceled":   BookingAppointmentInvoiceStatus_Canceled,
		"corrective": BookingAppointmentInvoiceStatus_Corrective,
		"draft":      BookingAppointmentInvoiceStatus_Draft,
		"open":       BookingAppointmentInvoiceStatus_Open,
		"paid":       BookingAppointmentInvoiceStatus_Paid,
		"reviewing":  BookingAppointmentInvoiceStatus_Reviewing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingAppointmentInvoiceStatus(input)
	return &out, nil
}
