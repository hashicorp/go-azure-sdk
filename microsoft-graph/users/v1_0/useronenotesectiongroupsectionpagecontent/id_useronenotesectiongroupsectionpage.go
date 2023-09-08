package useronenotesectiongroupsectionpagecontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupSectionPageId{}

// UserOnenoteSectionGroupSectionPageId is a struct representing the Resource ID for a User Onenote Section Group Section Page
type UserOnenoteSectionGroupSectionPageId struct {
	UserId           string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserOnenoteSectionGroupSectionPageID returns a new UserOnenoteSectionGroupSectionPageId struct
func NewUserOnenoteSectionGroupSectionPageID(userId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) UserOnenoteSectionGroupSectionPageId {
	return UserOnenoteSectionGroupSectionPageId{
		UserId:           userId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserOnenoteSectionGroupSectionPageID parses 'input' into a UserOnenoteSectionGroupSectionPageId
func ParseUserOnenoteSectionGroupSectionPageID(input string) (*UserOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionGroupSectionPageIDInsensitively(input string) (*UserOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionPageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionGroupSectionPageID checks that 'input' can be parsed as a User Onenote Section Group Section Page ID
func ValidateUserOnenoteSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section Group Section Page ID
func (id UserOnenoteSectionGroupSectionPageId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section Group Section Page ID
func (id UserOnenoteSectionGroupSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Section Group Section Page ID
func (id UserOnenoteSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Onenote Section Group Section Page (%s)", strings.Join(components, "\n"))
}
