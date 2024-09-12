package outlooktask

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOutlookTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateOutlookTask - Update outlooktask (deprecated). Change writable properties of an Outlook task. The
// completedDateTime property can be set by the complete action, or explicitly by a PATCH operation. If you use PATCH to
// set completedDateTime, make sure you set status to completed as well. By default, this operation (and the POST, GET,
// and complete task operations) returns date-related properties in UTC. You can use the Prefer: outlook.timezone header
// to have all the date-related properties in the response represented in a time zone different than UTC.
func (c OutlookTaskClient) UpdateOutlookTask(ctx context.Context, id beta.MeOutlookTaskId, input beta.OutlookTask) (result UpdateOutlookTaskOperationResponse, err error) {
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
