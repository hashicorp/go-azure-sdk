package portalrevision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PortalRevisionId{}

// PortalRevisionId is a struct representing the Resource ID for a Portal Revision
type PortalRevisionId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	PortalRevisionId  string
}

// NewPortalRevisionID returns a new PortalRevisionId struct
func NewPortalRevisionID(subscriptionId string, resourceGroupName string, serviceName string, portalRevisionId string) PortalRevisionId {
	return PortalRevisionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		PortalRevisionId:  portalRevisionId,
	}
}

// ParsePortalRevisionID parses 'input' into a PortalRevisionId
func ParsePortalRevisionID(input string) (*PortalRevisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PortalRevisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PortalRevisionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.PortalRevisionId, ok = parsed.Parsed["portalRevisionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "portalRevisionId", *parsed)
	}

	return &id, nil
}

// ParsePortalRevisionIDInsensitively parses 'input' case-insensitively into a PortalRevisionId
// note: this method should only be used for API response data and not user input
func ParsePortalRevisionIDInsensitively(input string) (*PortalRevisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PortalRevisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PortalRevisionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.PortalRevisionId, ok = parsed.Parsed["portalRevisionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "portalRevisionId", *parsed)
	}

	return &id, nil
}

// ValidatePortalRevisionID checks that 'input' can be parsed as a Portal Revision ID
func ValidatePortalRevisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePortalRevisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Portal Revision ID
func (id PortalRevisionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/portalRevisions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.PortalRevisionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Portal Revision ID
func (id PortalRevisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticPortalRevisions", "portalRevisions", "portalRevisions"),
		resourceids.UserSpecifiedSegment("portalRevisionId", "portalRevisionIdValue"),
	}
}

// String returns a human-readable description of this Portal Revision ID
func (id PortalRevisionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Portal Revision: %q", id.PortalRevisionId),
	}
	return fmt.Sprintf("Portal Revision (%s)", strings.Join(components, "\n"))
}
