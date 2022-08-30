package dsccompilationjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CompilationjobId{}

// CompilationjobId is a struct representing the Resource ID for a Compilationjob
type CompilationjobId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	CompilationJobName    string
}

// NewCompilationjobID returns a new CompilationjobId struct
func NewCompilationjobID(subscriptionId string, resourceGroupName string, automationAccountName string, compilationJobName string) CompilationjobId {
	return CompilationjobId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		CompilationJobName:    compilationJobName,
	}
}

// ParseCompilationjobID parses 'input' into a CompilationjobId
func ParseCompilationjobID(input string) (*CompilationjobId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationjobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationjobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.CompilationJobName, ok = parsed.Parsed["compilationJobName"]; !ok {
		return nil, fmt.Errorf("the segment 'compilationJobName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCompilationjobIDInsensitively parses 'input' case-insensitively into a CompilationjobId
// note: this method should only be used for API response data and not user input
func ParseCompilationjobIDInsensitively(input string) (*CompilationjobId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationjobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationjobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.CompilationJobName, ok = parsed.Parsed["compilationJobName"]; !ok {
		return nil, fmt.Errorf("the segment 'compilationJobName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCompilationjobID checks that 'input' can be parsed as a Compilationjob ID
func ValidateCompilationjobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCompilationjobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Compilationjob ID
func (id CompilationjobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/compilationjobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.CompilationJobName)
}

// Segments returns a slice of Resource ID Segments which comprise this Compilationjob ID
func (id CompilationjobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticCompilationjobs", "compilationjobs", "compilationjobs"),
		resourceids.UserSpecifiedSegment("compilationJobName", "compilationJobValue"),
	}
}

// String returns a human-readable description of this Compilationjob ID
func (id CompilationjobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Compilation Job Name: %q", id.CompilationJobName),
	}
	return fmt.Sprintf("Compilationjob (%s)", strings.Join(components, "\n"))
}
