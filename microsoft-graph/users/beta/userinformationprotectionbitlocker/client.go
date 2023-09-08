package userinformationprotectionbitlocker

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionBitlockerClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionBitlockerClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionBitlockerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionbitlocker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionBitlockerClient: %+v", err)
	}

	return &UserInformationProtectionBitlockerClient{
		Client: client,
	}, nil
}
