package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTrustedIdProviderWithAccountParameters struct {
	Name       string                                    `json:"name"`
	Properties CreateOrUpdateTrustedIdProviderProperties `json:"properties"`
}
