package groupsitelistsubscription

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListSubscriptionId{}

// GroupSiteListSubscriptionId is a struct representing the Resource ID for a Group Site List Subscription
type GroupSiteListSubscriptionId struct {
	GroupId        string
	SiteId         string
	ListId         string
	SubscriptionId string
}

// NewGroupSiteListSubscriptionID returns a new GroupSiteListSubscriptionId struct
func NewGroupSiteListSubscriptionID(groupId string, siteId string, listId string, subscriptionId string) GroupSiteListSubscriptionId {
	return GroupSiteListSubscriptionId{
		GroupId:        groupId,
		SiteId:         siteId,
		ListId:         listId,
		SubscriptionId: subscriptionId,
	}
}

// ParseGroupSiteListSubscriptionID parses 'input' into a GroupSiteListSubscriptionId
func ParseGroupSiteListSubscriptionID(input string) (*GroupSiteListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListSubscriptionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteListSubscriptionIDInsensitively parses 'input' case-insensitively into a GroupSiteListSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteListSubscriptionIDInsensitively(input string) (*GroupSiteListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteListSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteListSubscriptionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ListId, ok = parsed.Parsed["listId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "listId", *parsed)
	}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteListSubscriptionID checks that 'input' can be parsed as a Group Site List Subscription ID
func ValidateGroupSiteListSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteListSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site List Subscription ID
func (id GroupSiteListSubscriptionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site List Subscription ID
func (id GroupSiteListSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listIdValue"),
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionIdValue"),
	}
}

// String returns a human-readable description of this Group Site List Subscription ID
func (id GroupSiteListSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Group Site List Subscription (%s)", strings.Join(components, "\n"))
}
