package administrativeunit

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAdministrativeUnitMemberObjectRequest struct {
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}
