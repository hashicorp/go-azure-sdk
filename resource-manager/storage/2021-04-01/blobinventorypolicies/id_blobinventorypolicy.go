package blobinventorypolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BlobInventoryPolicyId{}

// BlobInventoryPolicyId is a struct representing the Resource ID for a Blob Inventory Policy
type BlobInventoryPolicyId struct {
	SubscriptionId          string
	ResourceGroupName       string
	AccountName             string
	BlobInventoryPolicyName BlobInventoryPolicyName
}

// NewBlobInventoryPolicyID returns a new BlobInventoryPolicyId struct
func NewBlobInventoryPolicyID(subscriptionId string, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName) BlobInventoryPolicyId {
	return BlobInventoryPolicyId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		AccountName:             accountName,
		BlobInventoryPolicyName: blobInventoryPolicyName,
	}
}

// ParseBlobInventoryPolicyID parses 'input' into a BlobInventoryPolicyId
func ParseBlobInventoryPolicyID(input string) (*BlobInventoryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(BlobInventoryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BlobInventoryPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["blobInventoryPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'blobInventoryPolicyName' was not found in the resource id %q", input)
		}

		blobInventoryPolicyName, err := parseBlobInventoryPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.BlobInventoryPolicyName = *blobInventoryPolicyName
	}

	return &id, nil
}

// ParseBlobInventoryPolicyIDInsensitively parses 'input' case-insensitively into a BlobInventoryPolicyId
// note: this method should only be used for API response data and not user input
func ParseBlobInventoryPolicyIDInsensitively(input string) (*BlobInventoryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(BlobInventoryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BlobInventoryPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["blobInventoryPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'blobInventoryPolicyName' was not found in the resource id %q", input)
		}

		blobInventoryPolicyName, err := parseBlobInventoryPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.BlobInventoryPolicyName = *blobInventoryPolicyName
	}

	return &id, nil
}

// ValidateBlobInventoryPolicyID checks that 'input' can be parsed as a Blob Inventory Policy ID
func ValidateBlobInventoryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBlobInventoryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Blob Inventory Policy ID
func (id BlobInventoryPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/inventoryPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.BlobInventoryPolicyName))
}

// Segments returns a slice of Resource ID Segments which comprise this Blob Inventory Policy ID
func (id BlobInventoryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticInventoryPolicies", "inventoryPolicies", "inventoryPolicies"),
		resourceids.ConstantSegment("blobInventoryPolicyName", PossibleValuesForBlobInventoryPolicyName(), "default"),
	}
}

// String returns a human-readable description of this Blob Inventory Policy ID
func (id BlobInventoryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Blob Inventory Policy Name: %q", string(id.BlobInventoryPolicyName)),
	}
	return fmt.Sprintf("Blob Inventory Policy (%s)", strings.Join(components, "\n"))
}
