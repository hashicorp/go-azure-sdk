package userpendingaccessreviewinstancedecisioninstancestagedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

// UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId is a struct representing the Resource ID for a User Pending Access Review Instance Decision Instance Stage Decision Insight
type UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId struct {
	UserId                              string
	AccessReviewInstanceId              string
	AccessReviewInstanceDecisionItemId  string
	AccessReviewStageId                 string
	AccessReviewInstanceDecisionItemId1 string
	GovernanceInsightId                 string
}

// NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID returns a new UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId struct
func NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string, accessReviewInstanceDecisionItemId1 string, governanceInsightId string) UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId {
	return UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{
		UserId:                              userId,
		AccessReviewInstanceId:              accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId:  accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                 accessReviewStageId,
		AccessReviewInstanceDecisionItemId1: accessReviewInstanceDecisionItemId1,
		GovernanceInsightId:                 governanceInsightId,
	}
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID parses 'input' into a UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

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

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively(input string) (*UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

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

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID checks that 'input' can be parsed as a User Pending Access Review Instance Decision Instance Stage Decision Insight ID
func ValidateUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId1, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Decision Instance Stage Decision Insight ID
func (id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item Id 1: %q", id.AccessReviewInstanceDecisionItemId1),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Decision Instance Stage Decision Insight (%s)", strings.Join(components, "\n"))
}
