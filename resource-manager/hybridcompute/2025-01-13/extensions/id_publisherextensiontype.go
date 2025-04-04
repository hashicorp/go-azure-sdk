package extensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PublisherExtensionTypeId{})
}

var _ resourceids.ResourceId = &PublisherExtensionTypeId{}

// PublisherExtensionTypeId is a struct representing the Resource ID for a Publisher Extension Type
type PublisherExtensionTypeId struct {
	SubscriptionId    string
	LocationName      string
	PublisherName     string
	ExtensionTypeName string
}

// NewPublisherExtensionTypeID returns a new PublisherExtensionTypeId struct
func NewPublisherExtensionTypeID(subscriptionId string, locationName string, publisherName string, extensionTypeName string) PublisherExtensionTypeId {
	return PublisherExtensionTypeId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		PublisherName:     publisherName,
		ExtensionTypeName: extensionTypeName,
	}
}

// ParsePublisherExtensionTypeID parses 'input' into a PublisherExtensionTypeId
func ParsePublisherExtensionTypeID(input string) (*PublisherExtensionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PublisherExtensionTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublisherExtensionTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePublisherExtensionTypeIDInsensitively parses 'input' case-insensitively into a PublisherExtensionTypeId
// note: this method should only be used for API response data and not user input
func ParsePublisherExtensionTypeIDInsensitively(input string) (*PublisherExtensionTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PublisherExtensionTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublisherExtensionTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PublisherExtensionTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.PublisherName, ok = input.Parsed["publisherName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "publisherName", input)
	}

	if id.ExtensionTypeName, ok = input.Parsed["extensionTypeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionTypeName", input)
	}

	return nil
}

// ValidatePublisherExtensionTypeID checks that 'input' can be parsed as a Publisher Extension Type ID
func ValidatePublisherExtensionTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePublisherExtensionTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Publisher Extension Type ID
func (id PublisherExtensionTypeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.HybridCompute/locations/%s/publishers/%s/extensionTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.PublisherName, id.ExtensionTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Publisher Extension Type ID
func (id PublisherExtensionTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherName"),
		resourceids.StaticSegment("staticExtensionTypes", "extensionTypes", "extensionTypes"),
		resourceids.UserSpecifiedSegment("extensionTypeName", "extensionTypeName"),
	}
}

// String returns a human-readable description of this Publisher Extension Type ID
func (id PublisherExtensionTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Extension Type Name: %q", id.ExtensionTypeName),
	}
	return fmt.Sprintf("Publisher Extension Type (%s)", strings.Join(components, "\n"))
}
