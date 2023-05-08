// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Outputs
{

    [OutputType]
    public sealed class BucketLifecycleRuleNoncurrentVersionTransition
    {
        /// <summary>
        /// Specifies the number of days noncurrent object versions transition.
        /// </summary>
        public readonly int Days;
        /// <summary>
        /// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
        /// </summary>
        public readonly string StorageClass;

        [OutputConstructor]
        private BucketLifecycleRuleNoncurrentVersionTransition(
            int days,

            string storageClass)
        {
            Days = days;
            StorageClass = storageClass;
        }
    }
}
