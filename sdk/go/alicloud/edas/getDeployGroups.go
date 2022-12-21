// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of EDAS deploy groups in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.82.0+
func GetDeployGroups(ctx *pulumi.Context, args *GetDeployGroupsArgs, opts ...pulumi.InvokeOption) (*GetDeployGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDeployGroupsResult
	err := ctx.Invoke("alicloud:edas/getDeployGroups:getDeployGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeployGroups.
type GetDeployGroupsArgs struct {
	// ID of the EDAS application.
	AppId string `pulumi:"appId"`
	// A regex string to filter results by the deploy group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getDeployGroups.
type GetDeployGroupsResult struct {
	// The ID of the application that you want to deploy.
	AppId string `pulumi:"appId"`
	// A list of consumer group ids.
	Groups []GetDeployGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of deploy group names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetDeployGroupsOutput(ctx *pulumi.Context, args GetDeployGroupsOutputArgs, opts ...pulumi.InvokeOption) GetDeployGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeployGroupsResult, error) {
			args := v.(GetDeployGroupsArgs)
			r, err := GetDeployGroups(ctx, &args, opts...)
			var s GetDeployGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeployGroupsResultOutput)
}

// A collection of arguments for invoking getDeployGroups.
type GetDeployGroupsOutputArgs struct {
	// ID of the EDAS application.
	AppId pulumi.StringInput `pulumi:"appId"`
	// A regex string to filter results by the deploy group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetDeployGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeployGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getDeployGroups.
type GetDeployGroupsResultOutput struct{ *pulumi.OutputState }

func (GetDeployGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeployGroupsResult)(nil)).Elem()
}

func (o GetDeployGroupsResultOutput) ToGetDeployGroupsResultOutput() GetDeployGroupsResultOutput {
	return o
}

func (o GetDeployGroupsResultOutput) ToGetDeployGroupsResultOutputWithContext(ctx context.Context) GetDeployGroupsResultOutput {
	return o
}

// The ID of the application that you want to deploy.
func (o GetDeployGroupsResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) string { return v.AppId }).(pulumi.StringOutput)
}

// A list of consumer group ids.
func (o GetDeployGroupsResultOutput) Groups() GetDeployGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) []GetDeployGroupsGroup { return v.Groups }).(GetDeployGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeployGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeployGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of deploy group names.
func (o GetDeployGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetDeployGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeployGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeployGroupsResultOutput{})
}