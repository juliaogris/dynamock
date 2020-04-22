package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ErrUnimpl = fmt.Errorf("not implemented")

// UnimplementedDB implements the dynamodbiface.DynamoDBAPI interface by
// returning an ErrUnimpl for every method call that can return an error.
type UnimplementedDB struct {
}

func (*UnimplementedDB) BatchGetItem(_ *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) BatchGetItemWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) BatchGetItemRequest(_ *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) BatchGetItemPages(_ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return ErrUnimpl
}

func (*UnimplementedDB) BatchGetItemPagesWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (*UnimplementedDB) BatchWriteItem(_ *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) BatchWriteItemWithContext(_ aws.Context, _ *dynamodb.BatchWriteItemInput, _ ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) BatchWriteItemRequest(_ *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) CreateBackup(_ *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateBackupWithContext(_ aws.Context, _ *dynamodb.CreateBackupInput, _ ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateBackupRequest(_ *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	return nil, nil
}

func (*UnimplementedDB) CreateGlobalTable(_ *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateGlobalTableWithContext(_ aws.Context, _ *dynamodb.CreateGlobalTableInput, _ ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateGlobalTableRequest(_ *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) CreateTable(_ *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateTableWithContext(_ aws.Context, _ *dynamodb.CreateTableInput, _ ...request.Option) (*dynamodb.CreateTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) CreateTableRequest(_ *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) DeleteBackup(_ *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteBackupWithContext(_ aws.Context, _ *dynamodb.DeleteBackupInput, _ ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteBackupRequest(_ *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	return nil, nil
}

func (*UnimplementedDB) DeleteItem(_ *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteItemWithContext(_ aws.Context, _ *dynamodb.DeleteItemInput, _ ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteItemRequest(_ *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) DeleteTable(_ *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteTableWithContext(_ aws.Context, _ *dynamodb.DeleteTableInput, _ ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DeleteTableRequest(_ *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeBackup(_ *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeBackupWithContext(_ aws.Context, _ *dynamodb.DescribeBackupInput, _ ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeBackupRequest(_ *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeContinuousBackups(_ *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.DescribeContinuousBackupsInput, _ ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeContinuousBackupsRequest(_ *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeContributorInsights(_ *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeContributorInsightsWithContext(_ aws.Context, _ *dynamodb.DescribeContributorInsightsInput, _ ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeContributorInsightsRequest(_ *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeEndpoints(_ *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeEndpointsWithContext(_ aws.Context, _ *dynamodb.DescribeEndpointsInput, _ ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeEndpointsRequest(_ *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeGlobalTable(_ *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeGlobalTableWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeGlobalTableRequest(_ *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeGlobalTableSettings(_ *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeGlobalTableSettingsRequest(_ *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeLimits(_ *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeLimitsWithContext(_ aws.Context, _ *dynamodb.DescribeLimitsInput, _ ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeLimitsRequest(_ *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeTable(_ *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTableWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTableRequest(_ *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeTableReplicaAutoScaling(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.DescribeTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTableReplicaAutoScalingRequest(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (*UnimplementedDB) DescribeTimeToLive(_ *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTimeToLiveWithContext(_ aws.Context, _ *dynamodb.DescribeTimeToLiveInput, _ ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) DescribeTimeToLiveRequest(_ *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	return nil, nil
}

func (*UnimplementedDB) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) GetItemWithContext(_ aws.Context, _ *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) GetItemRequest(_ *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) ListBackups(_ *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListBackupsWithContext(_ aws.Context, _ *dynamodb.ListBackupsInput, _ ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListBackupsRequest(_ *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	return nil, nil
}

func (*UnimplementedDB) ListContributorInsights(_ *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListContributorInsightsWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListContributorInsightsRequest(_ *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	return nil, nil
}

func (*UnimplementedDB) ListContributorInsightsPages(_ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	return ErrUnimpl
}

func (*UnimplementedDB) ListContributorInsightsPagesWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (*UnimplementedDB) ListGlobalTables(_ *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListGlobalTablesWithContext(_ aws.Context, _ *dynamodb.ListGlobalTablesInput, _ ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListGlobalTablesRequest(_ *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	return nil, nil
}

func (*UnimplementedDB) ListTables(_ *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListTablesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ ...request.Option) (*dynamodb.ListTablesOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListTablesRequest(_ *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return nil, nil
}

func (*UnimplementedDB) ListTablesPages(_ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool) error {
	return ErrUnimpl
}

func (*UnimplementedDB) ListTablesPagesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (*UnimplementedDB) ListTagsOfResource(_ *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListTagsOfResourceWithContext(_ aws.Context, _ *dynamodb.ListTagsOfResourceInput, _ ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ListTagsOfResourceRequest(_ *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	return nil, nil
}

func (*UnimplementedDB) PutItem(_ *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) PutItemWithContext(_ aws.Context, _ *dynamodb.PutItemInput, _ ...request.Option) (*dynamodb.PutItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) PutItemRequest(_ *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) Query(_ *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) QueryWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) QueryRequest(_ *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return nil, nil
}

func (*UnimplementedDB) QueryPages(_ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool) error {
	return ErrUnimpl
}

func (*UnimplementedDB) QueryPagesWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (*UnimplementedDB) RestoreTableFromBackup(_ *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) RestoreTableFromBackupWithContext(_ aws.Context, _ *dynamodb.RestoreTableFromBackupInput, _ ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) RestoreTableFromBackupRequest(_ *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return nil, nil
}

func (*UnimplementedDB) RestoreTableToPointInTime(_ *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) RestoreTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.RestoreTableToPointInTimeInput, _ ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) RestoreTableToPointInTimeRequest(_ *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return nil, nil
}

func (*UnimplementedDB) Scan(_ *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ScanWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ ...request.Option) (*dynamodb.ScanOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) ScanRequest(_ *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return nil, nil
}

func (*UnimplementedDB) ScanPages(_ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool) error {
	return ErrUnimpl
}

func (*UnimplementedDB) ScanPagesWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (*UnimplementedDB) TagResource(_ *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TagResourceWithContext(_ aws.Context, _ *dynamodb.TagResourceInput, _ ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TagResourceRequest(_ *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return nil, nil
}

func (*UnimplementedDB) TransactGetItems(_ *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TransactGetItemsWithContext(_ aws.Context, _ *dynamodb.TransactGetItemsInput, _ ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TransactGetItemsRequest(_ *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	return nil, nil
}

func (*UnimplementedDB) TransactWriteItems(_ *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TransactWriteItemsWithContext(_ aws.Context, _ *dynamodb.TransactWriteItemsInput, _ ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) TransactWriteItemsRequest(_ *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	return nil, nil
}

func (*UnimplementedDB) UntagResource(_ *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UntagResourceWithContext(_ aws.Context, _ *dynamodb.UntagResourceInput, _ ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UntagResourceRequest(_ *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateContinuousBackups(_ *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.UpdateContinuousBackupsInput, _ ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateContinuousBackupsRequest(_ *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateContributorInsights(_ *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateContributorInsightsWithContext(_ aws.Context, _ *dynamodb.UpdateContributorInsightsInput, _ ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateContributorInsightsRequest(_ *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateGlobalTable(_ *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateGlobalTableWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateGlobalTableRequest(_ *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateGlobalTableSettings(_ *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateGlobalTableSettingsRequest(_ *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateItem(_ *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateItemWithContext(_ aws.Context, _ *dynamodb.UpdateItemInput, _ ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateItemRequest(_ *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateTable(_ *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTableWithContext(_ aws.Context, _ *dynamodb.UpdateTableInput, _ ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTableRequest(_ *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateTableReplicaAutoScaling(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.UpdateTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTableReplicaAutoScalingRequest(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (*UnimplementedDB) UpdateTimeToLive(_ *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTimeToLiveWithContext(_ aws.Context, _ *dynamodb.UpdateTimeToLiveInput, _ ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (*UnimplementedDB) UpdateTimeToLiveRequest(_ *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return nil, nil
}

func (*UnimplementedDB) WaitUntilTableExists(_ *dynamodb.DescribeTableInput) error {
	return ErrUnimpl
}

func (*UnimplementedDB) WaitUntilTableExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	return ErrUnimpl
}

func (*UnimplementedDB) WaitUntilTableNotExists(_ *dynamodb.DescribeTableInput) error {
	return ErrUnimpl
}

func (*UnimplementedDB) WaitUntilTableNotExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	return ErrUnimpl
}
