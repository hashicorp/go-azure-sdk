package reservationrecommendationdetails

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BillingScopeId{}

// BillingScopeId is a struct representing the Resource ID for a Billing Scope
type BillingScopeId struct {
	BillingScope string
}

// NewBillingScopeID returns a new BillingScopeId struct
func NewBillingScopeID(billingScope string) BillingScopeId {
	return BillingScopeId{
		BillingScope: billingScope,
	}
}

// ParseBillingScopeID parses 'input' into a BillingScopeId
func ParseBillingScopeID(input string) (*BillingScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingScopeId{}

	if id.BillingScope, ok = parsed.Parsed["billingScope"]; !ok {
		return nil, fmt.Errorf("the segment 'billingScope' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBillingScopeIDInsensitively parses 'input' case-insensitively into a BillingScopeId
// note: this method should only be used for API response data and not user input
func ParseBillingScopeIDInsensitively(input string) (*BillingScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingScopeId{}

	if id.BillingScope, ok = parsed.Parsed["billingScope"]; !ok {
		return nil, fmt.Errorf("the segment 'billingScope' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBillingScopeID checks that 'input' can be parsed as a Billing Scope ID
func ValidateBillingScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Scope ID
func (id BillingScopeId) ID() string {
	fmtString := "/%s"
	return fmt.Sprintf(fmtString, id.BillingScope)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Scope ID
func (id BillingScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.UserSpecifiedSegment("billingScope", "billingScopeValue"),
	}
}

// String returns a human-readable description of this Billing Scope ID
func (id BillingScopeId) String() string {
	components := []string{
		fmt.Sprintf("Billing Scope: %q", id.BillingScope),
	}
	return fmt.Sprintf("Billing Scope (%s)", strings.Join(components, "\n"))
}
