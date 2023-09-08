package meinformationprotectiondatalosspreventionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionDataLossPreventionPolicyClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionDataLossPreventionPolicyClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionDataLossPreventionPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectiondatalosspreventionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionDataLossPreventionPolicyClient: %+v", err)
	}

	return &MeInformationProtectionDataLossPreventionPolicyClient{
		Client: client,
	}, nil
}
