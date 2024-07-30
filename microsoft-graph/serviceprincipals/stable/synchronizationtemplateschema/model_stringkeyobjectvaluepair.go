package synchronizationtemplateschema

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StringKeyObjectValuePair struct {
	Key       *string `json:"key,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
