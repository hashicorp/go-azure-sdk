package objectdatatypes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ObjectDataTypeId{}

// ObjectDataTypeId is a struct representing the Resource ID for a Object Data Type
type ObjectDataTypeId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	TypeName              string
}

// NewObjectDataTypeID returns a new ObjectDataTypeId struct
func NewObjectDataTypeID(subscriptionId string, resourceGroupName string, automationAccountName string, typeName string) ObjectDataTypeId {
	return ObjectDataTypeId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		TypeName:              typeName,
	}
}

// ParseObjectDataTypeID parses 'input' into a ObjectDataTypeId
func ParseObjectDataTypeID(input string) (*ObjectDataTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ObjectDataTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ObjectDataTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, fmt.Errorf("the segment 'typeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseObjectDataTypeIDInsensitively parses 'input' case-insensitively into a ObjectDataTypeId
// note: this method should only be used for API response data and not user input
func ParseObjectDataTypeIDInsensitively(input string) (*ObjectDataTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ObjectDataTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ObjectDataTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, fmt.Errorf("the segment 'typeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateObjectDataTypeID checks that 'input' can be parsed as a Object Data Type ID
func ValidateObjectDataTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseObjectDataTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Object Data Type ID
func (id ObjectDataTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/objectDataTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.TypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Object Data Type ID
func (id ObjectDataTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticObjectDataTypes", "objectDataTypes", "objectDataTypes"),
		resourceids.UserSpecifiedSegment("typeName", "typeValue"),
	}
}

// String returns a human-readable description of this Object Data Type ID
func (id ObjectDataTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Type Name: %q", id.TypeName),
	}
	return fmt.Sprintf("Object Data Type (%s)", strings.Join(components, "\n"))
}
