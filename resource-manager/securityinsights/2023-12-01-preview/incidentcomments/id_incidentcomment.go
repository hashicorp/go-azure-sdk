package incidentcomments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IncidentCommentId{})
}

var _ resourceids.ResourceId = &IncidentCommentId{}

// IncidentCommentId is a struct representing the Resource ID for a Incident Comment
type IncidentCommentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IncidentId        string
	IncidentCommentId string
}

// NewIncidentCommentID returns a new IncidentCommentId struct
func NewIncidentCommentID(subscriptionId string, resourceGroupName string, workspaceName string, incidentId string, incidentCommentId string) IncidentCommentId {
	return IncidentCommentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IncidentId:        incidentId,
		IncidentCommentId: incidentCommentId,
	}
}

// ParseIncidentCommentID parses 'input' into a IncidentCommentId
func ParseIncidentCommentID(input string) (*IncidentCommentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IncidentCommentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IncidentCommentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIncidentCommentIDInsensitively parses 'input' case-insensitively into a IncidentCommentId
// note: this method should only be used for API response data and not user input
func ParseIncidentCommentIDInsensitively(input string) (*IncidentCommentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IncidentCommentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IncidentCommentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IncidentCommentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.IncidentId, ok = input.Parsed["incidentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "incidentId", input)
	}

	if id.IncidentCommentId, ok = input.Parsed["incidentCommentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "incidentCommentId", input)
	}

	return nil
}

// ValidateIncidentCommentID checks that 'input' can be parsed as a Incident Comment ID
func ValidateIncidentCommentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIncidentCommentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Incident Comment ID
func (id IncidentCommentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/incidents/%s/comments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentId, id.IncidentCommentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Incident Comment ID
func (id IncidentCommentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticIncidents", "incidents", "incidents"),
		resourceids.UserSpecifiedSegment("incidentId", "incidentId"),
		resourceids.StaticSegment("staticComments", "comments", "comments"),
		resourceids.UserSpecifiedSegment("incidentCommentId", "incidentCommentId"),
	}
}

// String returns a human-readable description of this Incident Comment ID
func (id IncidentCommentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Incident: %q", id.IncidentId),
		fmt.Sprintf("Incident Comment: %q", id.IncidentCommentId),
	}
	return fmt.Sprintf("Incident Comment (%s)", strings.Join(components, "\n"))
}
