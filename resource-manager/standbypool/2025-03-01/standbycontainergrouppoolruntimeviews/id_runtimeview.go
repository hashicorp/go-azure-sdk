package standbycontainergrouppoolruntimeviews

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RuntimeViewId{})
}

var _ resourceids.ResourceId = &RuntimeViewId{}

// RuntimeViewId is a struct representing the Resource ID for a Runtime View
type RuntimeViewId struct {
	SubscriptionId                string
	ResourceGroupName             string
	StandbyContainerGroupPoolName string
	RuntimeViewName               string
}

// NewRuntimeViewID returns a new RuntimeViewId struct
func NewRuntimeViewID(subscriptionId string, resourceGroupName string, standbyContainerGroupPoolName string, runtimeViewName string) RuntimeViewId {
	return RuntimeViewId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		StandbyContainerGroupPoolName: standbyContainerGroupPoolName,
		RuntimeViewName:               runtimeViewName,
	}
}

// ParseRuntimeViewID parses 'input' into a RuntimeViewId
func ParseRuntimeViewID(input string) (*RuntimeViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RuntimeViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RuntimeViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRuntimeViewIDInsensitively parses 'input' case-insensitively into a RuntimeViewId
// note: this method should only be used for API response data and not user input
func ParseRuntimeViewIDInsensitively(input string) (*RuntimeViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RuntimeViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RuntimeViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RuntimeViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.StandbyContainerGroupPoolName, ok = input.Parsed["standbyContainerGroupPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "standbyContainerGroupPoolName", input)
	}

	if id.RuntimeViewName, ok = input.Parsed["runtimeViewName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runtimeViewName", input)
	}

	return nil
}

// ValidateRuntimeViewID checks that 'input' can be parsed as a Runtime View ID
func ValidateRuntimeViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRuntimeViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Runtime View ID
func (id RuntimeViewId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StandbyPool/standbyContainerGroupPools/%s/runtimeViews/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StandbyContainerGroupPoolName, id.RuntimeViewName)
}

// Segments returns a slice of Resource ID Segments which comprise this Runtime View ID
func (id RuntimeViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStandbyPool", "Microsoft.StandbyPool", "Microsoft.StandbyPool"),
		resourceids.StaticSegment("staticStandbyContainerGroupPools", "standbyContainerGroupPools", "standbyContainerGroupPools"),
		resourceids.UserSpecifiedSegment("standbyContainerGroupPoolName", "standbyContainerGroupPoolName"),
		resourceids.StaticSegment("staticRuntimeViews", "runtimeViews", "runtimeViews"),
		resourceids.UserSpecifiedSegment("runtimeViewName", "runtimeViewName"),
	}
}

// String returns a human-readable description of this Runtime View ID
func (id RuntimeViewId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Standby Container Group Pool Name: %q", id.StandbyContainerGroupPoolName),
		fmt.Sprintf("Runtime View Name: %q", id.RuntimeViewName),
	}
	return fmt.Sprintf("Runtime View (%s)", strings.Join(components, "\n"))
}
