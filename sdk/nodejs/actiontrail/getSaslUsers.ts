// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ALIKAFKA Sasl users in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.66.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const saslUsersDs = pulumi.output(alicloud.actiontrail.getSaslUsers({
 *     instanceId: "xxx",
 *     nameRegex: "username",
 *     outputFile: "saslUsers.txt",
 * }));
 *
 * export const firstSaslUsername = saslUsersDs.users[0].username;
 * ```
 */
export function getSaslUsers(args: GetSaslUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetSaslUsersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:actiontrail/getSaslUsers:getSaslUsers", {
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSaslUsers.
 */
export interface GetSaslUsersArgs {
    /**
     * ID of the ALIKAFKA Instance that owns the sasl users.
     */
    instanceId: string;
    /**
     * A regex string to filter results by the username.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getSaslUsers.
 */
export interface GetSaslUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly nameRegex?: string;
    /**
     * A list of sasl usernames.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of sasl users. Each element contains the following attributes:
     */
    readonly users: outputs.actiontrail.GetSaslUsersUser[];
}

export function getSaslUsersOutput(args: GetSaslUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSaslUsersResult> {
    return pulumi.output(args).apply(a => getSaslUsers(a, opts))
}

/**
 * A collection of arguments for invoking getSaslUsers.
 */
export interface GetSaslUsersOutputArgs {
    /**
     * ID of the ALIKAFKA Instance that owns the sasl users.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A regex string to filter results by the username.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}