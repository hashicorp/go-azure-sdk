package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DedicatedSQLminimalTlsSettingId{}

// DedicatedSQLminimalTlsSettingId is a struct representing the Resource ID for a Dedicated S Q Lminimal Tls Setting
type DedicatedSQLminimalTlsSettingId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	WorkspaceName                      string
	DedicatedSQLminimalTlsSettingsName string
}

// NewDedicatedSQLminimalTlsSettingID returns a new DedicatedSQLminimalTlsSettingId struct
func NewDedicatedSQLminimalTlsSettingID(subscriptionId string, resourceGroupName string, workspaceName string, dedicatedSQLminimalTlsSettingsName string) DedicatedSQLminimalTlsSettingId {
	return DedicatedSQLminimalTlsSettingId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		WorkspaceName:                      workspaceName,
		DedicatedSQLminimalTlsSettingsName: dedicatedSQLminimalTlsSettingsName,
	}
}

// ParseDedicatedSQLminimalTlsSettingID parses 'input' into a DedicatedSQLminimalTlsSettingId
func ParseDedicatedSQLminimalTlsSettingID(input string) (*DedicatedSQLminimalTlsSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DedicatedSQLminimalTlsSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DedicatedSQLminimalTlsSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.DedicatedSQLminimalTlsSettingsName, ok = parsed.Parsed["dedicatedSQLminimalTlsSettingsName"]; !ok {
		return nil, fmt.Errorf("the segment 'dedicatedSQLminimalTlsSettingsName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDedicatedSQLminimalTlsSettingIDInsensitively parses 'input' case-insensitively into a DedicatedSQLminimalTlsSettingId
// note: this method should only be used for API response data and not user input
func ParseDedicatedSQLminimalTlsSettingIDInsensitively(input string) (*DedicatedSQLminimalTlsSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DedicatedSQLminimalTlsSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DedicatedSQLminimalTlsSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.DedicatedSQLminimalTlsSettingsName, ok = parsed.Parsed["dedicatedSQLminimalTlsSettingsName"]; !ok {
		return nil, fmt.Errorf("the segment 'dedicatedSQLminimalTlsSettingsName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDedicatedSQLminimalTlsSettingID checks that 'input' can be parsed as a Dedicated S Q Lminimal Tls Setting ID
func ValidateDedicatedSQLminimalTlsSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDedicatedSQLminimalTlsSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dedicated S Q Lminimal Tls Setting ID
func (id DedicatedSQLminimalTlsSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/dedicatedSQLminimalTlsSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.DedicatedSQLminimalTlsSettingsName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dedicated S Q Lminimal Tls Setting ID
func (id DedicatedSQLminimalTlsSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticDedicatedSQLminimalTlsSettings", "dedicatedSQLminimalTlsSettings", "dedicatedSQLminimalTlsSettings"),
		resourceids.UserSpecifiedSegment("dedicatedSQLminimalTlsSettingsName", "dedicatedSQLminimalTlsSettingsValue"),
	}
}

// String returns a human-readable description of this Dedicated S Q Lminimal Tls Setting ID
func (id DedicatedSQLminimalTlsSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Dedicated S Q Lminimal Tls Settings Name: %q", id.DedicatedSQLminimalTlsSettingsName),
	}
	return fmt.Sprintf("Dedicated S Q Lminimal Tls Setting (%s)", strings.Join(components, "\n"))
}
