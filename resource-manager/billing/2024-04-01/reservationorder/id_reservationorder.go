package reservationorder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReservationOrderId{})
}

var _ resourceids.ResourceId = &ReservationOrderId{}

// ReservationOrderId is a struct representing the Resource ID for a Reservation Order
type ReservationOrderId struct {
	BillingAccountName string
	ReservationOrderId string
}

// NewReservationOrderID returns a new ReservationOrderId struct
func NewReservationOrderID(billingAccountName string, reservationOrderId string) ReservationOrderId {
	return ReservationOrderId{
		BillingAccountName: billingAccountName,
		ReservationOrderId: reservationOrderId,
	}
}

// ParseReservationOrderID parses 'input' into a ReservationOrderId
func ParseReservationOrderID(input string) (*ReservationOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReservationOrderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReservationOrderId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReservationOrderIDInsensitively parses 'input' case-insensitively into a ReservationOrderId
// note: this method should only be used for API response data and not user input
func ParseReservationOrderIDInsensitively(input string) (*ReservationOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReservationOrderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReservationOrderId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReservationOrderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.ReservationOrderId, ok = input.Parsed["reservationOrderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "reservationOrderId", input)
	}

	return nil
}

// ValidateReservationOrderID checks that 'input' can be parsed as a Reservation Order ID
func ValidateReservationOrderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReservationOrderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Reservation Order ID
func (id ReservationOrderId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/reservationOrders/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.ReservationOrderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Reservation Order ID
func (id ReservationOrderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticReservationOrders", "reservationOrders", "reservationOrders"),
		resourceids.UserSpecifiedSegment("reservationOrderId", "reservationOrderIdValue"),
	}
}

// String returns a human-readable description of this Reservation Order ID
func (id ReservationOrderId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Reservation Order: %q", id.ReservationOrderId),
	}
	return fmt.Sprintf("Reservation Order (%s)", strings.Join(components, "\n"))
}
