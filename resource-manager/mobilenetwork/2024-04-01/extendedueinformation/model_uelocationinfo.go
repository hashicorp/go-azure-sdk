package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeLocationInfo struct {
	LocationType string `json:"locationType"`
	Plmn         PlmnId `json:"plmn"`
	Tac          string `json:"tac"`
}
