package groupendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEndpointId{}

// GroupEndpointId is a struct representing the Resource ID for a Group Endpoint
type GroupEndpointId struct {
	GroupId    string
	EndpointId string
}

// NewGroupEndpointID returns a new GroupEndpointId struct
func NewGroupEndpointID(groupId string, endpointId string) GroupEndpointId {
	return GroupEndpointId{
		GroupId:    groupId,
		EndpointId: endpointId,
	}
}

// ParseGroupEndpointID parses 'input' into a GroupEndpointId
func ParseGroupEndpointID(input string) (*GroupEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEndpointId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EndpointId, ok = parsed.Parsed["endpointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointId", *parsed)
	}

	return &id, nil
}

// ParseGroupEndpointIDInsensitively parses 'input' case-insensitively into a GroupEndpointId
// note: this method should only be used for API response data and not user input
func ParseGroupEndpointIDInsensitively(input string) (*GroupEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEndpointId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EndpointId, ok = parsed.Parsed["endpointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointId", *parsed)
	}

	return &id, nil
}

// ValidateGroupEndpointID checks that 'input' can be parsed as a Group Endpoint ID
func ValidateGroupEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Endpoint ID
func (id GroupEndpointId) ID() string {
	fmtString := "/groups/%s/endpoints/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EndpointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Endpoint ID
func (id GroupEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEndpoints", "endpoints", "endpoints"),
		resourceids.UserSpecifiedSegment("endpointId", "endpointIdValue"),
	}
}

// String returns a human-readable description of this Group Endpoint ID
func (id GroupEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Endpoint: %q", id.EndpointId),
	}
	return fmt.Sprintf("Group Endpoint (%s)", strings.Join(components, "\n"))
}
