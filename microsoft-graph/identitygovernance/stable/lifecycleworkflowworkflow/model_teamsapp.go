package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsApp struct {
	AppDefinitions     *[]TeamsAppDefinition       `json:"appDefinitions,omitempty"`
	DisplayName        *string                     `json:"displayName,omitempty"`
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	ExternalId         *string                     `json:"externalId,omitempty"`
	Id                 *string                     `json:"id,omitempty"`
	ODataType          *string                     `json:"@odata.type,omitempty"`
}
