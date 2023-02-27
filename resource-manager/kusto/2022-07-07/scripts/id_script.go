// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scripts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScriptId{}

// ScriptId is a struct representing the Resource ID for a Script
type ScriptId struct {
	SubscriptionId    string
	ResourceGroupName string
	ClusterName       string
	DatabaseName      string
	ScriptName        string
}

// NewScriptID returns a new ScriptId struct
func NewScriptID(subscriptionId string, resourceGroupName string, clusterName string, databaseName string, scriptName string) ScriptId {
	return ScriptId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ClusterName:       clusterName,
		DatabaseName:      databaseName,
		ScriptName:        scriptName,
	}
}

// ParseScriptID parses 'input' into a ScriptId
func ParseScriptID(input string) (*ScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScriptId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, fmt.Errorf("the segment 'databaseName' was not found in the resource id %q", input)
	}

	if id.ScriptName, ok = parsed.Parsed["scriptName"]; !ok {
		return nil, fmt.Errorf("the segment 'scriptName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseScriptIDInsensitively parses 'input' case-insensitively into a ScriptId
// note: this method should only be used for API response data and not user input
func ParseScriptIDInsensitively(input string) (*ScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScriptId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, fmt.Errorf("the segment 'databaseName' was not found in the resource id %q", input)
	}

	if id.ScriptName, ok = parsed.Parsed["scriptName"]; !ok {
		return nil, fmt.Errorf("the segment 'scriptName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateScriptID checks that 'input' can be parsed as a Script ID
func ValidateScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Script ID
func (id ScriptId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Kusto/clusters/%s/databases/%s/scripts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.DatabaseName, id.ScriptName)
}

// Segments returns a slice of Resource ID Segments which comprise this Script ID
func (id ScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKusto", "Microsoft.Kusto", "Microsoft.Kusto"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticScripts", "scripts", "scripts"),
		resourceids.UserSpecifiedSegment("scriptName", "scriptValue"),
	}
}

// String returns a human-readable description of this Script ID
func (id ScriptId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Script Name: %q", id.ScriptName),
	}
	return fmt.Sprintf("Script (%s)", strings.Join(components, "\n"))
}
