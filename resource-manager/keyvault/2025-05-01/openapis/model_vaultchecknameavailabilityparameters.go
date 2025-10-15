package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultCheckNameAvailabilityParameters struct {
	Name string                                   `json:"name"`
	Type VaultCheckNameAvailabilityParametersType `json:"type"`
}
