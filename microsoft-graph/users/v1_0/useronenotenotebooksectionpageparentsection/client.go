package useronenotenotebooksectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionPageParentSectionClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionPageParentSectionClient{
		Client: client,
	}, nil
}
