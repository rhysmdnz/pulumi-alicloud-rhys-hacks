// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.
 *
 * For information about ECS Disk Attachment and how to use it, see [What is Disk Attachment](https://www.alibabacloud.com/help/en/doc-detail/25515.htm).
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new ECS disk-attachment and use it attach one disk to a new instance.
 * const ecsSg = new alicloud.ecs.SecurityGroup("ecsSg", {description: "New security group"});
 * const ecsDisk = new alicloud.ecs.EcsDisk("ecsDisk", {
 *     availabilityZone: "cn-beijing-a",
 *     size: 50,
 *     tags: {
 *         Name: "TerraformTest-disk",
 *     },
 * });
 * const ecsInstance = new alicloud.ecs.Instance("ecsInstance", {
 *     imageId: "ubuntu_18_04_64_20G_alibase_20190624.vhd",
 *     instanceType: "ecs.n4.small",
 *     availabilityZone: "cn-beijing-a",
 *     securityGroups: [ecsSg.id],
 *     instanceName: "Hello",
 *     internetChargeType: "PayByBandwidth",
 *     tags: {
 *         Name: "TerraformTest-instance",
 *     },
 * });
 * const ecsDiskAtt = new alicloud.ecs.EcsDiskAttachment("ecsDiskAtt", {
 *     diskId: ecsDisk.id,
 *     instanceId: ecsInstance.id,
 * });
 * ```
 *
 * ## Import
 *
 * The disk attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment example d-abc12345678:i-abc12355
 * ```
 */
export class EcsDiskAttachment extends pulumi.CustomResource {
    /**
     * Get an existing EcsDiskAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsDiskAttachmentState, opts?: pulumi.CustomResourceOptions): EcsDiskAttachment {
        return new EcsDiskAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment';

    /**
     * Returns true if the given object is an instance of EcsDiskAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsDiskAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsDiskAttachment.__pulumiType;
    }

    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    public readonly bootable!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly device!: pulumi.Output<string>;
    /**
     * ID of the Disk to be attached.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * ID of the Instance to attach to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of key pair
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    public readonly password!: pulumi.Output<string | undefined>;

    /**
     * Create a EcsDiskAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsDiskAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsDiskAttachmentArgs | EcsDiskAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsDiskAttachmentState | undefined;
            resourceInputs["bootable"] = state ? state.bootable : undefined;
            resourceInputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as EcsDiskAttachmentArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["bootable"] = args ? args.bootable : undefined;
            resourceInputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["device"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsDiskAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsDiskAttachment resources.
 */
export interface EcsDiskAttachmentState {
    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    bootable?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    device?: pulumi.Input<string>;
    /**
     * ID of the Disk to be attached.
     */
    diskId?: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of key pair
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsDiskAttachment resource.
 */
export interface EcsDiskAttachmentArgs {
    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    bootable?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * ID of the Disk to be attached.
     */
    diskId: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of key pair
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    password?: pulumi.Input<string>;
}