package mependingaccessreviewinstancestagedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageId{}

// MePendingAccessReviewInstanceStageId is a struct representing the Resource ID for a Me Pending Access Review Instance Stage
type MePendingAccessReviewInstanceStageId struct {
	AccessReviewInstanceId string
	AccessReviewStageId    string
}

// NewMePendingAccessReviewInstanceStageID returns a new MePendingAccessReviewInstanceStageId struct
func NewMePendingAccessReviewInstanceStageID(accessReviewInstanceId string, accessReviewStageId string) MePendingAccessReviewInstanceStageId {
	return MePendingAccessReviewInstanceStageId{
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewStageId:    accessReviewStageId,
	}
}

// ParseMePendingAccessReviewInstanceStageID parses 'input' into a MePendingAccessReviewInstanceStageId
func ParseMePendingAccessReviewInstanceStageID(input string) (*MePendingAccessReviewInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceStageIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceStageIDInsensitively(input string) (*MePendingAccessReviewInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceStageID checks that 'input' can be parsed as a Me Pending Access Review Instance Stage ID
func ValidateMePendingAccessReviewInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Stage ID
func (id MePendingAccessReviewInstanceStageId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Stage ID
func (id MePendingAccessReviewInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Stage ID
func (id MePendingAccessReviewInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Stage (%s)", strings.Join(components, "\n"))
}
