package fileservice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = FileServiceId{}

// FileServiceId is a struct representing the Resource ID for a File Service
type FileServiceId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	FileServicesName  FileServicesName
}

// NewFileServiceID returns a new FileServiceId struct
func NewFileServiceID(subscriptionId string, resourceGroupName string, accountName string, fileServicesName FileServicesName) FileServiceId {
	return FileServiceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		FileServicesName:  fileServicesName,
	}
}

// ParseFileServiceID parses 'input' into a FileServiceId
func ParseFileServiceID(input string) (*FileServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(FileServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FileServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["fileServicesName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'fileServicesName' was not found in the resource id %q", input)
		}

		fileServicesName, err := parseFileServicesName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.FileServicesName = *fileServicesName
	}

	return &id, nil
}

// ParseFileServiceIDInsensitively parses 'input' case-insensitively into a FileServiceId
// note: this method should only be used for API response data and not user input
func ParseFileServiceIDInsensitively(input string) (*FileServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(FileServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FileServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["fileServicesName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'fileServicesName' was not found in the resource id %q", input)
		}

		fileServicesName, err := parseFileServicesName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.FileServicesName = *fileServicesName
	}

	return &id, nil
}

// ValidateFileServiceID checks that 'input' can be parsed as a File Service ID
func ValidateFileServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFileServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted File Service ID
func (id FileServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/fileServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.FileServicesName))
}

// Segments returns a slice of Resource ID Segments which comprise this File Service ID
func (id FileServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticFileServices", "fileServices", "fileServices"),
		resourceids.ConstantSegment("fileServicesName", PossibleValuesForFileServicesName(), "default"),
	}
}

// String returns a human-readable description of this File Service ID
func (id FileServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("File Services Name: %q", string(id.FileServicesName)),
	}
	return fmt.Sprintf("File Service (%s)", strings.Join(components, "\n"))
}
