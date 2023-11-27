package smartgroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SmartGroupId{}

// SmartGroupId is a struct representing the Resource ID for a Smart Group
type SmartGroupId struct {
	SubscriptionId string
	SmartGroupId   string
}

// NewSmartGroupID returns a new SmartGroupId struct
func NewSmartGroupID(subscriptionId string, smartGroupId string) SmartGroupId {
	return SmartGroupId{
		SubscriptionId: subscriptionId,
		SmartGroupId:   smartGroupId,
	}
}

// ParseSmartGroupID parses 'input' into a SmartGroupId
func ParseSmartGroupID(input string) (*SmartGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(SmartGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SmartGroupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSmartGroupIDInsensitively parses 'input' case-insensitively into a SmartGroupId
// note: this method should only be used for API response data and not user input
func ParseSmartGroupIDInsensitively(input string) (*SmartGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(SmartGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SmartGroupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SmartGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.SmartGroupId, ok = input.Parsed["smartGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "smartGroupId", input)
	}

	return nil
}

// ValidateSmartGroupID checks that 'input' can be parsed as a Smart Group ID
func ValidateSmartGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSmartGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Smart Group ID
func (id SmartGroupId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.AlertsManagement/smartGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.SmartGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Smart Group ID
func (id SmartGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAlertsManagement", "Microsoft.AlertsManagement", "Microsoft.AlertsManagement"),
		resourceids.StaticSegment("staticSmartGroups", "smartGroups", "smartGroups"),
		resourceids.UserSpecifiedSegment("smartGroupId", "smartGroupIdValue"),
	}
}

// String returns a human-readable description of this Smart Group ID
func (id SmartGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Smart Group: %q", id.SmartGroupId),
	}
	return fmt.Sprintf("Smart Group (%s)", strings.Join(components, "\n"))
}
