package blobservice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BlobServiceId{}

// BlobServiceId is a struct representing the Resource ID for a Blob Service
type BlobServiceId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	BlobServicesName  BlobServicesName
}

// NewBlobServiceID returns a new BlobServiceId struct
func NewBlobServiceID(subscriptionId string, resourceGroupName string, accountName string, blobServicesName BlobServicesName) BlobServiceId {
	return BlobServiceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		BlobServicesName:  blobServicesName,
	}
}

// ParseBlobServiceID parses 'input' into a BlobServiceId
func ParseBlobServiceID(input string) (*BlobServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(BlobServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BlobServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["blobServicesName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'blobServicesName' was not found in the resource id %q", input)
		}

		blobServicesName, err := parseBlobServicesName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.BlobServicesName = *blobServicesName
	}

	return &id, nil
}

// ParseBlobServiceIDInsensitively parses 'input' case-insensitively into a BlobServiceId
// note: this method should only be used for API response data and not user input
func ParseBlobServiceIDInsensitively(input string) (*BlobServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(BlobServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BlobServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["blobServicesName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'blobServicesName' was not found in the resource id %q", input)
		}

		blobServicesName, err := parseBlobServicesName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.BlobServicesName = *blobServicesName
	}

	return &id, nil
}

// ValidateBlobServiceID checks that 'input' can be parsed as a Blob Service ID
func ValidateBlobServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBlobServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Blob Service ID
func (id BlobServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.BlobServicesName))
}

// Segments returns a slice of Resource ID Segments which comprise this Blob Service ID
func (id BlobServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticBlobServices", "blobServices", "blobServices"),
		resourceids.ConstantSegment("blobServicesName", PossibleValuesForBlobServicesName(), "default"),
	}
}

// String returns a human-readable description of this Blob Service ID
func (id BlobServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Blob Services Name: %q", string(id.BlobServicesName)),
	}
	return fmt.Sprintf("Blob Service (%s)", strings.Join(components, "\n"))
}
