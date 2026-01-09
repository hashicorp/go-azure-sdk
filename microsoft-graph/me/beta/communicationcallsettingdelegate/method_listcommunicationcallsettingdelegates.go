package communicationcallsettingdelegate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCommunicationCallSettingDelegatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DelegationSettings
}

type ListCommunicationCallSettingDelegatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DelegationSettings
}

type ListCommunicationCallSettingDelegatesOperationOptions struct {
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

func DefaultListCommunicationCallSettingDelegatesOperationOptions() ListCommunicationCallSettingDelegatesOperationOptions {
	return ListCommunicationCallSettingDelegatesOperationOptions{}
}

func (o ListCommunicationCallSettingDelegatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCommunicationCallSettingDelegatesOperationOptions) ToOData() *odata.Query {
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

func (o ListCommunicationCallSettingDelegatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCommunicationCallSettingDelegatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCommunicationCallSettingDelegatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCommunicationCallSettingDelegates - List delegates. Get a list of all delegates for a user.
func (c CommunicationCallSettingDelegateClient) ListCommunicationCallSettingDelegates(ctx context.Context, options ListCommunicationCallSettingDelegatesOperationOptions) (result ListCommunicationCallSettingDelegatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCommunicationCallSettingDelegatesCustomPager{},
		Path:          "/me/communications/callSettings/delegates",
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
		Values *[]beta.DelegationSettings `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCommunicationCallSettingDelegatesComplete retrieves all the results into a single object
func (c CommunicationCallSettingDelegateClient) ListCommunicationCallSettingDelegatesComplete(ctx context.Context, options ListCommunicationCallSettingDelegatesOperationOptions) (ListCommunicationCallSettingDelegatesCompleteResult, error) {
	return c.ListCommunicationCallSettingDelegatesCompleteMatchingPredicate(ctx, options, DelegationSettingsOperationPredicate{})
}

// ListCommunicationCallSettingDelegatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CommunicationCallSettingDelegateClient) ListCommunicationCallSettingDelegatesCompleteMatchingPredicate(ctx context.Context, options ListCommunicationCallSettingDelegatesOperationOptions, predicate DelegationSettingsOperationPredicate) (result ListCommunicationCallSettingDelegatesCompleteResult, err error) {
	items := make([]beta.DelegationSettings, 0)

	resp, err := c.ListCommunicationCallSettingDelegates(ctx, options)
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

	result = ListCommunicationCallSettingDelegatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
