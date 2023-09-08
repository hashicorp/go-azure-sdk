package useronenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
