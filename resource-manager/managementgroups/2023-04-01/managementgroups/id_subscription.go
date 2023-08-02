package managementgroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SubscriptionId{}

// SubscriptionId is a struct representing the Resource ID for a Subscription
type SubscriptionId struct {
	GroupId        string
	SubscriptionId string
}

// NewSubscriptionID returns a new SubscriptionId struct
func NewSubscriptionID(groupId string, subscriptionId string) SubscriptionId {
	return SubscriptionId{
		GroupId:        groupId,
		SubscriptionId: subscriptionId,
	}
}

// ParseSubscriptionID parses 'input' into a SubscriptionId
func ParseSubscriptionID(input string) (*SubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SubscriptionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	return &id, nil
}

// ParseSubscriptionIDInsensitively parses 'input' case-insensitively into a SubscriptionId
// note: this method should only be used for API response data and not user input
func ParseSubscriptionIDInsensitively(input string) (*SubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SubscriptionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	return &id, nil
}

// ValidateSubscriptionID checks that 'input' can be parsed as a Subscription ID
func ValidateSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Subscription ID
func (id SubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Subscription ID
func (id SubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
	}
}

// String returns a human-readable description of this Subscription ID
func (id SubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Subscription (%s)", strings.Join(components, "\n"))
}
