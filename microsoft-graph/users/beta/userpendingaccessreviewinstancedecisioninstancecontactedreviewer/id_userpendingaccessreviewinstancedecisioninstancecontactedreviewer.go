package userpendingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

// UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId is a struct representing the Resource ID for a User Pending Access Review Instance Decision Instance Contacted Reviewer
type UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID returns a new UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId struct
func NewUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId {
	return UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID parses 'input' into a UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId
func ParseUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID(input string) (*UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID checks that 'input' can be parsed as a User Pending Access Review Instance Decision Instance Contacted Reviewer ID
func ValidateUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticContactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
