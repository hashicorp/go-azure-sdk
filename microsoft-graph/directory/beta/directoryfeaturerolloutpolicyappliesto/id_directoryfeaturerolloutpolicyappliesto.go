package directoryfeaturerolloutpolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryFeatureRolloutPolicyAppliesToId{}

// DirectoryFeatureRolloutPolicyAppliesToId is a struct representing the Resource ID for a Directory Feature Rollout Policy Applies To
type DirectoryFeatureRolloutPolicyAppliesToId struct {
	FeatureRolloutPolicyId string
	DirectoryObjectId      string
}

// NewDirectoryFeatureRolloutPolicyAppliesToID returns a new DirectoryFeatureRolloutPolicyAppliesToId struct
func NewDirectoryFeatureRolloutPolicyAppliesToID(featureRolloutPolicyId string, directoryObjectId string) DirectoryFeatureRolloutPolicyAppliesToId {
	return DirectoryFeatureRolloutPolicyAppliesToId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
		DirectoryObjectId:      directoryObjectId,
	}
}

// ParseDirectoryFeatureRolloutPolicyAppliesToID parses 'input' into a DirectoryFeatureRolloutPolicyAppliesToId
func ParseDirectoryFeatureRolloutPolicyAppliesToID(input string) (*DirectoryFeatureRolloutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryFeatureRolloutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryFeatureRolloutPolicyAppliesToId{}

	if id.FeatureRolloutPolicyId, ok = parsed.Parsed["featureRolloutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryFeatureRolloutPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a DirectoryFeatureRolloutPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParseDirectoryFeatureRolloutPolicyAppliesToIDInsensitively(input string) (*DirectoryFeatureRolloutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryFeatureRolloutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryFeatureRolloutPolicyAppliesToId{}

	if id.FeatureRolloutPolicyId, ok = parsed.Parsed["featureRolloutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryFeatureRolloutPolicyAppliesToID checks that 'input' can be parsed as a Directory Feature Rollout Policy Applies To ID
func ValidateDirectoryFeatureRolloutPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryFeatureRolloutPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Feature Rollout Policy Applies To ID
func (id DirectoryFeatureRolloutPolicyAppliesToId) ID() string {
	fmtString := "/directory/featureRolloutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Feature Rollout Policy Applies To ID
func (id DirectoryFeatureRolloutPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticFeatureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Directory Feature Rollout Policy Applies To ID
func (id DirectoryFeatureRolloutPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Feature Rollout Policy Applies To (%s)", strings.Join(components, "\n"))
}
