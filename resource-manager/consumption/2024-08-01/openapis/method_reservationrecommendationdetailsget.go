package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ReservationRecommendationDetailsModel
}

type ReservationRecommendationDetailsGetOperationOptions struct {
	Filter         *string
	LookBackPeriod *LookBackPeriod
	Product        *string
	Region         *string
	Scope          *Scope
	Term           *Term
}

func DefaultReservationRecommendationDetailsGetOperationOptions() ReservationRecommendationDetailsGetOperationOptions {
	return ReservationRecommendationDetailsGetOperationOptions{}
}

func (o ReservationRecommendationDetailsGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationRecommendationDetailsGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationRecommendationDetailsGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.LookBackPeriod != nil {
		out.Append("lookBackPeriod", fmt.Sprintf("%v", *o.LookBackPeriod))
	}
	if o.Product != nil {
		out.Append("product", fmt.Sprintf("%v", *o.Product))
	}
	if o.Region != nil {
		out.Append("region", fmt.Sprintf("%v", *o.Region))
	}
	if o.Scope != nil {
		out.Append("scope", fmt.Sprintf("%v", *o.Scope))
	}
	if o.Term != nil {
		out.Append("term", fmt.Sprintf("%v", *o.Term))
	}
	return &out
}

// ReservationRecommendationDetailsGet ...
func (c OpenapisClient) ReservationRecommendationDetailsGet(ctx context.Context, id commonids.ScopeId, options ReservationRecommendationDetailsGetOperationOptions) (result ReservationRecommendationDetailsGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationRecommendationDetails", id.ID()),
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

	var model ReservationRecommendationDetailsModel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
