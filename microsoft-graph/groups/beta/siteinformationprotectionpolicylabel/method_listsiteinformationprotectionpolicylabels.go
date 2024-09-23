package siteinformationprotectionpolicylabel

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

type ListSiteInformationProtectionPolicyLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionLabel
}

type ListSiteInformationProtectionPolicyLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionLabel
}

type ListSiteInformationProtectionPolicyLabelsOperationOptions struct {
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

func DefaultListSiteInformationProtectionPolicyLabelsOperationOptions() ListSiteInformationProtectionPolicyLabelsOperationOptions {
	return ListSiteInformationProtectionPolicyLabelsOperationOptions{}
}

func (o ListSiteInformationProtectionPolicyLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteInformationProtectionPolicyLabelsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteInformationProtectionPolicyLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteInformationProtectionPolicyLabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionPolicyLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionPolicyLabels - Get labels from groups
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabels(ctx context.Context, id beta.GroupIdSiteId, options ListSiteInformationProtectionPolicyLabelsOperationOptions) (result ListSiteInformationProtectionPolicyLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteInformationProtectionPolicyLabelsCustomPager{},
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels", id.ID()),
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
		Values *[]beta.InformationProtectionLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteInformationProtectionPolicyLabelsComplete retrieves all the results into a single object
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabelsComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSiteInformationProtectionPolicyLabelsOperationOptions) (ListSiteInformationProtectionPolicyLabelsCompleteResult, error) {
	return c.ListSiteInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx, id, options, InformationProtectionLabelOperationPredicate{})
}

// ListSiteInformationProtectionPolicyLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionPolicyLabelClient) ListSiteInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSiteInformationProtectionPolicyLabelsOperationOptions, predicate InformationProtectionLabelOperationPredicate) (result ListSiteInformationProtectionPolicyLabelsCompleteResult, err error) {
	items := make([]beta.InformationProtectionLabel, 0)

	resp, err := c.ListSiteInformationProtectionPolicyLabels(ctx, id, options)
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

	result = ListSiteInformationProtectionPolicyLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
