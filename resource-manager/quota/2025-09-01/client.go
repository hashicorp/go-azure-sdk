package v2025_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/currentquotalimitbases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/currentusagesbases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalimitlists"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasenforcementstatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasentities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequeststatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotaallocationrequeststatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotarequeststatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/submittedresourcerequeststatuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/subscriptionquotaallocationslists"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CurrentQuotaLimitBases                *currentquotalimitbases.CurrentQuotaLimitBasesClient
	CurrentUsagesBases                    *currentusagesbases.CurrentUsagesBasesClient
	GroupQuotaLimitLists                  *groupquotalimitlists.GroupQuotaLimitListsClient
	GroupQuotaSubscriptionIds             *groupquotasubscriptionids.GroupQuotaSubscriptionIdsClient
	GroupQuotaSubscriptionRequestStatuses *groupquotasubscriptionrequeststatuses.GroupQuotaSubscriptionRequestStatusesClient
	GroupQuotasEnforcementStatuses        *groupquotasenforcementstatuses.GroupQuotasEnforcementStatusesClient
	GroupQuotasEntities                   *groupquotasentities.GroupQuotasEntitiesClient
	QuotaAllocationRequestStatuses        *quotaallocationrequeststatuses.QuotaAllocationRequestStatusesClient
	QuotaRequestStatus                    *quotarequeststatus.QuotaRequestStatusClient
	SubmittedResourceRequestStatuses      *submittedresourcerequeststatuses.SubmittedResourceRequestStatusesClient
	SubscriptionQuotaAllocationsLists     *subscriptionquotaallocationslists.SubscriptionQuotaAllocationsListsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	currentQuotaLimitBasesClient, err := currentquotalimitbases.NewCurrentQuotaLimitBasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CurrentQuotaLimitBases client: %+v", err)
	}
	configureFunc(currentQuotaLimitBasesClient.Client)

	currentUsagesBasesClient, err := currentusagesbases.NewCurrentUsagesBasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CurrentUsagesBases client: %+v", err)
	}
	configureFunc(currentUsagesBasesClient.Client)

	groupQuotaLimitListsClient, err := groupquotalimitlists.NewGroupQuotaLimitListsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaLimitLists client: %+v", err)
	}
	configureFunc(groupQuotaLimitListsClient.Client)

	groupQuotaSubscriptionIdsClient, err := groupquotasubscriptionids.NewGroupQuotaSubscriptionIdsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaSubscriptionIds client: %+v", err)
	}
	configureFunc(groupQuotaSubscriptionIdsClient.Client)

	groupQuotaSubscriptionRequestStatusesClient, err := groupquotasubscriptionrequeststatuses.NewGroupQuotaSubscriptionRequestStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotaSubscriptionRequestStatuses client: %+v", err)
	}
	configureFunc(groupQuotaSubscriptionRequestStatusesClient.Client)

	groupQuotasEnforcementStatusesClient, err := groupquotasenforcementstatuses.NewGroupQuotasEnforcementStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotasEnforcementStatuses client: %+v", err)
	}
	configureFunc(groupQuotasEnforcementStatusesClient.Client)

	groupQuotasEntitiesClient, err := groupquotasentities.NewGroupQuotasEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupQuotasEntities client: %+v", err)
	}
	configureFunc(groupQuotasEntitiesClient.Client)

	quotaAllocationRequestStatusesClient, err := quotaallocationrequeststatuses.NewQuotaAllocationRequestStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaAllocationRequestStatuses client: %+v", err)
	}
	configureFunc(quotaAllocationRequestStatusesClient.Client)

	quotaRequestStatusClient, err := quotarequeststatus.NewQuotaRequestStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaRequestStatus client: %+v", err)
	}
	configureFunc(quotaRequestStatusClient.Client)

	submittedResourceRequestStatusesClient, err := submittedresourcerequeststatuses.NewSubmittedResourceRequestStatusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubmittedResourceRequestStatuses client: %+v", err)
	}
	configureFunc(submittedResourceRequestStatusesClient.Client)

	subscriptionQuotaAllocationsListsClient, err := subscriptionquotaallocationslists.NewSubscriptionQuotaAllocationsListsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionQuotaAllocationsLists client: %+v", err)
	}
	configureFunc(subscriptionQuotaAllocationsListsClient.Client)

	return &Client{
		CurrentQuotaLimitBases:                currentQuotaLimitBasesClient,
		CurrentUsagesBases:                    currentUsagesBasesClient,
		GroupQuotaLimitLists:                  groupQuotaLimitListsClient,
		GroupQuotaSubscriptionIds:             groupQuotaSubscriptionIdsClient,
		GroupQuotaSubscriptionRequestStatuses: groupQuotaSubscriptionRequestStatusesClient,
		GroupQuotasEnforcementStatuses:        groupQuotasEnforcementStatusesClient,
		GroupQuotasEntities:                   groupQuotasEntitiesClient,
		QuotaAllocationRequestStatuses:        quotaAllocationRequestStatusesClient,
		QuotaRequestStatus:                    quotaRequestStatusClient,
		SubmittedResourceRequestStatuses:      submittedResourceRequestStatusesClient,
		SubscriptionQuotaAllocationsLists:     subscriptionQuotaAllocationsListsClient,
	}, nil
}
