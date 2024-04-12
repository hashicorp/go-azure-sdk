package connectiongateways

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConnectionGatewayId{})
}

var _ resourceids.ResourceId = &ConnectionGatewayId{}

// ConnectionGatewayId is a struct representing the Resource ID for a Connection Gateway
type ConnectionGatewayId struct {
	SubscriptionId        string
	ResourceGroupName     string
	ConnectionGatewayName string
}

// NewConnectionGatewayID returns a new ConnectionGatewayId struct
func NewConnectionGatewayID(subscriptionId string, resourceGroupName string, connectionGatewayName string) ConnectionGatewayId {
	return ConnectionGatewayId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		ConnectionGatewayName: connectionGatewayName,
	}
}

// ParseConnectionGatewayID parses 'input' into a ConnectionGatewayId
func ParseConnectionGatewayID(input string) (*ConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionGatewayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionGatewayId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConnectionGatewayIDInsensitively parses 'input' case-insensitively into a ConnectionGatewayId
// note: this method should only be used for API response data and not user input
func ParseConnectionGatewayIDInsensitively(input string) (*ConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionGatewayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionGatewayId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConnectionGatewayId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ConnectionGatewayName, ok = input.Parsed["connectionGatewayName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectionGatewayName", input)
	}

	return nil
}

// ValidateConnectionGatewayID checks that 'input' can be parsed as a Connection Gateway ID
func ValidateConnectionGatewayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectionGatewayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connection Gateway ID
func (id ConnectionGatewayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/connectionGateways/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ConnectionGatewayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Connection Gateway ID
func (id ConnectionGatewayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticConnectionGateways", "connectionGateways", "connectionGateways"),
		resourceids.UserSpecifiedSegment("connectionGatewayName", "connectionGatewayValue"),
	}
}

// String returns a human-readable description of this Connection Gateway ID
func (id ConnectionGatewayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Connection Gateway Name: %q", id.ConnectionGatewayName),
	}
	return fmt.Sprintf("Connection Gateway (%s)", strings.Join(components, "\n"))
}
