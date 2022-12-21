// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ga Additional Certificates of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.150.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/ga"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ga.GetAdditionalCertificates(ctx, &ga.GetAdditionalCertificatesArgs{
//				AcceleratorId: "example_value",
//				ListenerId:    "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("gaAdditionalCertificateId1", ids.Certificates[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAdditionalCertificates(ctx *pulumi.Context, args *GetAdditionalCertificatesArgs, opts ...pulumi.InvokeOption) (*GetAdditionalCertificatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAdditionalCertificatesResult
	err := ctx.Invoke("alicloud:ga/getAdditionalCertificates:getAdditionalCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAdditionalCertificates.
type GetAdditionalCertificatesArgs struct {
	// The ID of the GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// A list of Additional Certificate IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the listener. Only HTTPS listeners support this parameter.
	ListenerId string  `pulumi:"listenerId"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAdditionalCertificates.
type GetAdditionalCertificatesResult struct {
	AcceleratorId string                                 `pulumi:"acceleratorId"`
	Certificates  []GetAdditionalCertificatesCertificate `pulumi:"certificates"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	ListenerId string   `pulumi:"listenerId"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetAdditionalCertificatesOutput(ctx *pulumi.Context, args GetAdditionalCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetAdditionalCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAdditionalCertificatesResult, error) {
			args := v.(GetAdditionalCertificatesArgs)
			r, err := GetAdditionalCertificates(ctx, &args, opts...)
			var s GetAdditionalCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAdditionalCertificatesResultOutput)
}

// A collection of arguments for invoking getAdditionalCertificates.
type GetAdditionalCertificatesOutputArgs struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringInput `pulumi:"acceleratorId"`
	// A list of Additional Certificate IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the listener. Only HTTPS listeners support this parameter.
	ListenerId pulumi.StringInput    `pulumi:"listenerId"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAdditionalCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdditionalCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getAdditionalCertificates.
type GetAdditionalCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetAdditionalCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdditionalCertificatesResult)(nil)).Elem()
}

func (o GetAdditionalCertificatesResultOutput) ToGetAdditionalCertificatesResultOutput() GetAdditionalCertificatesResultOutput {
	return o
}

func (o GetAdditionalCertificatesResultOutput) ToGetAdditionalCertificatesResultOutputWithContext(ctx context.Context) GetAdditionalCertificatesResultOutput {
	return o
}

func (o GetAdditionalCertificatesResultOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) string { return v.AcceleratorId }).(pulumi.StringOutput)
}

func (o GetAdditionalCertificatesResultOutput) Certificates() GetAdditionalCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) []GetAdditionalCertificatesCertificate { return v.Certificates }).(GetAdditionalCertificatesCertificateArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAdditionalCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAdditionalCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAdditionalCertificatesResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

func (o GetAdditionalCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdditionalCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAdditionalCertificatesResultOutput{})
}