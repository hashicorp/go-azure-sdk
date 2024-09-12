package informationprotectionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPolicyClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionPolicyClient: %+v", err)
	}

	return &InformationProtectionPolicyClient{
		Client: client,
	}, nil
}
