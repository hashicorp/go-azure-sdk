package userpendingaccessreviewinstancedecisioninstancestagedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInstanceStageId{}

// UserPendingAccessReviewInstanceDecisionInstanceStageId is a struct representing the Resource ID for a User Pending Access Review Instance Decision Instance Stage
type UserPendingAccessReviewInstanceDecisionInstanceStageId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewUserPendingAccessReviewInstanceDecisionInstanceStageID returns a new UserPendingAccessReviewInstanceDecisionInstanceStageId struct
func NewUserPendingAccessReviewInstanceDecisionInstanceStageID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string) UserPendingAccessReviewInstanceDecisionInstanceStageId {
	return UserPendingAccessReviewInstanceDecisionInstanceStageId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageID parses 'input' into a UserPendingAccessReviewInstanceDecisionInstanceStageId
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageID(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceDecisionInstanceStageID checks that 'input' can be parsed as a User Pending Access Review Instance Decision Instance Stage ID
func ValidateUserPendingAccessReviewInstanceDecisionInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision Instance Stage ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision Instance Stage ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision Instance Stage ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision Instance Stage (%s)", strings.Join(components, "\n"))
}
