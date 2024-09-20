package devops

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AzureDevOpsOrgId{})
}

var _ resourceids.ResourceId = &AzureDevOpsOrgId{}

// AzureDevOpsOrgId is a struct representing the Resource ID for a Azure Dev Ops Org
type AzureDevOpsOrgId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
	AzureDevOpsOrgName    string
}

// NewAzureDevOpsOrgID returns a new AzureDevOpsOrgId struct
func NewAzureDevOpsOrgID(subscriptionId string, resourceGroupName string, securityConnectorName string, azureDevOpsOrgName string) AzureDevOpsOrgId {
	return AzureDevOpsOrgId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
		AzureDevOpsOrgName:    azureDevOpsOrgName,
	}
}

// ParseAzureDevOpsOrgID parses 'input' into a AzureDevOpsOrgId
func ParseAzureDevOpsOrgID(input string) (*AzureDevOpsOrgId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AzureDevOpsOrgId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AzureDevOpsOrgId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAzureDevOpsOrgIDInsensitively parses 'input' case-insensitively into a AzureDevOpsOrgId
// note: this method should only be used for API response data and not user input
func ParseAzureDevOpsOrgIDInsensitively(input string) (*AzureDevOpsOrgId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AzureDevOpsOrgId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AzureDevOpsOrgId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AzureDevOpsOrgId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SecurityConnectorName, ok = input.Parsed["securityConnectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityConnectorName", input)
	}

	if id.AzureDevOpsOrgName, ok = input.Parsed["azureDevOpsOrgName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "azureDevOpsOrgName", input)
	}

	return nil
}

// ValidateAzureDevOpsOrgID checks that 'input' can be parsed as a Azure Dev Ops Org ID
func ValidateAzureDevOpsOrgID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAzureDevOpsOrgID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Azure Dev Ops Org ID
func (id AzureDevOpsOrgId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s/devops/default/azureDevOpsOrgs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName, id.AzureDevOpsOrgName)
}

// Segments returns a slice of Resource ID Segments which comprise this Azure Dev Ops Org ID
func (id AzureDevOpsOrgId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityConnectors", "securityConnectors", "securityConnectors"),
		resourceids.UserSpecifiedSegment("securityConnectorName", "securityConnectorName"),
		resourceids.StaticSegment("staticDevops", "devops", "devops"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticAzureDevOpsOrgs", "azureDevOpsOrgs", "azureDevOpsOrgs"),
		resourceids.UserSpecifiedSegment("azureDevOpsOrgName", "orgName"),
	}
}

// String returns a human-readable description of this Azure Dev Ops Org ID
func (id AzureDevOpsOrgId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
		fmt.Sprintf("Azure Dev Ops Org Name: %q", id.AzureDevOpsOrgName),
	}
	return fmt.Sprintf("Azure Dev Ops Org (%s)", strings.Join(components, "\n"))
}
