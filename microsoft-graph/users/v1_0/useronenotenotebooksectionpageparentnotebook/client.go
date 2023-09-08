package useronenotenotebooksectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionPageParentNotebookClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
