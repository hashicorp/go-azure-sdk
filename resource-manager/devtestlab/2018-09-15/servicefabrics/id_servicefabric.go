// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabrics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ServiceFabricId{}

// ServiceFabricId is a struct representing the Resource ID for a Service Fabric
type ServiceFabricId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	UserName          string
	ServiceFabricName string
}

// NewServiceFabricID returns a new ServiceFabricId struct
func NewServiceFabricID(subscriptionId string, resourceGroupName string, labName string, userName string, serviceFabricName string) ServiceFabricId {
	return ServiceFabricId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		UserName:          userName,
		ServiceFabricName: serviceFabricName,
	}
}

// ParseServiceFabricID parses 'input' into a ServiceFabricId
func ParseServiceFabricID(input string) (*ServiceFabricId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceFabricId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceFabricId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.UserName, ok = parsed.Parsed["userName"]; !ok {
		return nil, fmt.Errorf("the segment 'userName' was not found in the resource id %q", input)
	}

	if id.ServiceFabricName, ok = parsed.Parsed["serviceFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceFabricName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseServiceFabricIDInsensitively parses 'input' case-insensitively into a ServiceFabricId
// note: this method should only be used for API response data and not user input
func ParseServiceFabricIDInsensitively(input string) (*ServiceFabricId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceFabricId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceFabricId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.UserName, ok = parsed.Parsed["userName"]; !ok {
		return nil, fmt.Errorf("the segment 'userName' was not found in the resource id %q", input)
	}

	if id.ServiceFabricName, ok = parsed.Parsed["serviceFabricName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceFabricName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateServiceFabricID checks that 'input' can be parsed as a Service Fabric ID
func ValidateServiceFabricID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceFabricID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Fabric ID
func (id ServiceFabricId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/users/%s/serviceFabrics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.UserName, id.ServiceFabricName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Fabric ID
func (id ServiceFabricId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userName", "userValue"),
		resourceids.StaticSegment("staticServiceFabrics", "serviceFabrics", "serviceFabrics"),
		resourceids.UserSpecifiedSegment("serviceFabricName", "serviceFabricValue"),
	}
}

// String returns a human-readable description of this Service Fabric ID
func (id ServiceFabricId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("User Name: %q", id.UserName),
		fmt.Sprintf("Service Fabric Name: %q", id.ServiceFabricName),
	}
	return fmt.Sprintf("Service Fabric (%s)", strings.Join(components, "\n"))
}
