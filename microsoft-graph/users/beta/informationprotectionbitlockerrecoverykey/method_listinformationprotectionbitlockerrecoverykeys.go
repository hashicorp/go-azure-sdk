package informationprotectionbitlockerrecoverykey

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

type ListInformationProtectionBitlockerRecoveryKeysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.BitlockerRecoveryKey
}

type ListInformationProtectionBitlockerRecoveryKeysCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.BitlockerRecoveryKey
}

type ListInformationProtectionBitlockerRecoveryKeysCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionBitlockerRecoveryKeysCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionBitlockerRecoveryKeys ...
func (c InformationProtectionBitlockerRecoveryKeyClient) ListInformationProtectionBitlockerRecoveryKeys(ctx context.Context, id UserId) (result ListInformationProtectionBitlockerRecoveryKeysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInformationProtectionBitlockerRecoveryKeysCustomPager{},
		Path:       fmt.Sprintf("%s/informationProtection/bitlocker/recoveryKeys", id.ID()),
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
		Values *[]beta.BitlockerRecoveryKey `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInformationProtectionBitlockerRecoveryKeysComplete retrieves all the results into a single object
func (c InformationProtectionBitlockerRecoveryKeyClient) ListInformationProtectionBitlockerRecoveryKeysComplete(ctx context.Context, id UserId) (ListInformationProtectionBitlockerRecoveryKeysCompleteResult, error) {
	return c.ListInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate(ctx, id, BitlockerRecoveryKeyOperationPredicate{})
}

// ListInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionBitlockerRecoveryKeyClient) ListInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate(ctx context.Context, id UserId, predicate BitlockerRecoveryKeyOperationPredicate) (result ListInformationProtectionBitlockerRecoveryKeysCompleteResult, err error) {
	items := make([]beta.BitlockerRecoveryKey, 0)

	resp, err := c.ListInformationProtectionBitlockerRecoveryKeys(ctx, id)
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

	result = ListInformationProtectionBitlockerRecoveryKeysCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
