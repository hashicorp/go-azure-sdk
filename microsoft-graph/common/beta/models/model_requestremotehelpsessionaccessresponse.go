package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestRemoteHelpSessionAccessResponse struct {
	ODataType           *string `json:"@odata.type,omitempty"`
	PubSubEncryption    *string `json:"pubSubEncryption,omitempty"`
	PubSubEncryptionKey *string `json:"pubSubEncryptionKey,omitempty"`
	SessionKey          *string `json:"sessionKey,omitempty"`
}
