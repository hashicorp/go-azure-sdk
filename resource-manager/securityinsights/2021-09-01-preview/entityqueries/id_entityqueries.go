package entityqueries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EntityQueriesId{}

// EntityQueriesId is a struct representing the Resource ID for a Entity Queries
type EntityQueriesId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EntityQueryId     string
}

// NewEntityQueriesID returns a new EntityQueriesId struct
func NewEntityQueriesID(subscriptionId string, resourceGroupName string, workspaceName string, entityQueryId string) EntityQueriesId {
	return EntityQueriesId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EntityQueryId:     entityQueryId,
	}
}

// ParseEntityQueriesID parses 'input' into a EntityQueriesId
func ParseEntityQueriesID(input string) (*EntityQueriesId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntityQueriesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntityQueriesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityQueryId, ok = parsed.Parsed["entityQueryId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityQueryId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEntityQueriesIDInsensitively parses 'input' case-insensitively into a EntityQueriesId
// note: this method should only be used for API response data and not user input
func ParseEntityQueriesIDInsensitively(input string) (*EntityQueriesId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntityQueriesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntityQueriesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EntityQueryId, ok = parsed.Parsed["entityQueryId"]; !ok {
		return nil, fmt.Errorf("the segment 'entityQueryId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEntityQueriesID checks that 'input' can be parsed as a Entity Queries ID
func ValidateEntityQueriesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEntityQueriesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Entity Queries ID
func (id EntityQueriesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/entityQueries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EntityQueryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Entity Queries ID
func (id EntityQueriesId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticEntityQueries", "entityQueries", "entityQueries"),
		resourceids.UserSpecifiedSegment("entityQueryId", "entityQueryIdValue"),
	}
}

// String returns a human-readable description of this Entity Queries ID
func (id EntityQueriesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Entity Query: %q", id.EntityQueryId),
	}
	return fmt.Sprintf("Entity Queries (%s)", strings.Join(components, "\n"))
}
