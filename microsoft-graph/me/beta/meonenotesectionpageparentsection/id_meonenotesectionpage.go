package meonenotesectionpageparentsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteSectionPageId{}

// MeOnenoteSectionPageId is a struct representing the Resource ID for a Me Onenote Section Page
type MeOnenoteSectionPageId struct {
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteSectionPageID returns a new MeOnenoteSectionPageId struct
func NewMeOnenoteSectionPageID(onenoteSectionId string, onenotePageId string) MeOnenoteSectionPageId {
	return MeOnenoteSectionPageId{
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteSectionPageID parses 'input' into a MeOnenoteSectionPageId
func ParseMeOnenoteSectionPageID(input string) (*MeOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionPageId{}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteSectionPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionPageIDInsensitively(input string) (*MeOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionPageId{}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteSectionPageID checks that 'input' can be parsed as a Me Onenote Section Page ID
func ValidateMeOnenoteSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Page ID
func (id MeOnenoteSectionPageId) ID() string {
	fmtString := "/me/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Page ID
func (id MeOnenoteSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Section Page ID
func (id MeOnenoteSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Section Page (%s)", strings.Join(components, "\n"))
}
