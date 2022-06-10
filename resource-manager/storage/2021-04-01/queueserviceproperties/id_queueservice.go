package queueserviceproperties

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QueueServiceId{}

// QueueServiceId is a struct representing the Resource ID for a Queue Service
type QueueServiceId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	QueueServiceName  QueueServiceName
}

// NewQueueServiceID returns a new QueueServiceId struct
func NewQueueServiceID(subscriptionId string, resourceGroupName string, accountName string, queueServiceName QueueServiceName) QueueServiceId {
	return QueueServiceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		QueueServiceName:  queueServiceName,
	}
}

// ParseQueueServiceID parses 'input' into a QueueServiceId
func ParseQueueServiceID(input string) (*QueueServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueueServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QueueServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["queueServiceName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'queueServiceName' was not found in the resource id %q", input)
		}

		queueServiceName, err := parseQueueServiceName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.QueueServiceName = *queueServiceName
	}

	return &id, nil
}

// ParseQueueServiceIDInsensitively parses 'input' case-insensitively into a QueueServiceId
// note: this method should only be used for API response data and not user input
func ParseQueueServiceIDInsensitively(input string) (*QueueServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueueServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QueueServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["queueServiceName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'queueServiceName' was not found in the resource id %q", input)
		}

		queueServiceName, err := parseQueueServiceName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.QueueServiceName = *queueServiceName
	}

	return &id, nil
}

// ValidateQueueServiceID checks that 'input' can be parsed as a Queue Service ID
func ValidateQueueServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQueueServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Queue Service ID
func (id QueueServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/queueServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.QueueServiceName))
}

// Segments returns a slice of Resource ID Segments which comprise this Queue Service ID
func (id QueueServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticQueueServices", "queueServices", "queueServices"),
		resourceids.ConstantSegment("queueServiceName", PossibleValuesForQueueServiceName(), "default"),
	}
}

// String returns a human-readable description of this Queue Service ID
func (id QueueServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Queue Service Name: %q", string(id.QueueServiceName)),
	}
	return fmt.Sprintf("Queue Service (%s)", strings.Join(components, "\n"))
}
