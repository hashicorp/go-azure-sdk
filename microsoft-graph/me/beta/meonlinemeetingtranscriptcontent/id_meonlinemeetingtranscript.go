package meonlinemeetingtranscriptcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingTranscriptId{}

// MeOnlineMeetingTranscriptId is a struct representing the Resource ID for a Me Online Meeting Transcript
type MeOnlineMeetingTranscriptId struct {
	OnlineMeetingId  string
	CallTranscriptId string
}

// NewMeOnlineMeetingTranscriptID returns a new MeOnlineMeetingTranscriptId struct
func NewMeOnlineMeetingTranscriptID(onlineMeetingId string, callTranscriptId string) MeOnlineMeetingTranscriptId {
	return MeOnlineMeetingTranscriptId{
		OnlineMeetingId:  onlineMeetingId,
		CallTranscriptId: callTranscriptId,
	}
}

// ParseMeOnlineMeetingTranscriptID parses 'input' into a MeOnlineMeetingTranscriptId
func ParseMeOnlineMeetingTranscriptID(input string) (*MeOnlineMeetingTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingTranscriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingTranscriptId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallTranscriptId, ok = parsed.Parsed["callTranscriptId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingTranscriptIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingTranscriptId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingTranscriptIDInsensitively(input string) (*MeOnlineMeetingTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingTranscriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingTranscriptId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.CallTranscriptId, ok = parsed.Parsed["callTranscriptId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingTranscriptID checks that 'input' can be parsed as a Me Online Meeting Transcript ID
func ValidateMeOnlineMeetingTranscriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingTranscriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Transcript ID
func (id MeOnlineMeetingTranscriptId) ID() string {
	fmtString := "/me/onlineMeetings/%s/transcripts/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.CallTranscriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Transcript ID
func (id MeOnlineMeetingTranscriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticTranscripts", "transcripts", "transcripts"),
		resourceids.UserSpecifiedSegment("callTranscriptId", "callTranscriptIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Transcript ID
func (id MeOnlineMeetingTranscriptId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Transcript: %q", id.CallTranscriptId),
	}
	return fmt.Sprintf("Me Online Meeting Transcript (%s)", strings.Join(components, "\n"))
}
