package reservationsummaries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReservationOrderId{}

// ReservationOrderId is a struct representing the Resource ID for a Reservation Order
type ReservationOrderId struct {
	ReservationOrderId string
}

// NewReservationOrderID returns a new ReservationOrderId struct
func NewReservationOrderID(reservationOrderId string) ReservationOrderId {
	return ReservationOrderId{
		ReservationOrderId: reservationOrderId,
	}
}

// ParseReservationOrderID parses 'input' into a ReservationOrderId
func ParseReservationOrderID(input string) (*ReservationOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReservationOrderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReservationOrderId{}

	if id.ReservationOrderId, ok = parsed.Parsed["reservationOrderId"]; !ok {
		return nil, fmt.Errorf("the segment 'reservationOrderId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReservationOrderIDInsensitively parses 'input' case-insensitively into a ReservationOrderId
// note: this method should only be used for API response data and not user input
func ParseReservationOrderIDInsensitively(input string) (*ReservationOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReservationOrderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReservationOrderId{}

	if id.ReservationOrderId, ok = parsed.Parsed["reservationOrderId"]; !ok {
		return nil, fmt.Errorf("the segment 'reservationOrderId' was not found in the resource id %q", input)
	}

	return &id, nil
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
	fmtString := "/providers/Microsoft.Capacity/reservationOrders/%s"
	return fmt.Sprintf(fmtString, id.ReservationOrderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Reservation Order ID
func (id ReservationOrderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCapacity", "Microsoft.Capacity", "Microsoft.Capacity"),
		resourceids.StaticSegment("staticReservationOrders", "reservationOrders", "reservationOrders"),
		resourceids.UserSpecifiedSegment("reservationOrderId", "reservationOrderIdValue"),
	}
}

// String returns a human-readable description of this Reservation Order ID
func (id ReservationOrderId) String() string {
	components := []string{
		fmt.Sprintf("Reservation Order: %q", id.ReservationOrderId),
	}
	return fmt.Sprintf("Reservation Order (%s)", strings.Join(components, "\n"))
}
