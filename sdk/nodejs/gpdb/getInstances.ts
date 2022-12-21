// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the AnalyticDB for PostgreSQL instances of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.47.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.gpdb.getInstances({});
 * export const gpdbDbInstanceId1 = ids.then(ids => ids.instances?[0]?.id);
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:gpdb/getInstances:getInstances", {
        "availabilityZone": args.availabilityZone,
        "dbInstanceCategories": args.dbInstanceCategories,
        "dbInstanceModes": args.dbInstanceModes,
        "description": args.description,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "instanceNetworkType": args.instanceNetworkType,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "status": args.status,
        "tags": args.tags,
        "vswitchId": args.vswitchId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Instance availability zone.
     */
    availabilityZone?: string;
    /**
     * The db instance categories.
     */
    dbInstanceCategories?: string;
    /**
     * The db instance modes.
     */
    dbInstanceModes?: string;
    /**
     * The description of the instance.
     */
    description?: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * The ids list of AnalyticDB for PostgreSQL instances.
     */
    ids?: string[];
    /**
     * The network type of the instance.
     */
    instanceNetworkType?: string;
    /**
     * A regex string to apply to the instance name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the enterprise resource group to which the instance belongs.
     */
    resourceGroupId?: string;
    /**
     * The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
     */
    status?: string;
    /**
     * The tags of the instance.
     */
    tags?: {[key: string]: any};
    /**
     * The vswitch id.
     */
    vswitchId?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly availabilityZone?: string;
    readonly dbInstanceCategories?: string;
    readonly dbInstanceModes?: string;
    readonly description?: string;
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceNetworkType?: string;
    readonly instances: outputs.gpdb.GetInstancesInstance[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly resourceGroupId?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    readonly vswitchId?: string;
}

export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * Instance availability zone.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The db instance categories.
     */
    dbInstanceCategories?: pulumi.Input<string>;
    /**
     * The db instance modes.
     */
    dbInstanceModes?: pulumi.Input<string>;
    /**
     * The description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * The ids list of AnalyticDB for PostgreSQL instances.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network type of the instance.
     */
    instanceNetworkType?: pulumi.Input<string>;
    /**
     * A regex string to apply to the instance name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the enterprise resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of the instance.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The vswitch id.
     */
    vswitchId?: pulumi.Input<string>;
}