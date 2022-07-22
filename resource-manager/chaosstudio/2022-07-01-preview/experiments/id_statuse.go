package experiments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = StatuseId{}

// StatuseId is a struct representing the Resource ID for a Statuse
type StatuseId struct {
	SubscriptionId    string
	ResourceGroupName string
	ExperimentName    string
	StatusId          string
}

// NewStatuseID returns a new StatuseId struct
func NewStatuseID(subscriptionId string, resourceGroupName string, experimentName string, statusId string) StatuseId {
	return StatuseId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ExperimentName:    experimentName,
		StatusId:          statusId,
	}
}

// ParseStatuseID parses 'input' into a StatuseId
func ParseStatuseID(input string) (*StatuseId, error) {
	parser := resourceids.NewParserFromResourceIdType(StatuseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StatuseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, fmt.Errorf("the segment 'experimentName' was not found in the resource id %q", input)
	}

	if id.StatusId, ok = parsed.Parsed["statusId"]; !ok {
		return nil, fmt.Errorf("the segment 'statusId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseStatuseIDInsensitively parses 'input' case-insensitively into a StatuseId
// note: this method should only be used for API response data and not user input
func ParseStatuseIDInsensitively(input string) (*StatuseId, error) {
	parser := resourceids.NewParserFromResourceIdType(StatuseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StatuseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, fmt.Errorf("the segment 'experimentName' was not found in the resource id %q", input)
	}

	if id.StatusId, ok = parsed.Parsed["statusId"]; !ok {
		return nil, fmt.Errorf("the segment 'statusId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateStatuseID checks that 'input' can be parsed as a Statuse ID
func ValidateStatuseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStatuseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Statuse ID
func (id StatuseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Chaos/experiments/%s/statuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ExperimentName, id.StatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Statuse ID
func (id StatuseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticExperiments", "experiments", "experiments"),
		resourceids.UserSpecifiedSegment("experimentName", "experimentValue"),
		resourceids.StaticSegment("staticStatuses", "statuses", "statuses"),
		resourceids.UserSpecifiedSegment("statusId", "statusIdValue"),
	}
}

// String returns a human-readable description of this Statuse ID
func (id StatuseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Experiment Name: %q", id.ExperimentName),
		fmt.Sprintf("Status: %q", id.StatusId),
	}
	return fmt.Sprintf("Statuse (%s)", strings.Join(components, "\n"))
}
