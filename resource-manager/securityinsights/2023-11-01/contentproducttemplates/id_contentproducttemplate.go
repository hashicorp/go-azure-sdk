package contentproducttemplates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ContentProductTemplateId{})
}

var _ resourceids.ResourceId = &ContentProductTemplateId{}

// ContentProductTemplateId is a struct representing the Resource ID for a Content Product Template
type ContentProductTemplateId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	TemplateId        string
}

// NewContentProductTemplateID returns a new ContentProductTemplateId struct
func NewContentProductTemplateID(subscriptionId string, resourceGroupName string, workspaceName string, templateId string) ContentProductTemplateId {
	return ContentProductTemplateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		TemplateId:        templateId,
	}
}

// ParseContentProductTemplateID parses 'input' into a ContentProductTemplateId
func ParseContentProductTemplateID(input string) (*ContentProductTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentProductTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentProductTemplateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContentProductTemplateIDInsensitively parses 'input' case-insensitively into a ContentProductTemplateId
// note: this method should only be used for API response data and not user input
func ParseContentProductTemplateIDInsensitively(input string) (*ContentProductTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentProductTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentProductTemplateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContentProductTemplateId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TemplateId, ok = input.Parsed["templateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "templateId", input)
	}

	return nil
}

// ValidateContentProductTemplateID checks that 'input' can be parsed as a Content Product Template ID
func ValidateContentProductTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentProductTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Product Template ID
func (id ContentProductTemplateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/contentProductTemplates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.TemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Product Template ID
func (id ContentProductTemplateId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticContentProductTemplates", "contentProductTemplates", "contentProductTemplates"),
		resourceids.UserSpecifiedSegment("templateId", "templateIdValue"),
	}
}

// String returns a human-readable description of this Content Product Template ID
func (id ContentProductTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Template: %q", id.TemplateId),
	}
	return fmt.Sprintf("Content Product Template (%s)", strings.Join(components, "\n"))
}
