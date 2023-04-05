package dscconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent ContentSourceType = "embeddedContent"
	ContentSourceTypeUri             ContentSourceType = "uri"
)

func PossibleValuesForContentSourceType() []string {
	return []string{
		string(ContentSourceTypeEmbeddedContent),
		string(ContentSourceTypeUri),
	}
}

type DscConfigurationProvisioningState string

const (
	DscConfigurationProvisioningStateSucceeded DscConfigurationProvisioningState = "Succeeded"
)

func PossibleValuesForDscConfigurationProvisioningState() []string {
	return []string{
		string(DscConfigurationProvisioningStateSucceeded),
	}
}

type DscConfigurationState string

const (
	DscConfigurationStateEdit      DscConfigurationState = "Edit"
	DscConfigurationStateNew       DscConfigurationState = "New"
	DscConfigurationStatePublished DscConfigurationState = "Published"
)

func PossibleValuesForDscConfigurationState() []string {
	return []string{
		string(DscConfigurationStateEdit),
		string(DscConfigurationStateNew),
		string(DscConfigurationStatePublished),
	}
}
