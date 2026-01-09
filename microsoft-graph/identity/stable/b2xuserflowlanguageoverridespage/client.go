package b2xuserflowlanguageoverridespage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowLanguageOverridesPageClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowLanguageOverridesPageClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowLanguageOverridesPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowlanguageoverridespage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowLanguageOverridesPageClient: %+v", err)
	}

	return &B2xUserFlowLanguageOverridesPageClient{
		Client: client,
	}, nil
}
