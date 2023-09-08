package meonenotenotebooksectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionPageParentSectionClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionPageParentSectionClient{
		Client: client,
	}, nil
}
