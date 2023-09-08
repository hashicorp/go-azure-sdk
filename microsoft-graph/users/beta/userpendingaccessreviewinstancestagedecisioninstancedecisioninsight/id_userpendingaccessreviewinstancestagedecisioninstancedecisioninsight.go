package userpendingaccessreviewinstancestagedecisioninstancedecisioninsight

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{}

// UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId is a struct representing the Resource ID for a User Pending Access Review Instance Stage Decision Instance Decision Insight
type UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId struct {
	UserId                              string
	AccessReviewInstanceId              string
	AccessReviewStageId                 string
	AccessReviewInstanceDecisionItemId  string
	AccessReviewInstanceDecisionItemId1 string
	GovernanceInsightId                 string
}

// NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID returns a new UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId struct
func NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(userId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewInstanceDecisionItemId1 string, governanceInsightId string) UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId {
	return UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{
		UserId:                              userId,
		AccessReviewInstanceId:              accessReviewInstanceId,
		AccessReviewStageId:                 accessReviewStageId,
		AccessReviewInstanceDecisionItemId:  accessReviewInstanceDecisionItemId,
		AccessReviewInstanceDecisionItemId1: accessReviewInstanceDecisionItemId1,
		GovernanceInsightId:                 governanceInsightId,
	}
}

// ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID parses 'input' into a UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId
func ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(input string) (*UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{}

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

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightIDInsensitively(input string) (*UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{}

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

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	if id.GovernanceInsightId, ok = parsed.Parsed["governanceInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID checks that 'input' can be parsed as a User Pending Access Review Instance Stage Decision Instance Decision Insight ID
func ValidateUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Stage Decision Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/instance/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewInstanceDecisionItemId1, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Stage Decision Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId1", "accessReviewInstanceDecisionItemId1Value"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Stage Decision Instance Decision Insight ID
func (id UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Instance Decision Item Id 1: %q", id.AccessReviewInstanceDecisionItemId1),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Stage Decision Instance Decision Insight (%s)", strings.Join(components, "\n"))
}
