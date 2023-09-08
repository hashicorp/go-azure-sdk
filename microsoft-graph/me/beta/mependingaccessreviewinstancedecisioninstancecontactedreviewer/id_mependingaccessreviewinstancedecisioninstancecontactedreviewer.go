package mependingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

// MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Decision Instance Contacted Reviewer
type MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID returns a new MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId struct
func NewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId {
	return MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId
func ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(input string) (*MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Decision Instance Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticContactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Decision Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
