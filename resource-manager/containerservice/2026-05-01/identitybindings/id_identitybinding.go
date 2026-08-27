package identitybindings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IdentityBindingId{})
}

var _ resourceids.ResourceId = &IdentityBindingId{}

// IdentityBindingId is a struct representing the Resource ID for a Identity Binding
type IdentityBindingId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedClusterName  string
	IdentityBindingName string
}

// NewIdentityBindingID returns a new IdentityBindingId struct
func NewIdentityBindingID(subscriptionId string, resourceGroupName string, managedClusterName string, identityBindingName string) IdentityBindingId {
	return IdentityBindingId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedClusterName:  managedClusterName,
		IdentityBindingName: identityBindingName,
	}
}

// ParseIdentityBindingID parses 'input' into a IdentityBindingId
func ParseIdentityBindingID(input string) (*IdentityBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityBindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityBindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityBindingIDInsensitively parses 'input' case-insensitively into a IdentityBindingId
// note: this method should only be used for API response data and not user input
func ParseIdentityBindingIDInsensitively(input string) (*IdentityBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityBindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityBindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityBindingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedClusterName, ok = input.Parsed["managedClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", input)
	}

	if id.IdentityBindingName, ok = input.Parsed["identityBindingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityBindingName", input)
	}

	return nil
}

// ValidateIdentityBindingID checks that 'input' can be parsed as a Identity Binding ID
func ValidateIdentityBindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityBindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Binding ID
func (id IdentityBindingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerService/managedClusters/%s/identityBindings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedClusterName, id.IdentityBindingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Binding ID
func (id IdentityBindingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftContainerService", "Microsoft.ContainerService", "Microsoft.ContainerService"),
		resourceids.StaticSegment("staticManagedClusters", "managedClusters", "managedClusters"),
		resourceids.UserSpecifiedSegment("managedClusterName", "managedClusterName"),
		resourceids.StaticSegment("staticIdentityBindings", "identityBindings", "identityBindings"),
		resourceids.UserSpecifiedSegment("identityBindingName", "identityBindingName"),
	}
}

// String returns a human-readable description of this Identity Binding ID
func (id IdentityBindingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Cluster Name: %q", id.ManagedClusterName),
		fmt.Sprintf("Identity Binding Name: %q", id.IdentityBindingName),
	}
	return fmt.Sprintf("Identity Binding (%s)", strings.Join(components, "\n"))
}
