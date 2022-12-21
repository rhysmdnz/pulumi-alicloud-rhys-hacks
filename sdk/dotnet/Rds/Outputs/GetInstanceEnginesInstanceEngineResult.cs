// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds.Outputs
{

    [OutputType]
    public sealed class GetInstanceEnginesInstanceEngineResult
    {
        /// <summary>
        /// DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "PPAS", "MariaDB". If not set, it will match all of engines.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        public readonly string EngineVersion;
        /// <summary>
        /// A list of Zone to launch the DB instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceEnginesInstanceEngineZoneIdResult> ZoneIds;

        [OutputConstructor]
        private GetInstanceEnginesInstanceEngineResult(
            string category,

            string engine,

            string engineVersion,

            ImmutableArray<Outputs.GetInstanceEnginesInstanceEngineZoneIdResult> zoneIds)
        {
            Category = category;
            Engine = engine;
            EngineVersion = engineVersion;
            ZoneIds = zoneIds;
        }
    }
}