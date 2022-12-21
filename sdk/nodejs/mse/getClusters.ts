// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of MSE Clusters in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.94.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = pulumi.output(alicloud.mse.getClusters({
 *     ids: ["mse-cn-0d9xxxx"],
 *     status: "INIT_SUCCESS",
 * }));
 *
 * export const clusterId = example.clusters[0].id;
 * ```
 */
export function getClusters(args?: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:mse/getClusters:getClusters", {
        "clusterAliasName": args.clusterAliasName,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "requestPars": args.requestPars,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    /**
     * The alias name of MSE Cluster.
     */
    clusterAliasName?: string;
    enableDetails?: boolean;
    /**
     * A list of MSE Cluster ids.
     */
    ids?: string[];
    /**
     * A regex string to filter the results by the cluster alias name.
     */
    nameRegex?: string;
    outputFile?: string;
    requestPars?: string;
    /**
     * The status of MSE Cluster. Valid: `DESTROY_FAILED`, `DESTROY_ING`, `DESTROY_SUCCESS`, `INIT_FAILED`, `INIT_ING`, `INIT_SUCCESS`, `INIT_TIME_OUT`, `RESTART_FAILED`, `RESTART_ING`, `RESTART_SUCCESS`, `SCALE_FAILED`, `SCALE_ING`, `SCALE_SUCCESS`
     */
    status?: string;
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    readonly clusterAliasName?: string;
    /**
     * A list of MSE Clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.mse.GetClustersCluster[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of MSE Cluster ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of MSE Cluster names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly requestPars?: string;
    /**
     * The status of MSE Cluster.
     */
    readonly status?: string;
}

export function getClustersOutput(args?: GetClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClustersResult> {
    return pulumi.output(args).apply(a => getClusters(a, opts))
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersOutputArgs {
    /**
     * The alias name of MSE Cluster.
     */
    clusterAliasName?: pulumi.Input<string>;
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of MSE Cluster ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter the results by the cluster alias name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    requestPars?: pulumi.Input<string>;
    /**
     * The status of MSE Cluster. Valid: `DESTROY_FAILED`, `DESTROY_ING`, `DESTROY_SUCCESS`, `INIT_FAILED`, `INIT_ING`, `INIT_SUCCESS`, `INIT_TIME_OUT`, `RESTART_FAILED`, `RESTART_ING`, `RESTART_SUCCESS`, `SCALE_FAILED`, `SCALE_ING`, `SCALE_SUCCESS`
     */
    status?: pulumi.Input<string>;
}