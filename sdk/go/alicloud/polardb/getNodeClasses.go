// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the PolarDB node classes resource available info of Alibaba Cloud.
//
// > **NOTE:** Available in v1.81.0+
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
//			resources, err := polardb.GetNodeClasses(ctx, &polardb.GetNodeClassesArgs{
//				PayType:   "PostPaid",
//				DbType:    pulumi.StringRef("MySQL"),
//				DbVersion: pulumi.StringRef("5.6"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("polardbNodeClasses", resources.Classes)
//			ctx.Export("polardbAvailableZoneId", resources.Classes[0].ZoneId)
//			return nil
//		})
//	}
//
// ```
func GetNodeClasses(ctx *pulumi.Context, args *GetNodeClassesArgs, opts ...pulumi.InvokeOption) (*GetNodeClassesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNodeClassesResult
	err := ctx.Invoke("alicloud:polardb/getNodeClasses:getNodeClasses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodeClasses.
type GetNodeClassesArgs struct {
	Category *string `pulumi:"category"`
	// The PolarDB node class type by the user.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	// Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If dbType is set, dbVersion also needs to be set.
	DbType *string `pulumi:"dbType"`
	// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If dbVersion is set, dbType also needs to be set.
	DbVersion  *string `pulumi:"dbVersion"`
	OutputFile *string `pulumi:"outputFile"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
	PayType string `pulumi:"payType"`
	// The Region to launch the PolarDB cluster.
	RegionId *string `pulumi:"regionId"`
	// The Zone to launch the PolarDB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getNodeClasses.
type GetNodeClassesResult struct {
	Category *string `pulumi:"category"`
	// A list of PolarDB node classes. Each element contains the following attributes:
	Classes []GetNodeClassesClass `pulumi:"classes"`
	// PolarDB node available class.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	DbType      *string `pulumi:"dbType"`
	DbVersion   *string `pulumi:"dbVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	PayType    string  `pulumi:"payType"`
	RegionId   *string `pulumi:"regionId"`
	// The Zone to launch the PolarDB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

func GetNodeClassesOutput(ctx *pulumi.Context, args GetNodeClassesOutputArgs, opts ...pulumi.InvokeOption) GetNodeClassesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNodeClassesResult, error) {
			args := v.(GetNodeClassesArgs)
			r, err := GetNodeClasses(ctx, &args, opts...)
			var s GetNodeClassesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNodeClassesResultOutput)
}

// A collection of arguments for invoking getNodeClasses.
type GetNodeClassesOutputArgs struct {
	Category pulumi.StringPtrInput `pulumi:"category"`
	// The PolarDB node class type by the user.
	DbNodeClass pulumi.StringPtrInput `pulumi:"dbNodeClass"`
	// Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If dbType is set, dbVersion also needs to be set.
	DbType pulumi.StringPtrInput `pulumi:"dbType"`
	// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If dbVersion is set, dbType also needs to be set.
	DbVersion  pulumi.StringPtrInput `pulumi:"dbVersion"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
	PayType pulumi.StringInput `pulumi:"payType"`
	// The Region to launch the PolarDB cluster.
	RegionId pulumi.StringPtrInput `pulumi:"regionId"`
	// The Zone to launch the PolarDB cluster.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetNodeClassesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodeClassesArgs)(nil)).Elem()
}

// A collection of values returned by getNodeClasses.
type GetNodeClassesResultOutput struct{ *pulumi.OutputState }

func (GetNodeClassesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodeClassesResult)(nil)).Elem()
}

func (o GetNodeClassesResultOutput) ToGetNodeClassesResultOutput() GetNodeClassesResultOutput {
	return o
}

func (o GetNodeClassesResultOutput) ToGetNodeClassesResultOutputWithContext(ctx context.Context) GetNodeClassesResultOutput {
	return o
}

func (o GetNodeClassesResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// A list of PolarDB node classes. Each element contains the following attributes:
func (o GetNodeClassesResultOutput) Classes() GetNodeClassesClassArrayOutput {
	return o.ApplyT(func(v GetNodeClassesResult) []GetNodeClassesClass { return v.Classes }).(GetNodeClassesClassArrayOutput)
}

// PolarDB node available class.
func (o GetNodeClassesResultOutput) DbNodeClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.DbNodeClass }).(pulumi.StringPtrOutput)
}

func (o GetNodeClassesResultOutput) DbType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.DbType }).(pulumi.StringPtrOutput)
}

func (o GetNodeClassesResultOutput) DbVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.DbVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNodeClassesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeClassesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNodeClassesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetNodeClassesResultOutput) PayType() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeClassesResult) string { return v.PayType }).(pulumi.StringOutput)
}

func (o GetNodeClassesResultOutput) RegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.RegionId }).(pulumi.StringPtrOutput)
}

// The Zone to launch the PolarDB cluster.
func (o GetNodeClassesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodeClassesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNodeClassesResultOutput{})
}
