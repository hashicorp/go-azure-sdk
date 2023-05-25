package syncmembers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySyncGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncMember
}

type ListBySyncGroupCompleteResult struct {
	Items []SyncMember
}

// ListBySyncGroup ...
func (c SyncMembersClient) ListBySyncGroup(ctx context.Context, id SyncGroupId) (result ListBySyncGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/syncMembers", id.ID()),
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
		Values *[]SyncMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySyncGroupComplete retrieves all the results into a single object
func (c SyncMembersClient) ListBySyncGroupComplete(ctx context.Context, id SyncGroupId) (ListBySyncGroupCompleteResult, error) {
	return c.ListBySyncGroupCompleteMatchingPredicate(ctx, id, SyncMemberOperationPredicate{})
}

// ListBySyncGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncMembersClient) ListBySyncGroupCompleteMatchingPredicate(ctx context.Context, id SyncGroupId, predicate SyncMemberOperationPredicate) (result ListBySyncGroupCompleteResult, err error) {
	items := make([]SyncMember, 0)

	resp, err := c.ListBySyncGroup(ctx, id)
	if err != nil {
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

	result = ListBySyncGroupCompleteResult{
		Items: items,
	}
	return
}
