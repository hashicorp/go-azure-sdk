package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingReminderRecipients string

const (
	BookingReminderRecipientsallAttendees BookingReminderRecipients = "AllAttendees"
	BookingReminderRecipientscustomer     BookingReminderRecipients = "Customer"
	BookingReminderRecipientsstaff        BookingReminderRecipients = "Staff"
)

func PossibleValuesForBookingReminderRecipients() []string {
	return []string{
		string(BookingReminderRecipientsallAttendees),
		string(BookingReminderRecipientscustomer),
		string(BookingReminderRecipientsstaff),
	}
}

func parseBookingReminderRecipients(input string) (*BookingReminderRecipients, error) {
	vals := map[string]BookingReminderRecipients{
		"allattendees": BookingReminderRecipientsallAttendees,
		"customer":     BookingReminderRecipientscustomer,
		"staff":        BookingReminderRecipientsstaff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingReminderRecipients(input)
	return &out, nil
}
