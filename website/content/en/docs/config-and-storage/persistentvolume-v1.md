---
weight: 5
title: "PersistentVolume"
description: "PersistentVolume (PV) is a storage resource provisioned by an administrator."
draft: false
collapsible: false
---
## PersistentVolume
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### PersistentVolume
PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
- **apiVersion**: v1
  
- **kind**: PersistentVolume
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolumespec" >}}">PersistentVolumeSpec</a>)
  Spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
- **status** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolumestatus" >}}">PersistentVolumeStatus</a>)
  Status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
### PersistentVolumeSpec
PersistentVolumeSpec is the specification of a persistent volume.
- **accessModes** ([]string)
  AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
- **capacity** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
- **claimRef** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
  ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
- **mountOptions** ([]string)
  A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
- **nodeAffinity** (VolumeNodeAffinity)
  NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
*VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.*
  - **nodeAffinity.required** (NodeSelector)
    Required specifies hard node constraints that must be met.
*A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.*
  - **nodeAffinity.required.nodeSelectorTerms** ([]NodeSelectorTerm), required
    Required. A list of node selector terms. The terms are ORed.
*A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.*
  - **nodeAffinity.required.nodeSelectorTerms.matchExpressions** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's labels.
  - **nodeAffinity.required.nodeSelectorTerms.matchFields** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's fields.
- **persistentVolumeReclaimPolicy** (string)
  What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
- **storageClassName** (string)
  Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
- **volumeMode** (string)
  volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
#### Local

- **hostPath** (HostPathVolumeSource)
  HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
*Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.*
  - **hostPath.path** (string), required
    Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
  - **hostPath.type** (string)
    Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
- **local** (LocalVolumeSource)
  Local represents directly-attached storage with node affinity
*Local represents directly-attached storage with node affinity (Beta feature)*
  - **local.path** (string), required
    The full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
  - **local.fsType** (string)
    Filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a fileystem if unspecified.
#### Persistent volumes

- **awsElasticBlockStore** (AWSElasticBlockStoreVolumeSource)
  AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
*Represents a Persistent Disk resource in AWS.*
*An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.*
  - **awsElasticBlockStore.volumeID** (string), required
    Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
  - **awsElasticBlockStore.fsType** (string)
    Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
  - **awsElasticBlockStore.partition** (int32)
    The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
  - **awsElasticBlockStore.readOnly** (boolean)
    Specify "true" to force and set the ReadOnly property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
- **azureDisk** (AzureDiskVolumeSource)
  AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
*AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.*
  - **azureDisk.diskName** (string), required
    The Name of the data disk in the blob storage
  - **azureDisk.diskURI** (string), required
    The URI the data disk in the blob storage
  - **azureDisk.cachingMode** (string)
    Host Caching mode: None, Read Only, Read Write.
  - **azureDisk.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
  - **azureDisk.kind** (string)
    Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
  - **azureDisk.readOnly** (boolean)
    Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
- **azureFile** (AzureFilePersistentVolumeSource)
  AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
*AzureFile represents an Azure File Service mount on the host and bind mount to the pod.*
  - **azureFile.secretName** (string), required
    the name of secret that contains Azure Storage Account Name and Key
  - **azureFile.shareName** (string), required
    Share Name
  - **azureFile.readOnly** (boolean)
    Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
  - **azureFile.secretNamespace** (string)
    the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
- **cephfs** (CephFSPersistentVolumeSource)
  CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
*Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.*
  - **cephfs.monitors** ([]string), required
    Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
  - **cephfs.path** (string)
    Optional: Used as the mounted root, rather than the full Ceph tree, default is /
  - **cephfs.readOnly** (boolean)
    Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
  - **cephfs.secretFile** (string)
    Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
  - **cephfs.secretRef** (SecretReference)
    Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **cephfs.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **cephfs.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **cephfs.user** (string)
    Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
- **cinder** (CinderPersistentVolumeSource)
  Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
*Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.*
  - **cinder.volumeID** (string), required
    volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
  - **cinder.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
  - **cinder.readOnly** (boolean)
    Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
  - **cinder.secretRef** (SecretReference)
    Optional: points to a secret object containing parameters used to connect to OpenStack.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **cinder.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **cinder.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
- **fc** (FCVolumeSource)
  FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
*Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.*
  - **fc.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
  - **fc.lun** (int32)
    Optional: FC target lun number
  - **fc.readOnly** (boolean)
    Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
  - **fc.targetWWNs** ([]string)
    Optional: FC target worldwide names (WWNs)
  - **fc.wwids** ([]string)
    Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
- **flexVolume** (FlexPersistentVolumeSource)
  FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
*FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.*
  - **flexVolume.driver** (string), required
    Driver is the name of the driver to use for this volume.
  - **flexVolume.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
  - **flexVolume.options** (map[string]string)
    Optional: Extra command options if any.
  - **flexVolume.readOnly** (boolean)
    Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
  - **flexVolume.secretRef** (SecretReference)
    Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **flexVolume.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **flexVolume.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
- **flocker** (FlockerVolumeSource)
  Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
*Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.*
  - **flocker.datasetName** (string)
    Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
  - **flocker.datasetUUID** (string)
    UUID of the dataset. This is unique identifier of a Flocker dataset
- **gcePersistentDisk** (GCEPersistentDiskVolumeSource)
  GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
*Represents a Persistent Disk resource in Google Compute Engine.*
*A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.*
  - **gcePersistentDisk.pdName** (string), required
    Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
  - **gcePersistentDisk.fsType** (string)
    Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
  - **gcePersistentDisk.partition** (int32)
    The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
  - **gcePersistentDisk.readOnly** (boolean)
    ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
- **glusterfs** (GlusterfsPersistentVolumeSource)
  Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
*Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.*
  - **glusterfs.endpoints** (string), required
    EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
  - **glusterfs.path** (string), required
    Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
  - **glusterfs.endpointsNamespace** (string)
    EndpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
  - **glusterfs.readOnly** (boolean)
    ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
- **iscsi** (ISCSIPersistentVolumeSource)
  ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
*ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.*
  - **iscsi.iqn** (string), required
    Target iSCSI Qualified Name.
  - **iscsi.lun** (int32), required
    iSCSI Target Lun number.
  - **iscsi.targetPortal** (string), required
    iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
  - **iscsi.chapAuthDiscovery** (boolean)
    whether support iSCSI Discovery CHAP authentication
  - **iscsi.chapAuthSession** (boolean)
    whether support iSCSI Session CHAP authentication
  - **iscsi.fsType** (string)
    Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
  - **iscsi.initiatorName** (string)
    Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
  - **iscsi.iscsiInterface** (string)
    iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
  - **iscsi.portals** ([]string)
    iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
  - **iscsi.readOnly** (boolean)
    ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
  - **iscsi.secretRef** (SecretReference)
    CHAP Secret for iSCSI target and initiator authentication
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **iscsi.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **iscsi.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
- **nfs** (NFSVolumeSource)
  NFS represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
*Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.*
  - **nfs.path** (string), required
    Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
  - **nfs.server** (string), required
    Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
  - **nfs.readOnly** (boolean)
    ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
- **photonPersistentDisk** (PhotonPersistentDiskVolumeSource)
  PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
*Represents a Photon Controller persistent disk resource.*
  - **photonPersistentDisk.pdID** (string), required
    ID that identifies Photon Controller persistent disk
  - **photonPersistentDisk.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
- **portworxVolume** (PortworxVolumeSource)
  PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
*PortworxVolumeSource represents a Portworx volume resource.*
  - **portworxVolume.volumeID** (string), required
    VolumeID uniquely identifies a Portworx volume
  - **portworxVolume.fsType** (string)
    FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
  - **portworxVolume.readOnly** (boolean)
    Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
- **quobyte** (QuobyteVolumeSource)
  Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
*Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.*
  - **quobyte.registry** (string), required
    Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
  - **quobyte.volume** (string), required
    Volume is a string that references an already created Quobyte volume by name.
  - **quobyte.group** (string)
    Group to map volume access to Default is no group
  - **quobyte.readOnly** (boolean)
    ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
  - **quobyte.tenant** (string)
    Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
  - **quobyte.user** (string)
    User to map volume access to Defaults to serivceaccount user
- **rbd** (RBDPersistentVolumeSource)
  RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
*Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.*
  - **rbd.image** (string), required
    The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
  - **rbd.monitors** ([]string), required
    A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
  - **rbd.fsType** (string)
    Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
  - **rbd.keyring** (string)
    Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
  - **rbd.pool** (string)
    The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
  - **rbd.readOnly** (boolean)
    ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
  - **rbd.secretRef** (SecretReference)
    SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **rbd.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **rbd.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **rbd.user** (string)
    The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
- **scaleIO** (ScaleIOPersistentVolumeSource)
  ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
*ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume*
  - **scaleIO.gateway** (string), required
    The host address of the ScaleIO API Gateway.
  - **scaleIO.secretRef** (SecretReference), required
    SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **scaleIO.secretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **scaleIO.secretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **scaleIO.system** (string), required
    The name of the storage system as configured in ScaleIO.
  - **scaleIO.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs"
  - **scaleIO.protectionDomain** (string)
    The name of the ScaleIO Protection Domain for the configured storage.
  - **scaleIO.readOnly** (boolean)
    Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
  - **scaleIO.sslEnabled** (boolean)
    Flag to enable/disable SSL communication with Gateway, default false
  - **scaleIO.storageMode** (string)
    Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
  - **scaleIO.storagePool** (string)
    The ScaleIO Storage Pool associated with the protection domain.
  - **scaleIO.volumeName** (string)
    The name of a volume already created in the ScaleIO system that is associated with this volume source.
- **storageos** (StorageOSPersistentVolumeSource)
  StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
*Represents a StorageOS persistent volume resource.*
  - **storageos.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
  - **storageos.readOnly** (boolean)
    Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
  - **storageos.secretRef** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
    SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.
  - **storageos.volumeName** (string)
    VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
  - **storageos.volumeNamespace** (string)
    VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
- **vsphereVolume** (VsphereVirtualDiskVolumeSource)
  VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
*Represents a vSphere volume resource.*
  - **vsphereVolume.volumePath** (string), required
    Path that identifies vSphere volume vmdk
  - **vsphereVolume.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
  - **vsphereVolume.storagePolicyID** (string)
    Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
  - **vsphereVolume.storagePolicyName** (string)
    Storage Policy Based Management (SPBM) profile name.
#### Beta level

- **csi** (CSIPersistentVolumeSource)
  CSI represents storage that is handled by an external CSI driver (Beta feature).
*Represents storage that is managed by an external CSI volume driver (Beta feature)*
  - **csi.driver** (string), required
    Driver is the name of the driver to use for this volume. Required.
  - **csi.volumeHandle** (string), required
    VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
  - **csi.controllerExpandSecretRef** (SecretReference)
    ControllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. This is an alpha field and requires enabling ExpandCSIVolumes feature gate. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **csi.controllerExpandSecretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **csi.controllerExpandSecretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **csi.controllerPublishSecretRef** (SecretReference)
    ControllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **csi.controllerPublishSecretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **csi.controllerPublishSecretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **csi.fsType** (string)
    Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
  - **csi.nodePublishSecretRef** (SecretReference)
    NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **csi.nodePublishSecretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **csi.nodePublishSecretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **csi.nodeStageSecretRef** (SecretReference)
    NodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
*SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace*
  - **csi.nodeStageSecretRef.name** (string)
    Name is unique within a namespace to reference a secret resource.
  - **csi.nodeStageSecretRef.namespace** (string)
    Namespace defines the space within which the secret name must be unique.
  - **csi.readOnly** (boolean)
    Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
  - **csi.volumeAttributes** (map[string]string)
    Attributes of the volume to publish.
### PersistentVolumeStatus
PersistentVolumeStatus is the current status of a persistent volume.
- **message** (string)
  A human-readable message indicating details about why the volume is in this state.
- **phase** (string)
  Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
- **reason** (string)
  Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
### PersistentVolumeList
PersistentVolumeList is a list of PersistentVolume items.
- **apiVersion**: v1
  
- **kind**: PersistentVolumeList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>), required
  List of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
### Operations
#### `get` read the specified PersistentVolume

##### HTTP Request
GET /api/v1/persistentvolumes/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
401: Unauthorized
#### `get` read status of the specified PersistentVolume

##### HTTP Request
GET /api/v1/persistentvolumes/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PersistentVolume

##### HTTP Request
GET /api/v1/persistentvolumes

##### Parameters
  - **?allowWatchBookmarks** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#allowwatchbookmarks" >}}">allowWatchBookmarks</a>
  - **?continue** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#continue" >}}">continue</a>
  - **?fieldSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldselector" >}}">fieldSelector</a>
  - **?labelSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#labelselector" >}}">labelSelector</a>
  - **?limit** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#limit" >}}">limit</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?resourceVersion** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversion" >}}">resourceVersion</a>
  - **?resourceVersionMatch** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversionmatch" >}}">resourceVersionMatch</a>
  - **?timeoutSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#timeoutseconds" >}}">timeoutSeconds</a>
  - **?watch** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#watch" >}}">watch</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolumelist" >}}">PersistentVolumeList</a>): OK
401: Unauthorized
#### `create` create a PersistentVolume

##### HTTP Request
POST /api/v1/persistentvolumes

##### Parameters
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): Created
202 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): Accepted
401: Unauthorized
#### `update` replace the specified PersistentVolume

##### HTTP Request
PUT /api/v1/persistentvolumes/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): Created
401: Unauthorized
#### `update` replace status of the specified PersistentVolume

##### HTTP Request
PUT /api/v1/persistentvolumes/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): Created
401: Unauthorized
#### `patch` partially update the specified PersistentVolume

##### HTTP Request
PATCH /api/v1/persistentvolumes/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **body** (<a href="{{< ref "/docs/common-definitions/patch-#patch" >}}">Patch</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?force** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#force" >}}">force</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified PersistentVolume

##### HTTP Request
PATCH /api/v1/persistentvolumes/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **body** (<a href="{{< ref "/docs/common-definitions/patch-#patch" >}}">Patch</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?force** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#force" >}}">force</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
401: Unauthorized
#### `delete` delete a PersistentVolume

##### HTTP Request
DELETE /api/v1/persistentvolumes/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolume
  - **body** (<a href="{{< ref "/docs/common-definitions/deleteoptions-#deleteoptions" >}}">DeleteOptions</a>)
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?gracePeriodSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#graceperiodseconds" >}}">gracePeriodSeconds</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?propagationPolicy** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#propagationpolicy" >}}">propagationPolicy</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): OK
202 (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolume" >}}">PersistentVolume</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of PersistentVolume

##### HTTP Request
DELETE /api/v1/persistentvolumes

##### Parameters
  - **body** (<a href="{{< ref "/docs/common-definitions/deleteoptions-#deleteoptions" >}}">DeleteOptions</a>)
    
  - **?continue** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#continue" >}}">continue</a>
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldselector" >}}">fieldSelector</a>
  - **?gracePeriodSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#graceperiodseconds" >}}">gracePeriodSeconds</a>
  - **?labelSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#labelselector" >}}">labelSelector</a>
  - **?limit** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#limit" >}}">limit</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?propagationPolicy** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#propagationpolicy" >}}">propagationPolicy</a>
  - **?resourceVersion** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversion" >}}">resourceVersion</a>
  - **?resourceVersionMatch** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversionmatch" >}}">resourceVersionMatch</a>
  - **?timeoutSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#timeoutseconds" >}}">timeoutSeconds</a>

##### Response
200 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): OK
401: Unauthorized
