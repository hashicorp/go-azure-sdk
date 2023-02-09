package datasetmapping

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DataSetMappingId{}

// DataSetMappingId is a struct representing the Resource ID for a Data Set Mapping
type DataSetMappingId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AccountName           string
	ShareSubscriptionName string
	DataSetMappingName    string
}

// NewDataSetMappingID returns a new DataSetMappingId struct
func NewDataSetMappingID(subscriptionId string, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) DataSetMappingId {
	return DataSetMappingId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AccountName:           accountName,
		ShareSubscriptionName: shareSubscriptionName,
		DataSetMappingName:    dataSetMappingName,
	}
}

// ParseDataSetMappingID parses 'input' into a DataSetMappingId
func ParseDataSetMappingID(input string) (*DataSetMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataSetMappingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataSetMappingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareSubscriptionName, ok = parsed.Parsed["shareSubscriptionName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareSubscriptionName' was not found in the resource id %q", input)
	}

	if id.DataSetMappingName, ok = parsed.Parsed["dataSetMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataSetMappingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDataSetMappingIDInsensitively parses 'input' case-insensitively into a DataSetMappingId
// note: this method should only be used for API response data and not user input
func ParseDataSetMappingIDInsensitively(input string) (*DataSetMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataSetMappingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataSetMappingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareSubscriptionName, ok = parsed.Parsed["shareSubscriptionName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareSubscriptionName' was not found in the resource id %q", input)
	}

	if id.DataSetMappingName, ok = parsed.Parsed["dataSetMappingName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataSetMappingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDataSetMappingID checks that 'input' can be parsed as a Data Set Mapping ID
func ValidateDataSetMappingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataSetMappingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Set Mapping ID
func (id DataSetMappingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataShare/accounts/%s/shareSubscriptions/%s/dataSetMappings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ShareSubscriptionName, id.DataSetMappingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Set Mapping ID
func (id DataSetMappingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticShareSubscriptions", "shareSubscriptions", "shareSubscriptions"),
		resourceids.UserSpecifiedSegment("shareSubscriptionName", "shareSubscriptionValue"),
		resourceids.StaticSegment("staticDataSetMappings", "dataSetMappings", "dataSetMappings"),
		resourceids.UserSpecifiedSegment("dataSetMappingName", "dataSetMappingValue"),
	}
}

// String returns a human-readable description of this Data Set Mapping ID
func (id DataSetMappingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Share Subscription Name: %q", id.ShareSubscriptionName),
		fmt.Sprintf("Data Set Mapping Name: %q", id.DataSetMappingName),
	}
	return fmt.Sprintf("Data Set Mapping (%s)", strings.Join(components, "\n"))
}
