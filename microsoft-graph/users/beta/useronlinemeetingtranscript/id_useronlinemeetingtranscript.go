package useronlinemeetingtranscript

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingTranscriptId{}

// UserOnlineMeetingTranscriptId is a struct representing the Resource ID for a User Online Meeting Transcript
type UserOnlineMeetingTranscriptId struct {
	UserId           string
	OnlineMeetingId  string
	CallTranscriptId string
}

// NewUserOnlineMeetingTranscriptID returns a new UserOnlineMeetingTranscriptId struct
func NewUserOnlineMeetingTranscriptID(userId string, onlineMeetingId string, callTranscriptId string) UserOnlineMeetingTranscriptId {
	return UserOnlineMeetingTranscriptId{
		UserId:           userId,
		OnlineMeetingId:  onlineMeetingId,
		CallTranscriptId: callTranscriptId,
	}
}

// ParseUserOnlineMeetingTranscriptID parses 'input' into a UserOnlineMeetingTranscriptId
func ParseUserOnlineMeetingTranscriptID(input string) (*UserOnlineMeetingTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingTranscriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingTranscriptId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallTranscriptId, ok = parsed.Parsed["callTranscriptId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingTranscriptIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingTranscriptId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingTranscriptIDInsensitively(input string) (*UserOnlineMeetingTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingTranscriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingTranscriptId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallTranscriptId, ok = parsed.Parsed["callTranscriptId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingTranscriptID checks that 'input' can be parsed as a User Online Meeting Transcript ID
func ValidateUserOnlineMeetingTranscriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingTranscriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Transcript ID
func (id UserOnlineMeetingTranscriptId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/transcripts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.CallTranscriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Transcript ID
func (id UserOnlineMeetingTranscriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticTranscripts", "transcripts", "transcripts"),
		resourceids.UserSpecifiedSegment("callTranscriptId", "callTranscriptIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Transcript ID
func (id UserOnlineMeetingTranscriptId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Transcript: %q", id.CallTranscriptId),
	}
	return fmt.Sprintf("User Online Meeting Transcript (%s)", strings.Join(components, "\n"))
}
