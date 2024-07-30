package informationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &InformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
