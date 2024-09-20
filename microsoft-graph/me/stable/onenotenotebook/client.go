package onenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookClient: %+v", err)
	}

	return &OnenoteNotebookClient{
		Client: client,
	}, nil
}
