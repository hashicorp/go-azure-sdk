package calendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &CalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
