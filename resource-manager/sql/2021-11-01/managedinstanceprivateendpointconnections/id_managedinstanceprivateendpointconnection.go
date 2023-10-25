package managedinstanceprivateendpointconnections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedInstancePrivateEndpointConnectionId{}

// ManagedInstancePrivateEndpointConnectionId is a struct representing the Resource ID for a Managed Instance Private Endpoint Connection
type ManagedInstancePrivateEndpointConnectionId struct {
	SubscriptionId                string
	ResourceGroupName             string
	ManagedInstanceName           string
	PrivateEndpointConnectionName string
}

// NewManagedInstancePrivateEndpointConnectionID returns a new ManagedInstancePrivateEndpointConnectionId struct
func NewManagedInstancePrivateEndpointConnectionID(subscriptionId string, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string) ManagedInstancePrivateEndpointConnectionId {
	return ManagedInstancePrivateEndpointConnectionId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		ManagedInstanceName:           managedInstanceName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
	}
}

// ParseManagedInstancePrivateEndpointConnectionID parses 'input' into a ManagedInstancePrivateEndpointConnectionId
func ParseManagedInstancePrivateEndpointConnectionID(input string) (*ManagedInstancePrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedInstancePrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedInstancePrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ParseManagedInstancePrivateEndpointConnectionIDInsensitively parses 'input' case-insensitively into a ManagedInstancePrivateEndpointConnectionId
// note: this method should only be used for API response data and not user input
func ParseManagedInstancePrivateEndpointConnectionIDInsensitively(input string) (*ManagedInstancePrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedInstancePrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedInstancePrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ValidateManagedInstancePrivateEndpointConnectionID checks that 'input' can be parsed as a Managed Instance Private Endpoint Connection ID
func ValidateManagedInstancePrivateEndpointConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstancePrivateEndpointConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Private Endpoint Connection ID
func (id ManagedInstancePrivateEndpointConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/privateEndpointConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.PrivateEndpointConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Private Endpoint Connection ID
func (id ManagedInstancePrivateEndpointConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionValue"),
	}
}

// String returns a human-readable description of this Managed Instance Private Endpoint Connection ID
func (id ManagedInstancePrivateEndpointConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
	}
	return fmt.Sprintf("Managed Instance Private Endpoint Connection (%s)", strings.Join(components, "\n"))
}
