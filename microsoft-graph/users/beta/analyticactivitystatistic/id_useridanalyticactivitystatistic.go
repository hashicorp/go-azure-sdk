package analyticactivitystatistic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAnalyticActivityStatisticId{}

// UserIdAnalyticActivityStatisticId is a struct representing the Resource ID for a User Id Analytic Activity Statistic
type UserIdAnalyticActivityStatisticId struct {
	UserId               string
	ActivityStatisticsId string
}

// NewUserIdAnalyticActivityStatisticID returns a new UserIdAnalyticActivityStatisticId struct
func NewUserIdAnalyticActivityStatisticID(userId string, activityStatisticsId string) UserIdAnalyticActivityStatisticId {
	return UserIdAnalyticActivityStatisticId{
		UserId:               userId,
		ActivityStatisticsId: activityStatisticsId,
	}
}

// ParseUserIdAnalyticActivityStatisticID parses 'input' into a UserIdAnalyticActivityStatisticId
func ParseUserIdAnalyticActivityStatisticID(input string) (*UserIdAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAnalyticActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAnalyticActivityStatisticIDInsensitively parses 'input' case-insensitively into a UserIdAnalyticActivityStatisticId
// note: this method should only be used for API response data and not user input
func ParseUserIdAnalyticActivityStatisticIDInsensitively(input string) (*UserIdAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAnalyticActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAnalyticActivityStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ActivityStatisticsId, ok = input.Parsed["activityStatisticsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", input)
	}

	return nil
}

// ValidateUserIdAnalyticActivityStatisticID checks that 'input' can be parsed as a User Id Analytic Activity Statistic ID
func ValidateUserIdAnalyticActivityStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAnalyticActivityStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Analytic Activity Statistic ID
func (id UserIdAnalyticActivityStatisticId) ID() string {
	fmtString := "/users/%s/analytics/activityStatistics/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ActivityStatisticsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Analytic Activity Statistic ID
func (id UserIdAnalyticActivityStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("activityStatistics", "activityStatistics", "activityStatistics"),
		resourceids.UserSpecifiedSegment("activityStatisticsId", "activityStatisticsIdValue"),
	}
}

// String returns a human-readable description of this User Id Analytic Activity Statistic ID
func (id UserIdAnalyticActivityStatisticId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Activity Statistics: %q", id.ActivityStatisticsId),
	}
	return fmt.Sprintf("User Id Analytic Activity Statistic (%s)", strings.Join(components, "\n"))
}
