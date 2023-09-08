package meonlinemeetingrecordingcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRecordingId{}

// MeOnlineMeetingRecordingId is a struct representing the Resource ID for a Me Online Meeting Recording
type MeOnlineMeetingRecordingId struct {
	OnlineMeetingId string
	CallRecordingId string
}

// NewMeOnlineMeetingRecordingID returns a new MeOnlineMeetingRecordingId struct
func NewMeOnlineMeetingRecordingID(onlineMeetingId string, callRecordingId string) MeOnlineMeetingRecordingId {
	return MeOnlineMeetingRecordingId{
		OnlineMeetingId: onlineMeetingId,
		CallRecordingId: callRecordingId,
	}
}

// ParseMeOnlineMeetingRecordingID parses 'input' into a MeOnlineMeetingRecordingId
func ParseMeOnlineMeetingRecordingID(input string) (*MeOnlineMeetingRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRecordingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRecordingId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallRecordingId, ok = parsed.Parsed["callRecordingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingRecordingIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingRecordingId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingRecordingIDInsensitively(input string) (*MeOnlineMeetingRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRecordingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRecordingId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallRecordingId, ok = parsed.Parsed["callRecordingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingRecordingID checks that 'input' can be parsed as a Me Online Meeting Recording ID
func ValidateMeOnlineMeetingRecordingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingRecordingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Recording ID
func (id MeOnlineMeetingRecordingId) ID() string {
	fmtString := "/me/onlineMeetings/%s/recordings/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.CallRecordingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Recording ID
func (id MeOnlineMeetingRecordingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRecordings", "recordings", "recordings"),
		resourceids.UserSpecifiedSegment("callRecordingId", "callRecordingIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Recording ID
func (id MeOnlineMeetingRecordingId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Recording: %q", id.CallRecordingId),
	}
	return fmt.Sprintf("Me Online Meeting Recording (%s)", strings.Join(components, "\n"))
}
