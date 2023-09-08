package grouponenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookClient: %+v", err)
	}

	return &GroupOnenoteNotebookClient{
		Client: client,
	}, nil
}
