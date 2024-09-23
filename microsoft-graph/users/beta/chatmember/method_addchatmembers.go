package chatmember

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

type AddChatMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type AddChatMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type AddChatMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAddChatMembersOperationOptions() AddChatMembersOperationOptions {
	return AddChatMembersOperationOptions{}
}

func (o AddChatMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddChatMembersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o AddChatMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddChatMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddChatMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddChatMembers - Invoke action add. Add multiple members in a single request to a team. The response provides details
// about which memberships could and couldn't be created.
func (c ChatMemberClient) AddChatMembers(ctx context.Context, id beta.UserIdChatId, input AddChatMembersRequest, options AddChatMembersOperationOptions) (result AddChatMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddChatMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members/add", id.ID()),
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

	temp := make([]beta.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// AddChatMembersComplete retrieves all the results into a single object
func (c ChatMemberClient) AddChatMembersComplete(ctx context.Context, id beta.UserIdChatId, input AddChatMembersRequest, options AddChatMembersOperationOptions) (AddChatMembersCompleteResult, error) {
	return c.AddChatMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// AddChatMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatMemberClient) AddChatMembersCompleteMatchingPredicate(ctx context.Context, id beta.UserIdChatId, input AddChatMembersRequest, options AddChatMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result AddChatMembersCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.AddChatMembers(ctx, id, input, options)
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

	result = AddChatMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
