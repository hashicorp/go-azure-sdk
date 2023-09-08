package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportRoot struct {
	AuthenticationMethods      *AuthenticationMethodsRoot `json:"authenticationMethods,omitempty"`
	DailyPrintUsageByPrinter   *[]PrintUsageByPrinter     `json:"dailyPrintUsageByPrinter,omitempty"`
	DailyPrintUsageByUser      *[]PrintUsageByUser        `json:"dailyPrintUsageByUser,omitempty"`
	MonthlyPrintUsageByPrinter *[]PrintUsageByPrinter     `json:"monthlyPrintUsageByPrinter,omitempty"`
	MonthlyPrintUsageByUser    *[]PrintUsageByUser        `json:"monthlyPrintUsageByUser,omitempty"`
	ODataType                  *string                    `json:"@odata.type,omitempty"`
	Security                   *SecurityReportsRoot       `json:"security,omitempty"`
}
