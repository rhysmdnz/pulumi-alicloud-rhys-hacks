// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleState, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/schedule:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly desiredCapacity!: pulumi.Output<number | undefined>;
    public readonly launchExpirationTime!: pulumi.Output<number | undefined>;
    public readonly launchTime!: pulumi.Output<string | undefined>;
    public readonly maxValue!: pulumi.Output<number | undefined>;
    public readonly minValue!: pulumi.Output<number | undefined>;
    public readonly recurrenceEndTime!: pulumi.Output<string>;
    public readonly recurrenceType!: pulumi.Output<string>;
    public readonly recurrenceValue!: pulumi.Output<string>;
    public readonly scalingGroupId!: pulumi.Output<string>;
    public readonly scheduledAction!: pulumi.Output<string | undefined>;
    public readonly scheduledTaskName!: pulumi.Output<string | undefined>;
    public readonly taskEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs | ScheduleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScheduleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desiredCapacity"] = state ? state.desiredCapacity : undefined;
            resourceInputs["launchExpirationTime"] = state ? state.launchExpirationTime : undefined;
            resourceInputs["launchTime"] = state ? state.launchTime : undefined;
            resourceInputs["maxValue"] = state ? state.maxValue : undefined;
            resourceInputs["minValue"] = state ? state.minValue : undefined;
            resourceInputs["recurrenceEndTime"] = state ? state.recurrenceEndTime : undefined;
            resourceInputs["recurrenceType"] = state ? state.recurrenceType : undefined;
            resourceInputs["recurrenceValue"] = state ? state.recurrenceValue : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            resourceInputs["scheduledAction"] = state ? state.scheduledAction : undefined;
            resourceInputs["scheduledTaskName"] = state ? state.scheduledTaskName : undefined;
            resourceInputs["taskEnabled"] = state ? state.taskEnabled : undefined;
        } else {
            const args = argsOrState as ScheduleArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desiredCapacity"] = args ? args.desiredCapacity : undefined;
            resourceInputs["launchExpirationTime"] = args ? args.launchExpirationTime : undefined;
            resourceInputs["launchTime"] = args ? args.launchTime : undefined;
            resourceInputs["maxValue"] = args ? args.maxValue : undefined;
            resourceInputs["minValue"] = args ? args.minValue : undefined;
            resourceInputs["recurrenceEndTime"] = args ? args.recurrenceEndTime : undefined;
            resourceInputs["recurrenceType"] = args ? args.recurrenceType : undefined;
            resourceInputs["recurrenceValue"] = args ? args.recurrenceValue : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["scheduledAction"] = args ? args.scheduledAction : undefined;
            resourceInputs["scheduledTaskName"] = args ? args.scheduledTaskName : undefined;
            resourceInputs["taskEnabled"] = args ? args.taskEnabled : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Schedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Schedule resources.
 */
export interface ScheduleState {
    description?: pulumi.Input<string>;
    desiredCapacity?: pulumi.Input<number>;
    launchExpirationTime?: pulumi.Input<number>;
    launchTime?: pulumi.Input<string>;
    maxValue?: pulumi.Input<number>;
    minValue?: pulumi.Input<number>;
    recurrenceEndTime?: pulumi.Input<string>;
    recurrenceType?: pulumi.Input<string>;
    recurrenceValue?: pulumi.Input<string>;
    scalingGroupId?: pulumi.Input<string>;
    scheduledAction?: pulumi.Input<string>;
    scheduledTaskName?: pulumi.Input<string>;
    taskEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    description?: pulumi.Input<string>;
    desiredCapacity?: pulumi.Input<number>;
    launchExpirationTime?: pulumi.Input<number>;
    launchTime?: pulumi.Input<string>;
    maxValue?: pulumi.Input<number>;
    minValue?: pulumi.Input<number>;
    recurrenceEndTime?: pulumi.Input<string>;
    recurrenceType?: pulumi.Input<string>;
    recurrenceValue?: pulumi.Input<string>;
    scalingGroupId?: pulumi.Input<string>;
    scheduledAction?: pulumi.Input<string>;
    scheduledTaskName?: pulumi.Input<string>;
    taskEnabled?: pulumi.Input<boolean>;
}