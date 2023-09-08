package userpendingaccessreviewinstancestagedecisioninstancedefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceStageDecisionId{}

// UserPendingAccessReviewInstanceStageDecisionId is a struct representing the Resource ID for a User Pending Access Review Instance Stage Decision
type UserPendingAccessReviewInstanceStageDecisionId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
}

// NewUserPendingAccessReviewInstanceStageDecisionID returns a new UserPendingAccessReviewInstanceStageDecisionId struct
func NewUserPendingAccessReviewInstanceStageDecisionID(userId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string) UserPendingAccessReviewInstanceStageDecisionId {
	return UserPendingAccessReviewInstanceStageDecisionId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseUserPendingAccessReviewInstanceStageDecisionID parses 'input' into a UserPendingAccessReviewInstanceStageDecisionId
func ParseUserPendingAccessReviewInstanceStageDecisionID(input string) (*UserPendingAccessReviewInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionId{}

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

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceStageDecisionIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceStageDecisionId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceStageDecisionIDInsensitively(input string) (*UserPendingAccessReviewInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionId{}

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

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceStageDecisionID checks that 'input' can be parsed as a User Pending Access Review Instance Stage Decision ID
func ValidateUserPendingAccessReviewInstanceStageDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceStageDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceStageDecisionId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceStageDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceStageDecisionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Stage Decision (%s)", strings.Join(components, "\n"))
}
