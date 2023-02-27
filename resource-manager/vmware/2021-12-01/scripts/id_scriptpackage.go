// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scripts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScriptPackageId{}

// ScriptPackageId is a struct representing the Resource ID for a Script Package
type ScriptPackageId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	ScriptPackageName string
}

// NewScriptPackageID returns a new ScriptPackageId struct
func NewScriptPackageID(subscriptionId string, resourceGroupName string, privateCloudName string, scriptPackageName string) ScriptPackageId {
	return ScriptPackageId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		ScriptPackageName: scriptPackageName,
	}
}

// ParseScriptPackageID parses 'input' into a ScriptPackageId
func ParseScriptPackageID(input string) (*ScriptPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScriptPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScriptPackageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.ScriptPackageName, ok = parsed.Parsed["scriptPackageName"]; !ok {
		return nil, fmt.Errorf("the segment 'scriptPackageName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseScriptPackageIDInsensitively parses 'input' case-insensitively into a ScriptPackageId
// note: this method should only be used for API response data and not user input
func ParseScriptPackageIDInsensitively(input string) (*ScriptPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScriptPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScriptPackageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.ScriptPackageName, ok = parsed.Parsed["scriptPackageName"]; !ok {
		return nil, fmt.Errorf("the segment 'scriptPackageName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateScriptPackageID checks that 'input' can be parsed as a Script Package ID
func ValidateScriptPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScriptPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Script Package ID
func (id ScriptPackageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/scriptPackages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ScriptPackageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Script Package ID
func (id ScriptPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticScriptPackages", "scriptPackages", "scriptPackages"),
		resourceids.UserSpecifiedSegment("scriptPackageName", "scriptPackageValue"),
	}
}

// String returns a human-readable description of this Script Package ID
func (id ScriptPackageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Script Package Name: %q", id.ScriptPackageName),
	}
	return fmt.Sprintf("Script Package (%s)", strings.Join(components, "\n"))
}
