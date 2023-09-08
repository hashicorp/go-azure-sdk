package meinformationprotectionthreatassessmentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionThreatAssessmentRequestClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionThreatAssessmentRequestClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionThreatAssessmentRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectionthreatassessmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionThreatAssessmentRequestClient: %+v", err)
	}

	return &MeInformationProtectionThreatAssessmentRequestClient{
		Client: client,
	}, nil
}
