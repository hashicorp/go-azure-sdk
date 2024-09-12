package informationprotectionpolicylabel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInformationProtectionPolicyLabelEvaluateRemovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionAction
}

type ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionAction
}

type ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions() ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions {
	return ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions{}
}

func (o ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInformationProtectionPolicyLabelEvaluateRemovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionPolicyLabelEvaluateRemovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionPolicyLabelEvaluateRemovals - Invoke action evaluateRemoval. Indicate to the consuming
// application what actions it should take to remove the label information. Given contentInfo as an input, which
// includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some
// combination of one of more of the following:
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateRemovals(ctx context.Context, input ListInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) (result ListInformationProtectionPolicyLabelEvaluateRemovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListInformationProtectionPolicyLabelEvaluateRemovalsCustomPager{},
		Path:          "/me/informationProtection/policy/labels/evaluateRemoval",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.InformationProtectionAction, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalInformationProtectionActionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.InformationProtectionAction (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListInformationProtectionPolicyLabelEvaluateRemovalsComplete retrieves all the results into a single object
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateRemovalsComplete(ctx context.Context, input ListInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) (ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult, error) {
	return c.ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate(ctx, input, options, InformationProtectionActionOperationPredicate{})
}

// ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate(ctx context.Context, input ListInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions, predicate InformationProtectionActionOperationPredicate) (result ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult, err error) {
	items := make([]beta.InformationProtectionAction, 0)

	resp, err := c.ListInformationProtectionPolicyLabelEvaluateRemovals(ctx, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
