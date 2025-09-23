package currentusagesbases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesObject struct {
	UsagesType *UsagesTypes `json:"usagesType,omitempty"`
	Value      int64        `json:"value"`
}
