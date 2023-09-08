package mecontactfoldercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderContactClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderContactClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderContactClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfoldercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderContactClient: %+v", err)
	}

	return &MeContactFolderContactClient{
		Client: client,
	}, nil
}
