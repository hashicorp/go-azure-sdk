package deletedwebapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LocationDeletedSiteId{})
}

var _ resourceids.ResourceId = &LocationDeletedSiteId{}

// LocationDeletedSiteId is a struct representing the Resource ID for a Location Deleted Site
type LocationDeletedSiteId struct {
	SubscriptionId string
	LocationName   string
	DeletedSiteId  string
}

// NewLocationDeletedSiteID returns a new LocationDeletedSiteId struct
func NewLocationDeletedSiteID(subscriptionId string, locationName string, deletedSiteId string) LocationDeletedSiteId {
	return LocationDeletedSiteId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		DeletedSiteId:  deletedSiteId,
	}
}

// ParseLocationDeletedSiteID parses 'input' into a LocationDeletedSiteId
func ParseLocationDeletedSiteID(input string) (*LocationDeletedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationDeletedSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationDeletedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLocationDeletedSiteIDInsensitively parses 'input' case-insensitively into a LocationDeletedSiteId
// note: this method should only be used for API response data and not user input
func ParseLocationDeletedSiteIDInsensitively(input string) (*LocationDeletedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationDeletedSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationDeletedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LocationDeletedSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.DeletedSiteId, ok = input.Parsed["deletedSiteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deletedSiteId", input)
	}

	return nil
}

// ValidateLocationDeletedSiteID checks that 'input' can be parsed as a Location Deleted Site ID
func ValidateLocationDeletedSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationDeletedSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location Deleted Site ID
func (id LocationDeletedSiteId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/locations/%s/deletedSites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.DeletedSiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Location Deleted Site ID
func (id LocationDeletedSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticDeletedSites", "deletedSites", "deletedSites"),
		resourceids.UserSpecifiedSegment("deletedSiteId", "deletedSiteIdValue"),
	}
}

// String returns a human-readable description of this Location Deleted Site ID
func (id LocationDeletedSiteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Deleted Site: %q", id.DeletedSiteId),
	}
	return fmt.Sprintf("Location Deleted Site (%s)", strings.Join(components, "\n"))
}
