package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/appcredentialsigninactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/applicationsignindetailedsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/authenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/authenticationmethoduserregistrationdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/credentialuserregistrationdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/dailyprintusage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/dailyprintusagebyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/dailyprintusagebyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/dailyprintusagesummariesbyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/dailyprintusagesummariesbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/healthmonitoring"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/healthmonitoringalert"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/healthmonitoringalertconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/monthlyprintusagebyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/monthlyprintusagebyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/monthlyprintusagesummariesbyprinter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/monthlyprintusagesummariesbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbilling"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingmanifest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationbilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationunbilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusagebilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusageunbilled"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/report"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/security"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/serviceactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/serviceprincipalsigninactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/sla"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/slaazureadauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/usercredentialusagedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdaily"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailyactiveuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailyauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailyinactiveuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailyinactiveusersbyapplication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailymfacompletion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailymfatelecomfraud"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailysignup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailysummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightdailyusercount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthly"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlyactiveuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlyauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlyinactiveuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlyinactiveusersbyapplication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlymfacompletion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlymfaregistereduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlyrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlysignup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/userinsightmonthlysummary"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppCredentialSignInActivity                  *appcredentialsigninactivity.AppCredentialSignInActivityClient
	ApplicationSignInDetailedSummary             *applicationsignindetailedsummary.ApplicationSignInDetailedSummaryClient
	AuthenticationMethod                         *authenticationmethod.AuthenticationMethodClient
	AuthenticationMethodUserRegistrationDetail   *authenticationmethoduserregistrationdetail.AuthenticationMethodUserRegistrationDetailClient
	CredentialUserRegistrationDetail             *credentialuserregistrationdetail.CredentialUserRegistrationDetailClient
	DailyPrintUsage                              *dailyprintusage.DailyPrintUsageClient
	DailyPrintUsageByPrinter                     *dailyprintusagebyprinter.DailyPrintUsageByPrinterClient
	DailyPrintUsageByUser                        *dailyprintusagebyuser.DailyPrintUsageByUserClient
	DailyPrintUsageSummariesByPrinter            *dailyprintusagesummariesbyprinter.DailyPrintUsageSummariesByPrinterClient
	DailyPrintUsageSummariesByUser               *dailyprintusagesummariesbyuser.DailyPrintUsageSummariesByUserClient
	HealthMonitoring                             *healthmonitoring.HealthMonitoringClient
	HealthMonitoringAlert                        *healthmonitoringalert.HealthMonitoringAlertClient
	HealthMonitoringAlertConfiguration           *healthmonitoringalertconfiguration.HealthMonitoringAlertConfigurationClient
	MonthlyPrintUsageByPrinter                   *monthlyprintusagebyprinter.MonthlyPrintUsageByPrinterClient
	MonthlyPrintUsageByUser                      *monthlyprintusagebyuser.MonthlyPrintUsageByUserClient
	MonthlyPrintUsageSummariesByPrinter          *monthlyprintusagesummariesbyprinter.MonthlyPrintUsageSummariesByPrinterClient
	MonthlyPrintUsageSummariesByUser             *monthlyprintusagesummariesbyuser.MonthlyPrintUsageSummariesByUserClient
	Partner                                      *partner.PartnerClient
	PartnerBilling                               *partnerbilling.PartnerBillingClient
	PartnerBillingManifest                       *partnerbillingmanifest.PartnerBillingManifestClient
	PartnerBillingOperation                      *partnerbillingoperation.PartnerBillingOperationClient
	PartnerBillingReconciliation                 *partnerbillingreconciliation.PartnerBillingReconciliationClient
	PartnerBillingReconciliationBilled           *partnerbillingreconciliationbilled.PartnerBillingReconciliationBilledClient
	PartnerBillingReconciliationUnbilled         *partnerbillingreconciliationunbilled.PartnerBillingReconciliationUnbilledClient
	PartnerBillingUsage                          *partnerbillingusage.PartnerBillingUsageClient
	PartnerBillingUsageBilled                    *partnerbillingusagebilled.PartnerBillingUsageBilledClient
	PartnerBillingUsageUnbilled                  *partnerbillingusageunbilled.PartnerBillingUsageUnbilledClient
	Report                                       *report.ReportClient
	Security                                     *security.SecurityClient
	ServiceActivity                              *serviceactivity.ServiceActivityClient
	ServicePrincipalSignInActivity               *serviceprincipalsigninactivity.ServicePrincipalSignInActivityClient
	Sla                                          *sla.SlaClient
	SlaAzureADAuthentication                     *slaazureadauthentication.SlaAzureADAuthenticationClient
	UserCredentialUsageDetail                    *usercredentialusagedetail.UserCredentialUsageDetailClient
	UserInsight                                  *userinsight.UserInsightClient
	UserInsightDaily                             *userinsightdaily.UserInsightDailyClient
	UserInsightDailyActiveUser                   *userinsightdailyactiveuser.UserInsightDailyActiveUserClient
	UserInsightDailyAuthentication               *userinsightdailyauthentication.UserInsightDailyAuthenticationClient
	UserInsightDailyInactiveUser                 *userinsightdailyinactiveuser.UserInsightDailyInactiveUserClient
	UserInsightDailyInactiveUsersByApplication   *userinsightdailyinactiveusersbyapplication.UserInsightDailyInactiveUsersByApplicationClient
	UserInsightDailyMfaCompletion                *userinsightdailymfacompletion.UserInsightDailyMfaCompletionClient
	UserInsightDailyMfaTelecomFraud              *userinsightdailymfatelecomfraud.UserInsightDailyMfaTelecomFraudClient
	UserInsightDailySignUp                       *userinsightdailysignup.UserInsightDailySignUpClient
	UserInsightDailySummary                      *userinsightdailysummary.UserInsightDailySummaryClient
	UserInsightDailyUserCount                    *userinsightdailyusercount.UserInsightDailyUserCountClient
	UserInsightMonthly                           *userinsightmonthly.UserInsightMonthlyClient
	UserInsightMonthlyActiveUser                 *userinsightmonthlyactiveuser.UserInsightMonthlyActiveUserClient
	UserInsightMonthlyAuthentication             *userinsightmonthlyauthentication.UserInsightMonthlyAuthenticationClient
	UserInsightMonthlyInactiveUser               *userinsightmonthlyinactiveuser.UserInsightMonthlyInactiveUserClient
	UserInsightMonthlyInactiveUsersByApplication *userinsightmonthlyinactiveusersbyapplication.UserInsightMonthlyInactiveUsersByApplicationClient
	UserInsightMonthlyMfaCompletion              *userinsightmonthlymfacompletion.UserInsightMonthlyMfaCompletionClient
	UserInsightMonthlyMfaRegisteredUser          *userinsightmonthlymfaregistereduser.UserInsightMonthlyMfaRegisteredUserClient
	UserInsightMonthlyRequest                    *userinsightmonthlyrequest.UserInsightMonthlyRequestClient
	UserInsightMonthlySignUp                     *userinsightmonthlysignup.UserInsightMonthlySignUpClient
	UserInsightMonthlySummary                    *userinsightmonthlysummary.UserInsightMonthlySummaryClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	appCredentialSignInActivityClient, err := appcredentialsigninactivity.NewAppCredentialSignInActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppCredentialSignInActivity client: %+v", err)
	}
	configureFunc(appCredentialSignInActivityClient.Client)

	applicationSignInDetailedSummaryClient, err := applicationsignindetailedsummary.NewApplicationSignInDetailedSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSignInDetailedSummary client: %+v", err)
	}
	configureFunc(applicationSignInDetailedSummaryClient.Client)

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

	credentialUserRegistrationDetailClient, err := credentialuserregistrationdetail.NewCredentialUserRegistrationDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CredentialUserRegistrationDetail client: %+v", err)
	}
	configureFunc(credentialUserRegistrationDetailClient.Client)

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

	dailyPrintUsageClient, err := dailyprintusage.NewDailyPrintUsageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DailyPrintUsage client: %+v", err)
	}
	configureFunc(dailyPrintUsageClient.Client)

	dailyPrintUsageSummariesByPrinterClient, err := dailyprintusagesummariesbyprinter.NewDailyPrintUsageSummariesByPrinterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DailyPrintUsageSummariesByPrinter client: %+v", err)
	}
	configureFunc(dailyPrintUsageSummariesByPrinterClient.Client)

	dailyPrintUsageSummariesByUserClient, err := dailyprintusagesummariesbyuser.NewDailyPrintUsageSummariesByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DailyPrintUsageSummariesByUser client: %+v", err)
	}
	configureFunc(dailyPrintUsageSummariesByUserClient.Client)

	healthMonitoringAlertClient, err := healthmonitoringalert.NewHealthMonitoringAlertClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HealthMonitoringAlert client: %+v", err)
	}
	configureFunc(healthMonitoringAlertClient.Client)

	healthMonitoringAlertConfigurationClient, err := healthmonitoringalertconfiguration.NewHealthMonitoringAlertConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HealthMonitoringAlertConfiguration client: %+v", err)
	}
	configureFunc(healthMonitoringAlertConfigurationClient.Client)

	healthMonitoringClient, err := healthmonitoring.NewHealthMonitoringClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HealthMonitoring client: %+v", err)
	}
	configureFunc(healthMonitoringClient.Client)

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

	monthlyPrintUsageSummariesByPrinterClient, err := monthlyprintusagesummariesbyprinter.NewMonthlyPrintUsageSummariesByPrinterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonthlyPrintUsageSummariesByPrinter client: %+v", err)
	}
	configureFunc(monthlyPrintUsageSummariesByPrinterClient.Client)

	monthlyPrintUsageSummariesByUserClient, err := monthlyprintusagesummariesbyuser.NewMonthlyPrintUsageSummariesByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonthlyPrintUsageSummariesByUser client: %+v", err)
	}
	configureFunc(monthlyPrintUsageSummariesByUserClient.Client)

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

	partnerBillingReconciliationUnbilledClient, err := partnerbillingreconciliationunbilled.NewPartnerBillingReconciliationUnbilledClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PartnerBillingReconciliationUnbilled client: %+v", err)
	}
	configureFunc(partnerBillingReconciliationUnbilledClient.Client)

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

	serviceActivityClient, err := serviceactivity.NewServiceActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceActivity client: %+v", err)
	}
	configureFunc(serviceActivityClient.Client)

	servicePrincipalSignInActivityClient, err := serviceprincipalsigninactivity.NewServicePrincipalSignInActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSignInActivity client: %+v", err)
	}
	configureFunc(servicePrincipalSignInActivityClient.Client)

	slaAzureADAuthenticationClient, err := slaazureadauthentication.NewSlaAzureADAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SlaAzureADAuthentication client: %+v", err)
	}
	configureFunc(slaAzureADAuthenticationClient.Client)

	slaClient, err := sla.NewSlaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sla client: %+v", err)
	}
	configureFunc(slaClient.Client)

	userCredentialUsageDetailClient, err := usercredentialusagedetail.NewUserCredentialUsageDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCredentialUsageDetail client: %+v", err)
	}
	configureFunc(userCredentialUsageDetailClient.Client)

	userInsightClient, err := userinsight.NewUserInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsight client: %+v", err)
	}
	configureFunc(userInsightClient.Client)

	userInsightDailyActiveUserClient, err := userinsightdailyactiveuser.NewUserInsightDailyActiveUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyActiveUser client: %+v", err)
	}
	configureFunc(userInsightDailyActiveUserClient.Client)

	userInsightDailyAuthenticationClient, err := userinsightdailyauthentication.NewUserInsightDailyAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyAuthentication client: %+v", err)
	}
	configureFunc(userInsightDailyAuthenticationClient.Client)

	userInsightDailyClient, err := userinsightdaily.NewUserInsightDailyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDaily client: %+v", err)
	}
	configureFunc(userInsightDailyClient.Client)

	userInsightDailyInactiveUserClient, err := userinsightdailyinactiveuser.NewUserInsightDailyInactiveUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyInactiveUser client: %+v", err)
	}
	configureFunc(userInsightDailyInactiveUserClient.Client)

	userInsightDailyInactiveUsersByApplicationClient, err := userinsightdailyinactiveusersbyapplication.NewUserInsightDailyInactiveUsersByApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyInactiveUsersByApplication client: %+v", err)
	}
	configureFunc(userInsightDailyInactiveUsersByApplicationClient.Client)

	userInsightDailyMfaCompletionClient, err := userinsightdailymfacompletion.NewUserInsightDailyMfaCompletionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyMfaCompletion client: %+v", err)
	}
	configureFunc(userInsightDailyMfaCompletionClient.Client)

	userInsightDailyMfaTelecomFraudClient, err := userinsightdailymfatelecomfraud.NewUserInsightDailyMfaTelecomFraudClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyMfaTelecomFraud client: %+v", err)
	}
	configureFunc(userInsightDailyMfaTelecomFraudClient.Client)

	userInsightDailySignUpClient, err := userinsightdailysignup.NewUserInsightDailySignUpClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailySignUp client: %+v", err)
	}
	configureFunc(userInsightDailySignUpClient.Client)

	userInsightDailySummaryClient, err := userinsightdailysummary.NewUserInsightDailySummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailySummary client: %+v", err)
	}
	configureFunc(userInsightDailySummaryClient.Client)

	userInsightDailyUserCountClient, err := userinsightdailyusercount.NewUserInsightDailyUserCountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightDailyUserCount client: %+v", err)
	}
	configureFunc(userInsightDailyUserCountClient.Client)

	userInsightMonthlyActiveUserClient, err := userinsightmonthlyactiveuser.NewUserInsightMonthlyActiveUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyActiveUser client: %+v", err)
	}
	configureFunc(userInsightMonthlyActiveUserClient.Client)

	userInsightMonthlyAuthenticationClient, err := userinsightmonthlyauthentication.NewUserInsightMonthlyAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyAuthentication client: %+v", err)
	}
	configureFunc(userInsightMonthlyAuthenticationClient.Client)

	userInsightMonthlyClient, err := userinsightmonthly.NewUserInsightMonthlyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthly client: %+v", err)
	}
	configureFunc(userInsightMonthlyClient.Client)

	userInsightMonthlyInactiveUserClient, err := userinsightmonthlyinactiveuser.NewUserInsightMonthlyInactiveUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyInactiveUser client: %+v", err)
	}
	configureFunc(userInsightMonthlyInactiveUserClient.Client)

	userInsightMonthlyInactiveUsersByApplicationClient, err := userinsightmonthlyinactiveusersbyapplication.NewUserInsightMonthlyInactiveUsersByApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyInactiveUsersByApplication client: %+v", err)
	}
	configureFunc(userInsightMonthlyInactiveUsersByApplicationClient.Client)

	userInsightMonthlyMfaCompletionClient, err := userinsightmonthlymfacompletion.NewUserInsightMonthlyMfaCompletionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyMfaCompletion client: %+v", err)
	}
	configureFunc(userInsightMonthlyMfaCompletionClient.Client)

	userInsightMonthlyMfaRegisteredUserClient, err := userinsightmonthlymfaregistereduser.NewUserInsightMonthlyMfaRegisteredUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyMfaRegisteredUser client: %+v", err)
	}
	configureFunc(userInsightMonthlyMfaRegisteredUserClient.Client)

	userInsightMonthlyRequestClient, err := userinsightmonthlyrequest.NewUserInsightMonthlyRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlyRequest client: %+v", err)
	}
	configureFunc(userInsightMonthlyRequestClient.Client)

	userInsightMonthlySignUpClient, err := userinsightmonthlysignup.NewUserInsightMonthlySignUpClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlySignUp client: %+v", err)
	}
	configureFunc(userInsightMonthlySignUpClient.Client)

	userInsightMonthlySummaryClient, err := userinsightmonthlysummary.NewUserInsightMonthlySummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightMonthlySummary client: %+v", err)
	}
	configureFunc(userInsightMonthlySummaryClient.Client)

	return &Client{
		AppCredentialSignInActivity:                  appCredentialSignInActivityClient,
		ApplicationSignInDetailedSummary:             applicationSignInDetailedSummaryClient,
		AuthenticationMethod:                         authenticationMethodClient,
		AuthenticationMethodUserRegistrationDetail:   authenticationMethodUserRegistrationDetailClient,
		CredentialUserRegistrationDetail:             credentialUserRegistrationDetailClient,
		DailyPrintUsage:                              dailyPrintUsageClient,
		DailyPrintUsageByPrinter:                     dailyPrintUsageByPrinterClient,
		DailyPrintUsageByUser:                        dailyPrintUsageByUserClient,
		DailyPrintUsageSummariesByPrinter:            dailyPrintUsageSummariesByPrinterClient,
		DailyPrintUsageSummariesByUser:               dailyPrintUsageSummariesByUserClient,
		HealthMonitoring:                             healthMonitoringClient,
		HealthMonitoringAlert:                        healthMonitoringAlertClient,
		HealthMonitoringAlertConfiguration:           healthMonitoringAlertConfigurationClient,
		MonthlyPrintUsageByPrinter:                   monthlyPrintUsageByPrinterClient,
		MonthlyPrintUsageByUser:                      monthlyPrintUsageByUserClient,
		MonthlyPrintUsageSummariesByPrinter:          monthlyPrintUsageSummariesByPrinterClient,
		MonthlyPrintUsageSummariesByUser:             monthlyPrintUsageSummariesByUserClient,
		Partner:                                      partnerClient,
		PartnerBilling:                               partnerBillingClient,
		PartnerBillingManifest:                       partnerBillingManifestClient,
		PartnerBillingOperation:                      partnerBillingOperationClient,
		PartnerBillingReconciliation:                 partnerBillingReconciliationClient,
		PartnerBillingReconciliationBilled:           partnerBillingReconciliationBilledClient,
		PartnerBillingReconciliationUnbilled:         partnerBillingReconciliationUnbilledClient,
		PartnerBillingUsage:                          partnerBillingUsageClient,
		PartnerBillingUsageBilled:                    partnerBillingUsageBilledClient,
		PartnerBillingUsageUnbilled:                  partnerBillingUsageUnbilledClient,
		Report:                                       reportClient,
		Security:                                     securityClient,
		ServiceActivity:                              serviceActivityClient,
		ServicePrincipalSignInActivity:               servicePrincipalSignInActivityClient,
		Sla:                                          slaClient,
		SlaAzureADAuthentication:                     slaAzureADAuthenticationClient,
		UserCredentialUsageDetail:                    userCredentialUsageDetailClient,
		UserInsight:                                  userInsightClient,
		UserInsightDaily:                             userInsightDailyClient,
		UserInsightDailyActiveUser:                   userInsightDailyActiveUserClient,
		UserInsightDailyAuthentication:               userInsightDailyAuthenticationClient,
		UserInsightDailyInactiveUser:                 userInsightDailyInactiveUserClient,
		UserInsightDailyInactiveUsersByApplication:   userInsightDailyInactiveUsersByApplicationClient,
		UserInsightDailyMfaCompletion:                userInsightDailyMfaCompletionClient,
		UserInsightDailyMfaTelecomFraud:              userInsightDailyMfaTelecomFraudClient,
		UserInsightDailySignUp:                       userInsightDailySignUpClient,
		UserInsightDailySummary:                      userInsightDailySummaryClient,
		UserInsightDailyUserCount:                    userInsightDailyUserCountClient,
		UserInsightMonthly:                           userInsightMonthlyClient,
		UserInsightMonthlyActiveUser:                 userInsightMonthlyActiveUserClient,
		UserInsightMonthlyAuthentication:             userInsightMonthlyAuthenticationClient,
		UserInsightMonthlyInactiveUser:               userInsightMonthlyInactiveUserClient,
		UserInsightMonthlyInactiveUsersByApplication: userInsightMonthlyInactiveUsersByApplicationClient,
		UserInsightMonthlyMfaCompletion:              userInsightMonthlyMfaCompletionClient,
		UserInsightMonthlyMfaRegisteredUser:          userInsightMonthlyMfaRegisteredUserClient,
		UserInsightMonthlyRequest:                    userInsightMonthlyRequestClient,
		UserInsightMonthlySignUp:                     userInsightMonthlySignUpClient,
		UserInsightMonthlySummary:                    userInsightMonthlySummaryClient,
	}, nil
}
