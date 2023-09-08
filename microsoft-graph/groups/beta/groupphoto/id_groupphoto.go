package groupphoto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPhotoId{}

// GroupPhotoId is a struct representing the Resource ID for a Group Photo
type GroupPhotoId struct {
	GroupId        string
	ProfilePhotoId string
}

// NewGroupPhotoID returns a new GroupPhotoId struct
func NewGroupPhotoID(groupId string, profilePhotoId string) GroupPhotoId {
	return GroupPhotoId{
		GroupId:        groupId,
		ProfilePhotoId: profilePhotoId,
	}
}

// ParseGroupPhotoID parses 'input' into a GroupPhotoId
func ParseGroupPhotoID(input string) (*GroupPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPhotoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPhotoId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ProfilePhotoId, ok = parsed.Parsed["profilePhotoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", *parsed)
	}

	return &id, nil
}

// ParseGroupPhotoIDInsensitively parses 'input' case-insensitively into a GroupPhotoId
// note: this method should only be used for API response data and not user input
func ParseGroupPhotoIDInsensitively(input string) (*GroupPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPhotoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPhotoId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ProfilePhotoId, ok = parsed.Parsed["profilePhotoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPhotoID checks that 'input' can be parsed as a Group Photo ID
func ValidateGroupPhotoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPhotoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Photo ID
func (id GroupPhotoId) ID() string {
	fmtString := "/groups/%s/photos/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ProfilePhotoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Photo ID
func (id GroupPhotoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPhotos", "photos", "photos"),
		resourceids.UserSpecifiedSegment("profilePhotoId", "profilePhotoIdValue"),
	}
}

// String returns a human-readable description of this Group Photo ID
func (id GroupPhotoId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Profile Photo: %q", id.ProfilePhotoId),
	}
	return fmt.Sprintf("Group Photo (%s)", strings.Join(components, "\n"))
}
