package meprofilepublication

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfilePublicationId{}

// MeProfilePublicationId is a struct representing the Resource ID for a Me Profile Publication
type MeProfilePublicationId struct {
	ItemPublicationId string
}

// NewMeProfilePublicationID returns a new MeProfilePublicationId struct
func NewMeProfilePublicationID(itemPublicationId string) MeProfilePublicationId {
	return MeProfilePublicationId{
		ItemPublicationId: itemPublicationId,
	}
}

// ParseMeProfilePublicationID parses 'input' into a MeProfilePublicationId
func ParseMeProfilePublicationID(input string) (*MeProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfilePublicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfilePublicationId{}

	if id.ItemPublicationId, ok = parsed.Parsed["itemPublicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPublicationId", *parsed)
	}

	return &id, nil
}

// ParseMeProfilePublicationIDInsensitively parses 'input' case-insensitively into a MeProfilePublicationId
// note: this method should only be used for API response data and not user input
func ParseMeProfilePublicationIDInsensitively(input string) (*MeProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfilePublicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfilePublicationId{}

	if id.ItemPublicationId, ok = parsed.Parsed["itemPublicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPublicationId", *parsed)
	}

	return &id, nil
}

// ValidateMeProfilePublicationID checks that 'input' can be parsed as a Me Profile Publication ID
func ValidateMeProfilePublicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfilePublicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Publication ID
func (id MeProfilePublicationId) ID() string {
	fmtString := "/me/profile/publications/%s"
	return fmt.Sprintf(fmtString, id.ItemPublicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Publication ID
func (id MeProfilePublicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticPublications", "publications", "publications"),
		resourceids.UserSpecifiedSegment("itemPublicationId", "itemPublicationIdValue"),
	}
}

// String returns a human-readable description of this Me Profile Publication ID
func (id MeProfilePublicationId) String() string {
	components := []string{
		fmt.Sprintf("Item Publication: %q", id.ItemPublicationId),
	}
	return fmt.Sprintf("Me Profile Publication (%s)", strings.Join(components, "\n"))
}
