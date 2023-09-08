package usermailfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderClient: %+v", err)
	}

	return &UserMailFolderClient{
		Client: client,
	}, nil
}
