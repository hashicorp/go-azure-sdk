package gates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&GateId{})
}

var _ resourceids.ResourceId = &GateId{}

// GateId is a struct representing the Resource ID for a Gate
type GateId struct {
	SubscriptionId    string
	ResourceGroupName string
	FleetName         string
	GateName          string
}

// NewGateID returns a new GateId struct
func NewGateID(subscriptionId string, resourceGroupName string, fleetName string, gateName string) GateId {
	return GateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FleetName:         fleetName,
		GateName:          gateName,
	}
}

// ParseGateID parses 'input' into a GateId
func ParseGateID(input string) (*GateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGateIDInsensitively parses 'input' case-insensitively into a GateId
// note: this method should only be used for API response data and not user input
func ParseGateIDInsensitively(input string) (*GateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.FleetName, ok = input.Parsed["fleetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fleetName", input)
	}

	if id.GateName, ok = input.Parsed["gateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "gateName", input)
	}

	return nil
}

// ValidateGateID checks that 'input' can be parsed as a Gate ID
func ValidateGateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Gate ID
func (id GateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerService/fleets/%s/gates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FleetName, id.GateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Gate ID
func (id GateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftContainerService", "Microsoft.ContainerService", "Microsoft.ContainerService"),
		resourceids.StaticSegment("staticFleets", "fleets", "fleets"),
		resourceids.UserSpecifiedSegment("fleetName", "fleetName"),
		resourceids.StaticSegment("staticGates", "gates", "gates"),
		resourceids.UserSpecifiedSegment("gateName", "gateName"),
	}
}

// String returns a human-readable description of this Gate ID
func (id GateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Fleet Name: %q", id.FleetName),
		fmt.Sprintf("Gate Name: %q", id.GateName),
	}
	return fmt.Sprintf("Gate (%s)", strings.Join(components, "\n"))
}
