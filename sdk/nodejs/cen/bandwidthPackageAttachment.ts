// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN bandwidth package attachment resource. The resource can be used to bind a bandwidth package to a specified CEN instance.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new bandwidth package attachment and use it to attach a bandwidth package to a new CEN
 * const cen = new alicloud.cen.Instance("cen", {description: "tf-testAccCenBandwidthPackageAttachmentDescription"});
 * const bwp = new alicloud.cen.BandwidthPackage("bwp", {
 *     bandwidth: 20,
 *     geographicRegionIds: [
 *         "China",
 *         "Asia-Pacific",
 *     ],
 * });
 * const foo = new alicloud.cen.BandwidthPackageAttachment("foo", {
 *     instanceId: cen.id,
 *     bandwidthPackageId: bwp.id,
 * });
 * ```
 *
 * ## Import
 *
 * CEN bandwidth package attachment resource can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment example bwp-abc123456
 * ```
 */
export class BandwidthPackageAttachment extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackageAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions): BandwidthPackageAttachment {
        return new BandwidthPackageAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment';

    /**
     * Returns true if the given object is an instance of BandwidthPackageAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackageAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackageAttachment.__pulumiType;
    }

    /**
     * The ID of the bandwidth package.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * The ID of the CEN.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPackageAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageAttachmentArgs | BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageAttachmentState | undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as BandwidthPackageAttachmentArgs | undefined;
            if ((!args || args.bandwidthPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthPackageId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPackageAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackageAttachment resources.
 */
export interface BandwidthPackageAttachmentState {
    /**
     * The ID of the bandwidth package.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPackageAttachment resource.
 */
export interface BandwidthPackageAttachmentArgs {
    /**
     * The ID of the bandwidth package.
     */
    bandwidthPackageId: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    instanceId: pulumi.Input<string>;
}