package connectordefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DataConnectorDefinitionId{})
}

var _ resourceids.ResourceId = &DataConnectorDefinitionId{}

// DataConnectorDefinitionId is a struct representing the Resource ID for a Data Connector Definition
type DataConnectorDefinitionId struct {
	SubscriptionId              string
	ResourceGroupName           string
	WorkspaceName               string
	DataConnectorDefinitionName string
}

// NewDataConnectorDefinitionID returns a new DataConnectorDefinitionId struct
func NewDataConnectorDefinitionID(subscriptionId string, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string) DataConnectorDefinitionId {
	return DataConnectorDefinitionId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		WorkspaceName:               workspaceName,
		DataConnectorDefinitionName: dataConnectorDefinitionName,
	}
}

// ParseDataConnectorDefinitionID parses 'input' into a DataConnectorDefinitionId
func ParseDataConnectorDefinitionID(input string) (*DataConnectorDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataConnectorDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataConnectorDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataConnectorDefinitionIDInsensitively parses 'input' case-insensitively into a DataConnectorDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDataConnectorDefinitionIDInsensitively(input string) (*DataConnectorDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataConnectorDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataConnectorDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataConnectorDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.DataConnectorDefinitionName, ok = input.Parsed["dataConnectorDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataConnectorDefinitionName", input)
	}

	return nil
}

// ValidateDataConnectorDefinitionID checks that 'input' can be parsed as a Data Connector Definition ID
func ValidateDataConnectorDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataConnectorDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Connector Definition ID
func (id DataConnectorDefinitionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/dataConnectorDefinitions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.DataConnectorDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Connector Definition ID
func (id DataConnectorDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticDataConnectorDefinitions", "dataConnectorDefinitions", "dataConnectorDefinitions"),
		resourceids.UserSpecifiedSegment("dataConnectorDefinitionName", "dataConnectorDefinitionName"),
	}
}

// String returns a human-readable description of this Data Connector Definition ID
func (id DataConnectorDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Data Connector Definition Name: %q", id.DataConnectorDefinitionName),
	}
	return fmt.Sprintf("Data Connector Definition (%s)", strings.Join(components, "\n"))
}
