package shares

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ShareId{}

// ShareId is a struct representing the Resource ID for a Share
type ShareId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	ShareName             string
}

// NewShareID returns a new ShareId struct
func NewShareID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, shareName string) ShareId {
	return ShareId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		ShareName:             shareName,
	}
}

// ParseShareID parses 'input' into a ShareId
func ParseShareID(input string) (*ShareId, error) {
	parser := resourceids.NewParserFromResourceIdType(ShareId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ShareId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.ShareName, ok = parsed.Parsed["shareName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseShareIDInsensitively parses 'input' case-insensitively into a ShareId
// note: this method should only be used for API response data and not user input
func ParseShareIDInsensitively(input string) (*ShareId, error) {
	parser := resourceids.NewParserFromResourceIdType(ShareId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ShareId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.ShareName, ok = parsed.Parsed["shareName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateShareID checks that 'input' can be parsed as a Share ID
func ValidateShareID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseShareID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Share ID
func (id ShareId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/shares/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.ShareName)
}

// Segments returns a slice of Resource ID Segments which comprise this Share ID
func (id ShareId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticShares", "shares", "shares"),
		resourceids.UserSpecifiedSegment("shareName", "shareValue"),
	}
}

// String returns a human-readable description of this Share ID
func (id ShareId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Share Name: %q", id.ShareName),
	}
	return fmt.Sprintf("Share (%s)", strings.Join(components, "\n"))
}
