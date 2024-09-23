package informationprotectiondatalosspreventionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionDataLossPreventionPolicyClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionDataLossPreventionPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectiondatalosspreventionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionDataLossPreventionPolicyClient: %+v", err)
	}

	return &InformationProtectionDataLossPreventionPolicyClient{
		Client: client,
	}, nil
}
