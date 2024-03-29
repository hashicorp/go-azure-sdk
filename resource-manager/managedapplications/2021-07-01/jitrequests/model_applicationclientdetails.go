package jitrequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationClientDetails struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	Oid           *string `json:"oid,omitempty"`
	Puid          *string `json:"puid,omitempty"`
}
