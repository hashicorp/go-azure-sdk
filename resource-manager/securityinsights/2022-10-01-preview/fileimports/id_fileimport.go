package fileimports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FileImportId{}

// FileImportId is a struct representing the Resource ID for a File Import
type FileImportId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	FileImportId      string
}

// NewFileImportID returns a new FileImportId struct
func NewFileImportID(subscriptionId string, resourceGroupName string, workspaceName string, fileImportId string) FileImportId {
	return FileImportId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		FileImportId:      fileImportId,
	}
}

// ParseFileImportID parses 'input' into a FileImportId
func ParseFileImportID(input string) (*FileImportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileImportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileImportId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFileImportIDInsensitively parses 'input' case-insensitively into a FileImportId
// note: this method should only be used for API response data and not user input
func ParseFileImportIDInsensitively(input string) (*FileImportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileImportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileImportId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FileImportId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.FileImportId, ok = input.Parsed["fileImportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fileImportId", input)
	}

	return nil
}

// ValidateFileImportID checks that 'input' can be parsed as a File Import ID
func ValidateFileImportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFileImportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted File Import ID
func (id FileImportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/fileImports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FileImportId)
}

// Segments returns a slice of Resource ID Segments which comprise this File Import ID
func (id FileImportId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticFileImports", "fileImports", "fileImports"),
		resourceids.UserSpecifiedSegment("fileImportId", "fileImportIdValue"),
	}
}

// String returns a human-readable description of this File Import ID
func (id FileImportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("File Import: %q", id.FileImportId),
	}
	return fmt.Sprintf("File Import (%s)", strings.Join(components, "\n"))
}
