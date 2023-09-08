package meactivityhistoryitemactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeActivityHistoryItemId{}

// MeActivityHistoryItemId is a struct representing the Resource ID for a Me Activity History Item
type MeActivityHistoryItemId struct {
	UserActivityId        string
	ActivityHistoryItemId string
}

// NewMeActivityHistoryItemID returns a new MeActivityHistoryItemId struct
func NewMeActivityHistoryItemID(userActivityId string, activityHistoryItemId string) MeActivityHistoryItemId {
	return MeActivityHistoryItemId{
		UserActivityId:        userActivityId,
		ActivityHistoryItemId: activityHistoryItemId,
	}
}

// ParseMeActivityHistoryItemID parses 'input' into a MeActivityHistoryItemId
func ParseMeActivityHistoryItemID(input string) (*MeActivityHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeActivityHistoryItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeActivityHistoryItemId{}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	if id.ActivityHistoryItemId, ok = parsed.Parsed["activityHistoryItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", *parsed)
	}

	return &id, nil
}

// ParseMeActivityHistoryItemIDInsensitively parses 'input' case-insensitively into a MeActivityHistoryItemId
// note: this method should only be used for API response data and not user input
func ParseMeActivityHistoryItemIDInsensitively(input string) (*MeActivityHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeActivityHistoryItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeActivityHistoryItemId{}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	if id.ActivityHistoryItemId, ok = parsed.Parsed["activityHistoryItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", *parsed)
	}

	return &id, nil
}

// ValidateMeActivityHistoryItemID checks that 'input' can be parsed as a Me Activity History Item ID
func ValidateMeActivityHistoryItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeActivityHistoryItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Activity History Item ID
func (id MeActivityHistoryItemId) ID() string {
	fmtString := "/me/activities/%s/historyItems/%s"
	return fmt.Sprintf(fmtString, id.UserActivityId, id.ActivityHistoryItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Activity History Item ID
func (id MeActivityHistoryItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityIdValue"),
		resourceids.StaticSegment("staticHistoryItems", "historyItems", "historyItems"),
		resourceids.UserSpecifiedSegment("activityHistoryItemId", "activityHistoryItemIdValue"),
	}
}

// String returns a human-readable description of this Me Activity History Item ID
func (id MeActivityHistoryItemId) String() string {
	components := []string{
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
		fmt.Sprintf("Activity History Item: %q", id.ActivityHistoryItemId),
	}
	return fmt.Sprintf("Me Activity History Item (%s)", strings.Join(components, "\n"))
}
