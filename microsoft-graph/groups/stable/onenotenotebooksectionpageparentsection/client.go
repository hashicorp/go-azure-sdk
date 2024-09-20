package onenotenotebooksectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionPageParentSectionClient: %+v", err)
	}

	return &OnenoteNotebookSectionPageParentSectionClient{
		Client: client,
	}, nil
}
