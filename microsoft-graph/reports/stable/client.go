package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/authenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/authenticationmethoduserregistrationdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/dailyprintusagebyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/dailyprintusagebyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/monthlyprintusagebyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/monthlyprintusagebyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/report"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/security"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AuthenticationMethod                       *authenticationmethod.AuthenticationMethodClient
	AuthenticationMethodUserRegistrationDetail *authenticationmethoduserregistrationdetail.AuthenticationMethodUserRegistrationDetailClient
	DailyPrintUsageByPrinter                   *dailyprintusagebyprinter.DailyPrintUsageByPrinterClient
	DailyPrintUsageByUser                      *dailyprintusagebyuser.DailyPrintUsageByUserClient
	MonthlyPrintUsageByPrinter                 *monthlyprintusagebyprinter.MonthlyPrintUsageByPrinterClient
	MonthlyPrintUsageByUser                    *monthlyprintusagebyuser.MonthlyPrintUsageByUserClient
	Report                                     *report.ReportClient
	Security                                   *security.SecurityClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	authenticationMethodClient, err := authenticationmethod.NewAuthenticationMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethod client: %+v", err)
	}
	configureFunc(authenticationMethodClient.Client)

	authenticationMethodUserRegistrationDetailClient, err := authenticationmethoduserregistrationdetail.NewAuthenticationMethodUserRegistrationDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodUserRegistrationDetail client: %+v", err)
	}
	configureFunc(authenticationMethodUserRegistrationDetailClient.Client)

	dailyPrintUsageByPrinterClient, err := dailyprintusagebyprinter.NewDailyPrintUsageByPrinterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DailyPrintUsageByPrinter client: %+v", err)
	}
	configureFunc(dailyPrintUsageByPrinterClient.Client)

	dailyPrintUsageByUserClient, err := dailyprintusagebyuser.NewDailyPrintUsageByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DailyPrintUsageByUser client: %+v", err)
	}
	configureFunc(dailyPrintUsageByUserClient.Client)

	monthlyPrintUsageByPrinterClient, err := monthlyprintusagebyprinter.NewMonthlyPrintUsageByPrinterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonthlyPrintUsageByPrinter client: %+v", err)
	}
	configureFunc(monthlyPrintUsageByPrinterClient.Client)

	monthlyPrintUsageByUserClient, err := monthlyprintusagebyuser.NewMonthlyPrintUsageByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonthlyPrintUsageByUser client: %+v", err)
	}
	configureFunc(monthlyPrintUsageByUserClient.Client)

	reportClient, err := report.NewReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Report client: %+v", err)
	}
	configureFunc(reportClient.Client)

	securityClient, err := security.NewSecurityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Security client: %+v", err)
	}
	configureFunc(securityClient.Client)

	return &Client{
		AuthenticationMethod:                       authenticationMethodClient,
		AuthenticationMethodUserRegistrationDetail: authenticationMethodUserRegistrationDetailClient,
		DailyPrintUsageByPrinter:                   dailyPrintUsageByPrinterClient,
		DailyPrintUsageByUser:                      dailyPrintUsageByUserClient,
		MonthlyPrintUsageByPrinter:                 monthlyPrintUsageByPrinterClient,
		MonthlyPrintUsageByUser:                    monthlyPrintUsageByUserClient,
		Report:                                     reportClient,
		Security:                                   securityClient,
	}, nil
}
