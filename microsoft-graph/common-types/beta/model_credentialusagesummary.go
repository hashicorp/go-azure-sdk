package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummary struct {
	AuthMethod              *CredentialUsageSummaryAuthMethod `json:"authMethod,omitempty"`
	FailureActivityCount    *int64                            `json:"failureActivityCount,omitempty"`
	Feature                 *CredentialUsageSummaryFeature    `json:"feature,omitempty"`
	Id                      *string                           `json:"id,omitempty"`
	ODataType               *string                           `json:"@odata.type,omitempty"`
	SuccessfulActivityCount *int64                            `json:"successfulActivityCount,omitempty"`
}
