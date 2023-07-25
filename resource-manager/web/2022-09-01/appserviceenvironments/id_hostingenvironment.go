package appserviceenvironments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HostingEnvironmentId{}

// HostingEnvironmentId is a struct representing the Resource ID for a Hosting Environment
type HostingEnvironmentId struct {
	SubscriptionId         string
	ResourceGroupName      string
	HostingEnvironmentName string
}

// NewHostingEnvironmentID returns a new HostingEnvironmentId struct
func NewHostingEnvironmentID(subscriptionId string, resourceGroupName string, hostingEnvironmentName string) HostingEnvironmentId {
	return HostingEnvironmentId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		HostingEnvironmentName: hostingEnvironmentName,
	}
}

// ParseHostingEnvironmentID parses 'input' into a HostingEnvironmentId
func ParseHostingEnvironmentID(input string) (*HostingEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	return &id, nil
}

// ParseHostingEnvironmentIDInsensitively parses 'input' case-insensitively into a HostingEnvironmentId
// note: this method should only be used for API response data and not user input
func ParseHostingEnvironmentIDInsensitively(input string) (*HostingEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	return &id, nil
}

// ValidateHostingEnvironmentID checks that 'input' can be parsed as a Hosting Environment ID
func ValidateHostingEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostingEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hosting Environment ID
func (id HostingEnvironmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/hostingEnvironments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HostingEnvironmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hosting Environment ID
func (id HostingEnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticHostingEnvironments", "hostingEnvironments", "hostingEnvironments"),
		resourceids.UserSpecifiedSegment("hostingEnvironmentName", "hostingEnvironmentValue"),
	}
}

// String returns a human-readable description of this Hosting Environment ID
func (id HostingEnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Hosting Environment Name: %q", id.HostingEnvironmentName),
	}
	return fmt.Sprintf("Hosting Environment (%s)", strings.Join(components, "\n"))
}
