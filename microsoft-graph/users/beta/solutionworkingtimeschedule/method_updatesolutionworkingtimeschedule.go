package solutionworkingtimeschedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSolutionWorkingTimeScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSolutionWorkingTimeScheduleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSolutionWorkingTimeScheduleOperationOptions() UpdateSolutionWorkingTimeScheduleOperationOptions {
	return UpdateSolutionWorkingTimeScheduleOperationOptions{}
}

func (o UpdateSolutionWorkingTimeScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSolutionWorkingTimeScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSolutionWorkingTimeScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSolutionWorkingTimeSchedule - Update the navigation property workingTimeSchedule in users
func (c SolutionWorkingTimeScheduleClient) UpdateSolutionWorkingTimeSchedule(ctx context.Context, id beta.UserId, input beta.WorkingTimeSchedule, options UpdateSolutionWorkingTimeScheduleOperationOptions) (result UpdateSolutionWorkingTimeScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/solutions/workingTimeSchedule", id.ID()),
		RetryFunc:     options.RetryFunc,
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
