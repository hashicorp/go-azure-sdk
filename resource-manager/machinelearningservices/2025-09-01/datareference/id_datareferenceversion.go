package datareference

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DataReferenceVersionId{})
}

var _ resourceids.ResourceId = &DataReferenceVersionId{}

// DataReferenceVersionId is a struct representing the Resource ID for a Data Reference Version
type DataReferenceVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	RegistryName      string
	DataReferenceName string
	VersionName       string
}

// NewDataReferenceVersionID returns a new DataReferenceVersionId struct
func NewDataReferenceVersionID(subscriptionId string, resourceGroupName string, registryName string, dataReferenceName string, versionName string) DataReferenceVersionId {
	return DataReferenceVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RegistryName:      registryName,
		DataReferenceName: dataReferenceName,
		VersionName:       versionName,
	}
}

// ParseDataReferenceVersionID parses 'input' into a DataReferenceVersionId
func ParseDataReferenceVersionID(input string) (*DataReferenceVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataReferenceVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataReferenceVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataReferenceVersionIDInsensitively parses 'input' case-insensitively into a DataReferenceVersionId
// note: this method should only be used for API response data and not user input
func ParseDataReferenceVersionIDInsensitively(input string) (*DataReferenceVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataReferenceVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataReferenceVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataReferenceVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RegistryName, ok = input.Parsed["registryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryName", input)
	}

	if id.DataReferenceName, ok = input.Parsed["dataReferenceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataReferenceName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateDataReferenceVersionID checks that 'input' can be parsed as a Data Reference Version ID
func ValidateDataReferenceVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataReferenceVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Reference Version ID
func (id DataReferenceVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/registries/%s/dataReferences/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RegistryName, id.DataReferenceName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Reference Version ID
func (id DataReferenceVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticRegistries", "registries", "registries"),
		resourceids.UserSpecifiedSegment("registryName", "registryName"),
		resourceids.StaticSegment("staticDataReferences", "dataReferences", "dataReferences"),
		resourceids.UserSpecifiedSegment("dataReferenceName", "dataReferenceName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Data Reference Version ID
func (id DataReferenceVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Registry Name: %q", id.RegistryName),
		fmt.Sprintf("Data Reference Name: %q", id.DataReferenceName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Data Reference Version (%s)", strings.Join(components, "\n"))
}
