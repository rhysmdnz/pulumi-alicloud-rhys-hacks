// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the VServer groups related to a server load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/slb"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "slbservergroups"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/16"),
//				ZoneId:      pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "defaultApplicationLoadBalancer", &slb.ApplicationLoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				VswitchId:        defaultSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = slb.NewServerGroup(ctx, "defaultServerGroup", &slb.ServerGroupArgs{
//				LoadBalancerId: defaultApplicationLoadBalancer.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			sampleDs := slb.GetServerGroupsOutput(ctx, slb.GetServerGroupsOutputArgs{
//				LoadBalancerId: defaultApplicationLoadBalancer.ID(),
//			}, nil)
//			ctx.Export("firstSlbServerGroupId", sampleDs.ApplyT(func(sampleDs slb.GetServerGroupsResult) (string, error) {
//				return sampleDs.SlbServerGroups[0].Id, nil
//			}).(pulumi.StringOutput))
//			return nil
//		})
//	}
//
// ```
func GetServerGroups(ctx *pulumi.Context, args *GetServerGroupsArgs, opts ...pulumi.InvokeOption) (*GetServerGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetServerGroupsResult
	err := ctx.Invoke("alicloud:slb/getServerGroups:getServerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerGroups.
type GetServerGroupsArgs struct {
	// A list of VServer group IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by VServer group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getServerGroups.
type GetServerGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB VServer groups IDs.
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	// A list of SLB VServer groups names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of SLB VServer groups. Each element contains the following attributes:
	SlbServerGroups []GetServerGroupsSlbServerGroup `pulumi:"slbServerGroups"`
}

func GetServerGroupsOutput(ctx *pulumi.Context, args GetServerGroupsOutputArgs, opts ...pulumi.InvokeOption) GetServerGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerGroupsResult, error) {
			args := v.(GetServerGroupsArgs)
			r, err := GetServerGroups(ctx, &args, opts...)
			var s GetServerGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerGroupsResultOutput)
}

// A collection of arguments for invoking getServerGroups.
type GetServerGroupsOutputArgs struct {
	// A list of VServer group IDs to filter results.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the SLB.
	LoadBalancerId pulumi.StringInput `pulumi:"loadBalancerId"`
	// A regex string to filter results by VServer group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetServerGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getServerGroups.
type GetServerGroupsResultOutput struct{ *pulumi.OutputState }

func (GetServerGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupsResult)(nil)).Elem()
}

func (o GetServerGroupsResultOutput) ToGetServerGroupsResultOutput() GetServerGroupsResultOutput {
	return o
}

func (o GetServerGroupsResultOutput) ToGetServerGroupsResultOutputWithContext(ctx context.Context) GetServerGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SLB VServer groups IDs.
func (o GetServerGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServerGroupsResultOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerGroupsResult) string { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o GetServerGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of SLB VServer groups names.
func (o GetServerGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServerGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of SLB VServer groups. Each element contains the following attributes:
func (o GetServerGroupsResultOutput) SlbServerGroups() GetServerGroupsSlbServerGroupArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []GetServerGroupsSlbServerGroup { return v.SlbServerGroups }).(GetServerGroupsSlbServerGroupArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerGroupsResultOutput{})
}