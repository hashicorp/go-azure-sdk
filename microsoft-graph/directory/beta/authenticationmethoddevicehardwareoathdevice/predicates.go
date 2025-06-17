package authenticationmethoddevicehardwareoathdevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type HardwareOathTokenAuthenticationMethodDeviceOperationPredicate struct {
}

func (p HardwareOathTokenAuthenticationMethodDeviceOperationPredicate) Matches(input beta.HardwareOathTokenAuthenticationMethodDevice) bool {

	return true
}
