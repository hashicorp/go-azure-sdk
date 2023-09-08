package useronenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
