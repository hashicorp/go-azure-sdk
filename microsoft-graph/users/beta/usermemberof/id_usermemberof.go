package usermemberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMemberOfId{}

// UserMemberOfId is a struct representing the Resource ID for a User Member Of
type UserMemberOfId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserMemberOfID returns a new UserMemberOfId struct
func NewUserMemberOfID(userId string, directoryObjectId string) UserMemberOfId {
	return UserMemberOfId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserMemberOfID parses 'input' into a UserMemberOfId
func ParseUserMemberOfID(input string) (*UserMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserMemberOfIDInsensitively parses 'input' case-insensitively into a UserMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserMemberOfIDInsensitively(input string) (*UserMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserMemberOfID checks that 'input' can be parsed as a User Member Of ID
func ValidateUserMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Member Of ID
func (id UserMemberOfId) ID() string {
	fmtString := "/users/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Member Of ID
func (id UserMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Member Of ID
func (id UserMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Member Of (%s)", strings.Join(components, "\n"))
}
