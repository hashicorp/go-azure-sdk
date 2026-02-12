package openidmetadatadiscovery

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIDConfigurationResponse struct {
	ClaimsSupported                  *[]string `json:"claims_supported,omitempty"`
	IdTokenSigningAlgValuesSupported *[]string `json:"id_token_signing_alg_values_supported,omitempty"`
	Issuer                           *string   `json:"issuer,omitempty"`
	JwksUri                          *string   `json:"jwks_uri,omitempty"`
	ResponseTypesSupported           *[]string `json:"response_types_supported,omitempty"`
	RevocationEndpoint               *string   `json:"revocation_endpoint,omitempty"`
}
