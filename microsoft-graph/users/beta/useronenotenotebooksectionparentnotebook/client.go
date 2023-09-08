package useronenotenotebooksectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionParentNotebookClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionParentNotebookClient{
		Client: client,
	}, nil
}
