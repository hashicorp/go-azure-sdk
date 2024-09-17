package codeversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CodeVersionId{})
}

var _ resourceids.ResourceId = &CodeVersionId{}

// CodeVersionId is a struct representing the Resource ID for a Code Version
type CodeVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	CodeName          string
	VersionName       string
}

// NewCodeVersionID returns a new CodeVersionId struct
func NewCodeVersionID(subscriptionId string, resourceGroupName string, workspaceName string, codeName string, versionName string) CodeVersionId {
	return CodeVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		CodeName:          codeName,
		VersionName:       versionName,
	}
}

// ParseCodeVersionID parses 'input' into a CodeVersionId
func ParseCodeVersionID(input string) (*CodeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CodeVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CodeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCodeVersionIDInsensitively parses 'input' case-insensitively into a CodeVersionId
// note: this method should only be used for API response data and not user input
func ParseCodeVersionIDInsensitively(input string) (*CodeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CodeVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CodeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CodeVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CodeName, ok = input.Parsed["codeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "codeName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateCodeVersionID checks that 'input' can be parsed as a Code Version ID
func ValidateCodeVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCodeVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Code Version ID
func (id CodeVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/codes/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.CodeName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Code Version ID
func (id CodeVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticCodes", "codes", "codes"),
		resourceids.UserSpecifiedSegment("codeName", "codeValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Code Version ID
func (id CodeVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Code Name: %q", id.CodeName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Code Version (%s)", strings.Join(components, "\n"))
}
