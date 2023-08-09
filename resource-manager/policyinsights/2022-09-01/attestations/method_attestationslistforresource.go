package attestations

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

type AttestationsListForResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Attestation
}

type AttestationsListForResourceCompleteResult struct {
	Items []Attestation
}

type AttestationsListForResourceOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultAttestationsListForResourceOperationOptions() AttestationsListForResourceOperationOptions {
	return AttestationsListForResourceOperationOptions{}
}

func (o AttestationsListForResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AttestationsListForResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AttestationsListForResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// AttestationsListForResource ...
func (c AttestationsClient) AttestationsListForResource(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions) (result AttestationsListForResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/attestations", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Attestation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AttestationsListForResourceComplete retrieves all the results into a single object
func (c AttestationsClient) AttestationsListForResourceComplete(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions) (AttestationsListForResourceCompleteResult, error) {
	return c.AttestationsListForResourceCompleteMatchingPredicate(ctx, id, options, AttestationOperationPredicate{})
}

// AttestationsListForResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AttestationsClient) AttestationsListForResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions, predicate AttestationOperationPredicate) (result AttestationsListForResourceCompleteResult, err error) {
	items := make([]Attestation, 0)

	resp, err := c.AttestationsListForResource(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = AttestationsListForResourceCompleteResult{
		Items: items,
	}
	return
}
