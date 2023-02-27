// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2019_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/consumerinvitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/dataset"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/datasetmapping"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/emailregistration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/invitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/share"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/sharesubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/synchronizationsetting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/trigger"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountClient := account.NewAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&accountClient.Client)

	consumerInvitationClient := consumerinvitation.NewConsumerInvitationClientWithBaseURI(endpoint)
	configureAuthFunc(&consumerInvitationClient.Client)

	dataSetClient := dataset.NewDataSetClientWithBaseURI(endpoint)
	configureAuthFunc(&dataSetClient.Client)

	dataSetMappingClient := datasetmapping.NewDataSetMappingClientWithBaseURI(endpoint)
	configureAuthFunc(&dataSetMappingClient.Client)

	emailRegistrationClient := emailregistration.NewEmailRegistrationClientWithBaseURI(endpoint)
	configureAuthFunc(&emailRegistrationClient.Client)

	invitationClient := invitation.NewInvitationClientWithBaseURI(endpoint)
	configureAuthFunc(&invitationClient.Client)

	shareClient := share.NewShareClientWithBaseURI(endpoint)
	configureAuthFunc(&shareClient.Client)

	shareSubscriptionClient := sharesubscription.NewShareSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&shareSubscriptionClient.Client)

	synchronizationSettingClient := synchronizationsetting.NewSynchronizationSettingClientWithBaseURI(endpoint)
	configureAuthFunc(&synchronizationSettingClient.Client)

	triggerClient := trigger.NewTriggerClientWithBaseURI(endpoint)
	configureAuthFunc(&triggerClient.Client)

	return Client{
		Account:                &accountClient,
		ConsumerInvitation:     &consumerInvitationClient,
		DataSet:                &dataSetClient,
		DataSetMapping:         &dataSetMappingClient,
		EmailRegistration:      &emailRegistrationClient,
		Invitation:             &invitationClient,
		Share:                  &shareClient,
		ShareSubscription:      &shareSubscriptionClient,
		SynchronizationSetting: &synchronizationSettingClient,
		Trigger:                &triggerClient,
	}
}
