package groupgrouplifecyclepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupGroupLifecyclePolicyId{}

// GroupGroupLifecyclePolicyId is a struct representing the Resource ID for a Group Group Lifecycle Policy
type GroupGroupLifecyclePolicyId struct {
	GroupId                string
	GroupLifecyclePolicyId string
}

// NewGroupGroupLifecyclePolicyID returns a new GroupGroupLifecyclePolicyId struct
func NewGroupGroupLifecyclePolicyID(groupId string, groupLifecyclePolicyId string) GroupGroupLifecyclePolicyId {
	return GroupGroupLifecyclePolicyId{
		GroupId:                groupId,
		GroupLifecyclePolicyId: groupLifecyclePolicyId,
	}
}

// ParseGroupGroupLifecyclePolicyID parses 'input' into a GroupGroupLifecyclePolicyId
func ParseGroupGroupLifecyclePolicyID(input string) (*GroupGroupLifecyclePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupGroupLifecyclePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupGroupLifecyclePolicyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.GroupLifecyclePolicyId, ok = parsed.Parsed["groupLifecyclePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupLifecyclePolicyId", *parsed)
	}

	return &id, nil
}

// ParseGroupGroupLifecyclePolicyIDInsensitively parses 'input' case-insensitively into a GroupGroupLifecyclePolicyId
// note: this method should only be used for API response data and not user input
func ParseGroupGroupLifecyclePolicyIDInsensitively(input string) (*GroupGroupLifecyclePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupGroupLifecyclePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupGroupLifecyclePolicyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.GroupLifecyclePolicyId, ok = parsed.Parsed["groupLifecyclePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupLifecyclePolicyId", *parsed)
	}

	return &id, nil
}

// ValidateGroupGroupLifecyclePolicyID checks that 'input' can be parsed as a Group Group Lifecycle Policy ID
func ValidateGroupGroupLifecyclePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupGroupLifecyclePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Group Lifecycle Policy ID
func (id GroupGroupLifecyclePolicyId) ID() string {
	fmtString := "/groups/%s/groupLifecyclePolicies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.GroupLifecyclePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Group Lifecycle Policy ID
func (id GroupGroupLifecyclePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticGroupLifecyclePolicies", "groupLifecyclePolicies", "groupLifecyclePolicies"),
		resourceids.UserSpecifiedSegment("groupLifecyclePolicyId", "groupLifecyclePolicyIdValue"),
	}
}

// String returns a human-readable description of this Group Group Lifecycle Policy ID
func (id GroupGroupLifecyclePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Group Lifecycle Policy: %q", id.GroupLifecyclePolicyId),
	}
	return fmt.Sprintf("Group Group Lifecycle Policy (%s)", strings.Join(components, "\n"))
}
