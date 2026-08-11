package subvolumeinfos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubvolumesListByVolumeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubvolumeInfo
}

type SubvolumesListByVolumeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubvolumeInfo
}

type SubvolumesListByVolumeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SubvolumesListByVolumeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SubvolumesListByVolume ...
func (c SubvolumeInfosClient) SubvolumesListByVolume(ctx context.Context, id VolumeId) (result SubvolumesListByVolumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SubvolumesListByVolumeCustomPager{},
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

// SubvolumesListByVolumeComplete retrieves all the results into a single object
func (c SubvolumeInfosClient) SubvolumesListByVolumeComplete(ctx context.Context, id VolumeId) (SubvolumesListByVolumeCompleteResult, error) {
	return c.SubvolumesListByVolumeCompleteMatchingPredicate(ctx, id, SubvolumeInfoOperationPredicate{})
}

// SubvolumesListByVolumeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubvolumeInfosClient) SubvolumesListByVolumeCompleteMatchingPredicate(ctx context.Context, id VolumeId, predicate SubvolumeInfoOperationPredicate) (result SubvolumesListByVolumeCompleteResult, err error) {
	items := make([]SubvolumeInfo, 0)

	resp, err := c.SubvolumesListByVolume(ctx, id)
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

	result = SubvolumesListByVolumeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
