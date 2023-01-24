package dsccompilationjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CompilationJobId{}

// CompilationJobId is a struct representing the Resource ID for a Compilation Job
type CompilationJobId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	CompilationJobName    string
}

// NewCompilationJobID returns a new CompilationJobId struct
func NewCompilationJobID(subscriptionId string, resourceGroupName string, automationAccountName string, compilationJobName string) CompilationJobId {
	return CompilationJobId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		CompilationJobName:    compilationJobName,
	}
}

// ParseCompilationJobID parses 'input' into a CompilationJobId
func ParseCompilationJobID(input string) (*CompilationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationJobId{}

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

// ParseCompilationJobIDInsensitively parses 'input' case-insensitively into a CompilationJobId
// note: this method should only be used for API response data and not user input
func ParseCompilationJobIDInsensitively(input string) (*CompilationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationJobId{}

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

// ValidateCompilationJobID checks that 'input' can be parsed as a Compilation Job ID
func ValidateCompilationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCompilationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Compilation Job ID
func (id CompilationJobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/compilationJobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.CompilationJobName)
}

// Segments returns a slice of Resource ID Segments which comprise this Compilation Job ID
func (id CompilationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticCompilationJobs", "compilationJobs", "compilationJobs"),
		resourceids.UserSpecifiedSegment("compilationJobName", "compilationJobValue"),
	}
}

// String returns a human-readable description of this Compilation Job ID
func (id CompilationJobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Compilation Job Name: %q", id.CompilationJobName),
	}
	return fmt.Sprintf("Compilation Job (%s)", strings.Join(components, "\n"))
}
