// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bookmarks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BookmarkId{}

// BookmarkId is a struct representing the Resource ID for a Bookmark
type BookmarkId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	BookmarkId        string
}

// NewBookmarkID returns a new BookmarkId struct
func NewBookmarkID(subscriptionId string, resourceGroupName string, workspaceName string, bookmarkId string) BookmarkId {
	return BookmarkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		BookmarkId:        bookmarkId,
	}
}

// ParseBookmarkID parses 'input' into a BookmarkId
func ParseBookmarkID(input string) (*BookmarkId, error) {
	parser := resourceids.NewParserFromResourceIdType(BookmarkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BookmarkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.BookmarkId, ok = parsed.Parsed["bookmarkId"]; !ok {
		return nil, fmt.Errorf("the segment 'bookmarkId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBookmarkIDInsensitively parses 'input' case-insensitively into a BookmarkId
// note: this method should only be used for API response data and not user input
func ParseBookmarkIDInsensitively(input string) (*BookmarkId, error) {
	parser := resourceids.NewParserFromResourceIdType(BookmarkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BookmarkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.BookmarkId, ok = parsed.Parsed["bookmarkId"]; !ok {
		return nil, fmt.Errorf("the segment 'bookmarkId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBookmarkID checks that 'input' can be parsed as a Bookmark ID
func ValidateBookmarkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBookmarkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bookmark ID
func (id BookmarkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/bookmarks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BookmarkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Bookmark ID
func (id BookmarkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticBookmarks", "bookmarks", "bookmarks"),
		resourceids.UserSpecifiedSegment("bookmarkId", "bookmarkIdValue"),
	}
}

// String returns a human-readable description of this Bookmark ID
func (id BookmarkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Bookmark: %q", id.BookmarkId),
	}
	return fmt.Sprintf("Bookmark (%s)", strings.Join(components, "\n"))
}
