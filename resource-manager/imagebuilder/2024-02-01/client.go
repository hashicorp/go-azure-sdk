package v2024_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/imagebuilder/2024-02-01/triggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/imagebuilder/2024-02-01/virtualmachineimagetemplate"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Triggers                    *triggers.TriggersClient
	VirtualMachineImageTemplate *virtualmachineimagetemplate.VirtualMachineImageTemplateClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	triggersClient, err := triggers.NewTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Triggers client: %+v", err)
	}
	configureFunc(triggersClient.Client)

	virtualMachineImageTemplateClient, err := virtualmachineimagetemplate.NewVirtualMachineImageTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineImageTemplate client: %+v", err)
	}
	configureFunc(virtualMachineImageTemplateClient.Client)

	return &Client{
		Triggers:                    triggersClient,
		VirtualMachineImageTemplate: virtualMachineImageTemplateClient,
	}, nil
}
