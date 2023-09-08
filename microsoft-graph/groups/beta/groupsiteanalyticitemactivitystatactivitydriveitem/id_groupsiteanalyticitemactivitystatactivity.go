package groupsiteanalyticitemactivitystatactivitydriveitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteAnalyticItemActivityStatActivityId{}

// GroupSiteAnalyticItemActivityStatActivityId is a struct representing the Resource ID for a Group Site Analytic Item Activity Stat Activity
type GroupSiteAnalyticItemActivityStatActivityId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewGroupSiteAnalyticItemActivityStatActivityID returns a new GroupSiteAnalyticItemActivityStatActivityId struct
func NewGroupSiteAnalyticItemActivityStatActivityID(groupId string, siteId string, itemActivityStatId string, itemActivityId string) GroupSiteAnalyticItemActivityStatActivityId {
	return GroupSiteAnalyticItemActivityStatActivityId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseGroupSiteAnalyticItemActivityStatActivityID parses 'input' into a GroupSiteAnalyticItemActivityStatActivityId
func ParseGroupSiteAnalyticItemActivityStatActivityID(input string) (*GroupSiteAnalyticItemActivityStatActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteAnalyticItemActivityStatActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteAnalyticItemActivityStatActivityId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ItemActivityStatId, ok = parsed.Parsed["itemActivityStatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", *parsed)
	}

	if id.ItemActivityId, ok = parsed.Parsed["itemActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteAnalyticItemActivityStatActivityIDInsensitively parses 'input' case-insensitively into a GroupSiteAnalyticItemActivityStatActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteAnalyticItemActivityStatActivityIDInsensitively(input string) (*GroupSiteAnalyticItemActivityStatActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteAnalyticItemActivityStatActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteAnalyticItemActivityStatActivityId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ItemActivityStatId, ok = parsed.Parsed["itemActivityStatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", *parsed)
	}

	if id.ItemActivityId, ok = parsed.Parsed["itemActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteAnalyticItemActivityStatActivityID checks that 'input' can be parsed as a Group Site Analytic Item Activity Stat Activity ID
func ValidateGroupSiteAnalyticItemActivityStatActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteAnalyticItemActivityStatActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Analytic Item Activity Stat Activity ID
func (id GroupSiteAnalyticItemActivityStatActivityId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Analytic Item Activity Stat Activity ID
func (id GroupSiteAnalyticItemActivityStatActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticAnalytics", "analytics", "analytics"),
		resourceids.StaticSegment("staticItemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatIdValue"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityIdValue"),
	}
}

// String returns a human-readable description of this Group Site Analytic Item Activity Stat Activity ID
func (id GroupSiteAnalyticItemActivityStatActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Group Site Analytic Item Activity Stat Activity (%s)", strings.Join(components, "\n"))
}
