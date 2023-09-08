package mecalendarcalendarviewexceptionoccurrenceattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount ...
func (c MeCalendarCalendarViewExceptionOccurrenceAttachmentClient) GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceId) (result GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attachments/$count", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
