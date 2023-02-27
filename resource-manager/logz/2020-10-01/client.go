package v2020_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/singlesignon"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/subaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/vmhost"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Monitors     *monitors.MonitorsClient
	SingleSignOn *singlesignon.SingleSignOnClient
	SubAccount   *subaccount.SubAccountClient
	TagRules     *tagrules.TagRulesClient
	VMHost       *vmhost.VMHostClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	monitorsClient := monitors.NewMonitorsClientWithBaseURI(endpoint)
	configureAuthFunc(&monitorsClient.Client)

	singleSignOnClient := singlesignon.NewSingleSignOnClientWithBaseURI(endpoint)
	configureAuthFunc(&singleSignOnClient.Client)

	subAccountClient := subaccount.NewSubAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&subAccountClient.Client)

	tagRulesClient := tagrules.NewTagRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&tagRulesClient.Client)

	vMHostClient := vmhost.NewVMHostClientWithBaseURI(endpoint)
	configureAuthFunc(&vMHostClient.Client)

	return Client{
		Monitors:     &monitorsClient,
		SingleSignOn: &singleSignOnClient,
		SubAccount:   &subAccountClient,
		TagRules:     &tagRulesClient,
		VMHost:       &vMHostClient,
	}
}
