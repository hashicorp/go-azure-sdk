package featurerolloutpolicyappliesto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListFeatureRolloutPolicyAppliesToRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListFeatureRolloutPolicyAppliesToRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListFeatureRolloutPolicyAppliesToRefsOperationOptions struct {
	Count    *bool
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Skip     *int64
	Top      *int64
}

func DefaultListFeatureRolloutPolicyAppliesToRefsOperationOptions() ListFeatureRolloutPolicyAppliesToRefsOperationOptions {
	return ListFeatureRolloutPolicyAppliesToRefsOperationOptions{}
}

func (o ListFeatureRolloutPolicyAppliesToRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListFeatureRolloutPolicyAppliesToRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListFeatureRolloutPolicyAppliesToRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListFeatureRolloutPolicyAppliesToRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFeatureRolloutPolicyAppliesToRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFeatureRolloutPolicyAppliesToRefs - Get ref of appliesTo from policies. Nullable. Specifies a list of
// directoryObject resources that feature is enabled for.
func (c FeatureRolloutPolicyAppliesToClient) ListFeatureRolloutPolicyAppliesToRefs(ctx context.Context, id stable.PolicyFeatureRolloutPolicyId, options ListFeatureRolloutPolicyAppliesToRefsOperationOptions) (result ListFeatureRolloutPolicyAppliesToRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListFeatureRolloutPolicyAppliesToRefsCustomPager{},
		Path:          fmt.Sprintf("%s/appliesTo/$ref", id.ID()),
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListFeatureRolloutPolicyAppliesToRefsComplete retrieves all the results into a single object
func (c FeatureRolloutPolicyAppliesToClient) ListFeatureRolloutPolicyAppliesToRefsComplete(ctx context.Context, id stable.PolicyFeatureRolloutPolicyId, options ListFeatureRolloutPolicyAppliesToRefsOperationOptions) (ListFeatureRolloutPolicyAppliesToRefsCompleteResult, error) {
	return c.ListFeatureRolloutPolicyAppliesToRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListFeatureRolloutPolicyAppliesToRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FeatureRolloutPolicyAppliesToClient) ListFeatureRolloutPolicyAppliesToRefsCompleteMatchingPredicate(ctx context.Context, id stable.PolicyFeatureRolloutPolicyId, options ListFeatureRolloutPolicyAppliesToRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListFeatureRolloutPolicyAppliesToRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListFeatureRolloutPolicyAppliesToRefs(ctx, id, options)
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

	result = ListFeatureRolloutPolicyAppliesToRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
