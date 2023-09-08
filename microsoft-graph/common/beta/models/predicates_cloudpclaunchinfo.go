package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCLaunchInfoOperationPredicate struct {
	CloudPCId                           *string
	CloudPCLaunchUrl                    *string
	ODataType                           *string
	Windows365SwitchCompatible          *bool
	Windows365SwitchNotCompatibleReason *string
}

func (p CloudPCLaunchInfoOperationPredicate) Matches(input CloudPCLaunchInfo) bool {

	if p.CloudPCId != nil && (input.CloudPCId == nil || *p.CloudPCId != *input.CloudPCId) {
		return false
	}

	if p.CloudPCLaunchUrl != nil && (input.CloudPCLaunchUrl == nil || *p.CloudPCLaunchUrl != *input.CloudPCLaunchUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Windows365SwitchCompatible != nil && (input.Windows365SwitchCompatible == nil || *p.Windows365SwitchCompatible != *input.Windows365SwitchCompatible) {
		return false
	}

	if p.Windows365SwitchNotCompatibleReason != nil && (input.Windows365SwitchNotCompatibleReason == nil || *p.Windows365SwitchNotCompatibleReason != *input.Windows365SwitchNotCompatibleReason) {
		return false
	}

	return true
}
