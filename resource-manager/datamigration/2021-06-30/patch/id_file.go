package patch

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&FileId{})
}

var _ resourceids.ResourceId = &FileId{}

// FileId is a struct representing the Resource ID for a File
type FileId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ProjectName       string
	FileName          string
}

// NewFileID returns a new FileId struct
func NewFileID(subscriptionId string, resourceGroupName string, serviceName string, projectName string, fileName string) FileId {
	return FileId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ProjectName:       projectName,
		FileName:          fileName,
	}
}

// ParseFileID parses 'input' into a FileId
func ParseFileID(input string) (*FileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFileIDInsensitively parses 'input' case-insensitively into a FileId
// note: this method should only be used for API response data and not user input
func ParseFileIDInsensitively(input string) (*FileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServiceName, ok = input.Parsed["serviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceName", input)
	}

	if id.ProjectName, ok = input.Parsed["projectName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectName", input)
	}

	if id.FileName, ok = input.Parsed["fileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fileName", input)
	}

	return nil
}

// ValidateFileID checks that 'input' can be parsed as a File ID
func ValidateFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted File ID
func (id FileId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataMigration/services/%s/projects/%s/files/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ProjectName, id.FileName)
}

// Segments returns a slice of Resource ID Segments which comprise this File ID
func (id FileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.UserSpecifiedSegment("resourceGroupName", "groupName"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticServices", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceName"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectName"),
		resourceids.StaticSegment("staticFiles", "files", "files"),
		resourceids.UserSpecifiedSegment("fileName", "fileName"),
	}
}

// String returns a human-readable description of this File ID
func (id FileId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("File Name: %q", id.FileName),
	}
	return fmt.Sprintf("File (%s)", strings.Join(components, "\n"))
}
