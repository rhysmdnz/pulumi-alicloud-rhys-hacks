// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the server certificate list.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/slb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sampleDs, err := slb.GetServerCertificates(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstSlbServerCertificateId", sampleDs.Certificates[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetServerCertificates(ctx *pulumi.Context, args *GetServerCertificatesArgs, opts ...pulumi.InvokeOption) (*GetServerCertificatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetServerCertificatesResult
	err := ctx.Invoke("alicloud:slb/getServerCertificates:getServerCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerCertificates.
type GetServerCertificatesArgs struct {
	// A list of server certificates IDs to filter results.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by server certificate name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which the slb server certificates belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getServerCertificates.
type GetServerCertificatesResult struct {
	// A list of SLB server certificates. Each element contains the following attributes:
	Certificates []GetServerCertificatesCertificate `pulumi:"certificates"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB server certificates IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of SLB server certificates names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The Id of resource group which the slb server certificates belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// (Available in v1.66.0+) A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetServerCertificatesOutput(ctx *pulumi.Context, args GetServerCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetServerCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerCertificatesResult, error) {
			args := v.(GetServerCertificatesArgs)
			r, err := GetServerCertificates(ctx, &args, opts...)
			var s GetServerCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerCertificatesResultOutput)
}

// A collection of arguments for invoking getServerCertificates.
type GetServerCertificatesOutputArgs struct {
	// A list of server certificates IDs to filter results.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by server certificate name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Id of resource group which the slb server certificates belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetServerCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getServerCertificates.
type GetServerCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetServerCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerCertificatesResult)(nil)).Elem()
}

func (o GetServerCertificatesResultOutput) ToGetServerCertificatesResultOutput() GetServerCertificatesResultOutput {
	return o
}

func (o GetServerCertificatesResultOutput) ToGetServerCertificatesResultOutputWithContext(ctx context.Context) GetServerCertificatesResultOutput {
	return o
}

// A list of SLB server certificates. Each element contains the following attributes:
func (o GetServerCertificatesResultOutput) Certificates() GetServerCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) []GetServerCertificatesCertificate { return v.Certificates }).(GetServerCertificatesCertificateArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SLB server certificates IDs.
func (o GetServerCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServerCertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of SLB server certificates names.
func (o GetServerCertificatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServerCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the slb server certificates belongs.
func (o GetServerCertificatesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// (Available in v1.66.0+) A mapping of tags to assign to the resource.
func (o GetServerCertificatesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetServerCertificatesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerCertificatesResultOutput{})
}