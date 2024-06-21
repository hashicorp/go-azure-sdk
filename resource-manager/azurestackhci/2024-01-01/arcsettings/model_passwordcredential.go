package arcsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredential struct {
	EndDateTime   *string `json:"endDateTime,omitempty"`
	KeyId         *string `json:"keyId,omitempty"`
	SecretText    *string `json:"secretText,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}
