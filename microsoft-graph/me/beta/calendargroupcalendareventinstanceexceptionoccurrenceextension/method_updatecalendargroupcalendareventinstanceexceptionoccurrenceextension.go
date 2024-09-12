package calendargroupcalendareventinstanceexceptionoccurrenceextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension - Update the navigation property extensions in
// me
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient) UpdateCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension(ctx context.Context, id beta.MeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId, input beta.Extension) (result UpdateCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
