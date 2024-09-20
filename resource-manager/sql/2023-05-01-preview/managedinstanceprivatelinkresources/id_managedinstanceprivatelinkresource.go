package managedinstanceprivatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagedInstancePrivateLinkResourceId{})
}

var _ resourceids.ResourceId = &ManagedInstancePrivateLinkResourceId{}

// ManagedInstancePrivateLinkResourceId is a struct representing the Resource ID for a Managed Instance Private Link Resource
type ManagedInstancePrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedInstanceName     string
	PrivateLinkResourceName string
}

// NewManagedInstancePrivateLinkResourceID returns a new ManagedInstancePrivateLinkResourceId struct
func NewManagedInstancePrivateLinkResourceID(subscriptionId string, resourceGroupName string, managedInstanceName string, privateLinkResourceName string) ManagedInstancePrivateLinkResourceId {
	return ManagedInstancePrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedInstanceName:     managedInstanceName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParseManagedInstancePrivateLinkResourceID parses 'input' into a ManagedInstancePrivateLinkResourceId
func ParseManagedInstancePrivateLinkResourceID(input string) (*ManagedInstancePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstancePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstancePrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedInstancePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a ManagedInstancePrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParseManagedInstancePrivateLinkResourceIDInsensitively(input string) (*ManagedInstancePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstancePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstancePrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedInstancePrivateLinkResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.PrivateLinkResourceName, ok = input.Parsed["privateLinkResourceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateLinkResourceName", input)
	}

	return nil
}

// ValidateManagedInstancePrivateLinkResourceID checks that 'input' can be parsed as a Managed Instance Private Link Resource ID
func ValidateManagedInstancePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstancePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceName"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "groupName"),
	}
}

// String returns a human-readable description of this Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Managed Instance Private Link Resource (%s)", strings.Join(components, "\n"))
}
