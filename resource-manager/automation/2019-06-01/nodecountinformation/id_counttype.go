package nodecountinformation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CountTypeId{}

// CountTypeId is a struct representing the Resource ID for a Count Type
type CountTypeId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	CountType             CountType
}

// NewCountTypeID returns a new CountTypeId struct
func NewCountTypeID(subscriptionId string, resourceGroupName string, automationAccountName string, countType CountType) CountTypeId {
	return CountTypeId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		CountType:             countType,
	}
}

// ParseCountTypeID parses 'input' into a CountTypeId
func ParseCountTypeID(input string) (*CountTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(CountTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CountTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["countType"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'countType' was not found in the resource id %q", input)
		}

		countType, err := parseCountType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.CountType = *countType
	}

	return &id, nil
}

// ParseCountTypeIDInsensitively parses 'input' case-insensitively into a CountTypeId
// note: this method should only be used for API response data and not user input
func ParseCountTypeIDInsensitively(input string) (*CountTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(CountTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CountTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["countType"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'countType' was not found in the resource id %q", input)
		}

		countType, err := parseCountType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.CountType = *countType
	}

	return &id, nil
}

// ValidateCountTypeID checks that 'input' can be parsed as a Count Type ID
func ValidateCountTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCountTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Count Type ID
func (id CountTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/nodecounts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, string(id.CountType))
}

// Segments returns a slice of Resource ID Segments which comprise this Count Type ID
func (id CountTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticNodecounts", "nodecounts", "nodecounts"),
		resourceids.ConstantSegment("countType", PossibleValuesForCountType(), "nodeconfiguration"),
	}
}

// String returns a human-readable description of this Count Type ID
func (id CountTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Count Type: %q", string(id.CountType)),
	}
	return fmt.Sprintf("Count Type (%s)", strings.Join(components, "\n"))
}
