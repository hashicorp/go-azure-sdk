package pricesheet20220601

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BillingPeriodId{}

// BillingPeriodId is a struct representing the Resource ID for a Billing Period
type BillingPeriodId struct {
	SubscriptionId    string
	BillingPeriodName string
}

// NewBillingPeriodID returns a new BillingPeriodId struct
func NewBillingPeriodID(subscriptionId string, billingPeriodName string) BillingPeriodId {
	return BillingPeriodId{
		SubscriptionId:    subscriptionId,
		BillingPeriodName: billingPeriodName,
	}
}

// ParseBillingPeriodID parses 'input' into a BillingPeriodId
func ParseBillingPeriodID(input string) (*BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingPeriodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingPeriodId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.BillingPeriodName, ok = parsed.Parsed["billingPeriodName"]; !ok {
		return nil, fmt.Errorf("the segment 'billingPeriodName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBillingPeriodIDInsensitively parses 'input' case-insensitively into a BillingPeriodId
// note: this method should only be used for API response data and not user input
func ParseBillingPeriodIDInsensitively(input string) (*BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingPeriodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingPeriodId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.BillingPeriodName, ok = parsed.Parsed["billingPeriodName"]; !ok {
		return nil, fmt.Errorf("the segment 'billingPeriodName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBillingPeriodID checks that 'input' can be parsed as a Billing Period ID
func ValidateBillingPeriodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingPeriodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Period ID
func (id BillingPeriodId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Billing/billingPeriods/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.BillingPeriodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Period ID
func (id BillingPeriodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingPeriods", "billingPeriods", "billingPeriods"),
		resourceids.UserSpecifiedSegment("billingPeriodName", "billingPeriodValue"),
	}
}

// String returns a human-readable description of this Billing Period ID
func (id BillingPeriodId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Billing Period Name: %q", id.BillingPeriodName),
	}
	return fmt.Sprintf("Billing Period (%s)", strings.Join(components, "\n"))
}
