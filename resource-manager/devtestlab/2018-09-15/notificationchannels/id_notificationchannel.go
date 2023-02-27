// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationchannels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = NotificationChannelId{}

// NotificationChannelId is a struct representing the Resource ID for a Notification Channel
type NotificationChannelId struct {
	SubscriptionId          string
	ResourceGroupName       string
	LabName                 string
	NotificationChannelName string
}

// NewNotificationChannelID returns a new NotificationChannelId struct
func NewNotificationChannelID(subscriptionId string, resourceGroupName string, labName string, notificationChannelName string) NotificationChannelId {
	return NotificationChannelId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		LabName:                 labName,
		NotificationChannelName: notificationChannelName,
	}
}

// ParseNotificationChannelID parses 'input' into a NotificationChannelId
func ParseNotificationChannelID(input string) (*NotificationChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(NotificationChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NotificationChannelId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.NotificationChannelName, ok = parsed.Parsed["notificationChannelName"]; !ok {
		return nil, fmt.Errorf("the segment 'notificationChannelName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseNotificationChannelIDInsensitively parses 'input' case-insensitively into a NotificationChannelId
// note: this method should only be used for API response data and not user input
func ParseNotificationChannelIDInsensitively(input string) (*NotificationChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(NotificationChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NotificationChannelId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.NotificationChannelName, ok = parsed.Parsed["notificationChannelName"]; !ok {
		return nil, fmt.Errorf("the segment 'notificationChannelName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateNotificationChannelID checks that 'input' can be parsed as a Notification Channel ID
func ValidateNotificationChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNotificationChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Notification Channel ID
func (id NotificationChannelId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/notificationChannels/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.NotificationChannelName)
}

// Segments returns a slice of Resource ID Segments which comprise this Notification Channel ID
func (id NotificationChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticNotificationChannels", "notificationChannels", "notificationChannels"),
		resourceids.UserSpecifiedSegment("notificationChannelName", "notificationChannelValue"),
	}
}

// String returns a human-readable description of this Notification Channel ID
func (id NotificationChannelId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Notification Channel Name: %q", id.NotificationChannelName),
	}
	return fmt.Sprintf("Notification Channel (%s)", strings.Join(components, "\n"))
}
