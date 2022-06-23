package entityrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EntitiesId{}

// EntitiesId is a struct representing the Resource ID for a Entities
type EntitiesId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EntityId          string
}

// NewEntitiesID returns a new EntitiesId struct
func NewEntitiesID(subscriptionId string, resourceGroupName string, workspaceName string, entityId string) EntitiesId {
	return EntitiesId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EntityId:          entityId,
	}
}

// ParseEntitiesID parses 'input' into a EntitiesId
func ParseEntitiesID(input string) (*EntitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntitiesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntitiesId{}

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

	return &id, nil
}

// ParseEntitiesIDInsensitively parses 'input' case-insensitively into a EntitiesId
// note: this method should only be used for API response data and not user input
func ParseEntitiesIDInsensitively(input string) (*EntitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(EntitiesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EntitiesId{}

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

	return &id, nil
}

// ValidateEntitiesID checks that 'input' can be parsed as a Entities ID
func ValidateEntitiesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEntitiesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Entities ID
func (id EntitiesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/entities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EntityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Entities ID
func (id EntitiesId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Entities ID
func (id EntitiesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Entity: %q", id.EntityId),
	}
	return fmt.Sprintf("Entities (%s)", strings.Join(components, "\n"))
}
