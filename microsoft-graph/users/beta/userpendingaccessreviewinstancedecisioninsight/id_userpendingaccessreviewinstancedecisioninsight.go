package userpendingaccessreviewinstancedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInsightId{}

// UserPendingAccessReviewInstanceDecisionInsightId is a struct representing the Resource ID for a User Pending Access Review Instance Decision Insight
type UserPendingAccessReviewInstanceDecisionInsightId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewUserPendingAccessReviewInstanceDecisionInsightID returns a new UserPendingAccessReviewInstanceDecisionInsightId struct
func NewUserPendingAccessReviewInstanceDecisionInsightID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) UserPendingAccessReviewInstanceDecisionInsightId {
	return UserPendingAccessReviewInstanceDecisionInsightId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionInsightID parses 'input' into a UserPendingAccessReviewInstanceDecisionInsightId
func ParseUserPendingAccessReviewInstanceDecisionInsightID(input string) (*UserPendingAccessReviewInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInsightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceDecisionInsightIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionInsightIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInsightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceDecisionInsightID checks that 'input' can be parsed as a User Pending Access Review Instance Decision Insight ID
func ValidateUserPendingAccessReviewInstanceDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInsightId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision Insight (%s)", strings.Join(components, "\n"))
}
