// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of ALIKAFKA Sasl users in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.66.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/actiontrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/actiontrail"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			saslUsersDs, err := actiontrail.GetSaslUsers(ctx, &actiontrail.GetSaslUsersArgs{
//				InstanceId: "xxx",
//				NameRegex:  pulumi.StringRef("username"),
//				OutputFile: pulumi.StringRef("saslUsers.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstSaslUsername", saslUsersDs.Users[0].Username)
//			return nil
//		})
//	}
//
// ```
func GetSaslUsers(ctx *pulumi.Context, args *GetSaslUsersArgs, opts ...pulumi.InvokeOption) (*GetSaslUsersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSaslUsersResult
	err := ctx.Invoke("alicloud:actiontrail/getSaslUsers:getSaslUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSaslUsers.
type GetSaslUsersArgs struct {
	// ID of the ALIKAFKA Instance that owns the sasl users.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by the username.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getSaslUsers.
type GetSaslUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of sasl usernames.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of sasl users. Each element contains the following attributes:
	Users []GetSaslUsersUser `pulumi:"users"`
}

func GetSaslUsersOutput(ctx *pulumi.Context, args GetSaslUsersOutputArgs, opts ...pulumi.InvokeOption) GetSaslUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSaslUsersResult, error) {
			args := v.(GetSaslUsersArgs)
			r, err := GetSaslUsers(ctx, &args, opts...)
			var s GetSaslUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSaslUsersResultOutput)
}

// A collection of arguments for invoking getSaslUsers.
type GetSaslUsersOutputArgs struct {
	// ID of the ALIKAFKA Instance that owns the sasl users.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by the username.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetSaslUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSaslUsersArgs)(nil)).Elem()
}

// A collection of values returned by getSaslUsers.
type GetSaslUsersResultOutput struct{ *pulumi.OutputState }

func (GetSaslUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSaslUsersResult)(nil)).Elem()
}

func (o GetSaslUsersResultOutput) ToGetSaslUsersResultOutput() GetSaslUsersResultOutput {
	return o
}

func (o GetSaslUsersResultOutput) ToGetSaslUsersResultOutputWithContext(ctx context.Context) GetSaslUsersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSaslUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSaslUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSaslUsersResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSaslUsersResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSaslUsersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSaslUsersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of sasl usernames.
func (o GetSaslUsersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSaslUsersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSaslUsersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSaslUsersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of sasl users. Each element contains the following attributes:
func (o GetSaslUsersResultOutput) Users() GetSaslUsersUserArrayOutput {
	return o.ApplyT(func(v GetSaslUsersResult) []GetSaslUsersUser { return v.Users }).(GetSaslUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSaslUsersResultOutput{})
}