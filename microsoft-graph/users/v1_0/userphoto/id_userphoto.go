package userphoto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPhotoId{}

// UserPhotoId is a struct representing the Resource ID for a User Photo
type UserPhotoId struct {
	UserId         string
	ProfilePhotoId string
}

// NewUserPhotoID returns a new UserPhotoId struct
func NewUserPhotoID(userId string, profilePhotoId string) UserPhotoId {
	return UserPhotoId{
		UserId:         userId,
		ProfilePhotoId: profilePhotoId,
	}
}

// ParseUserPhotoID parses 'input' into a UserPhotoId
func ParseUserPhotoID(input string) (*UserPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPhotoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPhotoId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ProfilePhotoId, ok = parsed.Parsed["profilePhotoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", *parsed)
	}

	return &id, nil
}

// ParseUserPhotoIDInsensitively parses 'input' case-insensitively into a UserPhotoId
// note: this method should only be used for API response data and not user input
func ParseUserPhotoIDInsensitively(input string) (*UserPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPhotoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPhotoId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ProfilePhotoId, ok = parsed.Parsed["profilePhotoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", *parsed)
	}

	return &id, nil
}

// ValidateUserPhotoID checks that 'input' can be parsed as a User Photo ID
func ValidateUserPhotoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPhotoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Photo ID
func (id UserPhotoId) ID() string {
	fmtString := "/users/%s/photos/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ProfilePhotoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Photo ID
func (id UserPhotoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPhotos", "photos", "photos"),
		resourceids.UserSpecifiedSegment("profilePhotoId", "profilePhotoIdValue"),
	}
}

// String returns a human-readable description of this User Photo ID
func (id UserPhotoId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Profile Photo: %q", id.ProfilePhotoId),
	}
	return fmt.Sprintf("User Photo (%s)", strings.Join(components, "\n"))
}
