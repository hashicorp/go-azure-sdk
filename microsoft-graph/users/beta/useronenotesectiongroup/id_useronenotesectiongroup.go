package useronenotesectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupId{}

// UserOnenoteSectionGroupId is a struct representing the Resource ID for a User Onenote Section Group
type UserOnenoteSectionGroupId struct {
	UserId         string
	SectionGroupId string
}

// NewUserOnenoteSectionGroupID returns a new UserOnenoteSectionGroupId struct
func NewUserOnenoteSectionGroupID(userId string, sectionGroupId string) UserOnenoteSectionGroupId {
	return UserOnenoteSectionGroupId{
		UserId:         userId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseUserOnenoteSectionGroupID parses 'input' into a UserOnenoteSectionGroupId
func ParseUserOnenoteSectionGroupID(input string) (*UserOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a UserOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteSectionGroupIDInsensitively(input string) (*UserOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteSectionGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteSectionGroupID checks that 'input' can be parsed as a User Onenote Section Group ID
func ValidateUserOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Section Group ID
func (id UserOnenoteSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Section Group ID
func (id UserOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Section Group ID
func (id UserOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("User Onenote Section Group (%s)", strings.Join(components, "\n"))
}
