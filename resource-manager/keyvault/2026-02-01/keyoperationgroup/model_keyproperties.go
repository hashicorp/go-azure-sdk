package keyoperationgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyProperties struct {
	Attributes        *KeyAttributes         `json:"attributes,omitempty"`
	CurveName         *JsonWebKeyCurveName   `json:"curveName,omitempty"`
	KeyOps            *[]JsonWebKeyOperation `json:"keyOps,omitempty"`
	KeySize           *int64                 `json:"keySize,omitempty"`
	KeyUri            *string                `json:"keyUri,omitempty"`
	KeyUriWithVersion *string                `json:"keyUriWithVersion,omitempty"`
	Kty               *JsonWebKeyType        `json:"kty,omitempty"`
	ReleasePolicy     *KeyReleasePolicy      `json:"release_policy,omitempty"`
	RotationPolicy    *RotationPolicy        `json:"rotationPolicy,omitempty"`
}
