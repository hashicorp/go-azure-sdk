package allowedconnections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConnectionTypeId{}

// ConnectionTypeId is a struct representing the Resource ID for a Connection Type
type ConnectionTypeId struct {
	SubscriptionId    string
	ResourceGroupName string
	LocationName      string
	ConnectionType    ConnectionType
}

// NewConnectionTypeID returns a new ConnectionTypeId struct
func NewConnectionTypeID(subscriptionId string, resourceGroupName string, locationName string, connectionType ConnectionType) ConnectionTypeId {
	return ConnectionTypeId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LocationName:      locationName,
		ConnectionType:    connectionType,
	}
}

// ParseConnectionTypeID parses 'input' into a ConnectionTypeId
func ParseConnectionTypeID(input string) (*ConnectionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConnectionTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConnectionTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if v, ok := parsed.Parsed["connectionType"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "connectionType", *parsed)
		}

		connectionType, err := parseConnectionType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ConnectionType = *connectionType
	}

	return &id, nil
}

// ParseConnectionTypeIDInsensitively parses 'input' case-insensitively into a ConnectionTypeId
// note: this method should only be used for API response data and not user input
func ParseConnectionTypeIDInsensitively(input string) (*ConnectionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConnectionTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConnectionTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if v, ok := parsed.Parsed["connectionType"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "connectionType", *parsed)
		}

		connectionType, err := parseConnectionType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ConnectionType = *connectionType
	}

	return &id, nil
}

// ValidateConnectionTypeID checks that 'input' can be parsed as a Connection Type ID
func ValidateConnectionTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectionTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connection Type ID
func (id ConnectionTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/allowedConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, string(id.ConnectionType))
}

// Segments returns a slice of Resource ID Segments which comprise this Connection Type ID
func (id ConnectionTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticAllowedConnections", "allowedConnections", "allowedConnections"),
		resourceids.ConstantSegment("connectionType", PossibleValuesForConnectionType(), "External"),
	}
}

// String returns a human-readable description of this Connection Type ID
func (id ConnectionTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Connection Type: %q", string(id.ConnectionType)),
	}
	return fmt.Sprintf("Connection Type (%s)", strings.Join(components, "\n"))
}
