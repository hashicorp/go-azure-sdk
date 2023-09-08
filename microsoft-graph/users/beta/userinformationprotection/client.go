package userinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionClient: %+v", err)
	}

	return &UserInformationProtectionClient{
		Client: client,
	}, nil
}
