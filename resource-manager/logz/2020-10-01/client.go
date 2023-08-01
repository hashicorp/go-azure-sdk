package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/singlesignon"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/subaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/vmhost"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Monitors     *monitors.MonitorsClient
	SingleSignOn *singlesignon.SingleSignOnClient
	SubAccount   *subaccount.SubAccountClient
	TagRules     *tagrules.TagRulesClient
	VMHost       *vmhost.VMHostClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	monitorsClient, err := monitors.NewMonitorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Monitors client: %+v", err)
	}
	configureFunc(monitorsClient.Client)

	singleSignOnClient, err := singlesignon.NewSingleSignOnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SingleSignOn client: %+v", err)
	}
	configureFunc(singleSignOnClient.Client)

	subAccountClient, err := subaccount.NewSubAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubAccount client: %+v", err)
	}
	configureFunc(subAccountClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	vMHostClient, err := vmhost.NewVMHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMHost client: %+v", err)
	}
	configureFunc(vMHostClient.Client)

	return &Client{
		Monitors:     monitorsClient,
		SingleSignOn: singleSignOnClient,
		SubAccount:   subAccountClient,
		TagRules:     tagRulesClient,
		VMHost:       vMHostClient,
	}, nil
}
