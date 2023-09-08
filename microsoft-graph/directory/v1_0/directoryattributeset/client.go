package directoryattributeset

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAttributeSetClient struct {
	Client *msgraph.Client
}

func NewDirectoryAttributeSetClientWithBaseURI(api sdkEnv.Api) (*DirectoryAttributeSetClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryattributeset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryAttributeSetClient: %+v", err)
	}

	return &DirectoryAttributeSetClient{
		Client: client,
	}, nil
}
