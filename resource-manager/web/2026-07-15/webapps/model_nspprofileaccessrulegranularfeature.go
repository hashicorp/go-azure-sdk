package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NspProfileAccessRuleGranularFeature struct {
	Features     *[]string `json:"features,omitempty"`
	ResourceType *string   `json:"resourceType,omitempty"`
}
