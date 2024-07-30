package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAuthorizationSystemTypeAction struct {
	ActionType    *AzureAuthorizationSystemTypeActionActionType `json:"actionType,omitempty"`
	ExternalId    *string                                       `json:"externalId,omitempty"`
	Id            *string                                       `json:"id,omitempty"`
	ODataType     *string                                       `json:"@odata.type,omitempty"`
	ResourceTypes *[]string                                     `json:"resourceTypes,omitempty"`
	Service       *AuthorizationSystemTypeService               `json:"service,omitempty"`
	Severity      *AzureAuthorizationSystemTypeActionSeverity   `json:"severity,omitempty"`
}
