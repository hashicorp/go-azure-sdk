package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevation struct {
	CertificatePayload *string                                    `json:"certificatePayload,omitempty"`
	CompanyName        *string                                    `json:"companyName,omitempty"`
	DeviceId           *string                                    `json:"deviceId,omitempty"`
	DeviceName         *string                                    `json:"deviceName,omitempty"`
	ElevationType      *PrivilegeManagementElevationElevationType `json:"elevationType,omitempty"`
	EventDateTime      *string                                    `json:"eventDateTime,omitempty"`
	FileDescription    *string                                    `json:"fileDescription,omitempty"`
	FilePath           *string                                    `json:"filePath,omitempty"`
	FileVersion        *string                                    `json:"fileVersion,omitempty"`
	Hash               *string                                    `json:"hash,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	InternalName       *string                                    `json:"internalName,omitempty"`
	Justification      *string                                    `json:"justification,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	ParentProcessName  *string                                    `json:"parentProcessName,omitempty"`
	PolicyId           *string                                    `json:"policyId,omitempty"`
	PolicyName         *string                                    `json:"policyName,omitempty"`
	ProcessType        *PrivilegeManagementElevationProcessType   `json:"processType,omitempty"`
	ProductName        *string                                    `json:"productName,omitempty"`
	Result             *int64                                     `json:"result,omitempty"`
	RuleId             *string                                    `json:"ruleId,omitempty"`
	Upn                *string                                    `json:"upn,omitempty"`
	UserType           *PrivilegeManagementElevationUserType      `json:"userType,omitempty"`
}
