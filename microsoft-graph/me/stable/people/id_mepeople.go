package people

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePeopleId{}

// MePeopleId is a struct representing the Resource ID for a Me People
type MePeopleId struct {
	PersonId string
}

// NewMePeopleID returns a new MePeopleId struct
func NewMePeopleID(personId string) MePeopleId {
	return MePeopleId{
		PersonId: personId,
	}
}

// ParseMePeopleID parses 'input' into a MePeopleId
func ParseMePeopleID(input string) (*MePeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePeopleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePeopleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePeopleIDInsensitively parses 'input' case-insensitively into a MePeopleId
// note: this method should only be used for API response data and not user input
func ParseMePeopleIDInsensitively(input string) (*MePeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePeopleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePeopleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePeopleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonId, ok = input.Parsed["personId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personId", input)
	}

	return nil
}

// ValidateMePeopleID checks that 'input' can be parsed as a Me People ID
func ValidateMePeopleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePeopleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me People ID
func (id MePeopleId) ID() string {
	fmtString := "/me/people/%s"
	return fmt.Sprintf(fmtString, id.PersonId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me People ID
func (id MePeopleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("people", "people", "people"),
		resourceids.UserSpecifiedSegment("personId", "personIdValue"),
	}
}

// String returns a human-readable description of this Me People ID
func (id MePeopleId) String() string {
	components := []string{
		fmt.Sprintf("Person: %q", id.PersonId),
	}
	return fmt.Sprintf("Me People (%s)", strings.Join(components, "\n"))
}
