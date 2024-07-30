package drive

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

type ListDrivesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Drive
}

type ListDrivesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Drive
}

type ListDrivesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDrivesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDrives ...
func (c DriveClient) ListDrives(ctx context.Context, id UserId) (result ListDrivesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDrivesCustomPager{},
		Path:       fmt.Sprintf("%s/drives", id.ID()),
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
		Values *[]beta.Drive `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDrivesComplete retrieves all the results into a single object
func (c DriveClient) ListDrivesComplete(ctx context.Context, id UserId) (ListDrivesCompleteResult, error) {
	return c.ListDrivesCompleteMatchingPredicate(ctx, id, DriveOperationPredicate{})
}

// ListDrivesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveClient) ListDrivesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate DriveOperationPredicate) (result ListDrivesCompleteResult, err error) {
	items := make([]beta.Drive, 0)

	resp, err := c.ListDrives(ctx, id)
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

	result = ListDrivesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
