package contenttemplates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ContentTemplateId{})
}

var _ resourceids.ResourceId = &ContentTemplateId{}

// ContentTemplateId is a struct representing the Resource ID for a Content Template
type ContentTemplateId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	TemplateId        string
}

// NewContentTemplateID returns a new ContentTemplateId struct
func NewContentTemplateID(subscriptionId string, resourceGroupName string, workspaceName string, templateId string) ContentTemplateId {
	return ContentTemplateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		TemplateId:        templateId,
	}
}

// ParseContentTemplateID parses 'input' into a ContentTemplateId
func ParseContentTemplateID(input string) (*ContentTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContentTemplateIDInsensitively parses 'input' case-insensitively into a ContentTemplateId
// note: this method should only be used for API response data and not user input
func ParseContentTemplateIDInsensitively(input string) (*ContentTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContentTemplateId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateContentTemplateID checks that 'input' can be parsed as a Content Template ID
func ValidateContentTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Template ID
func (id ContentTemplateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/contentTemplates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.TemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Template ID
func (id ContentTemplateId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticContentTemplates", "contentTemplates", "contentTemplates"),
		resourceids.UserSpecifiedSegment("templateId", "templateIdValue"),
	}
}

// String returns a human-readable description of this Content Template ID
func (id ContentTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Template: %q", id.TemplateId),
	}
	return fmt.Sprintf("Content Template (%s)", strings.Join(components, "\n"))
}
