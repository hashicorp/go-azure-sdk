package endpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type URLSigningParamIdentifier struct {
	ParamIndicator ParamIndicator `json:"paramIndicator"`
	ParamName      string         `json:"paramName"`
}
