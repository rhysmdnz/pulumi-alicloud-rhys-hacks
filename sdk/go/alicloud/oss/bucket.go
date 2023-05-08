// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a oss bucket and set its attribution.
//
// > **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
//
// ## Example Usage
//
// # Private Bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-acl", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-acl"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Static Website
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-website", &oss.BucketArgs{
//				Bucket: pulumi.String("bucket-170309-website"),
//				Website: &oss.BucketWebsiteArgs{
//					ErrorDocument: pulumi.String("error.html"),
//					IndexDocument: pulumi.String("index.html"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Enable Logging
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-target", &oss.BucketArgs{
//				Bucket: pulumi.String("bucket-170309-acl"),
//				Acl:    pulumi.String("public-read"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-logging", &oss.BucketArgs{
//				Bucket: pulumi.String("bucket-170309-logging"),
//				Logging: &oss.BucketLoggingArgs{
//					TargetBucket: bucket_target.ID(),
//					TargetPrefix: pulumi.String("log/"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Referer configuration
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-referer", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-referer"),
//				RefererConfig: &oss.BucketRefererConfigArgs{
//					AllowEmpty: pulumi.Bool(false),
//					Referers: pulumi.StringArray{
//						pulumi.String("http://www.aliyun.com"),
//						pulumi.String("https://www.aliyun.com"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set lifecycle rule
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-lifecycle", &oss.BucketArgs{
//				Acl:    pulumi.String("public-read"),
//				Bucket: pulumi.String("bucket-170309-lifecycle"),
//				LifecycleRules: oss.BucketLifecycleRuleArray{
//					&oss.BucketLifecycleRuleArgs{
//						AbortMultipartUploads: oss.BucketLifecycleRuleAbortMultipartUploadArray{
//							&oss.BucketLifecycleRuleAbortMultipartUploadArgs{
//								Days: pulumi.Int(128),
//							},
//						},
//						Enabled: pulumi.Bool(true),
//						Id:      pulumi.String("rule-abort-multipart-upload"),
//						Prefix:  pulumi.String("path3/"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-versioning-lifecycle", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-lifecycle"),
//				LifecycleRules: oss.BucketLifecycleRuleArray{
//					&oss.BucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Expirations: oss.BucketLifecycleRuleExpirationArray{
//							&oss.BucketLifecycleRuleExpirationArgs{
//								ExpiredObjectDeleteMarker: pulumi.Bool(true),
//							},
//						},
//						Id: pulumi.String("rule-versioning"),
//						NoncurrentVersionExpirations: oss.BucketLifecycleRuleNoncurrentVersionExpirationArray{
//							&oss.BucketLifecycleRuleNoncurrentVersionExpirationArgs{
//								Days: pulumi.Int(240),
//							},
//						},
//						NoncurrentVersionTransitions: oss.BucketLifecycleRuleNoncurrentVersionTransitionArray{
//							&oss.BucketLifecycleRuleNoncurrentVersionTransitionArgs{
//								Days:         pulumi.Int(180),
//								StorageClass: pulumi.String("Archive"),
//							},
//							&oss.BucketLifecycleRuleNoncurrentVersionTransitionArgs{
//								Days:         pulumi.Int(60),
//								StorageClass: pulumi.String("IA"),
//							},
//						},
//						Prefix: pulumi.String("path1/"),
//					},
//				},
//				Versioning: &oss.BucketVersioningArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set bucket policy
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-policy", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-policy"),
//				Policy: pulumi.String(fmt.Sprintf(`  {"Statement":
//	      [{"Action":
//	          ["oss:PutObject", "oss:GetObject", "oss:DeleteBucket"],
//	        "Effect":"Allow",
//	        "Resource":
//	            ["acs:oss:*:*:*"]}],
//	   "Version":"1"}
//
// `)),
//
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # IA Bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-storageclass", &oss.BucketArgs{
//				Bucket:       pulumi.String("bucket-170309-storageclass"),
//				StorageClass: pulumi.String("IA"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set bucket server-side encryption rule
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-sserule", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-sserule"),
//				ServerSideEncryptionRule: &oss.BucketServerSideEncryptionRuleArgs{
//					KmsMasterKeyId: pulumi.String("your kms key id"),
//					SseAlgorithm:   pulumi.String("KMS"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set bucket tags
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-tags", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-tags"),
//				Tags: pulumi.AnyMap{
//					"key1": pulumi.Any("value1"),
//					"key2": pulumi.Any("value2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Enable bucket versioning
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-versioning", &oss.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("bucket-170309-versioning"),
//				Versioning: &oss.BucketVersioningArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set bucket redundancy type
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-redundancytype", &oss.BucketArgs{
//				Bucket:         pulumi.String("bucket_name"),
//				RedundancyType: pulumi.String("ZRS"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Set bucket accelerate configuration
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucket(ctx, "bucket-accelerate", &oss.BucketArgs{
//				Bucket: pulumi.String("bucket_name"),
//				TransferAcceleration: &oss.BucketTransferAccelerationArgs{
//					Enabled: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OSS bucket can be imported using the bucket name, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
//
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrOutput `pulumi:"acl"`
	Bucket pulumi.StringPtrOutput `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringOutput `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringOutput `pulumi:"intranetEndpoint"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrOutput `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrOutput `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrOutput `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrOutput `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A transfer acceleration status of a bucket (documented below).
	TransferAcceleration BucketTransferAccelerationPtrOutput `pulumi:"transferAcceleration"`
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrOutput `pulumi:"versioning"`
	// A website object(documented below).
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Bucket
	err := ctx.RegisterResource("alicloud:oss/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("alicloud:oss/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate *string `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint *string `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint *string `pulumi:"intranetEndpoint"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging *BucketLogging `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner *string `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A transfer acceleration status of a bucket (documented below).
	TransferAcceleration *BucketTransferAcceleration `pulumi:"transferAcceleration"`
	// A state of versioning (documented below).
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayInput
	// The creation date of the bucket.
	CreationDate pulumi.StringPtrInput
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// The bucket owner.
	Owner pulumi.StringPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrInput
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapInput
	// A transfer acceleration status of a bucket (documented below).
	TransferAcceleration BucketTransferAccelerationPtrInput
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrInput
	// A website object(documented below).
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging *BucketLogging `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A transfer acceleration status of a bucket (documented below).
	TransferAcceleration *BucketTransferAcceleration `pulumi:"transferAcceleration"`
	// A state of versioning (documented below).
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrInput
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapInput
	// A transfer acceleration status of a bucket (documented below).
	TransferAcceleration BucketTransferAccelerationPtrInput
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrInput
	// A website object(documented below).
	Website BucketWebsitePtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
func (o BucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

func (o BucketOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Bucket }).(pulumi.StringPtrOutput)
}

// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
func (o BucketOutput) CorsRules() BucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketCorsRuleArrayOutput { return v.CorsRules }).(BucketCorsRuleArrayOutput)
}

// The creation date of the bucket.
func (o BucketOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The extranet access endpoint of the bucket.
func (o BucketOutput) ExtranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ExtranetEndpoint }).(pulumi.StringOutput)
}

// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
func (o BucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// The intranet access endpoint of the bucket.
func (o BucketOutput) IntranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.IntranetEndpoint }).(pulumi.StringOutput)
}

// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
func (o BucketOutput) LifecycleRules() BucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(BucketLifecycleRuleArrayOutput)
}

// The location of the bucket.
func (o BucketOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
func (o BucketOutput) Logging() BucketLoggingPtrOutput {
	return o.ApplyT(func(v *Bucket) BucketLoggingPtrOutput { return v.Logging }).(BucketLoggingPtrOutput)
}

// The flag of using logging enable container. Defaults true.
//
// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
func (o BucketOutput) LoggingIsenable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.LoggingIsenable }).(pulumi.BoolPtrOutput)
}

// The bucket owner.
func (o BucketOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
func (o BucketOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
func (o BucketOutput) RedundancyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.RedundancyType }).(pulumi.StringPtrOutput)
}

// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
func (o BucketOutput) RefererConfig() BucketRefererConfigPtrOutput {
	return o.ApplyT(func(v *Bucket) BucketRefererConfigPtrOutput { return v.RefererConfig }).(BucketRefererConfigPtrOutput)
}

// A configuration of server-side encryption (documented below).
func (o BucketOutput) ServerSideEncryptionRule() BucketServerSideEncryptionRulePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketServerSideEncryptionRulePtrOutput { return v.ServerSideEncryptionRule }).(BucketServerSideEncryptionRulePtrOutput)
}

// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`, `CodeArchive`. ColdArchive is available in 1.203.0+.
func (o BucketOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
func (o BucketOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Bucket) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// A transfer acceleration status of a bucket (documented below).
func (o BucketOutput) TransferAcceleration() BucketTransferAccelerationPtrOutput {
	return o.ApplyT(func(v *Bucket) BucketTransferAccelerationPtrOutput { return v.TransferAcceleration }).(BucketTransferAccelerationPtrOutput)
}

// A state of versioning (documented below).
func (o BucketOutput) Versioning() BucketVersioningPtrOutput {
	return o.ApplyT(func(v *Bucket) BucketVersioningPtrOutput { return v.Versioning }).(BucketVersioningPtrOutput)
}

// A website object(documented below).
func (o BucketOutput) Website() BucketWebsitePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketWebsitePtrOutput { return v.Website }).(BucketWebsitePtrOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
