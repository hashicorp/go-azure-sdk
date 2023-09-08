package meprofileemail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfileEmailId{}

// MeProfileEmailId is a struct representing the Resource ID for a Me Profile Email
type MeProfileEmailId struct {
	ItemEmailId string
}

// NewMeProfileEmailID returns a new MeProfileEmailId struct
func NewMeProfileEmailID(itemEmailId string) MeProfileEmailId {
	return MeProfileEmailId{
		ItemEmailId: itemEmailId,
	}
}

// ParseMeProfileEmailID parses 'input' into a MeProfileEmailId
func ParseMeProfileEmailID(input string) (*MeProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileEmailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileEmailId{}

	if id.ItemEmailId, ok = parsed.Parsed["itemEmailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemEmailId", *parsed)
	}

	return &id, nil
}

// ParseMeProfileEmailIDInsensitively parses 'input' case-insensitively into a MeProfileEmailId
// note: this method should only be used for API response data and not user input
func ParseMeProfileEmailIDInsensitively(input string) (*MeProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileEmailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileEmailId{}

	if id.ItemEmailId, ok = parsed.Parsed["itemEmailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemEmailId", *parsed)
	}

	return &id, nil
}

// ValidateMeProfileEmailID checks that 'input' can be parsed as a Me Profile Email ID
func ValidateMeProfileEmailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileEmailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Email ID
func (id MeProfileEmailId) ID() string {
	fmtString := "/me/profile/emails/%s"
	return fmt.Sprintf(fmtString, id.ItemEmailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Email ID
func (id MeProfileEmailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticEmails", "emails", "emails"),
		resourceids.UserSpecifiedSegment("itemEmailId", "itemEmailIdValue"),
	}
}

// String returns a human-readable description of this Me Profile Email ID
func (id MeProfileEmailId) String() string {
	components := []string{
		fmt.Sprintf("Item Email: %q", id.ItemEmailId),
	}
	return fmt.Sprintf("Me Profile Email (%s)", strings.Join(components, "\n"))
}
