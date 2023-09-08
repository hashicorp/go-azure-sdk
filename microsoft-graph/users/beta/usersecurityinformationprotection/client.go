package usersecurityinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewUserSecurityInformationProtectionClientWithBaseURI(api sdkEnv.Api) (*UserSecurityInformationProtectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersecurityinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSecurityInformationProtectionClient: %+v", err)
	}

	return &UserSecurityInformationProtectionClient{
		Client: client,
	}, nil
}
