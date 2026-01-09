package joinedteamscheduleoffershiftrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamScheduleOfferShiftRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamScheduleOfferShiftRequestOperationOptions() UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions {
	return UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions{}
}

func (o UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamScheduleOfferShiftRequest - Update the navigation property offerShiftRequests in users
func (c JoinedTeamScheduleOfferShiftRequestClient) UpdateJoinedTeamScheduleOfferShiftRequest(ctx context.Context, id stable.UserIdJoinedTeamIdScheduleOfferShiftRequestId, input stable.OfferShiftRequest, options UpdateJoinedTeamScheduleOfferShiftRequestOperationOptions) (result UpdateJoinedTeamScheduleOfferShiftRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
