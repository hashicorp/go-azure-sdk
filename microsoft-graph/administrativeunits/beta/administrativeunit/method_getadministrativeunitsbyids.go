package administrativeunit

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

type GetAdministrativeUnitsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type GetAdministrativeUnitsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type GetAdministrativeUnitsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetAdministrativeUnitsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAdministrativeUnitsByIds ...
func (c AdministrativeUnitClient) GetAdministrativeUnitsByIds(ctx context.Context, input GetAdministrativeUnitsByIdsRequest) (result GetAdministrativeUnitsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetAdministrativeUnitsByIdsCustomPager{},
		Path:       "/administrativeUnits/getByIds",
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAdministrativeUnitsByIdsComplete retrieves all the results into a single object
func (c AdministrativeUnitClient) GetAdministrativeUnitsByIdsComplete(ctx context.Context, input GetAdministrativeUnitsByIdsRequest) (GetAdministrativeUnitsByIdsCompleteResult, error) {
	return c.GetAdministrativeUnitsByIdsCompleteMatchingPredicate(ctx, input, DirectoryObjectOperationPredicate{})
}

// GetAdministrativeUnitsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitClient) GetAdministrativeUnitsByIdsCompleteMatchingPredicate(ctx context.Context, input GetAdministrativeUnitsByIdsRequest, predicate DirectoryObjectOperationPredicate) (result GetAdministrativeUnitsByIdsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.GetAdministrativeUnitsByIds(ctx, input)
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

	result = GetAdministrativeUnitsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
