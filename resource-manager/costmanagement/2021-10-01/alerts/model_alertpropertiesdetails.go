package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertPropertiesDetails struct {
	Amount              *float64            `json:"amount,omitempty"`
	CompanyName         *string             `json:"companyName,omitempty"`
	ContactEmails       *[]string           `json:"contactEmails,omitempty"`
	ContactGroups       *[]string           `json:"contactGroups,omitempty"`
	ContactRoles        *[]string           `json:"contactRoles,omitempty"`
	CurrentSpend        *float64            `json:"currentSpend,omitempty"`
	DepartmentName      *string             `json:"departmentName,omitempty"`
	EnrollmentEndDate   *string             `json:"enrollmentEndDate,omitempty"`
	EnrollmentNumber    *string             `json:"enrollmentNumber,omitempty"`
	EnrollmentStartDate *string             `json:"enrollmentStartDate,omitempty"`
	InvoicingThreshold  *float64            `json:"invoicingThreshold,omitempty"`
	MeterFilter         *[]interface{}      `json:"meterFilter,omitempty"`
	Operator            *AlertOperator      `json:"operator,omitempty"`
	OverridingAlert     *string             `json:"overridingAlert,omitempty"`
	PeriodStartDate     *string             `json:"periodStartDate,omitempty"`
	ResourceFilter      *[]interface{}      `json:"resourceFilter,omitempty"`
	ResourceGroupFilter *[]interface{}      `json:"resourceGroupFilter,omitempty"`
	TagFilter           *interface{}        `json:"tagFilter,omitempty"`
	Threshold           *float64            `json:"threshold,omitempty"`
	TimeGrainType       *AlertTimeGrainType `json:"timeGrainType,omitempty"`
	TriggeredBy         *string             `json:"triggeredBy,omitempty"`
	Unit                *string             `json:"unit,omitempty"`
}
