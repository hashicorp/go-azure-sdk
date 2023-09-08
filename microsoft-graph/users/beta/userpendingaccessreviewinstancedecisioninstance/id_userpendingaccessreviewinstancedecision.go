package userpendingaccessreviewinstancedecisioninstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionId{}

// UserPendingAccessReviewInstanceDecisionId is a struct representing the Resource ID for a User Pending Access Review Instance Decision
type UserPendingAccessReviewInstanceDecisionId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
}

// NewUserPendingAccessReviewInstanceDecisionID returns a new UserPendingAccessReviewInstanceDecisionId struct
func NewUserPendingAccessReviewInstanceDecisionID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string) UserPendingAccessReviewInstanceDecisionId {
	return UserPendingAccessReviewInstanceDecisionId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionID parses 'input' into a UserPendingAccessReviewInstanceDecisionId
func ParseUserPendingAccessReviewInstanceDecisionID(input string) (*UserPendingAccessReviewInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceDecisionIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceDecisionID checks that 'input' can be parsed as a User Pending Access Review Instance Decision ID
func ValidateUserPendingAccessReviewInstanceDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision ID
func (id UserPendingAccessReviewInstanceDecisionId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision ID
func (id UserPendingAccessReviewInstanceDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision ID
func (id UserPendingAccessReviewInstanceDecisionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision (%s)", strings.Join(components, "\n"))
}
