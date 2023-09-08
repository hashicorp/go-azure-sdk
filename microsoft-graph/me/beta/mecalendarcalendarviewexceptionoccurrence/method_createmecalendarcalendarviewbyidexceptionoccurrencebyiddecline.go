package mecalendarcalendarviewexceptionoccurrence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDecline ...
func (c MeCalendarCalendarViewExceptionOccurrenceClient) CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDecline(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceId, input CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineRequest) (result CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/decline", id.ID()),
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
