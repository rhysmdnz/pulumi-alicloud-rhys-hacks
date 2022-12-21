// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `polardb.getClusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
// Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.
//
// > **NOTE:** Available in v1.66.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/polardb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/polardb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			polardbClustersDs, err := polardb.GetClusters(ctx, &polardb.GetClustersArgs{
//				DescriptionRegex: pulumi.StringRef("pc-\\w+"),
//				Status:           pulumi.StringRef("Running"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstPolardbClusterId", polardbClustersDs.Clusters[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetClustersResult
	err := ctx.Invoke("alicloud:polardb/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
	DbType *string `pulumi:"dbType"`
	// A regex string to filter results by cluster description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// A list of PolarDB cluster IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// status of the cluster.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// A list of PolarDB clusters. Each element contains the following attributes:
	Clusters []GetClustersCluster `pulumi:"clusters"`
	// `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
	DbType           *string `pulumi:"dbType"`
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// A list of RDS cluster descriptions.
	Descriptions []string `pulumi:"descriptions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of RDS cluster IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// Status of the cluster.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
}

func GetClustersOutput(ctx *pulumi.Context, args GetClustersOutputArgs, opts ...pulumi.InvokeOption) GetClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClustersResult, error) {
			args := v.(GetClustersArgs)
			r, err := GetClusters(ctx, &args, opts...)
			var s GetClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClustersResultOutput)
}

// A collection of arguments for invoking getClusters.
type GetClustersOutputArgs struct {
	// Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
	DbType pulumi.StringPtrInput `pulumi:"dbType"`
	// A regex string to filter results by cluster description.
	DescriptionRegex pulumi.StringPtrInput `pulumi:"descriptionRegex"`
	// A list of PolarDB cluster IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// status of the cluster.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersArgs)(nil)).Elem()
}

// A collection of values returned by getClusters.
type GetClustersResultOutput struct{ *pulumi.OutputState }

func (GetClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersResult)(nil)).Elem()
}

func (o GetClustersResultOutput) ToGetClustersResultOutput() GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ToGetClustersResultOutputWithContext(ctx context.Context) GetClustersResultOutput {
	return o
}

// A list of PolarDB clusters. Each element contains the following attributes:
func (o GetClustersResultOutput) Clusters() GetClustersClusterArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []GetClustersCluster { return v.Clusters }).(GetClustersClusterArrayOutput)
}

// `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
func (o GetClustersResultOutput) DbType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.DbType }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) DescriptionRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.DescriptionRegex }).(pulumi.StringPtrOutput)
}

// A list of RDS cluster descriptions.
func (o GetClustersResultOutput) Descriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Descriptions }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of RDS cluster IDs.
func (o GetClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Status of the cluster.
func (o GetClustersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetClustersResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersResultOutput{})
}