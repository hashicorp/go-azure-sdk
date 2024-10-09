package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGroupEntityProperties struct {
	AdditionalData    *map[string]interface{} `json:"additionalData,omitempty"`
	DistinguishedName *string                 `json:"distinguishedName,omitempty"`
	FriendlyName      *string                 `json:"friendlyName,omitempty"`
	ObjectGuid        *string                 `json:"objectGuid,omitempty"`
	Sid               *string                 `json:"sid,omitempty"`
}
