package msiximage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpandMsixImageList struct {
	NextLink *string            `json:"nextLink,omitempty"`
	Value    *[]ExpandMsixImage `json:"value,omitempty"`
}
