// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides CEN Transit Router Route Table Associations available to the user.[What is Cen Transit Router Route Table Associations](https://help.aliyun.com/document_detail/261243.html)
//
// > **NOTE:** Available in 1.126.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := cen.GetTransitRouterRouteTableAssociations(ctx, &cen.GetTransitRouterRouteTableAssociationsArgs{
//				TransitRouterRouteTableId: "rtb-id1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType", _default.Associations[0].ResourceType)
//			return nil
//		})
//	}
//
// ```
func GetTransitRouterRouteTableAssociations(ctx *pulumi.Context, args *GetTransitRouterRouteTableAssociationsArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterRouteTableAssociationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTransitRouterRouteTableAssociationsResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterRouteTableAssociations:getTransitRouterRouteTableAssociations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterRouteTableAssociations.
type GetTransitRouterRouteTableAssociationsArgs struct {
	// A list of CEN Transit Router Route Table Association IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the route table, including `Active`, `Associating`, `Dissociating`.
	Status *string `pulumi:"status"`
	// ID of the route table of the VPC or VBR.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

// A collection of values returned by getTransitRouterRouteTableAssociations.
type GetTransitRouterRouteTableAssociationsResult struct {
	// A list of CEN Transit Router Route Table Associations. Each element contains the following attributes:
	Associations []GetTransitRouterRouteTableAssociationsAssociation `pulumi:"associations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CEN Transit Router Route Table Association IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the route table.
	Status *string `pulumi:"status"`
	// ID of the transit router route table.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

func GetTransitRouterRouteTableAssociationsOutput(ctx *pulumi.Context, args GetTransitRouterRouteTableAssociationsOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterRouteTableAssociationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterRouteTableAssociationsResult, error) {
			args := v.(GetTransitRouterRouteTableAssociationsArgs)
			r, err := GetTransitRouterRouteTableAssociations(ctx, &args, opts...)
			var s GetTransitRouterRouteTableAssociationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterRouteTableAssociationsResultOutput)
}

// A collection of arguments for invoking getTransitRouterRouteTableAssociations.
type GetTransitRouterRouteTableAssociationsOutputArgs struct {
	// A list of CEN Transit Router Route Table Association IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The status of the route table, including `Active`, `Associating`, `Dissociating`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// ID of the route table of the VPC or VBR.
	TransitRouterRouteTableId pulumi.StringInput `pulumi:"transitRouterRouteTableId"`
}

func (GetTransitRouterRouteTableAssociationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterRouteTableAssociationsArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterRouteTableAssociations.
type GetTransitRouterRouteTableAssociationsResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterRouteTableAssociationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterRouteTableAssociationsResult)(nil)).Elem()
}

func (o GetTransitRouterRouteTableAssociationsResultOutput) ToGetTransitRouterRouteTableAssociationsResultOutput() GetTransitRouterRouteTableAssociationsResultOutput {
	return o
}

func (o GetTransitRouterRouteTableAssociationsResultOutput) ToGetTransitRouterRouteTableAssociationsResultOutputWithContext(ctx context.Context) GetTransitRouterRouteTableAssociationsResultOutput {
	return o
}

// A list of CEN Transit Router Route Table Associations. Each element contains the following attributes:
func (o GetTransitRouterRouteTableAssociationsResultOutput) Associations() GetTransitRouterRouteTableAssociationsAssociationArrayOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) []GetTransitRouterRouteTableAssociationsAssociation {
		return v.Associations
	}).(GetTransitRouterRouteTableAssociationsAssociationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterRouteTableAssociationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of CEN Transit Router Route Table Association IDs.
func (o GetTransitRouterRouteTableAssociationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterRouteTableAssociationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The status of the route table.
func (o GetTransitRouterRouteTableAssociationsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// ID of the transit router route table.
func (o GetTransitRouterRouteTableAssociationsResultOutput) TransitRouterRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterRouteTableAssociationsResult) string { return v.TransitRouterRouteTableId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterRouteTableAssociationsResultOutput{})
}