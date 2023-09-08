package userpendingaccessreviewinstancedecisioninstancestagedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{}

// UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId is a struct representing the Resource ID for a User Pending Access Review Instance Decision Instance Stage Decision
type UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId struct {
	UserId                              string
	AccessReviewInstanceId              string
	AccessReviewInstanceDecisionItemId  string
	AccessReviewStageId                 string
	AccessReviewInstanceDecisionItemId1 string
}

// NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID returns a new UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId struct
func NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string, accessReviewInstanceDecisionItemId1 string) UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId {
	return UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{
		UserId:                              userId,
		AccessReviewInstanceId:              accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId:  accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                 accessReviewStageId,
		AccessReviewInstanceDecisionItemId1: accessReviewInstanceDecisionItemId1,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID parses 'input' into a UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{}

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

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{}

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

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID checks that 'input' can be parsed as a User Pending Access Review Instance Decision Instance Stage Decision ID
func ValidateUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId1", "accessReviewInstanceDecisionItemId1Value"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision Instance Stage Decision ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item Id 1: %q", id.AccessReviewInstanceDecisionItemId1),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision Instance Stage Decision (%s)", strings.Join(components, "\n"))
}
