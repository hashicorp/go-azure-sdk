package standardoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SubscriptionResourceGroupId{})
}

var _ resourceids.ResourceId = &SubscriptionResourceGroupId{}

// SubscriptionResourceGroupId is a struct representing the Resource ID for a Subscription Resource Group
type SubscriptionResourceGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
}

// NewSubscriptionResourceGroupID returns a new SubscriptionResourceGroupId struct
func NewSubscriptionResourceGroupID(subscriptionId string, resourceGroupName string) SubscriptionResourceGroupId {
	return SubscriptionResourceGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
	}
}

// ParseSubscriptionResourceGroupID parses 'input' into a SubscriptionResourceGroupId
func ParseSubscriptionResourceGroupID(input string) (*SubscriptionResourceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionResourceGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionResourceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSubscriptionResourceGroupIDInsensitively parses 'input' case-insensitively into a SubscriptionResourceGroupId
// note: this method should only be used for API response data and not user input
func ParseSubscriptionResourceGroupIDInsensitively(input string) (*SubscriptionResourceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionResourceGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionResourceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SubscriptionResourceGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	return nil
}

// ValidateSubscriptionResourceGroupID checks that 'input' can be parsed as a Subscription Resource Group ID
func ValidateSubscriptionResourceGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubscriptionResourceGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Subscription Resource Group ID
func (id SubscriptionResourceGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Subscription Resource Group ID
func (id SubscriptionResourceGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.UserSpecifiedSegment("resourceGroupName", "resourceGroupName"),
	}
}

// String returns a human-readable description of this Subscription Resource Group ID
func (id SubscriptionResourceGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
	}
	return fmt.Sprintf("Subscription Resource Group (%s)", strings.Join(components, "\n"))
}
