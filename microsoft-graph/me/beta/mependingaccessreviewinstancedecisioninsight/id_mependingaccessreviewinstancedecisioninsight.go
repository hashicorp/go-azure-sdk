package mependingaccessreviewinstancedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInsightId{}

// MePendingAccessReviewInstanceDecisionInsightId is a struct representing the Resource ID for a Me Pending Access Review Instance Decision Insight
type MePendingAccessReviewInstanceDecisionInsightId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewMePendingAccessReviewInstanceDecisionInsightID returns a new MePendingAccessReviewInstanceDecisionInsightId struct
func NewMePendingAccessReviewInstanceDecisionInsightID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) MePendingAccessReviewInstanceDecisionInsightId {
	return MePendingAccessReviewInstanceDecisionInsightId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseMePendingAccessReviewInstanceDecisionInsightID parses 'input' into a MePendingAccessReviewInstanceDecisionInsightId
func ParseMePendingAccessReviewInstanceDecisionInsightID(input string) (*MePendingAccessReviewInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceDecisionInsightIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceDecisionInsightIDInsensitively(input string) (*MePendingAccessReviewInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceDecisionInsightID checks that 'input' can be parsed as a Me Pending Access Review Instance Decision Insight ID
func ValidateMePendingAccessReviewInstanceDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInsightId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Decision Insight (%s)", strings.Join(components, "\n"))
}
