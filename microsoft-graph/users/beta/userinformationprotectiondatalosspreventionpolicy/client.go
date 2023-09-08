package userinformationprotectiondatalosspreventionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionDataLossPreventionPolicyClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionDataLossPreventionPolicyClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionDataLossPreventionPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectiondatalosspreventionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionDataLossPreventionPolicyClient: %+v", err)
	}

	return &UserInformationProtectionDataLossPreventionPolicyClient{
		Client: client,
	}, nil
}
