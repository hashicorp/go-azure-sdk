package billingproperty

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionEnrollmentDetails struct {
	DepartmentDisplayName        *string `json:"departmentDisplayName,omitempty"`
	DepartmentId                 *string `json:"departmentId,omitempty"`
	EnrollmentAccountDisplayName *string `json:"enrollmentAccountDisplayName,omitempty"`
	EnrollmentAccountId          *string `json:"enrollmentAccountId,omitempty"`
	EnrollmentAccountStatus      *string `json:"enrollmentAccountStatus,omitempty"`
}
