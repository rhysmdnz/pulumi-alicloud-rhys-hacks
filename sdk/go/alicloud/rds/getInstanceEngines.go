// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the RDS instance engines resource available info of Alibaba Cloud.
//
// > **NOTE:** Available in v1.46.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/rds"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			resources, err := rds.GetInstanceEngines(ctx, &rds.GetInstanceEnginesArgs{
//				Engine:             pulumi.StringRef("MySQL"),
//				EngineVersion:      pulumi.StringRef("5.6"),
//				InstanceChargeType: pulumi.StringRef("PostPaid"),
//				OutputFile:         pulumi.StringRef("./engines.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstDbCategory", resources.InstanceEngines[0].Category)
//			return nil
//		})
//	}
//
// ```
func GetInstanceEngines(ctx *pulumi.Context, args *GetInstanceEnginesArgs, opts ...pulumi.InvokeOption) (*GetInstanceEnginesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceEnginesResult
	err := ctx.Invoke("alicloud:rds/getInstanceEngines:getInstanceEngines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceEngines.
type GetInstanceEnginesArgs struct {
	// DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
	Category *string `pulumi:"category"`
	// The DB instance storage space required by the user. Valid values: "cloudSsd", "localSsd", "cloudEssd", "cloudEssd2", "cloudEssd3".
	DbInstanceStorageType *string `pulumi:"dbInstanceStorageType"`
	// Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "MariaDB". If not set, it will match all of engines.
	Engine *string `pulumi:"engine"`
	// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Whether to show multi available zone. Default false to not show multi availability zone.
	MultiZone  *bool   `pulumi:"multiZone"`
	OutputFile *string `pulumi:"outputFile"`
	// The Zone to launch the DB instance.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getInstanceEngines.
type GetInstanceEnginesResult struct {
	// DB Instance category.
	Category              *string `pulumi:"category"`
	DbInstanceStorageType *string `pulumi:"dbInstanceStorageType"`
	// Database type.
	Engine *string `pulumi:"engine"`
	// DB Instance version.
	EngineVersion *string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of engines.
	Ids                []string `pulumi:"ids"`
	InstanceChargeType *string  `pulumi:"instanceChargeType"`
	// A list of Rds available resource. Each element contains the following attributes:
	InstanceEngines []GetInstanceEnginesInstanceEngine `pulumi:"instanceEngines"`
	MultiZone       *bool                              `pulumi:"multiZone"`
	OutputFile      *string                            `pulumi:"outputFile"`
	ZoneId          *string                            `pulumi:"zoneId"`
}

func GetInstanceEnginesOutput(ctx *pulumi.Context, args GetInstanceEnginesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceEnginesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceEnginesResult, error) {
			args := v.(GetInstanceEnginesArgs)
			r, err := GetInstanceEngines(ctx, &args, opts...)
			var s GetInstanceEnginesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceEnginesResultOutput)
}

// A collection of arguments for invoking getInstanceEngines.
type GetInstanceEnginesOutputArgs struct {
	// DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
	Category pulumi.StringPtrInput `pulumi:"category"`
	// The DB instance storage space required by the user. Valid values: "cloudSsd", "localSsd", "cloudEssd", "cloudEssd2", "cloudEssd3".
	DbInstanceStorageType pulumi.StringPtrInput `pulumi:"dbInstanceStorageType"`
	// Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "MariaDB". If not set, it will match all of engines.
	Engine pulumi.StringPtrInput `pulumi:"engine"`
	// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion pulumi.StringPtrInput `pulumi:"engineVersion"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput `pulumi:"instanceChargeType"`
	// Whether to show multi available zone. Default false to not show multi availability zone.
	MultiZone  pulumi.BoolPtrInput   `pulumi:"multiZone"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Zone to launch the DB instance.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetInstanceEnginesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceEnginesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceEngines.
type GetInstanceEnginesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceEnginesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceEnginesResult)(nil)).Elem()
}

func (o GetInstanceEnginesResultOutput) ToGetInstanceEnginesResultOutput() GetInstanceEnginesResultOutput {
	return o
}

func (o GetInstanceEnginesResultOutput) ToGetInstanceEnginesResultOutputWithContext(ctx context.Context) GetInstanceEnginesResultOutput {
	return o
}

// DB Instance category.
func (o GetInstanceEnginesResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetInstanceEnginesResultOutput) DbInstanceStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.DbInstanceStorageType }).(pulumi.StringPtrOutput)
}

// Database type.
func (o GetInstanceEnginesResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

// DB Instance version.
func (o GetInstanceEnginesResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceEnginesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of engines.
func (o GetInstanceEnginesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstanceEnginesResultOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

// A list of Rds available resource. Each element contains the following attributes:
func (o GetInstanceEnginesResultOutput) InstanceEngines() GetInstanceEnginesInstanceEngineArrayOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) []GetInstanceEnginesInstanceEngine { return v.InstanceEngines }).(GetInstanceEnginesInstanceEngineArrayOutput)
}

func (o GetInstanceEnginesResultOutput) MultiZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *bool { return v.MultiZone }).(pulumi.BoolPtrOutput)
}

func (o GetInstanceEnginesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstanceEnginesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceEnginesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceEnginesResultOutput{})
}
