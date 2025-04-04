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
	recaser.RegisterResourceId(&ExtensionTypeVersionId{})
}

var _ resourceids.ResourceId = &ExtensionTypeVersionId{}

// ExtensionTypeVersionId is a struct representing the Resource ID for a Extension Type Version
type ExtensionTypeVersionId struct {
	SubscriptionId    string
	LocationName      string
	PublisherName     string
	ExtensionTypeName string
	VersionName       string
}

// NewExtensionTypeVersionID returns a new ExtensionTypeVersionId struct
func NewExtensionTypeVersionID(subscriptionId string, locationName string, publisherName string, extensionTypeName string, versionName string) ExtensionTypeVersionId {
	return ExtensionTypeVersionId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		PublisherName:     publisherName,
		ExtensionTypeName: extensionTypeName,
		VersionName:       versionName,
	}
}

// ParseExtensionTypeVersionID parses 'input' into a ExtensionTypeVersionId
func ParseExtensionTypeVersionID(input string) (*ExtensionTypeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExtensionTypeVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExtensionTypeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExtensionTypeVersionIDInsensitively parses 'input' case-insensitively into a ExtensionTypeVersionId
// note: this method should only be used for API response data and not user input
func ParseExtensionTypeVersionIDInsensitively(input string) (*ExtensionTypeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExtensionTypeVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExtensionTypeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExtensionTypeVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateExtensionTypeVersionID checks that 'input' can be parsed as a Extension Type Version ID
func ValidateExtensionTypeVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExtensionTypeVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Extension Type Version ID
func (id ExtensionTypeVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.HybridCompute/locations/%s/publishers/%s/extensionTypes/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.PublisherName, id.ExtensionTypeName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Extension Type Version ID
func (id ExtensionTypeVersionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Extension Type Version ID
func (id ExtensionTypeVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Extension Type Name: %q", id.ExtensionTypeName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Extension Type Version (%s)", strings.Join(components, "\n"))
}
