package incidentteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IncidentId{}

// IncidentId is a struct representing the Resource ID for a Incident
type IncidentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IncidentId        string
}

// NewIncidentID returns a new IncidentId struct
func NewIncidentID(subscriptionId string, resourceGroupName string, workspaceName string, incidentId string) IncidentId {
	return IncidentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IncidentId:        incidentId,
	}
}

// ParseIncidentID parses 'input' into a IncidentId
func ParseIncidentID(input string) (*IncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(IncidentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IncidentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseIncidentIDInsensitively parses 'input' case-insensitively into a IncidentId
// note: this method should only be used for API response data and not user input
func ParseIncidentIDInsensitively(input string) (*IncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(IncidentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IncidentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateIncidentID checks that 'input' can be parsed as a Incident ID
func ValidateIncidentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIncidentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Incident ID
func (id IncidentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/incidents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Incident ID
func (id IncidentId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Incident ID
func (id IncidentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Incident: %q", id.IncidentId),
	}
	return fmt.Sprintf("Incident (%s)", strings.Join(components, "\n"))
}
