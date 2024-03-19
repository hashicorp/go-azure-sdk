package partnerdestinations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PartnerDestinationId{}

// PartnerDestinationId is a struct representing the Resource ID for a Partner Destination
type PartnerDestinationId struct {
	SubscriptionId         string
	ResourceGroupName      string
	PartnerDestinationName string
}

// NewPartnerDestinationID returns a new PartnerDestinationId struct
func NewPartnerDestinationID(subscriptionId string, resourceGroupName string, partnerDestinationName string) PartnerDestinationId {
	return PartnerDestinationId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		PartnerDestinationName: partnerDestinationName,
	}
}

// ParsePartnerDestinationID parses 'input' into a PartnerDestinationId
func ParsePartnerDestinationID(input string) (*PartnerDestinationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PartnerDestinationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PartnerDestinationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePartnerDestinationIDInsensitively parses 'input' case-insensitively into a PartnerDestinationId
// note: this method should only be used for API response data and not user input
func ParsePartnerDestinationIDInsensitively(input string) (*PartnerDestinationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PartnerDestinationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PartnerDestinationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PartnerDestinationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PartnerDestinationName, ok = input.Parsed["partnerDestinationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "partnerDestinationName", input)
	}

	return nil
}

// ValidatePartnerDestinationID checks that 'input' can be parsed as a Partner Destination ID
func ValidatePartnerDestinationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePartnerDestinationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Partner Destination ID
func (id PartnerDestinationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.EventGrid/partnerDestinations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PartnerDestinationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Partner Destination ID
func (id PartnerDestinationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftEventGrid", "Microsoft.EventGrid", "Microsoft.EventGrid"),
		resourceids.StaticSegment("staticPartnerDestinations", "partnerDestinations", "partnerDestinations"),
		resourceids.UserSpecifiedSegment("partnerDestinationName", "partnerDestinationValue"),
	}
}

// String returns a human-readable description of this Partner Destination ID
func (id PartnerDestinationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Partner Destination Name: %q", id.PartnerDestinationName),
	}
	return fmt.Sprintf("Partner Destination (%s)", strings.Join(components, "\n"))
}
