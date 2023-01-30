package storageaccountcredentials

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = StorageAccountCredentialId{}

// StorageAccountCredentialId is a struct representing the Resource ID for a Storage Account Credential
type StorageAccountCredentialId struct {
	SubscriptionId               string
	ResourceGroupName            string
	DataBoxEdgeDeviceName        string
	StorageAccountCredentialName string
}

// NewStorageAccountCredentialID returns a new StorageAccountCredentialId struct
func NewStorageAccountCredentialID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, storageAccountCredentialName string) StorageAccountCredentialId {
	return StorageAccountCredentialId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		DataBoxEdgeDeviceName:        dataBoxEdgeDeviceName,
		StorageAccountCredentialName: storageAccountCredentialName,
	}
}

// ParseStorageAccountCredentialID parses 'input' into a StorageAccountCredentialId
func ParseStorageAccountCredentialID(input string) (*StorageAccountCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StorageAccountCredentialId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.StorageAccountCredentialName, ok = parsed.Parsed["storageAccountCredentialName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageAccountCredentialName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseStorageAccountCredentialIDInsensitively parses 'input' case-insensitively into a StorageAccountCredentialId
// note: this method should only be used for API response data and not user input
func ParseStorageAccountCredentialIDInsensitively(input string) (*StorageAccountCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StorageAccountCredentialId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.StorageAccountCredentialName, ok = parsed.Parsed["storageAccountCredentialName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageAccountCredentialName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateStorageAccountCredentialID checks that 'input' can be parsed as a Storage Account Credential ID
func ValidateStorageAccountCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStorageAccountCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Storage Account Credential ID
func (id StorageAccountCredentialId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/storageAccountCredentials/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.StorageAccountCredentialName)
}

// Segments returns a slice of Resource ID Segments which comprise this Storage Account Credential ID
func (id StorageAccountCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticStorageAccountCredentials", "storageAccountCredentials", "storageAccountCredentials"),
		resourceids.UserSpecifiedSegment("storageAccountCredentialName", "storageAccountCredentialValue"),
	}
}

// String returns a human-readable description of this Storage Account Credential ID
func (id StorageAccountCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Storage Account Credential Name: %q", id.StorageAccountCredentialName),
	}
	return fmt.Sprintf("Storage Account Credential (%s)", strings.Join(components, "\n"))
}
