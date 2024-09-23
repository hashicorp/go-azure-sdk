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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbilling"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingmanifest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingreconciliation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingreconciliationbilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingusage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingusagebilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingusageunbilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/report"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/security"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AuthenticationMethod                       *authenticationmethod.AuthenticationMethodClient
	AuthenticationMethodUserRegistrationDetail *authenticationmethoduserregistrationdetail.AuthenticationMethodUserRegistrationDetailClient
	DailyPrintUsageByPrinter                   *dailyprintusagebyprinter.DailyPrintUsageByPrinterClient
	DailyPrintUsageByUser                      *dailyprintusagebyuser.DailyPrintUsageByUserClient
	MonthlyPrintUsageByPrinter                 *monthlyprintusagebyprinter.MonthlyPrintUsageByPrinterClient
	MonthlyPrintUsageByUser                    *monthlyprintusagebyuser.MonthlyPrintUsageByUserClient
	Partner                                    *partner.PartnerClient
	PartnerBilling                             *partnerbilling.PartnerBillingClient
	PartnerBillingManifest                     *partnerbillingmanifest.PartnerBillingManifestClient
	PartnerBillingOperation                    *partnerbillingoperation.PartnerBillingOperationClient
	PartnerBillingReconciliation               *partnerbillingreconciliation.PartnerBillingReconciliationClient
	PartnerBillingReconciliationBilled         *partnerbillingreconciliationbilled.PartnerBillingReconciliationBilledClient
	PartnerBillingUsage                        *partnerbillingusage.PartnerBillingUsageClient
	PartnerBillingUsageBilled                  *partnerbillingusagebilled.PartnerBillingUsageBilledClient
	PartnerBillingUsageUnbilled                *partnerbillingusageunbilled.PartnerBillingUsageUnbilledClient
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

	partnerBillingClient, err := partnerbilling.NewPartnerBillingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBilling client: %+v", err)
	}
	configureFunc(partnerBillingClient.Client)

	partnerBillingManifestClient, err := partnerbillingmanifest.NewPartnerBillingManifestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingManifest client: %+v", err)
	}
	configureFunc(partnerBillingManifestClient.Client)

	partnerBillingOperationClient, err := partnerbillingoperation.NewPartnerBillingOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingOperation client: %+v", err)
	}
	configureFunc(partnerBillingOperationClient.Client)

	partnerBillingReconciliationBilledClient, err := partnerbillingreconciliationbilled.NewPartnerBillingReconciliationBilledClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingReconciliationBilled client: %+v", err)
	}
	configureFunc(partnerBillingReconciliationBilledClient.Client)

	partnerBillingReconciliationClient, err := partnerbillingreconciliation.NewPartnerBillingReconciliationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingReconciliation client: %+v", err)
	}
	configureFunc(partnerBillingReconciliationClient.Client)

	partnerBillingUsageBilledClient, err := partnerbillingusagebilled.NewPartnerBillingUsageBilledClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingUsageBilled client: %+v", err)
	}
	configureFunc(partnerBillingUsageBilledClient.Client)

	partnerBillingUsageClient, err := partnerbillingusage.NewPartnerBillingUsageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingUsage client: %+v", err)
	}
	configureFunc(partnerBillingUsageClient.Client)

	partnerBillingUsageUnbilledClient, err := partnerbillingusageunbilled.NewPartnerBillingUsageUnbilledClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingUsageUnbilled client: %+v", err)
	}
	configureFunc(partnerBillingUsageUnbilledClient.Client)

	partnerClient, err := partner.NewPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Partner client: %+v", err)
	}
	configureFunc(partnerClient.Client)

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
		Partner:                                    partnerClient,
		PartnerBilling:                             partnerBillingClient,
		PartnerBillingManifest:                     partnerBillingManifestClient,
		PartnerBillingOperation:                    partnerBillingOperationClient,
		PartnerBillingReconciliation:               partnerBillingReconciliationClient,
		PartnerBillingReconciliationBilled:         partnerBillingReconciliationBilledClient,
		PartnerBillingUsage:                        partnerBillingUsageClient,
		PartnerBillingUsageBilled:                  partnerBillingUsageBilledClient,
		PartnerBillingUsageUnbilled:                partnerBillingUsageUnbilledClient,
		Report:                                     reportClient,
		Security:                                   securityClient,
	}, nil
}
