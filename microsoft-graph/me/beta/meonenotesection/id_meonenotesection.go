package meonenotesection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteSectionId{}

// MeOnenoteSectionId is a struct representing the Resource ID for a Me Onenote Section
type MeOnenoteSectionId struct {
	OnenoteSectionId string
}

// NewMeOnenoteSectionID returns a new MeOnenoteSectionId struct
func NewMeOnenoteSectionID(onenoteSectionId string) MeOnenoteSectionId {
	return MeOnenoteSectionId{
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteSectionID parses 'input' into a MeOnenoteSectionId
func ParseMeOnenoteSectionID(input string) (*MeOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionId{}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionIDInsensitively(input string) (*MeOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionId{}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteSectionID checks that 'input' can be parsed as a Me Onenote Section ID
func ValidateMeOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section ID
func (id MeOnenoteSectionId) ID() string {
	fmtString := "/me/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section ID
func (id MeOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Section ID
func (id MeOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Section (%s)", strings.Join(components, "\n"))
}
