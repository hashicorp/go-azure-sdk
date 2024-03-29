package fluxconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostBuildDefinition struct {
	Substitute     *map[string]string          `json:"substitute,omitempty"`
	SubstituteFrom *[]SubstituteFromDefinition `json:"substituteFrom,omitempty"`
}
