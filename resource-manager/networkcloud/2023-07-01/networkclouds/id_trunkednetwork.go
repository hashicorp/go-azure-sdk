package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TrunkedNetworkId{})
}

var _ resourceids.ResourceId = &TrunkedNetworkId{}

// TrunkedNetworkId is a struct representing the Resource ID for a Trunked Network
type TrunkedNetworkId struct {
	SubscriptionId     string
	ResourceGroupName  string
	TrunkedNetworkName string
}

// NewTrunkedNetworkID returns a new TrunkedNetworkId struct
func NewTrunkedNetworkID(subscriptionId string, resourceGroupName string, trunkedNetworkName string) TrunkedNetworkId {
	return TrunkedNetworkId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		TrunkedNetworkName: trunkedNetworkName,
	}
}

// ParseTrunkedNetworkID parses 'input' into a TrunkedNetworkId
func ParseTrunkedNetworkID(input string) (*TrunkedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TrunkedNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TrunkedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTrunkedNetworkIDInsensitively parses 'input' case-insensitively into a TrunkedNetworkId
// note: this method should only be used for API response data and not user input
func ParseTrunkedNetworkIDInsensitively(input string) (*TrunkedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TrunkedNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TrunkedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TrunkedNetworkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.TrunkedNetworkName, ok = input.Parsed["trunkedNetworkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "trunkedNetworkName", input)
	}

	return nil
}

// ValidateTrunkedNetworkID checks that 'input' can be parsed as a Trunked Network ID
func ValidateTrunkedNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTrunkedNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trunked Network ID
func (id TrunkedNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/trunkedNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.TrunkedNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Trunked Network ID
func (id TrunkedNetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticTrunkedNetworks", "trunkedNetworks", "trunkedNetworks"),
		resourceids.UserSpecifiedSegment("trunkedNetworkName", "trunkedNetworkValue"),
	}
}

// String returns a human-readable description of this Trunked Network ID
func (id TrunkedNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Trunked Network Name: %q", id.TrunkedNetworkName),
	}
	return fmt.Sprintf("Trunked Network (%s)", strings.Join(components, "\n"))
}
