package siteinformationprotectionpolicylabel

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

type ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionAction
}

type ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionAction
}

type ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions() ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions {
	return ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions{}
}

func (o ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionPolicyLabelEvaluateRemovals - Invoke action evaluateRemoval. Indicate to the consuming
// application what actions it should take to remove the label information. Given contentInfo as an input, which
// includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some
// combination of one of more of the following:
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabelEvaluateRemovals(ctx context.Context, id beta.GroupIdSiteId, input ListSiteInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) (result ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCustomPager{},
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels/evaluateRemoval", id.ID()),
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

// ListSiteInformationProtectionPolicyLabelEvaluateRemovalsComplete retrieves all the results into a single object
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabelEvaluateRemovalsComplete(ctx context.Context, id beta.GroupIdSiteId, input ListSiteInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions) (ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult, error) {
	return c.ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate(ctx, id, input, options, InformationProtectionActionOperationPredicate{})
}

// ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, input ListSiteInformationProtectionPolicyLabelEvaluateRemovalsRequest, options ListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions, predicate InformationProtectionActionOperationPredicate) (result ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult, err error) {
	items := make([]beta.InformationProtectionAction, 0)

	resp, err := c.ListSiteInformationProtectionPolicyLabelEvaluateRemovals(ctx, id, input, options)
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

	result = ListSiteInformationProtectionPolicyLabelEvaluateRemovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
