package registeredservers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServerUpdateProperties struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	Identity      *bool   `json:"identity,omitempty"`
}
