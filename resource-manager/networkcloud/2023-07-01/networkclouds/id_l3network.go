package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = L3NetworkId{}

// L3NetworkId is a struct representing the Resource ID for a L 3 Network
type L3NetworkId struct {
	SubscriptionId    string
	ResourceGroupName string
	L3NetworkName     string
}

// NewL3NetworkID returns a new L3NetworkId struct
func NewL3NetworkID(subscriptionId string, resourceGroupName string, l3NetworkName string) L3NetworkId {
	return L3NetworkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		L3NetworkName:     l3NetworkName,
	}
}

// ParseL3NetworkID parses 'input' into a L3NetworkId
func ParseL3NetworkID(input string) (*L3NetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(L3NetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := L3NetworkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.L3NetworkName, ok = parsed.Parsed["l3NetworkName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "l3NetworkName", *parsed)
	}

	return &id, nil
}

// ParseL3NetworkIDInsensitively parses 'input' case-insensitively into a L3NetworkId
// note: this method should only be used for API response data and not user input
func ParseL3NetworkIDInsensitively(input string) (*L3NetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(L3NetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := L3NetworkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.L3NetworkName, ok = parsed.Parsed["l3NetworkName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "l3NetworkName", *parsed)
	}

	return &id, nil
}

// ValidateL3NetworkID checks that 'input' can be parsed as a L 3 Network ID
func ValidateL3NetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseL3NetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted L 3 Network ID
func (id L3NetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/l3Networks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.L3NetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this L 3 Network ID
func (id L3NetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticL3Networks", "l3Networks", "l3Networks"),
		resourceids.UserSpecifiedSegment("l3NetworkName", "l3NetworkValue"),
	}
}

// String returns a human-readable description of this L 3 Network ID
func (id L3NetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("L 3 Network Name: %q", id.L3NetworkName),
	}
	return fmt.Sprintf("L 3 Network (%s)", strings.Join(components, "\n"))
}
