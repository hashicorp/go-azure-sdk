package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAuthorizationSystemTypeAction struct {
	ActionType    *AwsAuthorizationSystemTypeActionActionType `json:"actionType,omitempty"`
	ExternalId    *string                                     `json:"externalId,omitempty"`
	Id            *string                                     `json:"id,omitempty"`
	ODataType     *string                                     `json:"@odata.type,omitempty"`
	ResourceTypes *[]string                                   `json:"resourceTypes,omitempty"`
	Service       *AuthorizationSystemTypeService             `json:"service,omitempty"`
	Severity      *AwsAuthorizationSystemTypeActionSeverity   `json:"severity,omitempty"`
}
