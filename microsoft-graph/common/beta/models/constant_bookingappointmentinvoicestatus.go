package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointmentInvoiceStatus string

const (
	BookingAppointmentInvoiceStatuscanceled   BookingAppointmentInvoiceStatus = "Canceled"
	BookingAppointmentInvoiceStatuscorrective BookingAppointmentInvoiceStatus = "Corrective"
	BookingAppointmentInvoiceStatusdraft      BookingAppointmentInvoiceStatus = "Draft"
	BookingAppointmentInvoiceStatusopen       BookingAppointmentInvoiceStatus = "Open"
	BookingAppointmentInvoiceStatuspaid       BookingAppointmentInvoiceStatus = "Paid"
	BookingAppointmentInvoiceStatusreviewing  BookingAppointmentInvoiceStatus = "Reviewing"
)

func PossibleValuesForBookingAppointmentInvoiceStatus() []string {
	return []string{
		string(BookingAppointmentInvoiceStatuscanceled),
		string(BookingAppointmentInvoiceStatuscorrective),
		string(BookingAppointmentInvoiceStatusdraft),
		string(BookingAppointmentInvoiceStatusopen),
		string(BookingAppointmentInvoiceStatuspaid),
		string(BookingAppointmentInvoiceStatusreviewing),
	}
}

func parseBookingAppointmentInvoiceStatus(input string) (*BookingAppointmentInvoiceStatus, error) {
	vals := map[string]BookingAppointmentInvoiceStatus{
		"canceled":   BookingAppointmentInvoiceStatuscanceled,
		"corrective": BookingAppointmentInvoiceStatuscorrective,
		"draft":      BookingAppointmentInvoiceStatusdraft,
		"open":       BookingAppointmentInvoiceStatusopen,
		"paid":       BookingAppointmentInvoiceStatuspaid,
		"reviewing":  BookingAppointmentInvoiceStatusreviewing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingAppointmentInvoiceStatus(input)
	return &out, nil
}
