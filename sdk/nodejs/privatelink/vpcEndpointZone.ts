// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Zone resource.
 *
 * For information about Private Link Vpc Endpoint Zone and how to use it, see [What is Vpc Endpoint Zone](https://help.aliyun.com/document_detail/183561.html).
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.privatelink.VpcEndpointZone("example", {
 *     endpointId: "ep-gw8boxxxxx",
 *     vswitchId: "vsw-rtycxxxxx",
 *     zoneId: "eu-central-1a",
 * });
 * ```
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Zone can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointZone:VpcEndpointZone example <endpoint_id>:<zone_id>
 * ```
 */
export class VpcEndpointZone extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointZoneState, opts?: pulumi.CustomResourceOptions): VpcEndpointZone {
        return new VpcEndpointZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointZone:VpcEndpointZone';

    /**
     * Returns true if the given object is an instance of VpcEndpointZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointZone.__pulumiType;
    }

    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Vpc Endpoint.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * Status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The VSwitch id.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone Id.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointZoneArgs | VpcEndpointZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointZoneState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VpcEndpointZoneArgs | undefined;
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpointZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointZone resources.
 */
export interface VpcEndpointZoneState {
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Vpc Endpoint.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Status.
     */
    status?: pulumi.Input<string>;
    /**
     * The VSwitch id.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone Id.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointZone resource.
 */
export interface VpcEndpointZoneArgs {
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Vpc Endpoint.
     */
    endpointId: pulumi.Input<string>;
    /**
     * The VSwitch id.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * The Zone Id.
     */
    zoneId?: pulumi.Input<string>;
}