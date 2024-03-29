package billingroledefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPermissionsProperties struct {
	Actions    *[]string `json:"actions,omitempty"`
	NotActions *[]string `json:"notActions,omitempty"`
}
