// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RelationId{}

// RelationId is a struct representing the Resource ID for a Relation
type RelationId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EntityId          string
	RelationName      string
}

// NewRelationID returns a new RelationId struct
func NewRelationID(subscriptionId string, resourceGroupName string, workspaceName string, entityId string, relationName string) RelationId {
	return RelationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EntityId:          entityId,
		RelationName:      relationName,
	}
}

// ParseRelationID parses 'input' into a RelationId
func ParseRelationID(input string) (*RelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(RelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RelationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityId, ok = parsed.Parsed["entityId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityId' was not found in the resource id %q", input)
	}

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, fmt.Errorf("the segment 'relationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRelationIDInsensitively parses 'input' case-insensitively into a RelationId
// note: this method should only be used for API response data and not user input
func ParseRelationIDInsensitively(input string) (*RelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(RelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RelationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityId, ok = parsed.Parsed["entityId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityId' was not found in the resource id %q", input)
	}

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, fmt.Errorf("the segment 'relationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRelationID checks that 'input' can be parsed as a Relation ID
func ValidateRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Relation ID
func (id RelationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/entities/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EntityId, id.RelationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Relation ID
func (id RelationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticEntities", "entities", "entities"),
		resourceids.UserSpecifiedSegment("entityId", "entityIdValue"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("relationName", "relationValue"),
	}
}

// String returns a human-readable description of this Relation ID
func (id RelationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Entity: %q", id.EntityId),
		fmt.Sprintf("Relation Name: %q", id.RelationName),
	}
	return fmt.Sprintf("Relation (%s)", strings.Join(components, "\n"))
}
