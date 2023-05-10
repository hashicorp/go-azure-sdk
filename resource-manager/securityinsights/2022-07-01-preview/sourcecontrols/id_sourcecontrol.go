package sourcecontrols

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SourceControlId{}

// SourceControlId is a struct representing the Resource ID for a Source Control
type SourceControlId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SourceControlId   string
}

// NewSourceControlID returns a new SourceControlId struct
func NewSourceControlID(subscriptionId string, resourceGroupName string, workspaceName string, sourceControlId string) SourceControlId {
	return SourceControlId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SourceControlId:   sourceControlId,
	}
}

// ParseSourceControlID parses 'input' into a SourceControlId
func ParseSourceControlID(input string) (*SourceControlId, error) {
	parser := resourceids.NewParserFromResourceIdType(SourceControlId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SourceControlId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.SourceControlId, ok = parsed.Parsed["sourceControlId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlId", *parsed)
	}

	return &id, nil
}

// ParseSourceControlIDInsensitively parses 'input' case-insensitively into a SourceControlId
// note: this method should only be used for API response data and not user input
func ParseSourceControlIDInsensitively(input string) (*SourceControlId, error) {
	parser := resourceids.NewParserFromResourceIdType(SourceControlId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SourceControlId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.SourceControlId, ok = parsed.Parsed["sourceControlId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlId", *parsed)
	}

	return &id, nil
}

// ValidateSourceControlID checks that 'input' can be parsed as a Source Control ID
func ValidateSourceControlID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSourceControlID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Source Control ID
func (id SourceControlId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/sourceControls/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SourceControlId)
}

// Segments returns a slice of Resource ID Segments which comprise this Source Control ID
func (id SourceControlId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticSourceControls", "sourceControls", "sourceControls"),
		resourceids.UserSpecifiedSegment("sourceControlId", "sourceControlIdValue"),
	}
}

// String returns a human-readable description of this Source Control ID
func (id SourceControlId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Source Control: %q", id.SourceControlId),
	}
	return fmt.Sprintf("Source Control (%s)", strings.Join(components, "\n"))
}
