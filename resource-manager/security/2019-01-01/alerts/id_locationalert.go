package alerts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LocationAlertId{})
}

var _ resourceids.ResourceId = &LocationAlertId{}

// LocationAlertId is a struct representing the Resource ID for a Location Alert
type LocationAlertId struct {
	SubscriptionId    string
	ResourceGroupName string
	LocationName      string
	AlertName         string
}

// NewLocationAlertID returns a new LocationAlertId struct
func NewLocationAlertID(subscriptionId string, resourceGroupName string, locationName string, alertName string) LocationAlertId {
	return LocationAlertId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LocationName:      locationName,
		AlertName:         alertName,
	}
}

// ParseLocationAlertID parses 'input' into a LocationAlertId
func ParseLocationAlertID(input string) (*LocationAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationAlertId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationAlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLocationAlertIDInsensitively parses 'input' case-insensitively into a LocationAlertId
// note: this method should only be used for API response data and not user input
func ParseLocationAlertIDInsensitively(input string) (*LocationAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationAlertId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationAlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LocationAlertId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.AlertName, ok = input.Parsed["alertName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertName", input)
	}

	return nil
}

// ValidateLocationAlertID checks that 'input' can be parsed as a Location Alert ID
func ValidateLocationAlertID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationAlertID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location Alert ID
func (id LocationAlertId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/alerts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.AlertName)
}

// Segments returns a slice of Resource ID Segments which comprise this Location Alert ID
func (id LocationAlertId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticAlerts", "alerts", "alerts"),
		resourceids.UserSpecifiedSegment("alertName", "alertValue"),
	}
}

// String returns a human-readable description of this Location Alert ID
func (id LocationAlertId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Alert Name: %q", id.AlertName),
	}
	return fmt.Sprintf("Location Alert (%s)", strings.Join(components, "\n"))
}
