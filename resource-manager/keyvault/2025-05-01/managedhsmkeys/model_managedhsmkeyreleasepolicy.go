package managedhsmkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmKeyReleasePolicy struct {
	ContentType *string `json:"contentType,omitempty"`
	Data        *string `json:"data,omitempty"`
}
