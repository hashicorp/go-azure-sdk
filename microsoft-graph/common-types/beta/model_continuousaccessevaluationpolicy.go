package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousAccessEvaluationPolicy struct {
	Description *string   `json:"description,omitempty"`
	DisplayName *string   `json:"displayName,omitempty"`
	Groups      *[]string `json:"groups,omitempty"`
	Id          *string   `json:"id,omitempty"`
	IsEnabled   *bool     `json:"isEnabled,omitempty"`
	Migrate     *bool     `json:"migrate,omitempty"`
	ODataType   *string   `json:"@odata.type,omitempty"`
	Users       *[]string `json:"users,omitempty"`
}
