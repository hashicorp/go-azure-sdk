package v2025_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalimitrequest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalimits"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalocationsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotassubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotausages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotainformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotarequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/subscriptionquotaallocation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/subscriptionquotaallocationrequest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/usagesinformation"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GroupQuotaLimitRequest             *groupquotalimitrequest.GroupQuotaLimitRequestClient
	GroupQuotaLimits                   *groupquotalimits.GroupQuotaLimitsClient
	GroupQuotaLocationSettings         *groupquotalocationsettings.GroupQuotaLocationSettingsClient
	GroupQuotaSubscriptionRequests     *groupquotasubscriptionrequests.GroupQuotaSubscriptionRequestsClient
	GroupQuotaUsages                   *groupquotausages.GroupQuotaUsagesClient
	GroupQuotas                        *groupquotas.GroupQuotasClient
	GroupQuotasSubscriptions           *groupquotassubscriptions.GroupQuotasSubscriptionsClient
	QuotaInformation                   *quotainformation.QuotaInformationClient
	QuotaRequests                      *quotarequests.QuotaRequestsClient
	SubscriptionQuotaAllocation        *subscriptionquotaallocation.SubscriptionQuotaAllocationClient
	SubscriptionQuotaAllocationRequest *subscriptionquotaallocationrequest.SubscriptionQuotaAllocationRequestClient
	UsagesInformation                  *usagesinformation.UsagesInformationClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	groupQuotaLimitRequestClient, err := groupquotalimitrequest.NewGroupQuotaLimitRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaLimitRequest client: %+v", err)
	}
	configureFunc(groupQuotaLimitRequestClient.Client)

	groupQuotaLimitsClient, err := groupquotalimits.NewGroupQuotaLimitsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaLimits client: %+v", err)
	}
	configureFunc(groupQuotaLimitsClient.Client)

	groupQuotaLocationSettingsClient, err := groupquotalocationsettings.NewGroupQuotaLocationSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaLocationSettings client: %+v", err)
	}
	configureFunc(groupQuotaLocationSettingsClient.Client)

	groupQuotaSubscriptionRequestsClient, err := groupquotasubscriptionrequests.NewGroupQuotaSubscriptionRequestsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaSubscriptionRequests client: %+v", err)
	}
	configureFunc(groupQuotaSubscriptionRequestsClient.Client)

	groupQuotaUsagesClient, err := groupquotausages.NewGroupQuotaUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaUsages client: %+v", err)
	}
	configureFunc(groupQuotaUsagesClient.Client)

	groupQuotasClient, err := groupquotas.NewGroupQuotasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotas client: %+v", err)
	}
	configureFunc(groupQuotasClient.Client)

	groupQuotasSubscriptionsClient, err := groupquotassubscriptions.NewGroupQuotasSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotasSubscriptions client: %+v", err)
	}
	configureFunc(groupQuotasSubscriptionsClient.Client)

	quotaInformationClient, err := quotainformation.NewQuotaInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaInformation client: %+v", err)
	}
	configureFunc(quotaInformationClient.Client)

	quotaRequestsClient, err := quotarequests.NewQuotaRequestsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaRequests client: %+v", err)
	}
	configureFunc(quotaRequestsClient.Client)

	subscriptionQuotaAllocationClient, err := subscriptionquotaallocation.NewSubscriptionQuotaAllocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionQuotaAllocation client: %+v", err)
	}
	configureFunc(subscriptionQuotaAllocationClient.Client)

	subscriptionQuotaAllocationRequestClient, err := subscriptionquotaallocationrequest.NewSubscriptionQuotaAllocationRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionQuotaAllocationRequest client: %+v", err)
	}
	configureFunc(subscriptionQuotaAllocationRequestClient.Client)

	usagesInformationClient, err := usagesinformation.NewUsagesInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsagesInformation client: %+v", err)
	}
	configureFunc(usagesInformationClient.Client)

	return &Client{
		GroupQuotaLimitRequest:             groupQuotaLimitRequestClient,
		GroupQuotaLimits:                   groupQuotaLimitsClient,
		GroupQuotaLocationSettings:         groupQuotaLocationSettingsClient,
		GroupQuotaSubscriptionRequests:     groupQuotaSubscriptionRequestsClient,
		GroupQuotaUsages:                   groupQuotaUsagesClient,
		GroupQuotas:                        groupQuotasClient,
		GroupQuotasSubscriptions:           groupQuotasSubscriptionsClient,
		QuotaInformation:                   quotaInformationClient,
		QuotaRequests:                      quotaRequestsClient,
		SubscriptionQuotaAllocation:        subscriptionQuotaAllocationClient,
		SubscriptionQuotaAllocationRequest: subscriptionQuotaAllocationRequestClient,
		UsagesInformation:                  usagesInformationClient,
	}, nil
}
