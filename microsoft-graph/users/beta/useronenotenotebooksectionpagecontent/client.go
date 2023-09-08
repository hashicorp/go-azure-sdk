package useronenotenotebooksectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionPageContentClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionPageContentClient{
		Client: client,
	}, nil
}
