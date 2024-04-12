package resourcepools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ResourcePoolId{})
}

var _ resourceids.ResourceId = &ResourcePoolId{}

// ResourcePoolId is a struct representing the Resource ID for a Resource Pool
type ResourcePoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourcePoolName  string
}

// NewResourcePoolID returns a new ResourcePoolId struct
func NewResourcePoolID(subscriptionId string, resourceGroupName string, resourcePoolName string) ResourcePoolId {
	return ResourcePoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourcePoolName:  resourcePoolName,
	}
}

// ParseResourcePoolID parses 'input' into a ResourcePoolId
func ParseResourcePoolID(input string) (*ResourcePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourcePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourcePoolId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResourcePoolIDInsensitively parses 'input' case-insensitively into a ResourcePoolId
// note: this method should only be used for API response data and not user input
func ParseResourcePoolIDInsensitively(input string) (*ResourcePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourcePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourcePoolId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResourcePoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ResourcePoolName, ok = input.Parsed["resourcePoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourcePoolName", input)
	}

	return nil
}

// ValidateResourcePoolID checks that 'input' can be parsed as a Resource Pool ID
func ValidateResourcePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourcePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Pool ID
func (id ResourcePoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/resourcePools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourcePoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Pool ID
func (id ResourcePoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticResourcePools", "resourcePools", "resourcePools"),
		resourceids.UserSpecifiedSegment("resourcePoolName", "resourcePoolValue"),
	}
}

// String returns a human-readable description of this Resource Pool ID
func (id ResourcePoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Pool Name: %q", id.ResourcePoolName),
	}
	return fmt.Sprintf("Resource Pool (%s)", strings.Join(components, "\n"))
}
