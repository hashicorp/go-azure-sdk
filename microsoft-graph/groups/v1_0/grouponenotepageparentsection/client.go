package grouponenotepageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenotePageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupOnenotePageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupOnenotePageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotepageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenotePageParentSectionClient: %+v", err)
	}

	return &GroupOnenotePageParentSectionClient{
		Client: client,
	}, nil
}
