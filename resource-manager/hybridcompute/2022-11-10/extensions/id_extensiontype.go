// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ExtensionTypeId{}

// ExtensionTypeId is a struct representing the Resource ID for a Extension Type
type ExtensionTypeId struct {
	SubscriptionId    string
	LocationName      string
	PublisherName     string
	ExtensionTypeName string
}

// NewExtensionTypeID returns a new ExtensionTypeId struct
func NewExtensionTypeID(subscriptionId string, locationName string, publisherName string, extensionTypeName string) ExtensionTypeId {
	return ExtensionTypeId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		PublisherName:     publisherName,
		ExtensionTypeName: extensionTypeName,
	}
}

// ParseExtensionTypeID parses 'input' into a ExtensionTypeId
func ParseExtensionTypeID(input string) (*ExtensionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExtensionTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExtensionTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.ExtensionTypeName, ok = parsed.Parsed["extensionTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'extensionTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseExtensionTypeIDInsensitively parses 'input' case-insensitively into a ExtensionTypeId
// note: this method should only be used for API response data and not user input
func ParseExtensionTypeIDInsensitively(input string) (*ExtensionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExtensionTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExtensionTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.ExtensionTypeName, ok = parsed.Parsed["extensionTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'extensionTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateExtensionTypeID checks that 'input' can be parsed as a Extension Type ID
func ValidateExtensionTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExtensionTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Extension Type ID
func (id ExtensionTypeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.HybridCompute/locations/%s/publishers/%s/extensionTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.PublisherName, id.ExtensionTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Extension Type ID
func (id ExtensionTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherValue"),
		resourceids.StaticSegment("staticExtensionTypes", "extensionTypes", "extensionTypes"),
		resourceids.UserSpecifiedSegment("extensionTypeName", "extensionTypeValue"),
	}
}

// String returns a human-readable description of this Extension Type ID
func (id ExtensionTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Extension Type Name: %q", id.ExtensionTypeName),
	}
	return fmt.Sprintf("Extension Type (%s)", strings.Join(components, "\n"))
}
