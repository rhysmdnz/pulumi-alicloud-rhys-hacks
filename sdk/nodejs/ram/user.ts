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
 * // Create a new RAM user.
 * const user = new alicloud.ram.User("user", {
 *     comments: "yoyoyo",
 *     displayName: "user_display_name",
 *     email: "hello.uuu@aaa.com",
 *     force: true,
 *     mobile: "86-18688888888",
 * });
 * ```
 *
 * ## Import
 *
 * RAM user can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ram/user:User example 123456789xxx
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Email of the RAM user.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
     */
    public readonly mobile!: pulumi.Output<string | undefined>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["mobile"] = state ? state.mobile : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["mobile"] = args ? args.mobile : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
     */
    comments?: pulumi.Input<string>;
    /**
     * Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Email of the RAM user.
     */
    email?: pulumi.Input<string>;
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
     */
    mobile?: pulumi.Input<string>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
     */
    comments?: pulumi.Input<string>;
    /**
     * Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Email of the RAM user.
     */
    email?: pulumi.Input<string>;
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
     */
    mobile?: pulumi.Input<string>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    name?: pulumi.Input<string>;
}