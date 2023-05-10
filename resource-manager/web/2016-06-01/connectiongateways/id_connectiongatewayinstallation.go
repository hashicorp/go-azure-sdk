package connectiongateways

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConnectionGatewayInstallationId{}

// ConnectionGatewayInstallationId is a struct representing the Resource ID for a Connection Gateway Installation
type ConnectionGatewayInstallationId struct {
	SubscriptionId string
	LocationName   string
	GatewayId      string
}

// NewConnectionGatewayInstallationID returns a new ConnectionGatewayInstallationId struct
func NewConnectionGatewayInstallationID(subscriptionId string, locationName string, gatewayId string) ConnectionGatewayInstallationId {
	return ConnectionGatewayInstallationId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		GatewayId:      gatewayId,
	}
}

// ParseConnectionGatewayInstallationID parses 'input' into a ConnectionGatewayInstallationId
func ParseConnectionGatewayInstallationID(input string) (*ConnectionGatewayInstallationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConnectionGatewayInstallationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConnectionGatewayInstallationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.GatewayId, ok = parsed.Parsed["gatewayId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayId", *parsed)
	}

	return &id, nil
}

// ParseConnectionGatewayInstallationIDInsensitively parses 'input' case-insensitively into a ConnectionGatewayInstallationId
// note: this method should only be used for API response data and not user input
func ParseConnectionGatewayInstallationIDInsensitively(input string) (*ConnectionGatewayInstallationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConnectionGatewayInstallationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConnectionGatewayInstallationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.GatewayId, ok = parsed.Parsed["gatewayId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayId", *parsed)
	}

	return &id, nil
}

// ValidateConnectionGatewayInstallationID checks that 'input' can be parsed as a Connection Gateway Installation ID
func ValidateConnectionGatewayInstallationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectionGatewayInstallationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connection Gateway Installation ID
func (id ConnectionGatewayInstallationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/locations/%s/connectionGatewayInstallations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.GatewayId)
}

// Segments returns a slice of Resource ID Segments which comprise this Connection Gateway Installation ID
func (id ConnectionGatewayInstallationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticConnectionGatewayInstallations", "connectionGatewayInstallations", "connectionGatewayInstallations"),
		resourceids.UserSpecifiedSegment("gatewayId", "gatewayIdValue"),
	}
}

// String returns a human-readable description of this Connection Gateway Installation ID
func (id ConnectionGatewayInstallationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Gateway: %q", id.GatewayId),
	}
	return fmt.Sprintf("Connection Gateway Installation (%s)", strings.Join(components, "\n"))
}
