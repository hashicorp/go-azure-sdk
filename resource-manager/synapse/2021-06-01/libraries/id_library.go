package libraries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LibraryId{})
}

var _ resourceids.ResourceId = &LibraryId{}

// LibraryId is a struct representing the Resource ID for a Library
type LibraryId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	LibraryName       string
}

// NewLibraryID returns a new LibraryId struct
func NewLibraryID(subscriptionId string, resourceGroupName string, workspaceName string, libraryName string) LibraryId {
	return LibraryId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		LibraryName:       libraryName,
	}
}

// ParseLibraryID parses 'input' into a LibraryId
func ParseLibraryID(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLibraryIDInsensitively parses 'input' case-insensitively into a LibraryId
// note: this method should only be used for API response data and not user input
func ParseLibraryIDInsensitively(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LibraryId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LibraryName, ok = input.Parsed["libraryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "libraryName", input)
	}

	return nil
}

// ValidateLibraryID checks that 'input' can be parsed as a Library ID
func ValidateLibraryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLibraryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Library ID
func (id LibraryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/libraries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.LibraryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Library ID
func (id LibraryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticLibraries", "libraries", "libraries"),
		resourceids.UserSpecifiedSegment("libraryName", "libraryName"),
	}
}

// String returns a human-readable description of this Library ID
func (id LibraryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Library Name: %q", id.LibraryName),
	}
	return fmt.Sprintf("Library (%s)", strings.Join(components, "\n"))
}
