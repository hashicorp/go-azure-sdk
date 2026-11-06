package managedenvironments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagedEnvironmentPrivateEndpointConnectionId{})
}

var _ resourceids.ResourceId = &ManagedEnvironmentPrivateEndpointConnectionId{}

// ManagedEnvironmentPrivateEndpointConnectionId is a struct representing the Resource ID for a Managed Environment Private Endpoint Connection
type ManagedEnvironmentPrivateEndpointConnectionId struct {
	SubscriptionId                string
	ResourceGroupName             string
	ManagedEnvironmentName        string
	PrivateEndpointConnectionName string
}

// NewManagedEnvironmentPrivateEndpointConnectionID returns a new ManagedEnvironmentPrivateEndpointConnectionId struct
func NewManagedEnvironmentPrivateEndpointConnectionID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, privateEndpointConnectionName string) ManagedEnvironmentPrivateEndpointConnectionId {
	return ManagedEnvironmentPrivateEndpointConnectionId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		ManagedEnvironmentName:        managedEnvironmentName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
	}
}

// ParseManagedEnvironmentPrivateEndpointConnectionID parses 'input' into a ManagedEnvironmentPrivateEndpointConnectionId
func ParseManagedEnvironmentPrivateEndpointConnectionID(input string) (*ManagedEnvironmentPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedEnvironmentPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedEnvironmentPrivateEndpointConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedEnvironmentPrivateEndpointConnectionIDInsensitively parses 'input' case-insensitively into a ManagedEnvironmentPrivateEndpointConnectionId
// note: this method should only be used for API response data and not user input
func ParseManagedEnvironmentPrivateEndpointConnectionIDInsensitively(input string) (*ManagedEnvironmentPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedEnvironmentPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedEnvironmentPrivateEndpointConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedEnvironmentPrivateEndpointConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedEnvironmentName, ok = input.Parsed["managedEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedEnvironmentName", input)
	}

	if id.PrivateEndpointConnectionName, ok = input.Parsed["privateEndpointConnectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", input)
	}

	return nil
}

// ValidateManagedEnvironmentPrivateEndpointConnectionID checks that 'input' can be parsed as a Managed Environment Private Endpoint Connection ID
func ValidateManagedEnvironmentPrivateEndpointConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedEnvironmentPrivateEndpointConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Environment Private Endpoint Connection ID
func (id ManagedEnvironmentPrivateEndpointConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/privateEndpointConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.PrivateEndpointConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Environment Private Endpoint Connection ID
func (id ManagedEnvironmentPrivateEndpointConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentName"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionName"),
	}
}

// String returns a human-readable description of this Managed Environment Private Endpoint Connection ID
func (id ManagedEnvironmentPrivateEndpointConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
	}
	return fmt.Sprintf("Managed Environment Private Endpoint Connection (%s)", strings.Join(components, "\n"))
}
