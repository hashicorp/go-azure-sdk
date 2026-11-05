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
	recaser.RegisterResourceId(&ManagedEnvironmentPrivateLinkResourceId{})
}

var _ resourceids.ResourceId = &ManagedEnvironmentPrivateLinkResourceId{}

// ManagedEnvironmentPrivateLinkResourceId is a struct representing the Resource ID for a Managed Environment Private Link Resource
type ManagedEnvironmentPrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedEnvironmentName  string
	PrivateLinkResourceName string
}

// NewManagedEnvironmentPrivateLinkResourceID returns a new ManagedEnvironmentPrivateLinkResourceId struct
func NewManagedEnvironmentPrivateLinkResourceID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, privateLinkResourceName string) ManagedEnvironmentPrivateLinkResourceId {
	return ManagedEnvironmentPrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedEnvironmentName:  managedEnvironmentName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParseManagedEnvironmentPrivateLinkResourceID parses 'input' into a ManagedEnvironmentPrivateLinkResourceId
func ParseManagedEnvironmentPrivateLinkResourceID(input string) (*ManagedEnvironmentPrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedEnvironmentPrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedEnvironmentPrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedEnvironmentPrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a ManagedEnvironmentPrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParseManagedEnvironmentPrivateLinkResourceIDInsensitively(input string) (*ManagedEnvironmentPrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedEnvironmentPrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedEnvironmentPrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedEnvironmentPrivateLinkResourceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PrivateLinkResourceName, ok = input.Parsed["privateLinkResourceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateLinkResourceName", input)
	}

	return nil
}

// ValidateManagedEnvironmentPrivateLinkResourceID checks that 'input' can be parsed as a Managed Environment Private Link Resource ID
func ValidateManagedEnvironmentPrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedEnvironmentPrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Environment Private Link Resource ID
func (id ManagedEnvironmentPrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Environment Private Link Resource ID
func (id ManagedEnvironmentPrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentName"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "privateLinkResourceName"),
	}
}

// String returns a human-readable description of this Managed Environment Private Link Resource ID
func (id ManagedEnvironmentPrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Managed Environment Private Link Resource (%s)", strings.Join(components, "\n"))
}
