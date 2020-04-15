// Package dynafake provides a file-based dynamoDB fake called DB.
//
// DB implements the dynamodbiface.DynamoDBAPI interface.
// Use it for testing, local development and execution.
// Dynamodb docs:
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb
package dynafake

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ErrUnimpl = fmt.Errorf("not implemented")

// DB implements the dynamodbiface.DynamoDBAPI interface
type DB struct{}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) BatchGetItem(_ *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) BatchGetItemWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) BatchGetItemRequest(_ *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return nil, nil
}

func (db *DB) BatchGetItemPages(_ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return ErrUnimpl
}

func (db *DB) BatchGetItemPagesWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (db *DB) BatchWriteItem(_ *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) BatchWriteItemWithContext(_ aws.Context, _ *dynamodb.BatchWriteItemInput, _ ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) BatchWriteItemRequest(_ *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return nil, nil
}

func (db *DB) CreateBackup(_ *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateBackupWithContext(_ aws.Context, _ *dynamodb.CreateBackupInput, _ ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateBackupRequest(_ *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	return nil, nil
}

func (db *DB) CreateGlobalTable(_ *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateGlobalTableWithContext(_ aws.Context, _ *dynamodb.CreateGlobalTableInput, _ ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateGlobalTableRequest(_ *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	return nil, nil
}

func (db *DB) CreateTable(_ *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateTableWithContext(_ aws.Context, _ *dynamodb.CreateTableInput, _ ...request.Option) (*dynamodb.CreateTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) CreateTableRequest(_ *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return nil, nil
}

func (db *DB) DeleteBackup(_ *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteBackupWithContext(_ aws.Context, _ *dynamodb.DeleteBackupInput, _ ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteBackupRequest(_ *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	return nil, nil
}

func (db *DB) DeleteItem(_ *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteItemWithContext(_ aws.Context, _ *dynamodb.DeleteItemInput, _ ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteItemRequest(_ *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return nil, nil
}

func (db *DB) DeleteTable(_ *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteTableWithContext(_ aws.Context, _ *dynamodb.DeleteTableInput, _ ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DeleteTableRequest(_ *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return nil, nil
}

func (db *DB) DescribeBackup(_ *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeBackupWithContext(_ aws.Context, _ *dynamodb.DescribeBackupInput, _ ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeBackupRequest(_ *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	return nil, nil
}

func (db *DB) DescribeContinuousBackups(_ *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.DescribeContinuousBackupsInput, _ ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeContinuousBackupsRequest(_ *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	return nil, nil
}

func (db *DB) DescribeContributorInsights(_ *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeContributorInsightsWithContext(_ aws.Context, _ *dynamodb.DescribeContributorInsightsInput, _ ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeContributorInsightsRequest(_ *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	return nil, nil
}

func (db *DB) DescribeEndpoints(_ *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeEndpointsWithContext(_ aws.Context, _ *dynamodb.DescribeEndpointsInput, _ ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeEndpointsRequest(_ *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	return nil, nil
}

func (db *DB) DescribeGlobalTable(_ *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeGlobalTableWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeGlobalTableRequest(_ *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	return nil, nil
}

func (db *DB) DescribeGlobalTableSettings(_ *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeGlobalTableSettingsRequest(_ *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	return nil, nil
}

func (db *DB) DescribeLimits(_ *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeLimitsWithContext(_ aws.Context, _ *dynamodb.DescribeLimitsInput, _ ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeLimitsRequest(_ *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	return nil, nil
}

func (db *DB) DescribeTable(_ *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTableWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTableRequest(_ *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return nil, nil
}

func (db *DB) DescribeTableReplicaAutoScaling(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.DescribeTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTableReplicaAutoScalingRequest(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (db *DB) DescribeTimeToLive(_ *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTimeToLiveWithContext(_ aws.Context, _ *dynamodb.DescribeTimeToLiveInput, _ ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) DescribeTimeToLiveRequest(_ *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	return nil, nil
}

func (db *DB) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) GetItemWithContext(_ aws.Context, _ *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) GetItemRequest(_ *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return nil, nil
}

func (db *DB) ListBackups(_ *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListBackupsWithContext(_ aws.Context, _ *dynamodb.ListBackupsInput, _ ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListBackupsRequest(_ *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	return nil, nil
}

func (db *DB) ListContributorInsights(_ *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListContributorInsightsWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListContributorInsightsRequest(_ *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	return nil, nil
}

func (db *DB) ListContributorInsightsPages(_ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	return ErrUnimpl
}

func (db *DB) ListContributorInsightsPagesWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (db *DB) ListGlobalTables(_ *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListGlobalTablesWithContext(_ aws.Context, _ *dynamodb.ListGlobalTablesInput, _ ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListGlobalTablesRequest(_ *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	return nil, nil
}

func (db *DB) ListTables(_ *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListTablesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ ...request.Option) (*dynamodb.ListTablesOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListTablesRequest(_ *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return nil, nil
}

func (db *DB) ListTablesPages(_ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool) error {
	return ErrUnimpl
}

func (db *DB) ListTablesPagesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (db *DB) ListTagsOfResource(_ *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListTagsOfResourceWithContext(_ aws.Context, _ *dynamodb.ListTagsOfResourceInput, _ ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ListTagsOfResourceRequest(_ *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	return nil, nil
}

func (db *DB) PutItem(_ *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) PutItemWithContext(_ aws.Context, _ *dynamodb.PutItemInput, _ ...request.Option) (*dynamodb.PutItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) PutItemRequest(_ *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return nil, nil
}

func (db *DB) Query(_ *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) QueryWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) QueryRequest(_ *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return nil, nil
}

func (db *DB) QueryPages(_ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool) error {
	return ErrUnimpl
}

func (db *DB) QueryPagesWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (db *DB) RestoreTableFromBackup(_ *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) RestoreTableFromBackupWithContext(_ aws.Context, _ *dynamodb.RestoreTableFromBackupInput, _ ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) RestoreTableFromBackupRequest(_ *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return nil, nil
}

func (db *DB) RestoreTableToPointInTime(_ *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) RestoreTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.RestoreTableToPointInTimeInput, _ ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) RestoreTableToPointInTimeRequest(_ *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return nil, nil
}

func (db *DB) Scan(_ *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ScanWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ ...request.Option) (*dynamodb.ScanOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) ScanRequest(_ *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return nil, nil
}

func (db *DB) ScanPages(_ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool) error {
	return ErrUnimpl
}

func (db *DB) ScanPagesWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool, _ ...request.Option) error {
	return ErrUnimpl
}

func (db *DB) TagResource(_ *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TagResourceWithContext(_ aws.Context, _ *dynamodb.TagResourceInput, _ ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TagResourceRequest(_ *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return nil, nil
}

func (db *DB) TransactGetItems(_ *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TransactGetItemsWithContext(_ aws.Context, _ *dynamodb.TransactGetItemsInput, _ ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TransactGetItemsRequest(_ *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	return nil, nil
}

func (db *DB) TransactWriteItems(_ *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TransactWriteItemsWithContext(_ aws.Context, _ *dynamodb.TransactWriteItemsInput, _ ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) TransactWriteItemsRequest(_ *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	return nil, nil
}

func (db *DB) UntagResource(_ *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UntagResourceWithContext(_ aws.Context, _ *dynamodb.UntagResourceInput, _ ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UntagResourceRequest(_ *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return nil, nil
}

func (db *DB) UpdateContinuousBackups(_ *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.UpdateContinuousBackupsInput, _ ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateContinuousBackupsRequest(_ *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return nil, nil
}

func (db *DB) UpdateContributorInsights(_ *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateContributorInsightsWithContext(_ aws.Context, _ *dynamodb.UpdateContributorInsightsInput, _ ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateContributorInsightsRequest(_ *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	return nil, nil
}

func (db *DB) UpdateGlobalTable(_ *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateGlobalTableWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateGlobalTableRequest(_ *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return nil, nil
}

func (db *DB) UpdateGlobalTableSettings(_ *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateGlobalTableSettingsRequest(_ *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return nil, nil
}

func (db *DB) UpdateItem(_ *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateItemWithContext(_ aws.Context, _ *dynamodb.UpdateItemInput, _ ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateItemRequest(_ *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return nil, nil
}

func (db *DB) UpdateTable(_ *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTableWithContext(_ aws.Context, _ *dynamodb.UpdateTableInput, _ ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTableRequest(_ *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return nil, nil
}

func (db *DB) UpdateTableReplicaAutoScaling(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.UpdateTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTableReplicaAutoScalingRequest(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (db *DB) UpdateTimeToLive(_ *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTimeToLiveWithContext(_ aws.Context, _ *dynamodb.UpdateTimeToLiveInput, _ ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) UpdateTimeToLiveRequest(_ *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return nil, nil
}

func (db *DB) WaitUntilTableExists(_ *dynamodb.DescribeTableInput) error {
	return ErrUnimpl
}

func (db *DB) WaitUntilTableExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	return ErrUnimpl
}

func (db *DB) WaitUntilTableNotExists(_ *dynamodb.DescribeTableInput) error {
	return ErrUnimpl
}

func (db *DB) WaitUntilTableNotExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	return ErrUnimpl
}
