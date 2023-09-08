package memailfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderMessageClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderMessageClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderMessageClient: %+v", err)
	}

	return &MeMailFolderMessageClient{
		Client: client,
	}, nil
}
