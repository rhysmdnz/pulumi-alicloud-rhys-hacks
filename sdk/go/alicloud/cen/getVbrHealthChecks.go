// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides CEN VBR Health Checks available to the user.
//
// > **NOTE:** Available in 1.98.0+
func GetVbrHealthChecks(ctx *pulumi.Context, args *GetVbrHealthChecksArgs, opts ...pulumi.InvokeOption) (*GetVbrHealthChecksResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetVbrHealthChecksResult
	err := ctx.Invoke("alicloud:cen/getVbrHealthChecks:getVbrHealthChecks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVbrHealthChecks.
type GetVbrHealthChecksArgs struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId      *string `pulumi:"cenId"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the VBR instance.
	VbrInstanceId *string `pulumi:"vbrInstanceId"`
	// The User ID (UID) of the account to which the VBR instance belongs.
	VbrInstanceOwnerId *int `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region where the VBR instance is deployed.
	VbrInstanceRegionId string `pulumi:"vbrInstanceRegionId"`
}

// A collection of values returned by getVbrHealthChecks.
type GetVbrHealthChecksResult struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId *string `pulumi:"cenId"`
	// A list of CEN VBR Heath Checks. Each element contains the following attributes:
	Checks []GetVbrHealthChecksCheck `pulumi:"checks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the CEN VBR Heath Check IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the VBR instance.
	VbrInstanceId      *string `pulumi:"vbrInstanceId"`
	VbrInstanceOwnerId *int    `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region where the VBR instance is deployed.
	VbrInstanceRegionId string `pulumi:"vbrInstanceRegionId"`
}

func GetVbrHealthChecksOutput(ctx *pulumi.Context, args GetVbrHealthChecksOutputArgs, opts ...pulumi.InvokeOption) GetVbrHealthChecksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVbrHealthChecksResult, error) {
			args := v.(GetVbrHealthChecksArgs)
			r, err := GetVbrHealthChecks(ctx, &args, opts...)
			var s GetVbrHealthChecksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVbrHealthChecksResultOutput)
}

// A collection of arguments for invoking getVbrHealthChecks.
type GetVbrHealthChecksOutputArgs struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId      pulumi.StringPtrInput `pulumi:"cenId"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the VBR instance.
	VbrInstanceId pulumi.StringPtrInput `pulumi:"vbrInstanceId"`
	// The User ID (UID) of the account to which the VBR instance belongs.
	VbrInstanceOwnerId pulumi.IntPtrInput `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region where the VBR instance is deployed.
	VbrInstanceRegionId pulumi.StringInput `pulumi:"vbrInstanceRegionId"`
}

func (GetVbrHealthChecksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVbrHealthChecksArgs)(nil)).Elem()
}

// A collection of values returned by getVbrHealthChecks.
type GetVbrHealthChecksResultOutput struct{ *pulumi.OutputState }

func (GetVbrHealthChecksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVbrHealthChecksResult)(nil)).Elem()
}

func (o GetVbrHealthChecksResultOutput) ToGetVbrHealthChecksResultOutput() GetVbrHealthChecksResultOutput {
	return o
}

func (o GetVbrHealthChecksResultOutput) ToGetVbrHealthChecksResultOutputWithContext(ctx context.Context) GetVbrHealthChecksResultOutput {
	return o
}

// The ID of the Cloud Enterprise Network (CEN) instance.
func (o GetVbrHealthChecksResultOutput) CenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) *string { return v.CenId }).(pulumi.StringPtrOutput)
}

// A list of CEN VBR Heath Checks. Each element contains the following attributes:
func (o GetVbrHealthChecksResultOutput) Checks() GetVbrHealthChecksCheckArrayOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) []GetVbrHealthChecksCheck { return v.Checks }).(GetVbrHealthChecksCheckArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVbrHealthChecksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the CEN VBR Heath Check IDs.
func (o GetVbrHealthChecksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVbrHealthChecksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the VBR instance.
func (o GetVbrHealthChecksResultOutput) VbrInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) *string { return v.VbrInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetVbrHealthChecksResultOutput) VbrInstanceOwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) *int { return v.VbrInstanceOwnerId }).(pulumi.IntPtrOutput)
}

// The ID of the region where the VBR instance is deployed.
func (o GetVbrHealthChecksResultOutput) VbrInstanceRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVbrHealthChecksResult) string { return v.VbrInstanceRegionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVbrHealthChecksResultOutput{})
}