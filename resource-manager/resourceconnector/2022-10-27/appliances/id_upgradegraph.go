package appliances

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = UpgradeGraphId{}

// UpgradeGraphId is a struct representing the Resource ID for a Upgrade Graph
type UpgradeGraphId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	UpgradeGraph      string
}

// NewUpgradeGraphID returns a new UpgradeGraphId struct
func NewUpgradeGraphID(subscriptionId string, resourceGroupName string, resourceName string, upgradeGraph string) UpgradeGraphId {
	return UpgradeGraphId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
		UpgradeGraph:      upgradeGraph,
	}
}

// ParseUpgradeGraphID parses 'input' into a UpgradeGraphId
func ParseUpgradeGraphID(input string) (*UpgradeGraphId, error) {
	parser := resourceids.NewParserFromResourceIdType(UpgradeGraphId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UpgradeGraphId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.UpgradeGraph, ok = parsed.Parsed["upgradeGraph"]; !ok {
		return nil, fmt.Errorf("the segment 'upgradeGraph' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseUpgradeGraphIDInsensitively parses 'input' case-insensitively into a UpgradeGraphId
// note: this method should only be used for API response data and not user input
func ParseUpgradeGraphIDInsensitively(input string) (*UpgradeGraphId, error) {
	parser := resourceids.NewParserFromResourceIdType(UpgradeGraphId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UpgradeGraphId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.UpgradeGraph, ok = parsed.Parsed["upgradeGraph"]; !ok {
		return nil, fmt.Errorf("the segment 'upgradeGraph' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateUpgradeGraphID checks that 'input' can be parsed as a Upgrade Graph ID
func ValidateUpgradeGraphID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUpgradeGraphID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Upgrade Graph ID
func (id UpgradeGraphId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ResourceConnector/appliances/%s/upgradeGraphs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.UpgradeGraph)
}

// Segments returns a slice of Resource ID Segments which comprise this Upgrade Graph ID
func (id UpgradeGraphId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResourceConnector", "Microsoft.ResourceConnector", "Microsoft.ResourceConnector"),
		resourceids.StaticSegment("staticAppliances", "appliances", "appliances"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticUpgradeGraphs", "upgradeGraphs", "upgradeGraphs"),
		resourceids.UserSpecifiedSegment("upgradeGraph", "upgradeGraphValue"),
	}
}

// String returns a human-readable description of this Upgrade Graph ID
func (id UpgradeGraphId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Upgrade Graph: %q", id.UpgradeGraph),
	}
	return fmt.Sprintf("Upgrade Graph (%s)", strings.Join(components, "\n"))
}
