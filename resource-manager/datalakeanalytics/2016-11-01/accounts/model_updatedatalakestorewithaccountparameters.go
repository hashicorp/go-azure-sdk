package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDataLakeStoreWithAccountParameters struct {
	Name       string                         `json:"name"`
	Properties *UpdateDataLakeStoreProperties `json:"properties,omitempty"`
}
