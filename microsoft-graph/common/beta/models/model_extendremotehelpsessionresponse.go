package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendRemoteHelpSessionResponse struct {
	AcsHelperUserToken        *string `json:"acsHelperUserToken,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	PubSubHelperAccessUri     *string `json:"pubSubHelperAccessUri,omitempty"`
	SessionExpirationDateTime *string `json:"sessionExpirationDateTime,omitempty"`
	SessionKey                *string `json:"sessionKey,omitempty"`
}
