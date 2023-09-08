package userdrive

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDriveId{}

// UserDriveId is a struct representing the Resource ID for a User Drive
type UserDriveId struct {
	UserId  string
	DriveId string
}

// NewUserDriveID returns a new UserDriveId struct
func NewUserDriveID(userId string, driveId string) UserDriveId {
	return UserDriveId{
		UserId:  userId,
		DriveId: driveId,
	}
}

// ParseUserDriveID parses 'input' into a UserDriveId
func ParseUserDriveID(input string) (*UserDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDriveId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ParseUserDriveIDInsensitively parses 'input' case-insensitively into a UserDriveId
// note: this method should only be used for API response data and not user input
func ParseUserDriveIDInsensitively(input string) (*UserDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDriveId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ValidateUserDriveID checks that 'input' can be parsed as a User Drive ID
func ValidateUserDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Drive ID
func (id UserDriveId) ID() string {
	fmtString := "/users/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Drive ID
func (id UserDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDrives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveIdValue"),
	}
}

// String returns a human-readable description of this User Drive ID
func (id UserDriveId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("User Drive (%s)", strings.Join(components, "\n"))
}
