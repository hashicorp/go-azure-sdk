package staticsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteUserProvidedFunctionAppProperties struct {
	CreatedOn             *string `json:"createdOn,omitempty"`
	FunctionAppRegion     *string `json:"functionAppRegion,omitempty"`
	FunctionAppResourceId *string `json:"functionAppResourceId,omitempty"`
}
