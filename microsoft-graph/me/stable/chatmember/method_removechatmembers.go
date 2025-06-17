package chatmember

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

type RemoveChatMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ActionResultPart
}

type RemoveChatMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ActionResultPart
}

type RemoveChatMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveChatMembersOperationOptions() RemoveChatMembersOperationOptions {
	return RemoveChatMembersOperationOptions{}
}

func (o RemoveChatMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveChatMembersOperationOptions) ToOData() *odata.Query {
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

func (o RemoveChatMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveChatMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveChatMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveChatMembers - Invoke action remove. Remove multiple members from a team in a single request. The response
// provides details about which memberships could and couldn't be removed.
func (c ChatMemberClient) RemoveChatMembers(ctx context.Context, id stable.MeChatId, input RemoveChatMembersRequest, options RemoveChatMembersOperationOptions) (result RemoveChatMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveChatMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members/remove", id.ID()),
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

	temp := make([]stable.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// RemoveChatMembersComplete retrieves all the results into a single object
func (c ChatMemberClient) RemoveChatMembersComplete(ctx context.Context, id stable.MeChatId, input RemoveChatMembersRequest, options RemoveChatMembersOperationOptions) (RemoveChatMembersCompleteResult, error) {
	return c.RemoveChatMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// RemoveChatMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatMemberClient) RemoveChatMembersCompleteMatchingPredicate(ctx context.Context, id stable.MeChatId, input RemoveChatMembersRequest, options RemoveChatMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result RemoveChatMembersCompleteResult, err error) {
	items := make([]stable.ActionResultPart, 0)

	resp, err := c.RemoveChatMembers(ctx, id, input, options)
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

	result = RemoveChatMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
