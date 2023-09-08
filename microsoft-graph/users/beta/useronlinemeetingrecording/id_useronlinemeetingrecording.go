package useronlinemeetingrecording

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRecordingId{}

// UserOnlineMeetingRecordingId is a struct representing the Resource ID for a User Online Meeting Recording
type UserOnlineMeetingRecordingId struct {
	UserId          string
	OnlineMeetingId string
	CallRecordingId string
}

// NewUserOnlineMeetingRecordingID returns a new UserOnlineMeetingRecordingId struct
func NewUserOnlineMeetingRecordingID(userId string, onlineMeetingId string, callRecordingId string) UserOnlineMeetingRecordingId {
	return UserOnlineMeetingRecordingId{
		UserId:          userId,
		OnlineMeetingId: onlineMeetingId,
		CallRecordingId: callRecordingId,
	}
}

// ParseUserOnlineMeetingRecordingID parses 'input' into a UserOnlineMeetingRecordingId
func ParseUserOnlineMeetingRecordingID(input string) (*UserOnlineMeetingRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRecordingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRecordingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallRecordingId, ok = parsed.Parsed["callRecordingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingRecordingIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingRecordingId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingRecordingIDInsensitively(input string) (*UserOnlineMeetingRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRecordingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRecordingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallRecordingId, ok = parsed.Parsed["callRecordingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingRecordingID checks that 'input' can be parsed as a User Online Meeting Recording ID
func ValidateUserOnlineMeetingRecordingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingRecordingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Recording ID
func (id UserOnlineMeetingRecordingId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/recordings/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.CallRecordingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Recording ID
func (id UserOnlineMeetingRecordingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRecordings", "recordings", "recordings"),
		resourceids.UserSpecifiedSegment("callRecordingId", "callRecordingIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Recording ID
func (id UserOnlineMeetingRecordingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Recording: %q", id.CallRecordingId),
	}
	return fmt.Sprintf("User Online Meeting Recording (%s)", strings.Join(components, "\n"))
}
