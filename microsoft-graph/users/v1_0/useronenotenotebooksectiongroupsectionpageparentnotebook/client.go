package useronenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
