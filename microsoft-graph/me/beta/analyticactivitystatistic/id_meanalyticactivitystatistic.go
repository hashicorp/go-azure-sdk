package analyticactivitystatistic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAnalyticActivityStatisticId{}

// MeAnalyticActivityStatisticId is a struct representing the Resource ID for a Me Analytic Activity Statistic
type MeAnalyticActivityStatisticId struct {
	ActivityStatisticsId string
}

// NewMeAnalyticActivityStatisticID returns a new MeAnalyticActivityStatisticId struct
func NewMeAnalyticActivityStatisticID(activityStatisticsId string) MeAnalyticActivityStatisticId {
	return MeAnalyticActivityStatisticId{
		ActivityStatisticsId: activityStatisticsId,
	}
}

// ParseMeAnalyticActivityStatisticID parses 'input' into a MeAnalyticActivityStatisticId
func ParseMeAnalyticActivityStatisticID(input string) (*MeAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAnalyticActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAnalyticActivityStatisticIDInsensitively parses 'input' case-insensitively into a MeAnalyticActivityStatisticId
// note: this method should only be used for API response data and not user input
func ParseMeAnalyticActivityStatisticIDInsensitively(input string) (*MeAnalyticActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAnalyticActivityStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAnalyticActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAnalyticActivityStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActivityStatisticsId, ok = input.Parsed["activityStatisticsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", input)
	}

	return nil
}

// ValidateMeAnalyticActivityStatisticID checks that 'input' can be parsed as a Me Analytic Activity Statistic ID
func ValidateMeAnalyticActivityStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAnalyticActivityStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Analytic Activity Statistic ID
func (id MeAnalyticActivityStatisticId) ID() string {
	fmtString := "/me/analytics/activityStatistics/%s"
	return fmt.Sprintf(fmtString, id.ActivityStatisticsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Analytic Activity Statistic ID
func (id MeAnalyticActivityStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("activityStatistics", "activityStatistics", "activityStatistics"),
		resourceids.UserSpecifiedSegment("activityStatisticsId", "activityStatisticsIdValue"),
	}
}

// String returns a human-readable description of this Me Analytic Activity Statistic ID
func (id MeAnalyticActivityStatisticId) String() string {
	components := []string{
		fmt.Sprintf("Activity Statistics: %q", id.ActivityStatisticsId),
	}
	return fmt.Sprintf("Me Analytic Activity Statistic (%s)", strings.Join(components, "\n"))
}
