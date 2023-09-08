package useronenotenotebooksection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionClient{
		Client: client,
	}, nil
}
