package dynamock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireErrUnimpl(t *testing.T, err error) {
	t.Helper()
	requireErrIs(t, err, ErrUnimpl)
}

func TestUnimplmented(t *testing.T) { //nolint:funlen
	db := UnimplementedDB{}
	ctx := context.Background()
	var r, o interface{}

	_, err := db.BatchGetItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.BatchGetItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	_, err = db.BatchGetItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.BatchGetItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.BatchGetItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.BatchGetItemPages(nil, nil)
	requireErrUnimpl(t, err)

	err = db.BatchGetItemPagesWithContext(ctx, nil, nil)
	requireErrUnimpl(t, err)

	_, err = db.BatchWriteItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.BatchWriteItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.BatchWriteItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.CreateBackup(nil)
	requireErrUnimpl(t, err)

	_, err = db.CreateBackupWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.CreateBackupRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.CreateGlobalTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.CreateGlobalTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)
	r, o = db.CreateGlobalTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.CreateTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.CreateTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.CreateTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DeleteBackup(nil)
	requireErrUnimpl(t, err)

	_, err = db.DeleteBackupWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DeleteBackupRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DeleteItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.DeleteItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DeleteItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DeleteTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.DeleteTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DeleteTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeBackup(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeBackupWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeBackupRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeContinuousBackups(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeContinuousBackupsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeContinuousBackupsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeContributorInsights(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeContributorInsightsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeContributorInsightsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeEndpoints(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeEndpointsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeEndpointsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeGlobalTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeGlobalTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeGlobalTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeGlobalTableSettings(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeGlobalTableSettingsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeGlobalTableSettingsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeLimits(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeLimitsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeLimitsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeTableReplicaAutoScaling(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeTableReplicaAutoScalingWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeTableReplicaAutoScalingRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.DescribeTimeToLive(nil)
	requireErrUnimpl(t, err)

	_, err = db.DescribeTimeToLiveWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.DescribeTimeToLiveRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.GetItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.GetItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.GetItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.ListBackups(nil)
	requireErrUnimpl(t, err)

	_, err = db.ListBackupsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ListBackupsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.ListContributorInsights(nil)
	requireErrUnimpl(t, err)

	_, err = db.ListContributorInsightsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ListContributorInsightsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.ListContributorInsightsPages(nil, nil)
	requireErrUnimpl(t, err)

	err = db.ListContributorInsightsPagesWithContext(ctx, nil, nil)
	requireErrUnimpl(t, err)

	_, err = db.ListGlobalTables(nil)
	requireErrUnimpl(t, err)

	_, err = db.ListGlobalTablesWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ListGlobalTablesRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.ListTables(nil)
	requireErrUnimpl(t, err)

	_, err = db.ListTablesWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ListTablesRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.ListTablesPages(nil, nil)
	requireErrUnimpl(t, err)

	err = db.ListTablesPagesWithContext(ctx, nil, nil)
	requireErrUnimpl(t, err)

	_, err = db.ListTagsOfResource(nil)
	requireErrUnimpl(t, err)

	_, err = db.ListTagsOfResourceWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ListTagsOfResourceRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.PutItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.PutItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.PutItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.Query(nil)
	requireErrUnimpl(t, err)

	_, err = db.QueryWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.QueryRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.QueryPages(nil, nil)
	requireErrUnimpl(t, err)

	err = db.QueryPagesWithContext(ctx, nil, nil)
	requireErrUnimpl(t, err)

	_, err = db.RestoreTableFromBackup(nil)
	requireErrUnimpl(t, err)

	_, err = db.RestoreTableFromBackupWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.RestoreTableFromBackupRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.RestoreTableToPointInTime(nil)
	requireErrUnimpl(t, err)

	_, err = db.RestoreTableToPointInTimeWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.RestoreTableToPointInTimeRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.Scan(nil)
	requireErrUnimpl(t, err)

	_, err = db.ScanWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.ScanRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.ScanPages(nil, nil)
	requireErrUnimpl(t, err)

	err = db.ScanPagesWithContext(ctx, nil, nil)
	requireErrUnimpl(t, err)

	_, err = db.TagResource(nil)
	requireErrUnimpl(t, err)

	_, err = db.TagResourceWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.TagResourceRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.TransactGetItems(nil)
	requireErrUnimpl(t, err)

	_, err = db.TransactGetItemsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.TransactGetItemsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.TransactWriteItems(nil)
	requireErrUnimpl(t, err)

	_, err = db.TransactWriteItemsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.TransactWriteItemsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UntagResource(nil)
	requireErrUnimpl(t, err)

	_, err = db.UntagResourceWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UntagResourceRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateContinuousBackups(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateContinuousBackupsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateContinuousBackupsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateContributorInsights(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateContributorInsightsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateContributorInsightsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateGlobalTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateGlobalTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateGlobalTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateGlobalTableSettings(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateGlobalTableSettingsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateGlobalTableSettingsRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateItemRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateTable(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateTableWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateTableRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateTableReplicaAutoScaling(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateTableReplicaAutoScalingWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateTableReplicaAutoScalingRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	_, err = db.UpdateTimeToLive(nil)
	requireErrUnimpl(t, err)

	_, err = db.UpdateTimeToLiveWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	r, o = db.UpdateTimeToLiveRequest(nil)
	require.Nil(t, r)
	require.Nil(t, o)

	err = db.WaitUntilTableExists(nil)
	requireErrUnimpl(t, err)

	err = db.WaitUntilTableExistsWithContext(ctx, nil)
	requireErrUnimpl(t, err)

	err = db.WaitUntilTableNotExists(nil)
	requireErrUnimpl(t, err)

	err = db.WaitUntilTableNotExistsWithContext(ctx, nil)
	requireErrUnimpl(t, err)
}
