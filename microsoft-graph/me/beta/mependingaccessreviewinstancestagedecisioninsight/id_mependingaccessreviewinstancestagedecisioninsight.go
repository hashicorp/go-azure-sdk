package mependingaccessreviewinstancestagedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageDecisionInsightId{}

// MePendingAccessReviewInstanceStageDecisionInsightId is a struct representing the Resource ID for a Me Pending Access Review Instance Stage Decision Insight
type MePendingAccessReviewInstanceStageDecisionInsightId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewMePendingAccessReviewInstanceStageDecisionInsightID returns a new MePendingAccessReviewInstanceStageDecisionInsightId struct
func NewMePendingAccessReviewInstanceStageDecisionInsightID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) MePendingAccessReviewInstanceStageDecisionInsightId {
	return MePendingAccessReviewInstanceStageDecisionInsightId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseMePendingAccessReviewInstanceStageDecisionInsightID parses 'input' into a MePendingAccessReviewInstanceStageDecisionInsightId
func ParseMePendingAccessReviewInstanceStageDecisionInsightID(input string) (*MePendingAccessReviewInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceStageDecisionInsightIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceStageDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceStageDecisionInsightIDInsensitively(input string) (*MePendingAccessReviewInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceStageDecisionInsightID checks that 'input' can be parsed as a Me Pending Access Review Instance Stage Decision Insight ID
func ValidateMePendingAccessReviewInstanceStageDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceStageDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceStageDecisionInsightId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceStageDecisionInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceStageDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Stage Decision Insight (%s)", strings.Join(components, "\n"))
}
