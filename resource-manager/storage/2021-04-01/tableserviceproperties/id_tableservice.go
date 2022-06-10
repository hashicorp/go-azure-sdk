package tableserviceproperties

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = TableServiceId{}

// TableServiceId is a struct representing the Resource ID for a Table Service
type TableServiceId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	TableServiceName  TableServiceName
}

// NewTableServiceID returns a new TableServiceId struct
func NewTableServiceID(subscriptionId string, resourceGroupName string, accountName string, tableServiceName TableServiceName) TableServiceId {
	return TableServiceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		TableServiceName:  tableServiceName,
	}
}

// ParseTableServiceID parses 'input' into a TableServiceId
func ParseTableServiceID(input string) (*TableServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(TableServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TableServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["tableServiceName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'tableServiceName' was not found in the resource id %q", input)
		}

		tableServiceName, err := parseTableServiceName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.TableServiceName = *tableServiceName
	}

	return &id, nil
}

// ParseTableServiceIDInsensitively parses 'input' case-insensitively into a TableServiceId
// note: this method should only be used for API response data and not user input
func ParseTableServiceIDInsensitively(input string) (*TableServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(TableServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TableServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["tableServiceName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'tableServiceName' was not found in the resource id %q", input)
		}

		tableServiceName, err := parseTableServiceName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.TableServiceName = *tableServiceName
	}

	return &id, nil
}

// ValidateTableServiceID checks that 'input' can be parsed as a Table Service ID
func ValidateTableServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTableServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Table Service ID
func (id TableServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/tableServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.TableServiceName))
}

// Segments returns a slice of Resource ID Segments which comprise this Table Service ID
func (id TableServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticTableServices", "tableServices", "tableServices"),
		resourceids.ConstantSegment("tableServiceName", PossibleValuesForTableServiceName(), "default"),
	}
}

// String returns a human-readable description of this Table Service ID
func (id TableServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Table Service Name: %q", string(id.TableServiceName)),
	}
	return fmt.Sprintf("Table Service (%s)", strings.Join(components, "\n"))
}
