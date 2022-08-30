package activity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ActivitiesId{}

// ActivitiesId is a struct representing the Resource ID for a Activities
type ActivitiesId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	ModuleName            string
	ActivityName          string
}

// NewActivitiesID returns a new ActivitiesId struct
func NewActivitiesID(subscriptionId string, resourceGroupName string, automationAccountName string, moduleName string, activityName string) ActivitiesId {
	return ActivitiesId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		ModuleName:            moduleName,
		ActivityName:          activityName,
	}
}

// ParseActivitiesID parses 'input' into a ActivitiesId
func ParseActivitiesID(input string) (*ActivitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(ActivitiesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ActivitiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, fmt.Errorf("the segment 'moduleName' was not found in the resource id %q", input)
	}

	if id.ActivityName, ok = parsed.Parsed["activityName"]; !ok {
		return nil, fmt.Errorf("the segment 'activityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseActivitiesIDInsensitively parses 'input' case-insensitively into a ActivitiesId
// note: this method should only be used for API response data and not user input
func ParseActivitiesIDInsensitively(input string) (*ActivitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(ActivitiesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ActivitiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, fmt.Errorf("the segment 'moduleName' was not found in the resource id %q", input)
	}

	if id.ActivityName, ok = parsed.Parsed["activityName"]; !ok {
		return nil, fmt.Errorf("the segment 'activityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateActivitiesID checks that 'input' can be parsed as a Activities ID
func ValidateActivitiesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseActivitiesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Activities ID
func (id ActivitiesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/modules/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.ModuleName, id.ActivityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Activities ID
func (id ActivitiesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleValue"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("activityName", "activityValue"),
	}
}

// String returns a human-readable description of this Activities ID
func (id ActivitiesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Activity Name: %q", id.ActivityName),
	}
	return fmt.Sprintf("Activities (%s)", strings.Join(components, "\n"))
}
