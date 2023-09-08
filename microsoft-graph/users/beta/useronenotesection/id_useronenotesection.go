package useronenotesection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionId{}

// UserOnenoteSectionId is a struct representing the Resource ID for a User Onenote Section
type UserOnenoteSectionId struct {
	UserId           string
	OnenoteSectionId string
}

// NewUserOnenoteSectionID returns a new UserOnenoteSectionId struct
func NewUserOnenoteSectionID(userId string, onenoteSectionId string) UserOnenoteSectionId {
	return UserOnenoteSectionId{
		UserId:           userId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserOnenoteSectionID parses 'input' into a UserOnenoteSectionId
func ParseUserOnenoteSectionID(input string) (*UserOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionIDInsensitively(input string) (*UserOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionID checks that 'input' can be parsed as a User Onenote Section ID
func ValidateUserOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section ID
func (id UserOnenoteSectionId) ID() string {
	fmtString := "/users/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section ID
func (id UserOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Section ID
func (id UserOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Onenote Section (%s)", strings.Join(components, "\n"))
}
