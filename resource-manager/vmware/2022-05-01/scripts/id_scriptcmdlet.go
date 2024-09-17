package scripts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScriptCmdletId{})
}

var _ resourceids.ResourceId = &ScriptCmdletId{}

// ScriptCmdletId is a struct representing the Resource ID for a Script Cmdlet
type ScriptCmdletId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	ScriptPackageName string
	ScriptCmdletName  string
}

// NewScriptCmdletID returns a new ScriptCmdletId struct
func NewScriptCmdletID(subscriptionId string, resourceGroupName string, privateCloudName string, scriptPackageName string, scriptCmdletName string) ScriptCmdletId {
	return ScriptCmdletId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		ScriptPackageName: scriptPackageName,
		ScriptCmdletName:  scriptCmdletName,
	}
}

// ParseScriptCmdletID parses 'input' into a ScriptCmdletId
func ParseScriptCmdletID(input string) (*ScriptCmdletId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScriptCmdletId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptCmdletId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScriptCmdletIDInsensitively parses 'input' case-insensitively into a ScriptCmdletId
// note: this method should only be used for API response data and not user input
func ParseScriptCmdletIDInsensitively(input string) (*ScriptCmdletId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScriptCmdletId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptCmdletId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScriptCmdletId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ScriptCmdletName, ok = input.Parsed["scriptCmdletName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scriptCmdletName", input)
	}

	return nil
}

// ValidateScriptCmdletID checks that 'input' can be parsed as a Script Cmdlet ID
func ValidateScriptCmdletID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScriptCmdletID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Script Cmdlet ID
func (id ScriptCmdletId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/scriptPackages/%s/scriptCmdlets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ScriptPackageName, id.ScriptCmdletName)
}

// Segments returns a slice of Resource ID Segments which comprise this Script Cmdlet ID
func (id ScriptCmdletId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticScriptCmdlets", "scriptCmdlets", "scriptCmdlets"),
		resourceids.UserSpecifiedSegment("scriptCmdletName", "scriptCmdletValue"),
	}
}

// String returns a human-readable description of this Script Cmdlet ID
func (id ScriptCmdletId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Script Package Name: %q", id.ScriptPackageName),
		fmt.Sprintf("Script Cmdlet Name: %q", id.ScriptCmdletName),
	}
	return fmt.Sprintf("Script Cmdlet (%s)", strings.Join(components, "\n"))
}
