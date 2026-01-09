package communicationcallsettingdelegator

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

type ListCommunicationCallSettingDelegatorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DelegationSettings
}

type ListCommunicationCallSettingDelegatorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DelegationSettings
}

type ListCommunicationCallSettingDelegatorsOperationOptions struct {
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

func DefaultListCommunicationCallSettingDelegatorsOperationOptions() ListCommunicationCallSettingDelegatorsOperationOptions {
	return ListCommunicationCallSettingDelegatorsOperationOptions{}
}

func (o ListCommunicationCallSettingDelegatorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCommunicationCallSettingDelegatorsOperationOptions) ToOData() *odata.Query {
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

func (o ListCommunicationCallSettingDelegatorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCommunicationCallSettingDelegatorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCommunicationCallSettingDelegatorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCommunicationCallSettingDelegators - Get delegators from users. Represents the delegator settings.
func (c CommunicationCallSettingDelegatorClient) ListCommunicationCallSettingDelegators(ctx context.Context, id beta.UserId, options ListCommunicationCallSettingDelegatorsOperationOptions) (result ListCommunicationCallSettingDelegatorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCommunicationCallSettingDelegatorsCustomPager{},
		Path:          fmt.Sprintf("%s/communications/callSettings/delegators", id.ID()),
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

// ListCommunicationCallSettingDelegatorsComplete retrieves all the results into a single object
func (c CommunicationCallSettingDelegatorClient) ListCommunicationCallSettingDelegatorsComplete(ctx context.Context, id beta.UserId, options ListCommunicationCallSettingDelegatorsOperationOptions) (ListCommunicationCallSettingDelegatorsCompleteResult, error) {
	return c.ListCommunicationCallSettingDelegatorsCompleteMatchingPredicate(ctx, id, options, DelegationSettingsOperationPredicate{})
}

// ListCommunicationCallSettingDelegatorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CommunicationCallSettingDelegatorClient) ListCommunicationCallSettingDelegatorsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListCommunicationCallSettingDelegatorsOperationOptions, predicate DelegationSettingsOperationPredicate) (result ListCommunicationCallSettingDelegatorsCompleteResult, err error) {
	items := make([]beta.DelegationSettings, 0)

	resp, err := c.ListCommunicationCallSettingDelegators(ctx, id, options)
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

	result = ListCommunicationCallSettingDelegatorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
