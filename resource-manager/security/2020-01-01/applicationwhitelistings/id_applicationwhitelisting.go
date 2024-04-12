package applicationwhitelistings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ApplicationWhitelistingId{})
}

var _ resourceids.ResourceId = &ApplicationWhitelistingId{}

// ApplicationWhitelistingId is a struct representing the Resource ID for a Application Whitelisting
type ApplicationWhitelistingId struct {
	SubscriptionId              string
	LocationName                string
	ApplicationWhitelistingName string
}

// NewApplicationWhitelistingID returns a new ApplicationWhitelistingId struct
func NewApplicationWhitelistingID(subscriptionId string, locationName string, applicationWhitelistingName string) ApplicationWhitelistingId {
	return ApplicationWhitelistingId{
		SubscriptionId:              subscriptionId,
		LocationName:                locationName,
		ApplicationWhitelistingName: applicationWhitelistingName,
	}
}

// ParseApplicationWhitelistingID parses 'input' into a ApplicationWhitelistingId
func ParseApplicationWhitelistingID(input string) (*ApplicationWhitelistingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationWhitelistingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationWhitelistingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationWhitelistingIDInsensitively parses 'input' case-insensitively into a ApplicationWhitelistingId
// note: this method should only be used for API response data and not user input
func ParseApplicationWhitelistingIDInsensitively(input string) (*ApplicationWhitelistingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationWhitelistingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationWhitelistingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationWhitelistingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.ApplicationWhitelistingName, ok = input.Parsed["applicationWhitelistingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationWhitelistingName", input)
	}

	return nil
}

// ValidateApplicationWhitelistingID checks that 'input' can be parsed as a Application Whitelisting ID
func ValidateApplicationWhitelistingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationWhitelistingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Whitelisting ID
func (id ApplicationWhitelistingId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/locations/%s/applicationWhitelistings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.ApplicationWhitelistingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Whitelisting ID
func (id ApplicationWhitelistingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticApplicationWhitelistings", "applicationWhitelistings", "applicationWhitelistings"),
		resourceids.UserSpecifiedSegment("applicationWhitelistingName", "applicationWhitelistingValue"),
	}
}

// String returns a human-readable description of this Application Whitelisting ID
func (id ApplicationWhitelistingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Application Whitelisting Name: %q", id.ApplicationWhitelistingName),
	}
	return fmt.Sprintf("Application Whitelisting (%s)", strings.Join(components, "\n"))
}
