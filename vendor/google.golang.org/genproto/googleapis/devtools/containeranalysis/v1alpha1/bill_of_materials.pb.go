// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1alpha1/bill_of_materials.proto

/*
Package containeranalysis is a generated protocol buffer package.

It is generated from these files:
	google/devtools/containeranalysis/v1alpha1/bill_of_materials.proto
	google/devtools/containeranalysis/v1alpha1/containeranalysis.proto
	google/devtools/containeranalysis/v1alpha1/image_basis.proto
	google/devtools/containeranalysis/v1alpha1/package_vulnerability.proto
	google/devtools/containeranalysis/v1alpha1/provenance.proto
	google/devtools/containeranalysis/v1alpha1/source_context.proto

It has these top-level messages:
	PackageManager
	Occurrence
	Resource
	Note
	Deployable
	Discovery
	BuildType
	BuildSignature
	PgpSignedAttestation
	AttestationAuthority
	BuildDetails
	ScanConfig
	GetOccurrenceRequest
	ListOccurrencesRequest
	ListOccurrencesResponse
	DeleteOccurrenceRequest
	CreateOccurrenceRequest
	UpdateOccurrenceRequest
	GetNoteRequest
	GetOccurrenceNoteRequest
	ListNotesRequest
	ListNotesResponse
	DeleteNoteRequest
	CreateNoteRequest
	UpdateNoteRequest
	ListNoteOccurrencesRequest
	ListNoteOccurrencesResponse
	CreateOperationRequest
	UpdateOperationRequest
	OperationMetadata
	GetVulnzOccurrencesSummaryRequest
	GetVulnzOccurrencesSummaryResponse
	GetScanConfigRequest
	ListScanConfigsRequest
	ListScanConfigsResponse
	UpdateScanConfigRequest
	DockerImage
	VulnerabilityType
	BuildProvenance
	Source
	FileHashes
	Hash
	StorageSource
	RepoSource
	Command
	Artifact
	SourceContext
	AliasContext
	CloudRepoSourceContext
	GerritSourceContext
	GitSourceContext
	RepoId
	ProjectRepoId
*/
package containeranalysis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Instruction set architectures supported by various package managers.
type PackageManager_Architecture int32

const (
	// Unknown architecture
	PackageManager_ARCHITECTURE_UNSPECIFIED PackageManager_Architecture = 0
	// X86 architecture
	PackageManager_X86 PackageManager_Architecture = 1
	// X64 architecture
	PackageManager_X64 PackageManager_Architecture = 2
)

var PackageManager_Architecture_name = map[int32]string{
	0: "ARCHITECTURE_UNSPECIFIED",
	1: "X86",
	2: "X64",
}
var PackageManager_Architecture_value = map[string]int32{
	"ARCHITECTURE_UNSPECIFIED": 0,
	"X86": 1,
	"X64": 2,
}

func (x PackageManager_Architecture) String() string {
	return proto.EnumName(PackageManager_Architecture_name, int32(x))
}
func (PackageManager_Architecture) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// PackageManager provides metadata about available / installed packages.
type PackageManager struct {
}

func (m *PackageManager) Reset()                    { *m = PackageManager{} }
func (m *PackageManager) String() string            { return proto.CompactTextString(m) }
func (*PackageManager) ProtoMessage()               {}
func (*PackageManager) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// This represents a particular channel of distribution for a given package.
// e.g. Debian's jessie-backports dpkg mirror
type PackageManager_Distribution struct {
	// The cpe_uri in [cpe format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution
	// channel were built
	Architecture PackageManager_Architecture `protobuf:"varint,2,opt,name=architecture,enum=google.devtools.containeranalysis.v1alpha1.PackageManager_Architecture" json:"architecture,omitempty"`
	// The latest available version of this package in
	// this distribution channel.
	LatestVersion *VulnerabilityType_Version `protobuf:"bytes,3,opt,name=latest_version,json=latestVersion" json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `protobuf:"bytes,4,opt,name=maintainer" json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
}

func (m *PackageManager_Distribution) Reset()                    { *m = PackageManager_Distribution{} }
func (m *PackageManager_Distribution) String() string            { return proto.CompactTextString(m) }
func (*PackageManager_Distribution) ProtoMessage()               {}
func (*PackageManager_Distribution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *PackageManager_Distribution) GetCpeUri() string {
	if m != nil {
		return m.CpeUri
	}
	return ""
}

func (m *PackageManager_Distribution) GetArchitecture() PackageManager_Architecture {
	if m != nil {
		return m.Architecture
	}
	return PackageManager_ARCHITECTURE_UNSPECIFIED
}

func (m *PackageManager_Distribution) GetLatestVersion() *VulnerabilityType_Version {
	if m != nil {
		return m.LatestVersion
	}
	return nil
}

func (m *PackageManager_Distribution) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *PackageManager_Distribution) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PackageManager_Distribution) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// An occurrence of a particular package installation found within a
// system's filesystem.
// e.g. glibc was found in /var/lib/dpkg/status
type PackageManager_Location struct {
	// The cpe_uri in [cpe format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri" json:"cpe_uri,omitempty"`
	// The version installed at this location.
	Version *VulnerabilityType_Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *PackageManager_Location) Reset()                    { *m = PackageManager_Location{} }
func (m *PackageManager_Location) String() string            { return proto.CompactTextString(m) }
func (*PackageManager_Location) ProtoMessage()               {}
func (*PackageManager_Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *PackageManager_Location) GetCpeUri() string {
	if m != nil {
		return m.CpeUri
	}
	return ""
}

func (m *PackageManager_Location) GetVersion() *VulnerabilityType_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *PackageManager_Location) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// This represents a particular package that is distributed over
// various channels.
// e.g. glibc (aka libc6) is distributed by many, at various versions.
type PackageManager_Package struct {
	// The name of the package.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The various channels by which a package is distributed.
	Distribution []*PackageManager_Distribution `protobuf:"bytes,10,rep,name=distribution" json:"distribution,omitempty"`
}

func (m *PackageManager_Package) Reset()                    { *m = PackageManager_Package{} }
func (m *PackageManager_Package) String() string            { return proto.CompactTextString(m) }
func (*PackageManager_Package) ProtoMessage()               {}
func (*PackageManager_Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *PackageManager_Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageManager_Package) GetDistribution() []*PackageManager_Distribution {
	if m != nil {
		return m.Distribution
	}
	return nil
}

// This represents how a particular software package may be installed on
// a system.
type PackageManager_Installation struct {
	// Output only. The name of the installed package.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// All of the places within the filesystem versions of this package
	// have been found.
	Location []*PackageManager_Location `protobuf:"bytes,2,rep,name=location" json:"location,omitempty"`
}

func (m *PackageManager_Installation) Reset()                    { *m = PackageManager_Installation{} }
func (m *PackageManager_Installation) String() string            { return proto.CompactTextString(m) }
func (*PackageManager_Installation) ProtoMessage()               {}
func (*PackageManager_Installation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *PackageManager_Installation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageManager_Installation) GetLocation() []*PackageManager_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageManager)(nil), "google.devtools.containeranalysis.v1alpha1.PackageManager")
	proto.RegisterType((*PackageManager_Distribution)(nil), "google.devtools.containeranalysis.v1alpha1.PackageManager.Distribution")
	proto.RegisterType((*PackageManager_Location)(nil), "google.devtools.containeranalysis.v1alpha1.PackageManager.Location")
	proto.RegisterType((*PackageManager_Package)(nil), "google.devtools.containeranalysis.v1alpha1.PackageManager.Package")
	proto.RegisterType((*PackageManager_Installation)(nil), "google.devtools.containeranalysis.v1alpha1.PackageManager.Installation")
	proto.RegisterEnum("google.devtools.containeranalysis.v1alpha1.PackageManager_Architecture", PackageManager_Architecture_name, PackageManager_Architecture_value)
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1alpha1/bill_of_materials.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x8a, 0xd3, 0x4e,
	0x14, 0xc6, 0xff, 0x49, 0x97, 0x76, 0xf7, 0xb4, 0xff, 0x52, 0xe6, 0xc6, 0x10, 0x16, 0x29, 0x0b,
	0x42, 0xf1, 0x22, 0x61, 0x57, 0x59, 0x04, 0x41, 0xe8, 0x76, 0xbb, 0x6b, 0x41, 0xa5, 0xc4, 0x76,
	0x11, 0xbd, 0x08, 0xa7, 0xe9, 0x98, 0x0e, 0x3b, 0x9d, 0x09, 0x93, 0x49, 0xa1, 0xd7, 0xde, 0x89,
	0x0f, 0xe0, 0xb5, 0x0f, 0xa5, 0xaf, 0x23, 0x99, 0x24, 0x92, 0xb2, 0x2a, 0xbb, 0xac, 0x77, 0x27,
	0xf3, 0x85, 0xdf, 0xf9, 0xce, 0x77, 0x66, 0xe0, 0x2c, 0x96, 0x32, 0xe6, 0xd4, 0x5f, 0xd2, 0x8d,
	0x96, 0x92, 0xa7, 0x7e, 0x24, 0x85, 0x46, 0x26, 0xa8, 0x42, 0x81, 0x7c, 0x9b, 0xb2, 0xd4, 0xdf,
	0x1c, 0x23, 0x4f, 0x56, 0x78, 0xec, 0x2f, 0x18, 0xe7, 0xa1, 0xfc, 0x18, 0xae, 0x51, 0x53, 0xc5,
	0x90, 0xa7, 0x5e, 0xa2, 0xa4, 0x96, 0xe4, 0x71, 0xc1, 0xf0, 0x2a, 0x86, 0x77, 0x83, 0xe1, 0x55,
	0x0c, 0xf7, 0xb0, 0xec, 0x87, 0x09, 0xf3, 0x51, 0x08, 0xa9, 0x51, 0x33, 0x29, 0x4a, 0x92, 0x7b,
	0x71, 0x07, 0x37, 0x09, 0x46, 0xd7, 0x18, 0xd3, 0x70, 0x93, 0xf1, 0x5c, 0x5f, 0x30, 0xce, 0xf4,
	0xb6, 0xe0, 0x1c, 0xfd, 0x68, 0x42, 0x77, 0x5a, 0xe8, 0xaf, 0x51, 0x60, 0x4c, 0x95, 0xfb, 0xdd,
	0x86, 0xce, 0x39, 0x4b, 0xb5, 0x62, 0x8b, 0x2c, 0x6f, 0x49, 0x1e, 0x40, 0x2b, 0x4a, 0x68, 0x98,
	0x29, 0xe6, 0x58, 0x7d, 0x6b, 0x70, 0x10, 0x34, 0xa3, 0x84, 0xce, 0x15, 0x23, 0xd7, 0xd0, 0x41,
	0x15, 0xad, 0x98, 0xa6, 0x91, 0xce, 0x14, 0x75, 0xec, 0xbe, 0x35, 0xe8, 0x9e, 0x5c, 0x7a, 0xb7,
	0x9f, 0xd2, 0xdb, 0xed, 0xed, 0x0d, 0x6b, 0xb8, 0x60, 0x07, 0x4e, 0x38, 0x74, 0x39, 0x6a, 0x9a,
	0xea, 0x70, 0x43, 0x55, 0xca, 0xa4, 0x70, 0x1a, 0x7d, 0x6b, 0xd0, 0x3e, 0x19, 0xdf, 0xa5, 0xdd,
	0x55, 0x3d, 0x82, 0xd9, 0x36, 0xa1, 0xde, 0x55, 0x01, 0x0b, 0xfe, 0x2f, 0xe0, 0xe5, 0x27, 0x79,
	0x08, 0xb0, 0x46, 0x56, 0x72, 0x9c, 0x3d, 0x33, 0x76, 0xed, 0x84, 0xf4, 0xa0, 0x91, 0x29, 0xee,
	0x34, 0x8d, 0x90, 0x97, 0xa4, 0x0f, 0xed, 0x25, 0x4d, 0x23, 0xc5, 0x92, 0x3c, 0x34, 0xa7, 0x65,
	0x94, 0xfa, 0x91, 0xfb, 0xd5, 0x82, 0xfd, 0x57, 0x32, 0xc2, 0xbf, 0x87, 0x1a, 0x42, 0xab, 0x1a,
	0xd0, 0xfe, 0x97, 0x03, 0x56, 0x54, 0x42, 0x60, 0x2f, 0x41, 0xbd, 0x32, 0xf1, 0x1d, 0x04, 0xa6,
	0x76, 0x3f, 0x5b, 0xd0, 0x2a, 0x57, 0x91, 0xeb, 0x02, 0xd7, 0xb4, 0xb4, 0x65, 0xea, 0x7c, 0xd3,
	0xcb, 0xda, 0x95, 0x70, 0xa0, 0xdf, 0x18, 0xb4, 0xef, 0xb5, 0xe9, 0xfa, 0x0d, 0x0b, 0x76, 0xe0,
	0xee, 0x27, 0x0b, 0x3a, 0x13, 0x91, 0x6a, 0xe4, 0xbc, 0xc8, 0xea, 0x77, 0x8e, 0x42, 0xd8, 0xe7,
	0x65, 0x96, 0x8e, 0x6d, 0xdc, 0x8c, 0xee, 0xe1, 0xa6, 0x5a, 0x4b, 0xf0, 0x0b, 0x7a, 0xf4, 0x02,
	0x3a, 0xf5, 0xdb, 0x48, 0x0e, 0xc1, 0x19, 0x06, 0xa3, 0x97, 0x93, 0xd9, 0x78, 0x34, 0x9b, 0x07,
	0xe3, 0x70, 0xfe, 0xe6, 0xed, 0x74, 0x3c, 0x9a, 0x5c, 0x4c, 0xc6, 0xe7, 0xbd, 0xff, 0x48, 0x0b,
	0x1a, 0xef, 0x9e, 0x9d, 0xf6, 0x2c, 0x53, 0x9c, 0x3e, 0xed, 0xd9, 0x67, 0x5f, 0x2c, 0x78, 0x14,
	0xc9, 0x75, 0x65, 0xea, 0xcf, 0x5e, 0xa6, 0xd6, 0xfb, 0x0f, 0xe5, 0x4f, 0xb1, 0xe4, 0x28, 0x62,
	0x4f, 0xaa, 0xd8, 0x8f, 0xa9, 0x30, 0x2f, 0xd4, 0x2f, 0x24, 0x4c, 0x58, 0x7a, 0x9b, 0xc7, 0xfe,
	0xfc, 0x86, 0xf4, 0xcd, 0x6e, 0x5c, 0x8e, 0x86, 0x8b, 0xa6, 0xa1, 0x3d, 0xf9, 0x19, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0x4f, 0xa4, 0x56, 0xc7, 0x04, 0x00, 0x00,
}
