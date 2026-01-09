package claimsmappingpolicyappliesto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListClaimsMappingPolicyAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListClaimsMappingPolicyAppliesTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListClaimsMappingPolicyAppliesTosOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListClaimsMappingPolicyAppliesTosOperationOptions() ListClaimsMappingPolicyAppliesTosOperationOptions {
	return ListClaimsMappingPolicyAppliesTosOperationOptions{}
}

func (o ListClaimsMappingPolicyAppliesTosOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListClaimsMappingPolicyAppliesTosOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
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
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListClaimsMappingPolicyAppliesTosOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListClaimsMappingPolicyAppliesTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListClaimsMappingPolicyAppliesTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListClaimsMappingPolicyAppliesTos - List appliesTo. Get a list of directoryObject objects that a claimsMappingPolicy
// object has been applied to. The claimsMappingPolicy can only be applied to application and servicePrincipal
// resources.
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTos(ctx context.Context, id stable.PolicyClaimsMappingPolicyId, options ListClaimsMappingPolicyAppliesTosOperationOptions) (result ListClaimsMappingPolicyAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListClaimsMappingPolicyAppliesTosCustomPager{},
		Path:          fmt.Sprintf("%s/appliesTo", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListClaimsMappingPolicyAppliesTosComplete retrieves all the results into a single object
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTosComplete(ctx context.Context, id stable.PolicyClaimsMappingPolicyId, options ListClaimsMappingPolicyAppliesTosOperationOptions) (ListClaimsMappingPolicyAppliesTosCompleteResult, error) {
	return c.ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate(ctx context.Context, id stable.PolicyClaimsMappingPolicyId, options ListClaimsMappingPolicyAppliesTosOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListClaimsMappingPolicyAppliesTosCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListClaimsMappingPolicyAppliesTos(ctx, id, options)
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

	result = ListClaimsMappingPolicyAppliesTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
