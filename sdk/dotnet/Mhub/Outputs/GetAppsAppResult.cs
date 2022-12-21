// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mhub.Outputs
{

    [OutputType]
    public sealed class GetAppsAppResult
    {
        /// <summary>
        /// Application AppKey, which uniquely identifies an application when requested by the interface
        /// </summary>
        public readonly string AppKey;
        /// <summary>
        /// The Name of the App.
        /// </summary>
        public readonly string AppName;
        /// <summary>
        /// iOS application ID. Required when creating an iOS app. **NOTE:** Either `bundle_id` or `package_name` must be set.
        /// </summary>
        public readonly string BundleId;
        /// <summary>
        /// The CreateTime of the App.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Base64 string of picture.
        /// </summary>
        public readonly string EncodedIcon;
        /// <summary>
        /// The ID of the App.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
        /// </summary>
        public readonly string IndustryId;
        /// <summary>
        /// Android App package name.  **NOTE:** Either `bundle_id` or `package_name` must be set.
        /// </summary>
        public readonly string PackageName;
        /// <summary>
        /// The ID of the Product.
        /// </summary>
        public readonly string ProductId;
        /// <summary>
        /// The type of the App. Valid values: `Android` and `iOS`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppsAppResult(
            string appKey,

            string appName,

            string bundleId,

            string createTime,

            string encodedIcon,

            string id,

            string industryId,

            string packageName,

            string productId,

            string type)
        {
            AppKey = appKey;
            AppName = appName;
            BundleId = bundleId;
            CreateTime = createTime;
            EncodedIcon = encodedIcon;
            Id = id;
            IndustryId = industryId;
            PackageName = packageName;
            ProductId = productId;
            Type = type;
        }
    }
}