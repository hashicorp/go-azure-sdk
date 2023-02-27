// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package portalconfig

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PortalConfigId{}

// PortalConfigId is a struct representing the Resource ID for a Portal Config
type PortalConfigId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	PortalConfigId    string
}

// NewPortalConfigID returns a new PortalConfigId struct
func NewPortalConfigID(subscriptionId string, resourceGroupName string, serviceName string, portalConfigId string) PortalConfigId {
	return PortalConfigId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		PortalConfigId:    portalConfigId,
	}
}

// ParsePortalConfigID parses 'input' into a PortalConfigId
func ParsePortalConfigID(input string) (*PortalConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(PortalConfigId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PortalConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.PortalConfigId, ok = parsed.Parsed["portalConfigId"]; !ok {
		return nil, fmt.Errorf("the segment 'portalConfigId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePortalConfigIDInsensitively parses 'input' case-insensitively into a PortalConfigId
// note: this method should only be used for API response data and not user input
func ParsePortalConfigIDInsensitively(input string) (*PortalConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(PortalConfigId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PortalConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.PortalConfigId, ok = parsed.Parsed["portalConfigId"]; !ok {
		return nil, fmt.Errorf("the segment 'portalConfigId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePortalConfigID checks that 'input' can be parsed as a Portal Config ID
func ValidatePortalConfigID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePortalConfigID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Portal Config ID
func (id PortalConfigId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/portalConfigs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.PortalConfigId)
}

// Segments returns a slice of Resource ID Segments which comprise this Portal Config ID
func (id PortalConfigId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticPortalConfigs", "portalConfigs", "portalConfigs"),
		resourceids.UserSpecifiedSegment("portalConfigId", "portalConfigIdValue"),
	}
}

// String returns a human-readable description of this Portal Config ID
func (id PortalConfigId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Portal Config: %q", id.PortalConfigId),
	}
	return fmt.Sprintf("Portal Config (%s)", strings.Join(components, "\n"))
}
