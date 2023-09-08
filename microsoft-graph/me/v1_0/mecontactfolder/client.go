package mecontactfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderClient: %+v", err)
	}

	return &MeContactFolderClient{
		Client: client,
	}, nil
}
