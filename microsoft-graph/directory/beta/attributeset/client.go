package attributeset

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeSetClient struct {
	Client *msgraph.Client
}

func NewAttributeSetClientWithBaseURI(sdkApi sdkEnv.Api) (*AttributeSetClient, error) {
	client, err := msgraph.NewClient(sdkApi, "attributeset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttributeSetClient: %+v", err)
	}

	return &AttributeSetClient{
		Client: client,
	}, nil
}
