package meprofilecertification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileCertificationClient struct {
	Client *msgraph.Client
}

func NewMeProfileCertificationClientWithBaseURI(api sdkEnv.Api) (*MeProfileCertificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilecertification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileCertificationClient: %+v", err)
	}

	return &MeProfileCertificationClient{
		Client: client,
	}, nil
}
