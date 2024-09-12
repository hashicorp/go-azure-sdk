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

type ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionAction
}

type ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionAction
}

type ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions() ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions {
	return ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions{}
}

func (o ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInformationProtectionPolicyLabelEvaluateClassificationResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionPolicyLabelEvaluateClassificationResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionPolicyLabelEvaluateClassificationResults - Invoke action evaluateClassificationResults.
// Using classification results, compute the information protection label that should be applied and return the set of
// actions that must be taken to correctly label the information. This API is useful when a label should be set
// automatically based on classification of the file contents, rather than labeled directly by a user or service. To
// evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value
// pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the
// following:
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateClassificationResults(ctx context.Context, id beta.UserId, input ListInformationProtectionPolicyLabelEvaluateClassificationResultsRequest, options ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions) (result ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListInformationProtectionPolicyLabelEvaluateClassificationResultsCustomPager{},
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels/evaluateClassificationResults", id.ID()),
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

// ListInformationProtectionPolicyLabelEvaluateClassificationResultsComplete retrieves all the results into a single object
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateClassificationResultsComplete(ctx context.Context, id beta.UserId, input ListInformationProtectionPolicyLabelEvaluateClassificationResultsRequest, options ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions) (ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteResult, error) {
	return c.ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteMatchingPredicate(ctx, id, input, options, InformationProtectionActionOperationPredicate{})
}

// ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, input ListInformationProtectionPolicyLabelEvaluateClassificationResultsRequest, options ListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions, predicate InformationProtectionActionOperationPredicate) (result ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteResult, err error) {
	items := make([]beta.InformationProtectionAction, 0)

	resp, err := c.ListInformationProtectionPolicyLabelEvaluateClassificationResults(ctx, id, input, options)
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

	result = ListInformationProtectionPolicyLabelEvaluateClassificationResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
