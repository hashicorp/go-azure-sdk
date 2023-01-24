package sourcecontrolsyncjobstreams

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = StreamId{}

// StreamId is a struct representing the Resource ID for a Stream
type StreamId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	JobId                 string
	JobStreamId           string
}

// NewStreamID returns a new StreamId struct
func NewStreamID(subscriptionId string, resourceGroupName string, automationAccountName string, jobId string, jobStreamId string) StreamId {
	return StreamId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		JobId:                 jobId,
		JobStreamId:           jobStreamId,
	}
}

// ParseStreamID parses 'input' into a StreamId
func ParseStreamID(input string) (*StreamId, error) {
	parser := resourceids.NewParserFromResourceIdType(StreamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StreamId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobId' was not found in the resource id %q", input)
	}

	if id.JobStreamId, ok = parsed.Parsed["jobStreamId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobStreamId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseStreamIDInsensitively parses 'input' case-insensitively into a StreamId
// note: this method should only be used for API response data and not user input
func ParseStreamIDInsensitively(input string) (*StreamId, error) {
	parser := resourceids.NewParserFromResourceIdType(StreamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StreamId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobId' was not found in the resource id %q", input)
	}

	if id.JobStreamId, ok = parsed.Parsed["jobStreamId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobStreamId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateStreamID checks that 'input' can be parsed as a Stream ID
func ValidateStreamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStreamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Stream ID
func (id StreamId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/compilationJobs/%s/streams/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.JobId, id.JobStreamId)
}

// Segments returns a slice of Resource ID Segments which comprise this Stream ID
func (id StreamId) Segments() []resourceids.Segment {
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
		resourceids.UserSpecifiedSegment("jobId", "jobIdValue"),
		resourceids.StaticSegment("staticStreams", "streams", "streams"),
		resourceids.UserSpecifiedSegment("jobStreamId", "jobStreamIdValue"),
	}
}

// String returns a human-readable description of this Stream ID
func (id StreamId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Job: %q", id.JobId),
		fmt.Sprintf("Job Stream: %q", id.JobStreamId),
	}
	return fmt.Sprintf("Stream (%s)", strings.Join(components, "\n"))
}
