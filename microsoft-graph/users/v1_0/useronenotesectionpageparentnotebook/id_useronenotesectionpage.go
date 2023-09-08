package useronenotesectionpageparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionPageId{}

// UserOnenoteSectionPageId is a struct representing the Resource ID for a User Onenote Section Page
type UserOnenoteSectionPageId struct {
	UserId           string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserOnenoteSectionPageID returns a new UserOnenoteSectionPageId struct
func NewUserOnenoteSectionPageID(userId string, onenoteSectionId string, onenotePageId string) UserOnenoteSectionPageId {
	return UserOnenoteSectionPageId{
		UserId:           userId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserOnenoteSectionPageID parses 'input' into a UserOnenoteSectionPageId
func ParseUserOnenoteSectionPageID(input string) (*UserOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionPageIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionPageId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionPageIDInsensitively(input string) (*UserOnenoteSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionPageID checks that 'input' can be parsed as a User Onenote Section Page ID
func ValidateUserOnenoteSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section Page ID
func (id UserOnenoteSectionPageId) ID() string {
	fmtString := "/users/%s/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section Page ID
func (id UserOnenoteSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Section Page ID
func (id UserOnenoteSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Onenote Section Page (%s)", strings.Join(components, "\n"))
}
