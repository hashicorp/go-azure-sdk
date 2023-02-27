package entityqueries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = EntityQueryTemplateId{}

// EntityQueryTemplateId is a struct representing the Resource ID for a Entity Query Template
type EntityQueryTemplateId struct {
	SubscriptionId        string
	ResourceGroupName     string
	WorkspaceName         string
	EntityQueryTemplateId string
}

// NewEntityQueryTemplateID returns a new EntityQueryTemplateId struct
func NewEntityQueryTemplateID(subscriptionId string, resourceGroupName string, workspaceName string, entityQueryTemplateId string) EntityQueryTemplateId {
	return EntityQueryTemplateId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		WorkspaceName:         workspaceName,
		EntityQueryTemplateId: entityQueryTemplateId,
	}
}

// ParseEntityQueryTemplateID parses 'input' into a EntityQueryTemplateId
func ParseEntityQueryTemplateID(input string) (*EntityQueryTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntityQueryTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntityQueryTemplateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityQueryTemplateId, ok = parsed.Parsed["entityQueryTemplateId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityQueryTemplateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEntityQueryTemplateIDInsensitively parses 'input' case-insensitively into a EntityQueryTemplateId
// note: this method should only be used for API response data and not user input
func ParseEntityQueryTemplateIDInsensitively(input string) (*EntityQueryTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntityQueryTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntityQueryTemplateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityQueryTemplateId, ok = parsed.Parsed["entityQueryTemplateId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityQueryTemplateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEntityQueryTemplateID checks that 'input' can be parsed as a Entity Query Template ID
func ValidateEntityQueryTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEntityQueryTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Entity Query Template ID
func (id EntityQueryTemplateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/entityQueryTemplates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EntityQueryTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Entity Query Template ID
func (id EntityQueryTemplateId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticEntityQueryTemplates", "entityQueryTemplates", "entityQueryTemplates"),
		resourceids.UserSpecifiedSegment("entityQueryTemplateId", "entityQueryTemplateIdValue"),
	}
}

// String returns a human-readable description of this Entity Query Template ID
func (id EntityQueryTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Entity Query Template: %q", id.EntityQueryTemplateId),
	}
	return fmt.Sprintf("Entity Query Template (%s)", strings.Join(components, "\n"))
}
