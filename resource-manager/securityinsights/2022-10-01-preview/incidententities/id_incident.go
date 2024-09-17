package incidententities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IncidentId{})
}

var _ resourceids.ResourceId = &IncidentId{}

// IncidentId is a struct representing the Resource ID for a Incident
type IncidentId struct {
	SubscriptionId     string
	ResourceGroupName  string
	WorkspaceName      string
	IncidentIdentifier string
}

// NewIncidentID returns a new IncidentId struct
func NewIncidentID(subscriptionId string, resourceGroupName string, workspaceName string, incidentIdentifier string) IncidentId {
	return IncidentId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		WorkspaceName:      workspaceName,
		IncidentIdentifier: incidentIdentifier,
	}
}

// ParseIncidentID parses 'input' into a IncidentId
func ParseIncidentID(input string) (*IncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IncidentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IncidentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIncidentIDInsensitively parses 'input' case-insensitively into a IncidentId
// note: this method should only be used for API response data and not user input
func ParseIncidentIDInsensitively(input string) (*IncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IncidentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IncidentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IncidentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.IncidentIdentifier, ok = input.Parsed["incidentIdentifier"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "incidentIdentifier", input)
	}

	return nil
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentIdentifier)
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
		resourceids.UserSpecifiedSegment("incidentIdentifier", "incidentIdentifierValue"),
	}
}

// String returns a human-readable description of this Incident ID
func (id IncidentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Incident Identifier: %q", id.IncidentIdentifier),
	}
	return fmt.Sprintf("Incident (%s)", strings.Join(components, "\n"))
}
