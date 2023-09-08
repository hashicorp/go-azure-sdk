package memailfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderClient: %+v", err)
	}

	return &MeMailFolderClient{
		Client: client,
	}, nil
}
