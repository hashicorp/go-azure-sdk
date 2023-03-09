package jobdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpmErrorInfo struct {
	ErrorString     *string   `json:"errorString,omitempty"`
	Recommendations *[]string `json:"recommendations,omitempty"`
}
