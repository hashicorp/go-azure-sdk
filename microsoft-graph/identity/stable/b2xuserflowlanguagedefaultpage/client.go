package b2xuserflowlanguagedefaultpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowLanguageDefaultPageClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowLanguageDefaultPageClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowLanguageDefaultPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowlanguagedefaultpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowLanguageDefaultPageClient: %+v", err)
	}

	return &B2xUserFlowLanguageDefaultPageClient{
		Client: client,
	}, nil
}
