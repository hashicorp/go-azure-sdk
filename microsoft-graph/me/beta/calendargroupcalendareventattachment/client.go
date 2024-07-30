package calendargroupcalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
