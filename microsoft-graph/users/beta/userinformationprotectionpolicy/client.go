package userinformationprotectionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionPolicyClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionPolicyClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionPolicyClient: %+v", err)
	}

	return &UserInformationProtectionPolicyClient{
		Client: client,
	}, nil
}
