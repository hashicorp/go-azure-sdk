package groupsiteanalyticitemactivitystat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteAnalyticItemActivityStatId{}

// GroupSiteAnalyticItemActivityStatId is a struct representing the Resource ID for a Group Site Analytic Item Activity Stat
type GroupSiteAnalyticItemActivityStatId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
}

// NewGroupSiteAnalyticItemActivityStatID returns a new GroupSiteAnalyticItemActivityStatId struct
func NewGroupSiteAnalyticItemActivityStatID(groupId string, siteId string, itemActivityStatId string) GroupSiteAnalyticItemActivityStatId {
	return GroupSiteAnalyticItemActivityStatId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseGroupSiteAnalyticItemActivityStatID parses 'input' into a GroupSiteAnalyticItemActivityStatId
func ParseGroupSiteAnalyticItemActivityStatID(input string) (*GroupSiteAnalyticItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteAnalyticItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteAnalyticItemActivityStatId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ItemActivityStatId, ok = parsed.Parsed["itemActivityStatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteAnalyticItemActivityStatIDInsensitively parses 'input' case-insensitively into a GroupSiteAnalyticItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteAnalyticItemActivityStatIDInsensitively(input string) (*GroupSiteAnalyticItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteAnalyticItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteAnalyticItemActivityStatId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ItemActivityStatId, ok = parsed.Parsed["itemActivityStatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteAnalyticItemActivityStatID checks that 'input' can be parsed as a Group Site Analytic Item Activity Stat ID
func ValidateGroupSiteAnalyticItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteAnalyticItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Analytic Item Activity Stat ID
func (id GroupSiteAnalyticItemActivityStatId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Analytic Item Activity Stat ID
func (id GroupSiteAnalyticItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticAnalytics", "analytics", "analytics"),
		resourceids.StaticSegment("staticItemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatIdValue"),
	}
}

// String returns a human-readable description of this Group Site Analytic Item Activity Stat ID
func (id GroupSiteAnalyticItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Group Site Analytic Item Activity Stat (%s)", strings.Join(components, "\n"))
}
