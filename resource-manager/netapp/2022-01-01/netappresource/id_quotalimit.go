package netappresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QuotaLimitId{}

// QuotaLimitId is a struct representing the Resource ID for a Quota Limit
type QuotaLimitId struct {
	SubscriptionId string
	Location       string
	QuotaLimitName string
}

// NewQuotaLimitID returns a new QuotaLimitId struct
func NewQuotaLimitID(subscriptionId string, location string, quotaLimitName string) QuotaLimitId {
	return QuotaLimitId{
		SubscriptionId: subscriptionId,
		Location:       location,
		QuotaLimitName: quotaLimitName,
	}
}

// ParseQuotaLimitID parses 'input' into a QuotaLimitId
func ParseQuotaLimitID(input string) (*QuotaLimitId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaLimitId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaLimitId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.QuotaLimitName, ok = parsed.Parsed["quotaLimitName"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaLimitName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseQuotaLimitIDInsensitively parses 'input' case-insensitively into a QuotaLimitId
// note: this method should only be used for API response data and not user input
func ParseQuotaLimitIDInsensitively(input string) (*QuotaLimitId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaLimitId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaLimitId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.QuotaLimitName, ok = parsed.Parsed["quotaLimitName"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaLimitName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateQuotaLimitID checks that 'input' can be parsed as a Quota Limit ID
func ValidateQuotaLimitID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQuotaLimitID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Quota Limit ID
func (id QuotaLimitId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.NetApp/locations/%s/quotaLimits/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.QuotaLimitName)
}

// Segments returns a slice of Resource ID Segments which comprise this Quota Limit ID
func (id QuotaLimitId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticQuotaLimits", "quotaLimits", "quotaLimits"),
		resourceids.UserSpecifiedSegment("quotaLimitName", "quotaLimitValue"),
	}
}

// String returns a human-readable description of this Quota Limit ID
func (id QuotaLimitId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Quota Limit Name: %q", id.QuotaLimitName),
	}
	return fmt.Sprintf("Quota Limit (%s)", strings.Join(components, "\n"))
}
