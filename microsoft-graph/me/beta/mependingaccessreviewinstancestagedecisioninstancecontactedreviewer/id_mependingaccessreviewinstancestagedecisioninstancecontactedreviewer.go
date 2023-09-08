package mependingaccessreviewinstancestagedecisioninstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

// MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer
type MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID returns a new MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId struct
func NewMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId {
	return MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId
func ParseMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(input string) (*MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticContactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Stage Decision Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
