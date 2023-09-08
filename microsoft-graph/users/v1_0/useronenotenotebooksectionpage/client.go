package useronenotenotebooksectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionPageClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
