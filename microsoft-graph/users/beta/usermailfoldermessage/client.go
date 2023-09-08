package usermailfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderMessageClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderMessageClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderMessageClient: %+v", err)
	}

	return &UserMailFolderMessageClient{
		Client: client,
	}, nil
}
