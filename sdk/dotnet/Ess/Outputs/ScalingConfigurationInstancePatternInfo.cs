// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Outputs
{

    [OutputType]
    public sealed class ScalingConfigurationInstancePatternInfo
    {
        public readonly int? Cores;
        public readonly string? InstanceFamilyLevel;
        public readonly double? MaxPrice;
        public readonly double? Memory;

        [OutputConstructor]
        private ScalingConfigurationInstancePatternInfo(
            int? cores,

            string? instanceFamilyLevel,

            double? maxPrice,

            double? memory)
        {
            Cores = cores;
            InstanceFamilyLevel = instanceFamilyLevel;
            MaxPrice = maxPrice;
            Memory = memory;
        }
    }
}