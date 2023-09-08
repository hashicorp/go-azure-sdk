package meprofilewebsite

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfileWebsiteId{}

// MeProfileWebsiteId is a struct representing the Resource ID for a Me Profile Website
type MeProfileWebsiteId struct {
	PersonWebsiteId string
}

// NewMeProfileWebsiteID returns a new MeProfileWebsiteId struct
func NewMeProfileWebsiteID(personWebsiteId string) MeProfileWebsiteId {
	return MeProfileWebsiteId{
		PersonWebsiteId: personWebsiteId,
	}
}

// ParseMeProfileWebsiteID parses 'input' into a MeProfileWebsiteId
func ParseMeProfileWebsiteID(input string) (*MeProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileWebsiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileWebsiteId{}

	if id.PersonWebsiteId, ok = parsed.Parsed["personWebsiteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personWebsiteId", *parsed)
	}

	return &id, nil
}

// ParseMeProfileWebsiteIDInsensitively parses 'input' case-insensitively into a MeProfileWebsiteId
// note: this method should only be used for API response data and not user input
func ParseMeProfileWebsiteIDInsensitively(input string) (*MeProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileWebsiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileWebsiteId{}

	if id.PersonWebsiteId, ok = parsed.Parsed["personWebsiteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personWebsiteId", *parsed)
	}

	return &id, nil
}

// ValidateMeProfileWebsiteID checks that 'input' can be parsed as a Me Profile Website ID
func ValidateMeProfileWebsiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileWebsiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Website ID
func (id MeProfileWebsiteId) ID() string {
	fmtString := "/me/profile/websites/%s"
	return fmt.Sprintf(fmtString, id.PersonWebsiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Website ID
func (id MeProfileWebsiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticWebsites", "websites", "websites"),
		resourceids.UserSpecifiedSegment("personWebsiteId", "personWebsiteIdValue"),
	}
}

// String returns a human-readable description of this Me Profile Website ID
func (id MeProfileWebsiteId) String() string {
	components := []string{
		fmt.Sprintf("Person Website: %q", id.PersonWebsiteId),
	}
	return fmt.Sprintf("Me Profile Website (%s)", strings.Join(components, "\n"))
}
