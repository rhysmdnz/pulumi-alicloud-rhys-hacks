// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Route Tables owned by an Alibaba Cloud account.
 *
 * > **NOTE:** Available in 1.36.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "route-tables-datasource-example-name";
 *
 * const fooNetwork = new alicloud.vpc.Network("foo", {
 *     cidrBlock: "172.16.0.0/12",
 *     vpcName: name,
 * });
 * const fooRouteTable = new alicloud.vpc.RouteTable("foo", {
 *     description: name,
 *     routeTableName: name,
 *     vpcId: fooNetwork.id,
 * });
 * const fooRouteTables = fooRouteTable.id.apply(id => alicloud.vpc.getRouteTables({
 *     ids: [id],
 * }));
 *
 * export const routeTableIds = fooRouteTables.ids!;
 * ```
 */
export function getRouteTables(args?: GetRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTablesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:vpc/getRouteTables:getRouteTables", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "resourceGroupId": args.resourceGroupId,
        "routeTableName": args.routeTableName,
        "routerId": args.routerId,
        "routerType": args.routerType,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesArgs {
    /**
     * A list of Route Tables IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter route tables by name.
     */
    nameRegex?: string;
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The Id of resource group which route tables belongs.
     */
    resourceGroupId?: string;
    /**
     * The route table name.
     */
    routeTableName?: string;
    /**
     * The router ID.
     */
    routerId?: string;
    /**
     * The route type of route table. Valid values: `VRouter` and `VBR`.
     */
    routerType?: string;
    /**
     * The status of resource. Valid values: `Available` and `Pending`.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
    /**
     * Vpc id of the route table.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getRouteTables.
 */
export interface GetRouteTablesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Optional) A list of Route Tables IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Route Tables names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    /**
     * The Id of resource group which route tables belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * The route table name.
     */
    readonly routeTableName?: string;
    /**
     * Router Id of the route table.
     */
    readonly routerId?: string;
    /**
     * The route type.
     */
    readonly routerType?: string;
    /**
     * The status of route table.
     */
    readonly status?: string;
    /**
     * A list of Route Tables. Each element contains the following attributes:
     */
    readonly tables: outputs.vpc.GetRouteTablesTable[];
    readonly tags?: {[key: string]: any};
    readonly totalCount: number;
    /**
     * The VPC ID.
     */
    readonly vpcId?: string;
}

export function getRouteTablesOutput(args?: GetRouteTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteTablesResult> {
    return pulumi.output(args).apply(a => getRouteTables(a, opts))
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesOutputArgs {
    /**
     * A list of Route Tables IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter route tables by name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The Id of resource group which route tables belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The route table name.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * The router ID.
     */
    routerId?: pulumi.Input<string>;
    /**
     * The route type of route table. Valid values: `VRouter` and `VBR`.
     */
    routerType?: pulumi.Input<string>;
    /**
     * The status of resource. Valid values: `Available` and `Pending`.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Vpc id of the route table.
     */
    vpcId?: pulumi.Input<string>;
}