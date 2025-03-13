package inferencepool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScaleUnitConfiguration struct {
	DisablePublicEgress *bool     `json:"disablePublicEgress,omitempty"`
	Registries          *[]string `json:"registries,omitempty"`
}
