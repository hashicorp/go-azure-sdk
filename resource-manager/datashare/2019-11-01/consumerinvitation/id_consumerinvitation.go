package consumerinvitation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConsumerInvitationId{})
}

var _ resourceids.ResourceId = &ConsumerInvitationId{}

// ConsumerInvitationId is a struct representing the Resource ID for a Consumer Invitation
type ConsumerInvitationId struct {
	LocationName string
	InvitationId string
}

// NewConsumerInvitationID returns a new ConsumerInvitationId struct
func NewConsumerInvitationID(locationName string, invitationId string) ConsumerInvitationId {
	return ConsumerInvitationId{
		LocationName: locationName,
		InvitationId: invitationId,
	}
}

// ParseConsumerInvitationID parses 'input' into a ConsumerInvitationId
func ParseConsumerInvitationID(input string) (*ConsumerInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConsumerInvitationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConsumerInvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConsumerInvitationIDInsensitively parses 'input' case-insensitively into a ConsumerInvitationId
// note: this method should only be used for API response data and not user input
func ParseConsumerInvitationIDInsensitively(input string) (*ConsumerInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConsumerInvitationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConsumerInvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConsumerInvitationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.InvitationId, ok = input.Parsed["invitationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invitationId", input)
	}

	return nil
}

// ValidateConsumerInvitationID checks that 'input' can be parsed as a Consumer Invitation ID
func ValidateConsumerInvitationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConsumerInvitationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Consumer Invitation ID
func (id ConsumerInvitationId) ID() string {
	fmtString := "/providers/Microsoft.DataShare/locations/%s/consumerInvitations/%s"
	return fmt.Sprintf(fmtString, id.LocationName, id.InvitationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Consumer Invitation ID
func (id ConsumerInvitationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticConsumerInvitations", "consumerInvitations", "consumerInvitations"),
		resourceids.UserSpecifiedSegment("invitationId", "invitationId"),
	}
}

// String returns a human-readable description of this Consumer Invitation ID
func (id ConsumerInvitationId) String() string {
	components := []string{
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Invitation: %q", id.InvitationId),
	}
	return fmt.Sprintf("Consumer Invitation (%s)", strings.Join(components, "\n"))
}
