package namedvalue

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultLastAccessStatusContractProperties struct {
	Code         *string `json:"code,omitempty"`
	Message      *string `json:"message,omitempty"`
	TimeStampUtc *string `json:"timeStampUtc,omitempty"`
}
