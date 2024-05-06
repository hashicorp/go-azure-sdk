package subvolumes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVolumeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubvolumeInfo
}

type ListByVolumeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubvolumeInfo
}

// ListByVolume ...
func (c SubVolumesClient) ListByVolume(ctx context.Context, id VolumeId) (result ListByVolumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/subVolumes", id.ID()),
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
		Values *[]SubvolumeInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByVolumeComplete retrieves all the results into a single object
func (c SubVolumesClient) ListByVolumeComplete(ctx context.Context, id VolumeId) (ListByVolumeCompleteResult, error) {
	return c.ListByVolumeCompleteMatchingPredicate(ctx, id, SubvolumeInfoOperationPredicate{})
}

// ListByVolumeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubVolumesClient) ListByVolumeCompleteMatchingPredicate(ctx context.Context, id VolumeId, predicate SubvolumeInfoOperationPredicate) (result ListByVolumeCompleteResult, err error) {
	items := make([]SubvolumeInfo, 0)

	resp, err := c.ListByVolume(ctx, id)
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

	result = ListByVolumeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
