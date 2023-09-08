package groupsetting

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSettingId{}

// GroupSettingId is a struct representing the Resource ID for a Group Setting
type GroupSettingId struct {
	GroupId        string
	GroupSettingId string
}

// NewGroupSettingID returns a new GroupSettingId struct
func NewGroupSettingID(groupId string, groupSettingId string) GroupSettingId {
	return GroupSettingId{
		GroupId:        groupId,
		GroupSettingId: groupSettingId,
	}
}

// ParseGroupSettingID parses 'input' into a GroupSettingId
func ParseGroupSettingID(input string) (*GroupSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSettingId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.GroupSettingId, ok = parsed.Parsed["groupSettingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupSettingId", *parsed)
	}

	return &id, nil
}

// ParseGroupSettingIDInsensitively parses 'input' case-insensitively into a GroupSettingId
// note: this method should only be used for API response data and not user input
func ParseGroupSettingIDInsensitively(input string) (*GroupSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSettingId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.GroupSettingId, ok = parsed.Parsed["groupSettingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupSettingId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSettingID checks that 'input' can be parsed as a Group Setting ID
func ValidateGroupSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Setting ID
func (id GroupSettingId) ID() string {
	fmtString := "/groups/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.GroupSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Setting ID
func (id GroupSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSettings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("groupSettingId", "groupSettingIdValue"),
	}
}

// String returns a human-readable description of this Group Setting ID
func (id GroupSettingId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Group Setting: %q", id.GroupSettingId),
	}
	return fmt.Sprintf("Group Setting (%s)", strings.Join(components, "\n"))
}
