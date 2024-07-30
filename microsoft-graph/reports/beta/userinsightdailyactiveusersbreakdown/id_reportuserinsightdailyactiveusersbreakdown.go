package userinsightdailyactiveusersbreakdown

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyActiveUsersBreakdownId{}

// ReportUserInsightDailyActiveUsersBreakdownId is a struct representing the Resource ID for a Report User Insight Daily Active Users Breakdown
type ReportUserInsightDailyActiveUsersBreakdownId struct {
	ActiveUsersBreakdownMetricId string
}

// NewReportUserInsightDailyActiveUsersBreakdownID returns a new ReportUserInsightDailyActiveUsersBreakdownId struct
func NewReportUserInsightDailyActiveUsersBreakdownID(activeUsersBreakdownMetricId string) ReportUserInsightDailyActiveUsersBreakdownId {
	return ReportUserInsightDailyActiveUsersBreakdownId{
		ActiveUsersBreakdownMetricId: activeUsersBreakdownMetricId,
	}
}

// ParseReportUserInsightDailyActiveUsersBreakdownID parses 'input' into a ReportUserInsightDailyActiveUsersBreakdownId
func ParseReportUserInsightDailyActiveUsersBreakdownID(input string) (*ReportUserInsightDailyActiveUsersBreakdownId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyActiveUsersBreakdownId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyActiveUsersBreakdownId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyActiveUsersBreakdownIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyActiveUsersBreakdownId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyActiveUsersBreakdownIDInsensitively(input string) (*ReportUserInsightDailyActiveUsersBreakdownId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyActiveUsersBreakdownId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyActiveUsersBreakdownId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyActiveUsersBreakdownId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActiveUsersBreakdownMetricId, ok = input.Parsed["activeUsersBreakdownMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activeUsersBreakdownMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyActiveUsersBreakdownID checks that 'input' can be parsed as a Report User Insight Daily Active Users Breakdown ID
func ValidateReportUserInsightDailyActiveUsersBreakdownID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyActiveUsersBreakdownID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Active Users Breakdown ID
func (id ReportUserInsightDailyActiveUsersBreakdownId) ID() string {
	fmtString := "/reports/userInsights/daily/activeUsersBreakdown/%s"
	return fmt.Sprintf(fmtString, id.ActiveUsersBreakdownMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Active Users Breakdown ID
func (id ReportUserInsightDailyActiveUsersBreakdownId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("activeUsersBreakdown", "activeUsersBreakdown", "activeUsersBreakdown"),
		resourceids.UserSpecifiedSegment("activeUsersBreakdownMetricId", "activeUsersBreakdownMetricIdValue"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Active Users Breakdown ID
func (id ReportUserInsightDailyActiveUsersBreakdownId) String() string {
	components := []string{
		fmt.Sprintf("Active Users Breakdown Metric: %q", id.ActiveUsersBreakdownMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Active Users Breakdown (%s)", strings.Join(components, "\n"))
}
