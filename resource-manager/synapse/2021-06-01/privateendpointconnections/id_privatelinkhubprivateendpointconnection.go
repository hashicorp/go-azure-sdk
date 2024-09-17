package privateendpointconnections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PrivateLinkHubPrivateEndpointConnectionId{})
}

var _ resourceids.ResourceId = &PrivateLinkHubPrivateEndpointConnectionId{}

// PrivateLinkHubPrivateEndpointConnectionId is a struct representing the Resource ID for a Private Link Hub Private Endpoint Connection
type PrivateLinkHubPrivateEndpointConnectionId struct {
	SubscriptionId                string
	ResourceGroupName             string
	PrivateLinkHubName            string
	PrivateEndpointConnectionName string
}

// NewPrivateLinkHubPrivateEndpointConnectionID returns a new PrivateLinkHubPrivateEndpointConnectionId struct
func NewPrivateLinkHubPrivateEndpointConnectionID(subscriptionId string, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string) PrivateLinkHubPrivateEndpointConnectionId {
	return PrivateLinkHubPrivateEndpointConnectionId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		PrivateLinkHubName:            privateLinkHubName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
	}
}

// ParsePrivateLinkHubPrivateEndpointConnectionID parses 'input' into a PrivateLinkHubPrivateEndpointConnectionId
func ParsePrivateLinkHubPrivateEndpointConnectionID(input string) (*PrivateLinkHubPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PrivateLinkHubPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PrivateLinkHubPrivateEndpointConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePrivateLinkHubPrivateEndpointConnectionIDInsensitively parses 'input' case-insensitively into a PrivateLinkHubPrivateEndpointConnectionId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkHubPrivateEndpointConnectionIDInsensitively(input string) (*PrivateLinkHubPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PrivateLinkHubPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PrivateLinkHubPrivateEndpointConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PrivateLinkHubPrivateEndpointConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateLinkHubName, ok = input.Parsed["privateLinkHubName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateLinkHubName", input)
	}

	if id.PrivateEndpointConnectionName, ok = input.Parsed["privateEndpointConnectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", input)
	}

	return nil
}

// ValidatePrivateLinkHubPrivateEndpointConnectionID checks that 'input' can be parsed as a Private Link Hub Private Endpoint Connection ID
func ValidatePrivateLinkHubPrivateEndpointConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkHubPrivateEndpointConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link Hub Private Endpoint Connection ID
func (id PrivateLinkHubPrivateEndpointConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/privateLinkHubs/%s/privateEndpointConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateLinkHubName, id.PrivateEndpointConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link Hub Private Endpoint Connection ID
func (id PrivateLinkHubPrivateEndpointConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticPrivateLinkHubs", "privateLinkHubs", "privateLinkHubs"),
		resourceids.UserSpecifiedSegment("privateLinkHubName", "privateLinkHubValue"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionValue"),
	}
}

// String returns a human-readable description of this Private Link Hub Private Endpoint Connection ID
func (id PrivateLinkHubPrivateEndpointConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Link Hub Name: %q", id.PrivateLinkHubName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
	}
	return fmt.Sprintf("Private Link Hub Private Endpoint Connection (%s)", strings.Join(components, "\n"))
}
