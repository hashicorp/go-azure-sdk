package useranalyticactivitystatistic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAnalyticActivityStatisticId{}

// UserAnalyticActivityStatisticId is a struct representing the Resource ID for a User Analytic Activity Statistic
type UserAnalyticActivityStatisticId struct {
	UserId               string
	ActivityStatisticsId string
}

// NewUserAnalyticActivityStatisticID returns a new UserAnalyticActivityStatisticId struct
func NewUserAnalyticActivityStatisticID(userId string, activityStatisticsId string) UserAnalyticActivityStatisticId {
	return UserAnalyticActivityStatisticId{
		UserId:               userId,
		ActivityStatisticsId: activityStatisticsId,
	}
}

// ParseUserAnalyticActivityStatisticID parses 'input' into a UserAnalyticActivityStatisticId
func ParseUserAnalyticActivityStatisticID(input string) (*UserAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAnalyticActivityStatisticId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ActivityStatisticsId, ok = parsed.Parsed["activityStatisticsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", *parsed)
	}

	return &id, nil
}

// ParseUserAnalyticActivityStatisticIDInsensitively parses 'input' case-insensitively into a UserAnalyticActivityStatisticId
// note: this method should only be used for API response data and not user input
func ParseUserAnalyticActivityStatisticIDInsensitively(input string) (*UserAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAnalyticActivityStatisticId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ActivityStatisticsId, ok = parsed.Parsed["activityStatisticsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", *parsed)
	}

	return &id, nil
}

// ValidateUserAnalyticActivityStatisticID checks that 'input' can be parsed as a User Analytic Activity Statistic ID
func ValidateUserAnalyticActivityStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAnalyticActivityStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Analytic Activity Statistic ID
func (id UserAnalyticActivityStatisticId) ID() string {
	fmtString := "/users/%s/analytics/activityStatistics/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ActivityStatisticsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Analytic Activity Statistic ID
func (id UserAnalyticActivityStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAnalytics", "analytics", "analytics"),
		resourceids.StaticSegment("staticActivityStatistics", "activityStatistics", "activityStatistics"),
		resourceids.UserSpecifiedSegment("activityStatisticsId", "activityStatisticsIdValue"),
	}
}

// String returns a human-readable description of this User Analytic Activity Statistic ID
func (id UserAnalyticActivityStatisticId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Activity Statistics: %q", id.ActivityStatisticsId),
	}
	return fmt.Sprintf("User Analytic Activity Statistic (%s)", strings.Join(components, "\n"))
}
