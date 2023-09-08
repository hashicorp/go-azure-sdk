package useronenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookClient: %+v", err)
	}

	return &UserOnenoteNotebookClient{
		Client: client,
	}, nil
}
