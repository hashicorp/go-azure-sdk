package reservationdetails

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReservationId{}

// ReservationId is a struct representing the Resource ID for a Reservation
type ReservationId struct {
	ReservationOrderId string
	ReservationId      string
}

// NewReservationID returns a new ReservationId struct
func NewReservationID(reservationOrderId string, reservationId string) ReservationId {
	return ReservationId{
		ReservationOrderId: reservationOrderId,
		ReservationId:      reservationId,
	}
}

// ParseReservationID parses 'input' into a ReservationId
func ParseReservationID(input string) (*ReservationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReservationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReservationId{}

	if id.ReservationOrderId, ok = parsed.Parsed["reservationOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reservationOrderId", *parsed)
	}

	if id.ReservationId, ok = parsed.Parsed["reservationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reservationId", *parsed)
	}

	return &id, nil
}

// ParseReservationIDInsensitively parses 'input' case-insensitively into a ReservationId
// note: this method should only be used for API response data and not user input
func ParseReservationIDInsensitively(input string) (*ReservationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReservationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReservationId{}

	if id.ReservationOrderId, ok = parsed.Parsed["reservationOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reservationOrderId", *parsed)
	}

	if id.ReservationId, ok = parsed.Parsed["reservationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reservationId", *parsed)
	}

	return &id, nil
}

// ValidateReservationID checks that 'input' can be parsed as a Reservation ID
func ValidateReservationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReservationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Reservation ID
func (id ReservationId) ID() string {
	fmtString := "/providers/Microsoft.Capacity/reservationOrders/%s/reservations/%s"
	return fmt.Sprintf(fmtString, id.ReservationOrderId, id.ReservationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Reservation ID
func (id ReservationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCapacity", "Microsoft.Capacity", "Microsoft.Capacity"),
		resourceids.StaticSegment("staticReservationOrders", "reservationOrders", "reservationOrders"),
		resourceids.UserSpecifiedSegment("reservationOrderId", "reservationOrderIdValue"),
		resourceids.StaticSegment("staticReservations", "reservations", "reservations"),
		resourceids.UserSpecifiedSegment("reservationId", "reservationIdValue"),
	}
}

// String returns a human-readable description of this Reservation ID
func (id ReservationId) String() string {
	components := []string{
		fmt.Sprintf("Reservation Order: %q", id.ReservationOrderId),
		fmt.Sprintf("Reservation: %q", id.ReservationId),
	}
	return fmt.Sprintf("Reservation (%s)", strings.Join(components, "\n"))
}
