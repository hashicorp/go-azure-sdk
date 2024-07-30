package siteanalyticitemactivitystatactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdAnalyticItemActivityStatId{}

// GroupIdSiteIdAnalyticItemActivityStatId is a struct representing the Resource ID for a Group Id Site Id Analytic Item Activity Stat
type GroupIdSiteIdAnalyticItemActivityStatId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
}

// NewGroupIdSiteIdAnalyticItemActivityStatID returns a new GroupIdSiteIdAnalyticItemActivityStatId struct
func NewGroupIdSiteIdAnalyticItemActivityStatID(groupId string, siteId string, itemActivityStatId string) GroupIdSiteIdAnalyticItemActivityStatId {
	return GroupIdSiteIdAnalyticItemActivityStatId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseGroupIdSiteIdAnalyticItemActivityStatID parses 'input' into a GroupIdSiteIdAnalyticItemActivityStatId
func ParseGroupIdSiteIdAnalyticItemActivityStatID(input string) (*GroupIdSiteIdAnalyticItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdAnalyticItemActivityStatIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdAnalyticItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdAnalyticItemActivityStatIDInsensitively(input string) (*GroupIdSiteIdAnalyticItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdAnalyticItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdAnalyticItemActivityStatID checks that 'input' can be parsed as a Group Id Site Id Analytic Item Activity Stat ID
func ValidateGroupIdSiteIdAnalyticItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdAnalyticItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Analytic Item Activity Stat ID
func (id GroupIdSiteIdAnalyticItemActivityStatId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Analytic Item Activity Stat ID
func (id GroupIdSiteIdAnalyticItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatIdValue"),
	}
}

// String returns a human-readable description of this Group Id Site Id Analytic Item Activity Stat ID
func (id GroupIdSiteIdAnalyticItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Group Id Site Id Analytic Item Activity Stat (%s)", strings.Join(components, "\n"))
}
