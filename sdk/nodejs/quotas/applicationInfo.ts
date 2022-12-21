// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class ApplicationInfo extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationInfo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationInfoState, opts?: pulumi.CustomResourceOptions): ApplicationInfo {
        return new ApplicationInfo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:quotas/applicationInfo:ApplicationInfo';

    /**
     * Returns true if the given object is an instance of ApplicationInfo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationInfo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationInfo.__pulumiType;
    }

    public /*out*/ readonly approveValue!: pulumi.Output<string>;
    public readonly auditMode!: pulumi.Output<string | undefined>;
    public /*out*/ readonly auditReason!: pulumi.Output<string>;
    public readonly desireValue!: pulumi.Output<number>;
    public readonly dimensions!: pulumi.Output<outputs.quotas.ApplicationInfoDimension[] | undefined>;
    public /*out*/ readonly effectiveTime!: pulumi.Output<string>;
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    public readonly noticeType!: pulumi.Output<number | undefined>;
    public readonly productCode!: pulumi.Output<string>;
    public readonly quotaActionCode!: pulumi.Output<string>;
    public readonly quotaCategory!: pulumi.Output<string | undefined>;
    public /*out*/ readonly quotaDescription!: pulumi.Output<string>;
    public /*out*/ readonly quotaName!: pulumi.Output<string>;
    public /*out*/ readonly quotaUnit!: pulumi.Output<string>;
    public readonly reason!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ApplicationInfo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationInfoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationInfoArgs | ApplicationInfoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationInfoState | undefined;
            resourceInputs["approveValue"] = state ? state.approveValue : undefined;
            resourceInputs["auditMode"] = state ? state.auditMode : undefined;
            resourceInputs["auditReason"] = state ? state.auditReason : undefined;
            resourceInputs["desireValue"] = state ? state.desireValue : undefined;
            resourceInputs["dimensions"] = state ? state.dimensions : undefined;
            resourceInputs["effectiveTime"] = state ? state.effectiveTime : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["noticeType"] = state ? state.noticeType : undefined;
            resourceInputs["productCode"] = state ? state.productCode : undefined;
            resourceInputs["quotaActionCode"] = state ? state.quotaActionCode : undefined;
            resourceInputs["quotaCategory"] = state ? state.quotaCategory : undefined;
            resourceInputs["quotaDescription"] = state ? state.quotaDescription : undefined;
            resourceInputs["quotaName"] = state ? state.quotaName : undefined;
            resourceInputs["quotaUnit"] = state ? state.quotaUnit : undefined;
            resourceInputs["reason"] = state ? state.reason : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ApplicationInfoArgs | undefined;
            if ((!args || args.desireValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desireValue'");
            }
            if ((!args || args.productCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productCode'");
            }
            if ((!args || args.quotaActionCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'quotaActionCode'");
            }
            if ((!args || args.reason === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reason'");
            }
            resourceInputs["auditMode"] = args ? args.auditMode : undefined;
            resourceInputs["desireValue"] = args ? args.desireValue : undefined;
            resourceInputs["dimensions"] = args ? args.dimensions : undefined;
            resourceInputs["noticeType"] = args ? args.noticeType : undefined;
            resourceInputs["productCode"] = args ? args.productCode : undefined;
            resourceInputs["quotaActionCode"] = args ? args.quotaActionCode : undefined;
            resourceInputs["quotaCategory"] = args ? args.quotaCategory : undefined;
            resourceInputs["reason"] = args ? args.reason : undefined;
            resourceInputs["approveValue"] = undefined /*out*/;
            resourceInputs["auditReason"] = undefined /*out*/;
            resourceInputs["effectiveTime"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["quotaDescription"] = undefined /*out*/;
            resourceInputs["quotaName"] = undefined /*out*/;
            resourceInputs["quotaUnit"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationInfo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationInfo resources.
 */
export interface ApplicationInfoState {
    approveValue?: pulumi.Input<string>;
    auditMode?: pulumi.Input<string>;
    auditReason?: pulumi.Input<string>;
    desireValue?: pulumi.Input<number>;
    dimensions?: pulumi.Input<pulumi.Input<inputs.quotas.ApplicationInfoDimension>[]>;
    effectiveTime?: pulumi.Input<string>;
    expireTime?: pulumi.Input<string>;
    noticeType?: pulumi.Input<number>;
    productCode?: pulumi.Input<string>;
    quotaActionCode?: pulumi.Input<string>;
    quotaCategory?: pulumi.Input<string>;
    quotaDescription?: pulumi.Input<string>;
    quotaName?: pulumi.Input<string>;
    quotaUnit?: pulumi.Input<string>;
    reason?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationInfo resource.
 */
export interface ApplicationInfoArgs {
    auditMode?: pulumi.Input<string>;
    desireValue: pulumi.Input<number>;
    dimensions?: pulumi.Input<pulumi.Input<inputs.quotas.ApplicationInfoDimension>[]>;
    noticeType?: pulumi.Input<number>;
    productCode: pulumi.Input<string>;
    quotaActionCode: pulumi.Input<string>;
    quotaCategory?: pulumi.Input<string>;
    reason: pulumi.Input<string>;
}