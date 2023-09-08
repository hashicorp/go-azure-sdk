package useractivityhistoryitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserActivityHistoryItemId{}

// UserActivityHistoryItemId is a struct representing the Resource ID for a User Activity History Item
type UserActivityHistoryItemId struct {
	UserId                string
	UserActivityId        string
	ActivityHistoryItemId string
}

// NewUserActivityHistoryItemID returns a new UserActivityHistoryItemId struct
func NewUserActivityHistoryItemID(userId string, userActivityId string, activityHistoryItemId string) UserActivityHistoryItemId {
	return UserActivityHistoryItemId{
		UserId:                userId,
		UserActivityId:        userActivityId,
		ActivityHistoryItemId: activityHistoryItemId,
	}
}

// ParseUserActivityHistoryItemID parses 'input' into a UserActivityHistoryItemId
func ParseUserActivityHistoryItemID(input string) (*UserActivityHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserActivityHistoryItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserActivityHistoryItemId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	if id.ActivityHistoryItemId, ok = parsed.Parsed["activityHistoryItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", *parsed)
	}

	return &id, nil
}

// ParseUserActivityHistoryItemIDInsensitively parses 'input' case-insensitively into a UserActivityHistoryItemId
// note: this method should only be used for API response data and not user input
func ParseUserActivityHistoryItemIDInsensitively(input string) (*UserActivityHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserActivityHistoryItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserActivityHistoryItemId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	if id.ActivityHistoryItemId, ok = parsed.Parsed["activityHistoryItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", *parsed)
	}

	return &id, nil
}

// ValidateUserActivityHistoryItemID checks that 'input' can be parsed as a User Activity History Item ID
func ValidateUserActivityHistoryItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserActivityHistoryItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Activity History Item ID
func (id UserActivityHistoryItemId) ID() string {
	fmtString := "/users/%s/activities/%s/historyItems/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserActivityId, id.ActivityHistoryItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Activity History Item ID
func (id UserActivityHistoryItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityIdValue"),
		resourceids.StaticSegment("staticHistoryItems", "historyItems", "historyItems"),
		resourceids.UserSpecifiedSegment("activityHistoryItemId", "activityHistoryItemIdValue"),
	}
}

// String returns a human-readable description of this User Activity History Item ID
func (id UserActivityHistoryItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
		fmt.Sprintf("Activity History Item: %q", id.ActivityHistoryItemId),
	}
	return fmt.Sprintf("User Activity History Item (%s)", strings.Join(components, "\n"))
}
