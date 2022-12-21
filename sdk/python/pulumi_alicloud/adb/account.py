# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 db_cluster_id: pulumi.Input[str],
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[str] account_description: Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if kms_encrypted_password is not None:
            pulumi.set(__self__, "kms_encrypted_password", kms_encrypted_password)
        if kms_encryption_context is not None:
            pulumi.set(__self__, "kms_encryption_context", kms_encryption_context)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> pulumi.Input[str]:
        """
        The Id of cluster in which account belongs.
        """
        return pulumi.get(self, "db_cluster_id")

    @db_cluster_id.setter
    def db_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_cluster_id", value)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> Optional[pulumi.Input[str]]:
        """
        An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @kms_encrypted_password.setter
    def kms_encrypted_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encrypted_password", value)

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @kms_encryption_context.setter
    def kms_encryption_context(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "kms_encryption_context", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] account_description: Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if db_cluster_id is not None:
            pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if kms_encrypted_password is not None:
            pulumi.set(__self__, "kms_encrypted_password", kms_encrypted_password)
        if kms_encryption_context is not None:
            pulumi.set(__self__, "kms_encryption_context", kms_encryption_context)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of cluster in which account belongs.
        """
        return pulumi.get(self, "db_cluster_id")

    @db_cluster_id.setter
    def db_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_cluster_id", value)

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> Optional[pulumi.Input[str]]:
        """
        An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @kms_encrypted_password.setter
    def kms_encrypted_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encrypted_password", value)

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @kms_encryption_context.setter
    def kms_encryption_context(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "kms_encryption_context", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a [ADB](https://www.alibabacloud.com/help/product/92664.htm) account resource and used to manage databases.

        > **NOTE:** Available in v1.71.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        creation = config.get("creation")
        if creation is None:
            creation = "ADB"
        name = config.get("name")
        if name is None:
            name = "adbaccountmysql"
        default_zones = alicloud.get_zones(available_resource_creation=creation)
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.zones[0].id)
        cluster = alicloud.adb.Cluster("cluster",
            db_cluster_version="3.0",
            db_cluster_category="Cluster",
            db_node_class="C8",
            db_node_count=2,
            db_node_storage=200,
            pay_type="PostPaid",
            vswitch_id=default_switch.id,
            description=name)
        account = alicloud.adb.Account("account",
            db_cluster_id=cluster.id,
            account_name="tftestnormal",
            account_password="Test12345",
            account_description=name)
        ```

        ## Import

        ADB account can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:adb/account:Account example am-12345:tf_account
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a [ADB](https://www.alibabacloud.com/help/product/92664.htm) account resource and used to manage databases.

        > **NOTE:** Available in v1.71.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        creation = config.get("creation")
        if creation is None:
            creation = "ADB"
        name = config.get("name")
        if name is None:
            name = "adbaccountmysql"
        default_zones = alicloud.get_zones(available_resource_creation=creation)
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.zones[0].id)
        cluster = alicloud.adb.Cluster("cluster",
            db_cluster_version="3.0",
            db_cluster_category="Cluster",
            db_node_class="C8",
            db_node_count=2,
            db_node_storage=200,
            pay_type="PostPaid",
            vswitch_id=default_switch.id,
            description=name)
        account = alicloud.adb.Account("account",
            db_cluster_id=cluster.id,
            account_name="tftestnormal",
            account_password="Test12345",
            account_description=name)
        ```

        ## Import

        ADB account can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:adb/account:Account example am-12345:tf_account
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["account_description"] = account_description
            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["account_password"] = account_password
            if db_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_cluster_id'")
            __props__.__dict__["db_cluster_id"] = db_cluster_id
            __props__.__dict__["kms_encrypted_password"] = kms_encrypted_password
            __props__.__dict__["kms_encryption_context"] = kms_encryption_context
        super(Account, __self__).__init__(
            'alicloud:adb/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_description: Optional[pulumi.Input[str]] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            account_password: Optional[pulumi.Input[str]] = None,
            db_cluster_id: Optional[pulumi.Input[str]] = None,
            kms_encrypted_password: Optional[pulumi.Input[str]] = None,
            kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["account_description"] = account_description
        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["account_password"] = account_password
        __props__.__dict__["db_cluster_id"] = db_cluster_id
        __props__.__dict__["kms_encrypted_password"] = kms_encrypted_password
        __props__.__dict__["kms_encryption_context"] = kms_encryption_context
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> pulumi.Output[Optional[str]]:
        """
        Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        """
        return pulumi.get(self, "account_description")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Output[Optional[str]]:
        """
        Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        """
        return pulumi.get(self, "account_password")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> pulumi.Output[str]:
        """
        The Id of cluster in which account belongs.
        """
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> pulumi.Output[Optional[str]]:
        """
        An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")
