package quotabyperiodkeys

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PeriodId{}

// PeriodId is a struct representing the Resource ID for a Period
type PeriodId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	QuotaCounterKey   string
	QuotaPeriodKey    string
}

// NewPeriodID returns a new PeriodId struct
func NewPeriodID(subscriptionId string, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string) PeriodId {
	return PeriodId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		QuotaCounterKey:   quotaCounterKey,
		QuotaPeriodKey:    quotaPeriodKey,
	}
}

// ParsePeriodID parses 'input' into a PeriodId
func ParsePeriodID(input string) (*PeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(PeriodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PeriodId{}

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

	if id.QuotaPeriodKey, ok = parsed.Parsed["quotaPeriodKey"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaPeriodKey' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePeriodIDInsensitively parses 'input' case-insensitively into a PeriodId
// note: this method should only be used for API response data and not user input
func ParsePeriodIDInsensitively(input string) (*PeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(PeriodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PeriodId{}

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

	if id.QuotaPeriodKey, ok = parsed.Parsed["quotaPeriodKey"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaPeriodKey' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePeriodID checks that 'input' can be parsed as a Period ID
func ValidatePeriodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePeriodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Period ID
func (id PeriodId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/quotas/%s/periods/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.QuotaCounterKey, id.QuotaPeriodKey)
}

// Segments returns a slice of Resource ID Segments which comprise this Period ID
func (id PeriodId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticPeriods", "periods", "periods"),
		resourceids.UserSpecifiedSegment("quotaPeriodKey", "quotaPeriodKeyValue"),
	}
}

// String returns a human-readable description of this Period ID
func (id PeriodId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Quota Counter Key: %q", id.QuotaCounterKey),
		fmt.Sprintf("Quota Period Key: %q", id.QuotaPeriodKey),
	}
	return fmt.Sprintf("Period (%s)", strings.Join(components, "\n"))
}
