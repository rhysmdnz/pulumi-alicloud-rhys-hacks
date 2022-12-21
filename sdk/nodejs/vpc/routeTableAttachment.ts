// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "route-table-attachment-example-name";
 * const fooNetwork = new alicloud.vpc.Network("fooNetwork", {cidrBlock: "172.16.0.0/12"});
 * const default = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const fooSwitch = new alicloud.vpc.Switch("fooSwitch", {
 *     vpcId: fooNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: _default.then(_default => _default.zones?[0]?.id),
 * });
 * const fooRouteTable = new alicloud.vpc.RouteTable("fooRouteTable", {
 *     vpcId: fooNetwork.id,
 *     routeTableName: name,
 *     description: "route_table_attachment",
 * });
 * const fooRouteTableAttachment = new alicloud.vpc.RouteTableAttachment("fooRouteTableAttachment", {
 *     vswitchId: fooSwitch.id,
 *     routeTableId: fooRouteTable.id,
 * });
 * ```
 *
 * ## Import
 *
 * The route table attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/routeTableAttachment:RouteTableAttachment foo vtb-abc123456:vsw-abc123456
 * ```
 */
export class RouteTableAttachment extends pulumi.CustomResource {
    /**
     * Get an existing RouteTableAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableAttachmentState, opts?: pulumi.CustomResourceOptions): RouteTableAttachment {
        return new RouteTableAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routeTableAttachment:RouteTableAttachment';

    /**
     * Returns true if the given object is an instance of RouteTableAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTableAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTableAttachment.__pulumiType;
    }

    /**
     * The routeTableId of the route table attachment, the field can't be changed.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The vswitchId of the route table attachment, the field can't be changed.
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a RouteTableAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableAttachmentArgs | RouteTableAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableAttachmentState | undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as RouteTableAttachmentArgs | undefined;
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteTableAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTableAttachment resources.
 */
export interface RouteTableAttachmentState {
    /**
     * The routeTableId of the route table attachment, the field can't be changed.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The vswitchId of the route table attachment, the field can't be changed.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTableAttachment resource.
 */
export interface RouteTableAttachmentArgs {
    /**
     * The routeTableId of the route table attachment, the field can't be changed.
     */
    routeTableId: pulumi.Input<string>;
    /**
     * The vswitchId of the route table attachment, the field can't be changed.
     */
    vswitchId: pulumi.Input<string>;
}