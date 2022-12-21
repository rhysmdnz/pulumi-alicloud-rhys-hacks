// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about the current account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("currentAccountId", current.Id)
//			return nil
//		})
//	}
//
// ```
func GetAccount(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAccountResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAccountResult
	err := ctx.Invoke("alicloud:index/getAccount:getAccount", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	// Account ID (e.g. "1239306421830812"). It can be used to construct an ARN.
	Id string `pulumi:"id"`
}