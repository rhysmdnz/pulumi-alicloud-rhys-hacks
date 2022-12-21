// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the available instance types with the Cloud Phone (ECP) Instance of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.158.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/ecp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/ecp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := ecp.GetInstanceTypes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcpInstanceTypesInstanceType", _default.InstanceTypes[0].InstanceType)
//			return nil
//		})
//	}
//
// ```
func GetInstanceTypes(ctx *pulumi.Context, args *GetInstanceTypesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceTypesResult
	err := ctx.Invoke("alicloud:ecp/getInstanceTypes:getInstanceTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesArgs struct {
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                         `pulumi:"id"`
	InstanceTypes []GetInstanceTypesInstanceType `pulumi:"instanceTypes"`
	OutputFile    *string                        `pulumi:"outputFile"`
}

func GetInstanceTypesOutput(ctx *pulumi.Context, args GetInstanceTypesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceTypesResult, error) {
			args := v.(GetInstanceTypesArgs)
			r, err := GetInstanceTypes(ctx, &args, opts...)
			var s GetInstanceTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceTypesResultOutput)
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesOutputArgs struct {
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetInstanceTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesResult)(nil)).Elem()
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutput() GetInstanceTypesResultOutput {
	return o
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutputWithContext(ctx context.Context) GetInstanceTypesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceTypesResultOutput) InstanceTypes() GetInstanceTypesInstanceTypeArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []GetInstanceTypesInstanceType { return v.InstanceTypes }).(GetInstanceTypesInstanceTypeArrayOutput)
}

func (o GetInstanceTypesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceTypesResultOutput{})
}