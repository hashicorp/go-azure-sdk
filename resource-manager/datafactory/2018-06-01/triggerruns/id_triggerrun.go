package triggerruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = TriggerRunId{}

// TriggerRunId is a struct representing the Resource ID for a Trigger Run
type TriggerRunId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	TriggerName       string
	RunId             string
}

// NewTriggerRunID returns a new TriggerRunId struct
func NewTriggerRunID(subscriptionId string, resourceGroupName string, factoryName string, triggerName string, runId string) TriggerRunId {
	return TriggerRunId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		TriggerName:       triggerName,
		RunId:             runId,
	}
}

// ParseTriggerRunID parses 'input' into a TriggerRunId
func ParseTriggerRunID(input string) (*TriggerRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggerRunId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.TriggerName, ok = parsed.Parsed["triggerName"]; !ok {
		return nil, fmt.Errorf("the segment 'triggerName' was not found in the resource id %q", input)
	}

	if id.RunId, ok = parsed.Parsed["runId"]; !ok {
		return nil, fmt.Errorf("the segment 'runId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseTriggerRunIDInsensitively parses 'input' case-insensitively into a TriggerRunId
// note: this method should only be used for API response data and not user input
func ParseTriggerRunIDInsensitively(input string) (*TriggerRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggerRunId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.TriggerName, ok = parsed.Parsed["triggerName"]; !ok {
		return nil, fmt.Errorf("the segment 'triggerName' was not found in the resource id %q", input)
	}

	if id.RunId, ok = parsed.Parsed["runId"]; !ok {
		return nil, fmt.Errorf("the segment 'runId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateTriggerRunID checks that 'input' can be parsed as a Trigger Run ID
func ValidateTriggerRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger Run ID
func (id TriggerRunId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/triggers/%s/triggerRuns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.TriggerName, id.RunId)
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger Run ID
func (id TriggerRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerValue"),
		resourceids.StaticSegment("staticTriggerRuns", "triggerRuns", "triggerRuns"),
		resourceids.UserSpecifiedSegment("runId", "runIdValue"),
	}
}

// String returns a human-readable description of this Trigger Run ID
func (id TriggerRunId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Trigger Run (%s)", strings.Join(components, "\n"))
}
