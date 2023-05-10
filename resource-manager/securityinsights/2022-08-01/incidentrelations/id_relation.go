package incidentrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RelationId{}

// RelationId is a struct representing the Resource ID for a Relation
type RelationId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IncidentId        string
	RelationName      string
}

// NewRelationID returns a new RelationId struct
func NewRelationID(subscriptionId string, resourceGroupName string, workspaceName string, incidentId string, relationName string) RelationId {
	return RelationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IncidentId:        incidentId,
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
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "incidentId", *parsed)
	}

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationName", *parsed)
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
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "incidentId", *parsed)
	}

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relationName", *parsed)
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/incidents/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentId, id.RelationName)
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
		resourceids.StaticSegment("staticIncidents", "incidents", "incidents"),
		resourceids.UserSpecifiedSegment("incidentId", "incidentIdValue"),
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
		fmt.Sprintf("Incident: %q", id.IncidentId),
		fmt.Sprintf("Relation Name: %q", id.RelationName),
	}
	return fmt.Sprintf("Relation (%s)", strings.Join(components, "\n"))
}
