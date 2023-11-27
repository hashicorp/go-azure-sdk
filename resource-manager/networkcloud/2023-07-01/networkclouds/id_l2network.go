package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = L2NetworkId{}

// L2NetworkId is a struct representing the Resource ID for a L 2 Network
type L2NetworkId struct {
	SubscriptionId    string
	ResourceGroupName string
	L2NetworkName     string
}

// NewL2NetworkID returns a new L2NetworkId struct
func NewL2NetworkID(subscriptionId string, resourceGroupName string, l2NetworkName string) L2NetworkId {
	return L2NetworkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		L2NetworkName:     l2NetworkName,
	}
}

// ParseL2NetworkID parses 'input' into a L2NetworkId
func ParseL2NetworkID(input string) (*L2NetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(L2NetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := L2NetworkId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseL2NetworkIDInsensitively parses 'input' case-insensitively into a L2NetworkId
// note: this method should only be used for API response data and not user input
func ParseL2NetworkIDInsensitively(input string) (*L2NetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(L2NetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := L2NetworkId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *L2NetworkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.L2NetworkName, ok = input.Parsed["l2NetworkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "l2NetworkName", input)
	}

	return nil
}

// ValidateL2NetworkID checks that 'input' can be parsed as a L 2 Network ID
func ValidateL2NetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseL2NetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted L 2 Network ID
func (id L2NetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/l2Networks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.L2NetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this L 2 Network ID
func (id L2NetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticL2Networks", "l2Networks", "l2Networks"),
		resourceids.UserSpecifiedSegment("l2NetworkName", "l2NetworkValue"),
	}
}

// String returns a human-readable description of this L 2 Network ID
func (id L2NetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("L 2 Network Name: %q", id.L2NetworkName),
	}
	return fmt.Sprintf("L 2 Network (%s)", strings.Join(components, "\n"))
}
