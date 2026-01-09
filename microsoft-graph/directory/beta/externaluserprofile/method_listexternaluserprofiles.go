package externaluserprofile

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

type ListExternalUserProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ExternalUserProfile
}

type ListExternalUserProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ExternalUserProfile
}

type ListExternalUserProfilesOperationOptions struct {
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

func DefaultListExternalUserProfilesOperationOptions() ListExternalUserProfilesOperationOptions {
	return ListExternalUserProfilesOperationOptions{}
}

func (o ListExternalUserProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExternalUserProfilesOperationOptions) ToOData() *odata.Query {
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

func (o ListExternalUserProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExternalUserProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExternalUserProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExternalUserProfiles - List externalUserProfiles. Retrieve the properties of all externalUserProfiles
func (c ExternalUserProfileClient) ListExternalUserProfiles(ctx context.Context, options ListExternalUserProfilesOperationOptions) (result ListExternalUserProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExternalUserProfilesCustomPager{},
		Path:          "/directory/externalUserProfiles",
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
		Values *[]beta.ExternalUserProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExternalUserProfilesComplete retrieves all the results into a single object
func (c ExternalUserProfileClient) ListExternalUserProfilesComplete(ctx context.Context, options ListExternalUserProfilesOperationOptions) (ListExternalUserProfilesCompleteResult, error) {
	return c.ListExternalUserProfilesCompleteMatchingPredicate(ctx, options, ExternalUserProfileOperationPredicate{})
}

// ListExternalUserProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExternalUserProfileClient) ListExternalUserProfilesCompleteMatchingPredicate(ctx context.Context, options ListExternalUserProfilesOperationOptions, predicate ExternalUserProfileOperationPredicate) (result ListExternalUserProfilesCompleteResult, err error) {
	items := make([]beta.ExternalUserProfile, 0)

	resp, err := c.ListExternalUserProfiles(ctx, options)
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

	result = ListExternalUserProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
