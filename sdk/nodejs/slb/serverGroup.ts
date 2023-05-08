// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A virtual server group contains several ECS instances. The virtual server group can help you to define multiple listening dimension,
 * and to meet the personalized requirements of domain name and URL forwarding.
 *
 * > **NOTE:** One ECS instance can be added into multiple virtual server groups.
 *
 * > **NOTE:** One virtual server group can be attached with multiple listeners in one load balancer.
 *
 * > **NOTE:** One Classic and Internet load balancer, its virtual server group can add Classic and VPC ECS instances.
 *
 * > **NOTE:** One Classic and Intranet load balancer, its virtual server group can only add Classic ECS instances.
 *
 * > **NOTE:** One VPC load balancer, its virtual server group can only add the same VPC ECS instances.
 *
 * For information about server group and how to use it, see [Configure a server group](https://www.alibabacloud.com/help/en/doc-detail/35215.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const slbServerGroupName = config.get("slbServerGroupName") || "forSlbServerGroup";
 * const serverGroupZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const serverGroupNetwork = new alicloud.vpc.Network("serverGroupNetwork", {
 *     vpcName: slbServerGroupName,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const serverGroupSwitch = new alicloud.vpc.Switch("serverGroupSwitch", {
 *     vpcId: serverGroupNetwork.id,
 *     cidrBlock: "172.16.0.0/16",
 *     zoneId: serverGroupZones.then(serverGroupZones => serverGroupZones.zones?[0]?.id),
 *     vswitchName: slbServerGroupName,
 * });
 * const serverGroupApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("serverGroupApplicationLoadBalancer", {
 *     loadBalancerName: slbServerGroupName,
 *     vswitchId: serverGroupSwitch.id,
 *     instanceChargeType: "PayByCLCU",
 * });
 * const serverGroupServerGroup = new alicloud.slb.ServerGroup("serverGroupServerGroup", {loadBalancerId: serverGroupApplicationLoadBalancer.id});
 * ```
 * ## Block servers
 *
 * The servers mapping supports the following:
 *
 * * `serverIds` - (Required) A list backend server ID (ECS instance ID).
 * * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
 * * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
 * * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
 *
 * ## Import
 *
 * Load balancer backend server group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:slb/serverGroup:ServerGroup example abc123456
 * ```
 */
export class ServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerGroupState, opts?: pulumi.CustomResourceOptions): ServerGroup {
        return new ServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/serverGroup:ServerGroup';

    /**
     * Returns true if the given object is an instance of ServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerGroup.__pulumiType;
    }

    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    public readonly deleteProtectionValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     *
     * @deprecated Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
     */
    public readonly servers!: pulumi.Output<outputs.slb.ServerGroupServer[]>;

    /**
     * Create a ServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerGroupArgs | ServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerGroupState | undefined;
            resourceInputs["deleteProtectionValidation"] = state ? state.deleteProtectionValidation : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as ServerGroupArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            resourceInputs["deleteProtectionValidation"] = args ? args.deleteProtectionValidation : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerGroup resources.
 */
export interface ServerGroupState {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     *
     * @deprecated Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
     */
    servers?: pulumi.Input<pulumi.Input<inputs.slb.ServerGroupServer>[]>;
}

/**
 * The set of arguments for constructing a ServerGroup resource.
 */
export interface ServerGroupArgs {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     *
     * @deprecated Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
     */
    servers?: pulumi.Input<pulumi.Input<inputs.slb.ServerGroupServer>[]>;
}
