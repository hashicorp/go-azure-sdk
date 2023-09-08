package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetrieveRemoteHelpSessionResponse struct {
	AcsGroupId                *string `json:"acsGroupId,omitempty"`
	AcsHelperUserId           *string `json:"acsHelperUserId,omitempty"`
	AcsHelperUserToken        *string `json:"acsHelperUserToken,omitempty"`
	AcsSharerUserId           *string `json:"acsSharerUserId,omitempty"`
	DeviceName                *string `json:"deviceName,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	PubSubGroupId             *string `json:"pubSubGroupId,omitempty"`
	PubSubHelperAccessUri     *string `json:"pubSubHelperAccessUri,omitempty"`
	SessionExpirationDateTime *string `json:"sessionExpirationDateTime,omitempty"`
	SessionKey                *string `json:"sessionKey,omitempty"`
}
