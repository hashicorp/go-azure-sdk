package bookmarkrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BookmarkRelationId{})
}

var _ resourceids.ResourceId = &BookmarkRelationId{}

// BookmarkRelationId is a struct representing the Resource ID for a Bookmark Relation
type BookmarkRelationId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	BookmarkId        string
	RelationName      string
}

// NewBookmarkRelationID returns a new BookmarkRelationId struct
func NewBookmarkRelationID(subscriptionId string, resourceGroupName string, workspaceName string, bookmarkId string, relationName string) BookmarkRelationId {
	return BookmarkRelationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		BookmarkId:        bookmarkId,
		RelationName:      relationName,
	}
}

// ParseBookmarkRelationID parses 'input' into a BookmarkRelationId
func ParseBookmarkRelationID(input string) (*BookmarkRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BookmarkRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BookmarkRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBookmarkRelationIDInsensitively parses 'input' case-insensitively into a BookmarkRelationId
// note: this method should only be used for API response data and not user input
func ParseBookmarkRelationIDInsensitively(input string) (*BookmarkRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BookmarkRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BookmarkRelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BookmarkRelationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.BookmarkId, ok = input.Parsed["bookmarkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bookmarkId", input)
	}

	if id.RelationName, ok = input.Parsed["relationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationName", input)
	}

	return nil
}

// ValidateBookmarkRelationID checks that 'input' can be parsed as a Bookmark Relation ID
func ValidateBookmarkRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBookmarkRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bookmark Relation ID
func (id BookmarkRelationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/bookmarks/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BookmarkId, id.RelationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Bookmark Relation ID
func (id BookmarkRelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticBookmarks", "bookmarks", "bookmarks"),
		resourceids.UserSpecifiedSegment("bookmarkId", "bookmarkId"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationName", "relationName"),
	}
}

// String returns a human-readable description of this Bookmark Relation ID
func (id BookmarkRelationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Bookmark: %q", id.BookmarkId),
		fmt.Sprintf("Relation Name: %q", id.RelationName),
	}
	return fmt.Sprintf("Bookmark Relation (%s)", strings.Join(components, "\n"))
}
