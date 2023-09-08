package grouponenotenotebooksectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionPageParentSectionClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionPageParentSectionClient{
		Client: client,
	}, nil
}
