package botconnection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionSettingParameter struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}
