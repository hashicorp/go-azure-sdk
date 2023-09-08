package userpendingaccessreviewinstancestagedecisioninstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

// UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId is a struct representing the Resource ID for a User Pending Access Review Instance Stage Decision Instance Contacted Reviewer
type UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID returns a new UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId struct
func NewUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(userId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId {
	return UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID parses 'input' into a UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId
func ParseUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(input string) (*UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerIDInsensitively(input string) (*UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID checks that 'input' can be parsed as a User Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func ValidateUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
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

// String returns a human-readable description of this User Pending Access Review Instance Stage Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Stage Decision Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
