package mependingaccessreviewinstancedecisioninstancestagedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

// MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId is a struct representing the Resource ID for a Me Pending Access Review Instance Decision Instance Stage Decision Insight
type MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId struct {
	AccessReviewInstanceId              string
	AccessReviewInstanceDecisionItemId  string
	AccessReviewStageId                 string
	AccessReviewInstanceDecisionItemId1 string
	GovernanceInsightId                 string
}

// NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID returns a new MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId struct
func NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string, accessReviewInstanceDecisionItemId1 string, governanceInsightId string) MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId {
	return MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{
		AccessReviewInstanceId:              accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId:  accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                 accessReviewStageId,
		AccessReviewInstanceDecisionItemId1: accessReviewInstanceDecisionItemId1,
		GovernanceInsightId:                 governanceInsightId,
	}
}

// ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID parses 'input' into a MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
func ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(input string) (*MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively(input string) (*MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID checks that 'input' can be parsed as a Me Pending Access Review Instance Decision Instance Stage Decision Insight ID
func ValidateMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId1, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId1", "accessReviewInstanceDecisionItemId1Value"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item Id 1: %q", id.AccessReviewInstanceDecisionItemId1),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Decision Instance Stage Decision Insight (%s)", strings.Join(components, "\n"))
}
