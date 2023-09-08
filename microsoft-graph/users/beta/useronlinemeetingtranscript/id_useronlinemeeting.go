package useronlinemeetingtranscript

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingId{}

// UserOnlineMeetingId is a struct representing the Resource ID for a User Online Meeting
type UserOnlineMeetingId struct {
	UserId          string
	OnlineMeetingId string
}

// NewUserOnlineMeetingID returns a new UserOnlineMeetingId struct
func NewUserOnlineMeetingID(userId string, onlineMeetingId string) UserOnlineMeetingId {
	return UserOnlineMeetingId{
		UserId:          userId,
		OnlineMeetingId: onlineMeetingId,
	}
}

// ParseUserOnlineMeetingID parses 'input' into a UserOnlineMeetingId
func ParseUserOnlineMeetingID(input string) (*UserOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingIDInsensitively(input string) (*UserOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingID checks that 'input' can be parsed as a User Online Meeting ID
func ValidateUserOnlineMeetingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting ID
func (id UserOnlineMeetingId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting ID
func (id UserOnlineMeetingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting ID
func (id UserOnlineMeetingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
	}
	return fmt.Sprintf("User Online Meeting (%s)", strings.Join(components, "\n"))
}
