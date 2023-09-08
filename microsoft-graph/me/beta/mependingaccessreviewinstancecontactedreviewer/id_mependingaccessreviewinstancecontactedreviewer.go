package mependingaccessreviewinstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceContactedReviewerId{}

// MePendingAccessReviewInstanceContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Contacted Reviewer
type MePendingAccessReviewInstanceContactedReviewerId struct {
	AccessReviewInstanceId string
	AccessReviewReviewerId string
}

// NewMePendingAccessReviewInstanceContactedReviewerID returns a new MePendingAccessReviewInstanceContactedReviewerId struct
func NewMePendingAccessReviewInstanceContactedReviewerID(accessReviewInstanceId string, accessReviewReviewerId string) MePendingAccessReviewInstanceContactedReviewerId {
	return MePendingAccessReviewInstanceContactedReviewerId{
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewReviewerId: accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceContactedReviewerId
func ParseMePendingAccessReviewInstanceContactedReviewerID(input string) (*MePendingAccessReviewInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceContactedReviewerId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticContactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
