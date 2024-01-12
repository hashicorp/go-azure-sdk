package v2021_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/consumerinvitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/dataset"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/datasetmapping"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/emailregistration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/invitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/share"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/sharesubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/synchronizationsetting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/trigger"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Account                *account.AccountClient
	ConsumerInvitation     *consumerinvitation.ConsumerInvitationClient
	DataSet                *dataset.DataSetClient
	DataSetMapping         *datasetmapping.DataSetMappingClient
	EmailRegistration      *emailregistration.EmailRegistrationClient
	Invitation             *invitation.InvitationClient
	Share                  *share.ShareClient
	ShareSubscription      *sharesubscription.ShareSubscriptionClient
	SynchronizationSetting *synchronizationsetting.SynchronizationSettingClient
	Trigger                *trigger.TriggerClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountClient, err := account.NewAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Account client: %+v", err)
	}
	configureFunc(accountClient.Client)

	consumerInvitationClient, err := consumerinvitation.NewConsumerInvitationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConsumerInvitation client: %+v", err)
	}
	configureFunc(consumerInvitationClient.Client)

	dataSetClient, err := dataset.NewDataSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataSet client: %+v", err)
	}
	configureFunc(dataSetClient.Client)

	dataSetMappingClient, err := datasetmapping.NewDataSetMappingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataSetMapping client: %+v", err)
	}
	configureFunc(dataSetMappingClient.Client)

	emailRegistrationClient, err := emailregistration.NewEmailRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmailRegistration client: %+v", err)
	}
	configureFunc(emailRegistrationClient.Client)

	invitationClient, err := invitation.NewInvitationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Invitation client: %+v", err)
	}
	configureFunc(invitationClient.Client)

	shareClient, err := share.NewShareClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Share client: %+v", err)
	}
	configureFunc(shareClient.Client)

	shareSubscriptionClient, err := sharesubscription.NewShareSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ShareSubscription client: %+v", err)
	}
	configureFunc(shareSubscriptionClient.Client)

	synchronizationSettingClient, err := synchronizationsetting.NewSynchronizationSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationSetting client: %+v", err)
	}
	configureFunc(synchronizationSettingClient.Client)

	triggerClient, err := trigger.NewTriggerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Trigger client: %+v", err)
	}
	configureFunc(triggerClient.Client)

	return &Client{
		Account:                accountClient,
		ConsumerInvitation:     consumerInvitationClient,
		DataSet:                dataSetClient,
		DataSetMapping:         dataSetMappingClient,
		EmailRegistration:      emailRegistrationClient,
		Invitation:             invitationClient,
		Share:                  shareClient,
		ShareSubscription:      shareSubscriptionClient,
		SynchronizationSetting: synchronizationSettingClient,
		Trigger:                triggerClient,
	}, nil
}
