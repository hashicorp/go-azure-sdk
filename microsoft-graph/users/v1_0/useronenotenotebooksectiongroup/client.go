package useronenotenotebooksectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupClient{
		Client: client,
	}, nil
}
