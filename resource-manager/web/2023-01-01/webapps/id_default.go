package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DefaultId{}

// DefaultId is a struct representing the Resource ID for a Default
type DefaultId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	DefaultName       string
	KeyName           string
}

// NewDefaultID returns a new DefaultId struct
func NewDefaultID(subscriptionId string, resourceGroupName string, siteName string, defaultName string, keyName string) DefaultId {
	return DefaultId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		DefaultName:       defaultName,
		KeyName:           keyName,
	}
}

// ParseDefaultID parses 'input' into a DefaultId
func ParseDefaultID(input string) (*DefaultId, error) {
	parser := resourceids.NewParserFromResourceIdType(DefaultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DefaultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.DefaultName, ok = parsed.Parsed["defaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultName", *parsed)
	}

	if id.KeyName, ok = parsed.Parsed["keyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "keyName", *parsed)
	}

	return &id, nil
}

// ParseDefaultIDInsensitively parses 'input' case-insensitively into a DefaultId
// note: this method should only be used for API response data and not user input
func ParseDefaultIDInsensitively(input string) (*DefaultId, error) {
	parser := resourceids.NewParserFromResourceIdType(DefaultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DefaultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.DefaultName, ok = parsed.Parsed["defaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultName", *parsed)
	}

	if id.KeyName, ok = parsed.Parsed["keyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "keyName", *parsed)
	}

	return &id, nil
}

// ValidateDefaultID checks that 'input' can be parsed as a Default ID
func ValidateDefaultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDefaultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Default ID
func (id DefaultId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/host/default/%s/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.DefaultName, id.KeyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Default ID
func (id DefaultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticHost", "host", "host"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.UserSpecifiedSegment("defaultName", "defaultValue"),
		resourceids.UserSpecifiedSegment("keyName", "keyValue"),
	}
}

// String returns a human-readable description of this Default ID
func (id DefaultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Default Name: %q", id.DefaultName),
		fmt.Sprintf("Key Name: %q", id.KeyName),
	}
	return fmt.Sprintf("Default (%s)", strings.Join(components, "\n"))
}
