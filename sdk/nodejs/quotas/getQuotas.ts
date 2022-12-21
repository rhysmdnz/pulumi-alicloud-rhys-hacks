// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Quotas Quotas of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.115.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.quotas.getQuotas({
 *     productCode: "ecs",
 *     nameRegex: "专有宿主机总数量上限",
 * });
 * export const firstQuotasQuotaId = example.then(example => example.quotas?[0]?.id);
 * ```
 */
export function getQuotas(args: GetQuotasArgs, opts?: pulumi.InvokeOptions): Promise<GetQuotasResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:quotas/getQuotas:getQuotas", {
        "dimensions": args.dimensions,
        "groupCode": args.groupCode,
        "keyWord": args.keyWord,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "productCode": args.productCode,
        "quotaActionCode": args.quotaActionCode,
        "quotaCategory": args.quotaCategory,
        "sortField": args.sortField,
        "sortOrder": args.sortOrder,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuotas.
 */
export interface GetQuotasArgs {
    /**
     * The dimensions.
     */
    dimensions?: inputs.quotas.GetQuotasDimension[];
    /**
     * The group code.
     */
    groupCode?: string;
    /**
     * The key word.
     */
    keyWord?: string;
    /**
     * A regex string to filter results by Quota name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The product code.
     */
    productCode: string;
    /**
     * The quota action code.
     */
    quotaActionCode?: string;
    /**
     * The category of quota. Valid Values: `FlowControl` and `CommonQuota`.
     */
    quotaCategory?: string;
    /**
     * Cloud service ECS specification quota supports setting sorting fields. Valid Values: `TIME`, `TOTAL` and `RESERVED`.
     */
    sortField?: string;
    /**
     * Ranking of cloud service ECS specification quota support. Valid Values: `Ascending` and `Descending`.
     */
    sortOrder?: string;
}

/**
 * A collection of values returned by getQuotas.
 */
export interface GetQuotasResult {
    readonly dimensions?: outputs.quotas.GetQuotasDimension[];
    readonly groupCode?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly keyWord?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly productCode: string;
    readonly quotaActionCode?: string;
    readonly quotaCategory?: string;
    readonly quotas: outputs.quotas.GetQuotasQuota[];
    readonly sortField?: string;
    readonly sortOrder?: string;
}

export function getQuotasOutput(args: GetQuotasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQuotasResult> {
    return pulumi.output(args).apply(a => getQuotas(a, opts))
}

/**
 * A collection of arguments for invoking getQuotas.
 */
export interface GetQuotasOutputArgs {
    /**
     * The dimensions.
     */
    dimensions?: pulumi.Input<pulumi.Input<inputs.quotas.GetQuotasDimensionArgs>[]>;
    /**
     * The group code.
     */
    groupCode?: pulumi.Input<string>;
    /**
     * The key word.
     */
    keyWord?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Quota name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The product code.
     */
    productCode: pulumi.Input<string>;
    /**
     * The quota action code.
     */
    quotaActionCode?: pulumi.Input<string>;
    /**
     * The category of quota. Valid Values: `FlowControl` and `CommonQuota`.
     */
    quotaCategory?: pulumi.Input<string>;
    /**
     * Cloud service ECS specification quota supports setting sorting fields. Valid Values: `TIME`, `TOTAL` and `RESERVED`.
     */
    sortField?: pulumi.Input<string>;
    /**
     * Ranking of cloud service ECS specification quota support. Valid Values: `Ascending` and `Descending`.
     */
    sortOrder?: pulumi.Input<string>;
}