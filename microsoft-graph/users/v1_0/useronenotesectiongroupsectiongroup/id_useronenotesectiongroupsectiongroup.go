package useronenotesectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupSectionGroupId{}

// UserOnenoteSectionGroupSectionGroupId is a struct representing the Resource ID for a User Onenote Section Group Section Group
type UserOnenoteSectionGroupSectionGroupId struct {
	UserId          string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewUserOnenoteSectionGroupSectionGroupID returns a new UserOnenoteSectionGroupSectionGroupId struct
func NewUserOnenoteSectionGroupSectionGroupID(userId string, sectionGroupId string, sectionGroupId1 string) UserOnenoteSectionGroupSectionGroupId {
	return UserOnenoteSectionGroupSectionGroupId{
		UserId:          userId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseUserOnenoteSectionGroupSectionGroupID parses 'input' into a UserOnenoteSectionGroupSectionGroupId
func ParseUserOnenoteSectionGroupSectionGroupID(input string) (*UserOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionGroupSectionGroupIDInsensitively(input string) (*UserOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionGroupSectionGroupID checks that 'input' can be parsed as a User Onenote Section Group Section Group ID
func ValidateUserOnenoteSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section Group Section Group ID
func (id UserOnenoteSectionGroupSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section Group Section Group ID
func (id UserOnenoteSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this User Onenote Section Group Section Group ID
func (id UserOnenoteSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("User Onenote Section Group Section Group (%s)", strings.Join(components, "\n"))
}
