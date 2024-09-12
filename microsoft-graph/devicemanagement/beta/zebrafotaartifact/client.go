package zebrafotaartifact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaArtifactClient struct {
	Client *msgraph.Client
}

func NewZebraFotaArtifactClientWithBaseURI(sdkApi sdkEnv.Api) (*ZebraFotaArtifactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "zebrafotaartifact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ZebraFotaArtifactClient: %+v", err)
	}

	return &ZebraFotaArtifactClient{
		Client: client,
	}, nil
}
