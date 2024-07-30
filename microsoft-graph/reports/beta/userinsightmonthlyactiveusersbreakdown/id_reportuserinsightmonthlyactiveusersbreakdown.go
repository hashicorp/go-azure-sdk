package userinsightmonthlyactiveusersbreakdown

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyActiveUsersBreakdownId{}

// ReportUserInsightMonthlyActiveUsersBreakdownId is a struct representing the Resource ID for a Report User Insight Monthly Active Users Breakdown
type ReportUserInsightMonthlyActiveUsersBreakdownId struct {
	ActiveUsersBreakdownMetricId string
}

// NewReportUserInsightMonthlyActiveUsersBreakdownID returns a new ReportUserInsightMonthlyActiveUsersBreakdownId struct
func NewReportUserInsightMonthlyActiveUsersBreakdownID(activeUsersBreakdownMetricId string) ReportUserInsightMonthlyActiveUsersBreakdownId {
	return ReportUserInsightMonthlyActiveUsersBreakdownId{
		ActiveUsersBreakdownMetricId: activeUsersBreakdownMetricId,
	}
}

// ParseReportUserInsightMonthlyActiveUsersBreakdownID parses 'input' into a ReportUserInsightMonthlyActiveUsersBreakdownId
func ParseReportUserInsightMonthlyActiveUsersBreakdownID(input string) (*ReportUserInsightMonthlyActiveUsersBreakdownId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyActiveUsersBreakdownId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyActiveUsersBreakdownId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyActiveUsersBreakdownIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyActiveUsersBreakdownId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyActiveUsersBreakdownIDInsensitively(input string) (*ReportUserInsightMonthlyActiveUsersBreakdownId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyActiveUsersBreakdownId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyActiveUsersBreakdownId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyActiveUsersBreakdownId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActiveUsersBreakdownMetricId, ok = input.Parsed["activeUsersBreakdownMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activeUsersBreakdownMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyActiveUsersBreakdownID checks that 'input' can be parsed as a Report User Insight Monthly Active Users Breakdown ID
func ValidateReportUserInsightMonthlyActiveUsersBreakdownID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyActiveUsersBreakdownID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Active Users Breakdown ID
func (id ReportUserInsightMonthlyActiveUsersBreakdownId) ID() string {
	fmtString := "/reports/userInsights/monthly/activeUsersBreakdown/%s"
	return fmt.Sprintf(fmtString, id.ActiveUsersBreakdownMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Active Users Breakdown ID
func (id ReportUserInsightMonthlyActiveUsersBreakdownId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("activeUsersBreakdown", "activeUsersBreakdown", "activeUsersBreakdown"),
		resourceids.UserSpecifiedSegment("activeUsersBreakdownMetricId", "activeUsersBreakdownMetricIdValue"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Active Users Breakdown ID
func (id ReportUserInsightMonthlyActiveUsersBreakdownId) String() string {
	components := []string{
		fmt.Sprintf("Active Users Breakdown Metric: %q", id.ActiveUsersBreakdownMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Active Users Breakdown (%s)", strings.Join(components, "\n"))
}
