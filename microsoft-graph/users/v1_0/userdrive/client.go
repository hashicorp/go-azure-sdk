package userdrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDriveClient struct {
	Client *msgraph.Client
}

func NewUserDriveClientWithBaseURI(api sdkEnv.Api) (*UserDriveClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDriveClient: %+v", err)
	}

	return &UserDriveClient{
		Client: client,
	}, nil
}
