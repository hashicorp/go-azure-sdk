package b2xuserflowlanguage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowLanguageClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowLanguageClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowLanguageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowlanguage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowLanguageClient: %+v", err)
	}

	return &B2xUserFlowLanguageClient{
		Client: client,
	}, nil
}
