package solutionworkingtimeschedule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndSolutionWorkingTimeScheduleWorkingTimeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultEndSolutionWorkingTimeScheduleWorkingTimeOperationOptions() EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions {
	return EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions{}
}

func (o EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EndSolutionWorkingTimeScheduleWorkingTime - Invoke action endWorkingTime. Triggers the policies associated with the
// end of working hours.
func (c SolutionWorkingTimeScheduleClient) EndSolutionWorkingTimeScheduleWorkingTime(ctx context.Context, options EndSolutionWorkingTimeScheduleWorkingTimeOperationOptions) (result EndSolutionWorkingTimeScheduleWorkingTimeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/solutions/workingTimeSchedule/endWorkingTime",
		RetryFunc:     options.RetryFunc,
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
