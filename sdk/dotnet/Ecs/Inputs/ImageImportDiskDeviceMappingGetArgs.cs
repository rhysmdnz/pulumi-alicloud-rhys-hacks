// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Inputs
{

    public sealed class ImageImportDiskDeviceMappingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of disk N in the custom image.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Resolution size. You must ensure that the system disk space ≥ file system space. Ranges: When n = 1, the system disk: 5 ~ 500GiB, When n = 2 ~ 17, that is, data disk: 5 ~ 1000GiB, When temporary is introduced, the system automatically detects the size, which is subject to the detection result.
        /// </summary>
        [Input("diskImageSize")]
        public Input<int>? DiskImageSize { get; set; }

        /// <summary>
        /// Image format. Value range: When the `RAW`, `VHD`, `qcow2` is imported into the image, the system automatically detects the image format, whichever comes first.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Save the exported OSS bucket.
        /// </summary>
        [Input("ossBucket")]
        public Input<string>? OssBucket { get; set; }

        /// <summary>
        /// The file name of your OSS Object.
        /// </summary>
        [Input("ossObject")]
        public Input<string>? OssObject { get; set; }

        public ImageImportDiskDeviceMappingGetArgs()
        {
        }
        public static new ImageImportDiskDeviceMappingGetArgs Empty => new ImageImportDiskDeviceMappingGetArgs();
    }
}