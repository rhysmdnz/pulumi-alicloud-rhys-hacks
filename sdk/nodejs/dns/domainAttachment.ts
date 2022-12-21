// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const dns = new alicloud.dns.DomainAttachment("dns", {
 *     domainNames: [
 *         "test111.abc",
 *         "test222.abc",
 *     ],
 *     instanceId: "dns-cn-mp91lyq9xxxx",
 * });
 * ```
 *
 * ## Import
 *
 * DNS domain attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dns/domainAttachment:DomainAttachment example dns-cn-v0h1ldjhxxx
 * ```
 */
export class DomainAttachment extends pulumi.CustomResource {
    /**
     * Get an existing DomainAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainAttachmentState, opts?: pulumi.CustomResourceOptions): DomainAttachment {
        return new DomainAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/domainAttachment:DomainAttachment';

    /**
     * Returns true if the given object is an instance of DomainAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainAttachment.__pulumiType;
    }

    /**
     * The domain names bound to the DNS instance.
     */
    public readonly domainNames!: pulumi.Output<string[]>;
    /**
     * The id of the DNS instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a DomainAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainAttachmentArgs | DomainAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainAttachmentState | undefined;
            resourceInputs["domainNames"] = state ? state.domainNames : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as DomainAttachmentArgs | undefined;
            if ((!args || args.domainNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainNames'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["domainNames"] = args ? args.domainNames : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainAttachment resources.
 */
export interface DomainAttachmentState {
    /**
     * The domain names bound to the DNS instance.
     */
    domainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the DNS instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainAttachment resource.
 */
export interface DomainAttachmentArgs {
    /**
     * The domain names bound to the DNS instance.
     */
    domainNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the DNS instance.
     */
    instanceId: pulumi.Input<string>;
}