package akriconnectortemplate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsSecret struct {
	SecretAlias string `json:"secretAlias"`
	SecretKey   string `json:"secretKey"`
	SecretRef   string `json:"secretRef"`
}
