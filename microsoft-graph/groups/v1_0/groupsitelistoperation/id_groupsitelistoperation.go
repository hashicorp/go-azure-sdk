package groupsitelistoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListOperationId{}

// GroupSiteListOperationId is a struct representing the Resource ID for a Group Site List Operation
type GroupSiteListOperationId struct {
	GroupId                    string
	SiteId                     string
	ListId                     string
	RichLongRunningOperationId string
}

// NewGroupSiteListOperationID returns a new GroupSiteListOperationId struct
func NewGroupSiteListOperationID(groupId string, siteId string, listId string, richLongRunningOperationId string) GroupSiteListOperationId {
	return GroupSiteListOperationId{
		GroupId:                    groupId,
		SiteId:                     siteId,
		ListId:                     listId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseGroupSiteListOperationID parses 'input' into a GroupSiteListOperationId
func ParseGroupSiteListOperationID(input string) (*GroupSiteListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.RichLongRunningOperationId, ok = parsed.Parsed["richLongRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListOperationIDInsensitively parses 'input' case-insensitively into a GroupSiteListOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListOperationIDInsensitively(input string) (*GroupSiteListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.RichLongRunningOperationId, ok = parsed.Parsed["richLongRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListOperationID checks that 'input' can be parsed as a Group Site List Operation ID
func ValidateGroupSiteListOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Operation ID
func (id GroupSiteListOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Operation ID
func (id GroupSiteListOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Operation ID
func (id GroupSiteListOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Group Site List Operation (%s)", strings.Join(components, "\n"))
}
