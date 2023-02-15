package quotabycounterkeys

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QuotaId{}

// QuotaId is a struct representing the Resource ID for a Quota
type QuotaId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	QuotaCounterKey   string
}

// NewQuotaID returns a new QuotaId struct
func NewQuotaID(subscriptionId string, resourceGroupName string, serviceName string, quotaCounterKey string) QuotaId {
	return QuotaId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		QuotaCounterKey:   quotaCounterKey,
	}
}

// ParseQuotaID parses 'input' into a QuotaId
func ParseQuotaID(input string) (*QuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.QuotaCounterKey, ok = parsed.Parsed["quotaCounterKey"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaCounterKey' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseQuotaIDInsensitively parses 'input' case-insensitively into a QuotaId
// note: this method should only be used for API response data and not user input
func ParseQuotaIDInsensitively(input string) (*QuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.QuotaCounterKey, ok = parsed.Parsed["quotaCounterKey"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaCounterKey' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateQuotaID checks that 'input' can be parsed as a Quota ID
func ValidateQuotaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQuotaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Quota ID
func (id QuotaId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/quotas/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.QuotaCounterKey)
}

// Segments returns a slice of Resource ID Segments which comprise this Quota ID
func (id QuotaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticQuotas", "quotas", "quotas"),
		resourceids.UserSpecifiedSegment("quotaCounterKey", "quotaCounterKeyValue"),
	}
}

// String returns a human-readable description of this Quota ID
func (id QuotaId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Quota Counter Key: %q", id.QuotaCounterKey),
	}
	return fmt.Sprintf("Quota (%s)", strings.Join(components, "\n"))
}
