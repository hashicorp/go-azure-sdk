// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicerunners

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ServiceRunnerId{}

// ServiceRunnerId is a struct representing the Resource ID for a Service Runner
type ServiceRunnerId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	ServiceRunnerName string
}

// NewServiceRunnerID returns a new ServiceRunnerId struct
func NewServiceRunnerID(subscriptionId string, resourceGroupName string, labName string, serviceRunnerName string) ServiceRunnerId {
	return ServiceRunnerId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		ServiceRunnerName: serviceRunnerName,
	}
}

// ParseServiceRunnerID parses 'input' into a ServiceRunnerId
func ParseServiceRunnerID(input string) (*ServiceRunnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceRunnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceRunnerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.ServiceRunnerName, ok = parsed.Parsed["serviceRunnerName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceRunnerName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseServiceRunnerIDInsensitively parses 'input' case-insensitively into a ServiceRunnerId
// note: this method should only be used for API response data and not user input
func ParseServiceRunnerIDInsensitively(input string) (*ServiceRunnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceRunnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceRunnerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.ServiceRunnerName, ok = parsed.Parsed["serviceRunnerName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceRunnerName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateServiceRunnerID checks that 'input' can be parsed as a Service Runner ID
func ValidateServiceRunnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceRunnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Runner ID
func (id ServiceRunnerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/serviceRunners/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.ServiceRunnerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Runner ID
func (id ServiceRunnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticServiceRunners", "serviceRunners", "serviceRunners"),
		resourceids.UserSpecifiedSegment("serviceRunnerName", "serviceRunnerValue"),
	}
}

// String returns a human-readable description of this Service Runner ID
func (id ServiceRunnerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Service Runner Name: %q", id.ServiceRunnerName),
	}
	return fmt.Sprintf("Service Runner (%s)", strings.Join(components, "\n"))
}
