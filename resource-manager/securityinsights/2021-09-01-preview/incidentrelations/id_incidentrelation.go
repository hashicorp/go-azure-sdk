package incidentrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = IncidentRelationId{}

// IncidentRelationId is a struct representing the Resource ID for a Incident Relation
type IncidentRelationId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IncidentId        string
	RelationName      string
}

// NewIncidentRelationID returns a new IncidentRelationId struct
func NewIncidentRelationID(subscriptionId string, resourceGroupName string, workspaceName string, incidentId string, relationName string) IncidentRelationId {
	return IncidentRelationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IncidentId:        incidentId,
		RelationName:      relationName,
	}
}

// ParseIncidentRelationID parses 'input' into a IncidentRelationId
func ParseIncidentRelationID(input string) (*IncidentRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(IncidentRelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IncidentRelationId{}

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

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, fmt.Errorf("the segment 'relationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseIncidentRelationIDInsensitively parses 'input' case-insensitively into a IncidentRelationId
// note: this method should only be used for API response data and not user input
func ParseIncidentRelationIDInsensitively(input string) (*IncidentRelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(IncidentRelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IncidentRelationId{}

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

	if id.RelationName, ok = parsed.Parsed["relationName"]; !ok {
		return nil, fmt.Errorf("the segment 'relationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateIncidentRelationID checks that 'input' can be parsed as a Incident Relation ID
func ValidateIncidentRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIncidentRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Incident Relation ID
func (id IncidentRelationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/incidents/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentId, id.RelationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Incident Relation ID
func (id IncidentRelationId) Segments() []resourceids.Segment {
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

// String returns a human-readable description of this Incident Relation ID
func (id IncidentRelationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Incident: %q", id.IncidentId),
		fmt.Sprintf("Relation Name: %q", id.RelationName),
	}
	return fmt.Sprintf("Incident Relation (%s)", strings.Join(components, "\n"))
}
