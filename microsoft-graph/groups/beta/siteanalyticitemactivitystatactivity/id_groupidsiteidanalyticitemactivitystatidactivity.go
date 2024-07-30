package siteanalyticitemactivitystatactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdAnalyticItemActivityStatIdActivityId{}

// GroupIdSiteIdAnalyticItemActivityStatIdActivityId is a struct representing the Resource ID for a Group Id Site Id Analytic Item Activity Stat Id Activity
type GroupIdSiteIdAnalyticItemActivityStatIdActivityId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewGroupIdSiteIdAnalyticItemActivityStatIdActivityID returns a new GroupIdSiteIdAnalyticItemActivityStatIdActivityId struct
func NewGroupIdSiteIdAnalyticItemActivityStatIdActivityID(groupId string, siteId string, itemActivityStatId string, itemActivityId string) GroupIdSiteIdAnalyticItemActivityStatIdActivityId {
	return GroupIdSiteIdAnalyticItemActivityStatIdActivityId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseGroupIdSiteIdAnalyticItemActivityStatIdActivityID parses 'input' into a GroupIdSiteIdAnalyticItemActivityStatIdActivityId
func ParseGroupIdSiteIdAnalyticItemActivityStatIdActivityID(input string) (*GroupIdSiteIdAnalyticItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdAnalyticItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdAnalyticItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdAnalyticItemActivityStatIdActivityIDInsensitively(input string) (*GroupIdSiteIdAnalyticItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdAnalyticItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdAnalyticItemActivityStatIdActivityID checks that 'input' can be parsed as a Group Id Site Id Analytic Item Activity Stat Id Activity ID
func ValidateGroupIdSiteIdAnalyticItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdAnalyticItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Analytic Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticItemActivityStatIdActivityId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Analytic Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatIdValue"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Analytic Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Group Id Site Id Analytic Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
