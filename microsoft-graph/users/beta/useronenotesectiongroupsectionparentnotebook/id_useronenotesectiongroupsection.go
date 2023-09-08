package useronenotesectiongroupsectionparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupSectionId{}

// UserOnenoteSectionGroupSectionId is a struct representing the Resource ID for a User Onenote Section Group Section
type UserOnenoteSectionGroupSectionId struct {
	UserId           string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewUserOnenoteSectionGroupSectionID returns a new UserOnenoteSectionGroupSectionId struct
func NewUserOnenoteSectionGroupSectionID(userId string, sectionGroupId string, onenoteSectionId string) UserOnenoteSectionGroupSectionId {
	return UserOnenoteSectionGroupSectionId{
		UserId:           userId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserOnenoteSectionGroupSectionID parses 'input' into a UserOnenoteSectionGroupSectionId
func ParseUserOnenoteSectionGroupSectionID(input string) (*UserOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionGroupSectionIDInsensitively(input string) (*UserOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionGroupSectionID checks that 'input' can be parsed as a User Onenote Section Group Section ID
func ValidateUserOnenoteSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section Group Section ID
func (id UserOnenoteSectionGroupSectionId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section Group Section ID
func (id UserOnenoteSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Section Group Section ID
func (id UserOnenoteSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Onenote Section Group Section (%s)", strings.Join(components, "\n"))
}
