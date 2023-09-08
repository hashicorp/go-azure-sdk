package meonenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookClient: %+v", err)
	}

	return &MeOnenoteNotebookClient{
		Client: client,
	}, nil
}
