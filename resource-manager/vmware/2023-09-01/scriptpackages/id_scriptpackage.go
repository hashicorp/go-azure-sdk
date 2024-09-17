package scriptpackages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScriptPackageId{})
}

var _ resourceids.ResourceId = &ScriptPackageId{}

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
	parser := resourceids.NewParserFromResourceIdType(&ScriptPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScriptPackageIDInsensitively parses 'input' case-insensitively into a ScriptPackageId
// note: this method should only be used for API response data and not user input
func ParseScriptPackageIDInsensitively(input string) (*ScriptPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScriptPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScriptPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.ScriptPackageName, ok = input.Parsed["scriptPackageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scriptPackageName", input)
	}

	return nil
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
