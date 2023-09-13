// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/api/metric.proto

package metric

import (
	reflect "reflect"
	sync "sync"

	api "google.golang.org/genproto/googleapis/api"
	label "google.golang.org/genproto/googleapis/api/label"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The kind of measurement. It describes how the data is reported.
// For information on setting the start time and end time based on
// the MetricKind, see [TimeInterval][google.monitoring.v3.TimeInterval].
type MetricDescriptor_MetricKind int32

const (
	// Do not use this default value.
	MetricDescriptor_METRIC_KIND_UNSPECIFIED MetricDescriptor_MetricKind = 0
	// An instantaneous measurement of a value.
	MetricDescriptor_GAUGE MetricDescriptor_MetricKind = 1
	// The change in a value during a time interval.
	MetricDescriptor_DELTA MetricDescriptor_MetricKind = 2
	// A value accumulated over a time interval.  Cumulative
	// measurements in a time series should have the same start time
	// and increasing end times, until an event resets the cumulative
	// value to zero and sets a new start time for the following
	// points.
	MetricDescriptor_CUMULATIVE MetricDescriptor_MetricKind = 3
)

// Enum value maps for MetricDescriptor_MetricKind.
var (
	MetricDescriptor_MetricKind_name = map[int32]string{
		0: "METRIC_KIND_UNSPECIFIED",
		1: "GAUGE",
		2: "DELTA",
		3: "CUMULATIVE",
	}
	MetricDescriptor_MetricKind_value = map[string]int32{
		"METRIC_KIND_UNSPECIFIED": 0,
		"GAUGE":                   1,
		"DELTA":                   2,
		"CUMULATIVE":              3,
	}
)

func (x MetricDescriptor_MetricKind) Enum() *MetricDescriptor_MetricKind {
	p := new(MetricDescriptor_MetricKind)
	*p = x
	return p
}

func (x MetricDescriptor_MetricKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricDescriptor_MetricKind) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_metric_proto_enumTypes[0].Descriptor()
}

func (MetricDescriptor_MetricKind) Type() protoreflect.EnumType {
	return &file_google_api_metric_proto_enumTypes[0]
}

func (x MetricDescriptor_MetricKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricDescriptor_MetricKind.Descriptor instead.
func (MetricDescriptor_MetricKind) EnumDescriptor() ([]byte, []int) {
	return file_google_api_metric_proto_rawDescGZIP(), []int{0, 0}
}

// The value type of a metric.
type MetricDescriptor_ValueType int32

const (
	// Do not use this default value.
	MetricDescriptor_VALUE_TYPE_UNSPECIFIED MetricDescriptor_ValueType = 0
	// The value is a boolean.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_BOOL MetricDescriptor_ValueType = 1
	// The value is a signed 64-bit integer.
	MetricDescriptor_INT64 MetricDescriptor_ValueType = 2
	// The value is a double precision floating point number.
	MetricDescriptor_DOUBLE MetricDescriptor_ValueType = 3
	// The value is a text string.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_STRING MetricDescriptor_ValueType = 4
	// The value is a [`Distribution`][google.api.Distribution].
	MetricDescriptor_DISTRIBUTION MetricDescriptor_ValueType = 5
	// The value is money.
	MetricDescriptor_MONEY MetricDescriptor_ValueType = 6
)

// Enum value maps for MetricDescriptor_ValueType.
var (
	MetricDescriptor_ValueType_name = map[int32]string{
		0: "VALUE_TYPE_UNSPECIFIED",
		1: "BOOL",
		2: "INT64",
		3: "DOUBLE",
		4: "STRING",
		5: "DISTRIBUTION",
		6: "MONEY",
	}
	MetricDescriptor_ValueType_value = map[string]int32{
		"VALUE_TYPE_UNSPECIFIED": 0,
		"BOOL":                   1,
		"INT64":                  2,
		"DOUBLE":                 3,
		"STRING":                 4,
		"DISTRIBUTION":           5,
		"MONEY":                  6,
	}
)

func (x MetricDescriptor_ValueType) Enum() *MetricDescriptor_ValueType {
	p := new(MetricDescriptor_ValueType)
	*p = x
	return p
}

func (x MetricDescriptor_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricDescriptor_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_metric_proto_enumTypes[1].Descriptor()
}

func (MetricDescriptor_ValueType) Type() protoreflect.EnumType {
	return &file_google_api_metric_proto_enumTypes[1]
}

func (x MetricDescriptor_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricDescriptor_ValueType.Descriptor instead.
func (MetricDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_google_api_metric_proto_rawDescGZIP(), []int{0, 1}
}

// Defines a metric type and its schema. Once a metric descriptor is created,
// deleting or altering it stops data collection and makes the metric type's
// existing data unusable.
type MetricDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the metric descriptor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric type, including its DNS name prefix. The type is not
	// URL-encoded. All user-defined metric types have the DNS name
	// `custom.googleapis.com` or `external.googleapis.com`. Metric types should
	// use a natural hierarchical grouping. For example:
	//
	//	"custom.googleapis.com/invoice/paid/amount"
	//	"external.googleapis.com/prometheus/up"
	//	"appengine.googleapis.com/http/server/response_latencies"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	// The set of labels that can be used to describe a specific
	// instance of this metric type. For example, the
	// `appengine.googleapis.com/http/server/response_latencies` metric
	// type has a label for the HTTP response code, `response_code`, so
	// you can look at latencies for successful responses or just
	// for responses that failed.
	Labels []*label.LabelDescriptor `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of `metric_kind` and `value_type` might not be supported.
	MetricKind MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of `metric_kind` and `value_type` might not be supported.
	ValueType MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The units in which the metric value is reported. It is only applicable
	// if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`. The `unit`
	// defines the representation of the stored metric values.
	//
	// Different systems might scale the values to be more easily displayed (so a
	// value of `0.02kBy` _might_ be displayed as `20By`, and a value of
	// `3523kBy` _might_ be displayed as `3.5MBy`). However, if the `unit` is
	// `kBy`, then the value of the metric is always in thousands of bytes, no
	// matter how it might be displayed.
	//
	// If you want a custom metric to record the exact number of CPU-seconds used
	// by a job, you can create an `INT64 CUMULATIVE` metric whose `unit` is
	// `s{CPU}` (or equivalently `1s{CPU}` or just `s`). If the job uses 12,005
	// CPU-seconds, then the value is written as `12005`.
	//
	// Alternatively, if you want a custom metric to record data in a more
	// granular way, you can create a `DOUBLE CUMULATIVE` metric whose `unit` is
	// `ks{CPU}`, and then write the value `12.005` (which is `12005/1000`),
	// or use `Kis{CPU}` and write `11.723` (which is `12005/1024`).
	//
	// The supported units are a subset of [The Unified Code for Units of
	// Measure](https://unitsofmeasure.org/ucum.html) standard:
	//
	// **Basic units (UNIT)**
	//
	// * `bit`   bit
	// * `By`    byte
	// * `s`     second
	// * `min`   minute
	// * `h`     hour
	// * `d`     day
	// * `1`     dimensionless
	//
	// **Prefixes (PREFIX)**
	//
	// * `k`     kilo    (10^3)
	// * `M`     mega    (10^6)
	// * `G`     giga    (10^9)
	// * `T`     tera    (10^12)
	// * `P`     peta    (10^15)
	// * `E`     exa     (10^18)
	// * `Z`     zetta   (10^21)
	// * `Y`     yotta   (10^24)
	//
	// * `m`     milli   (10^-3)
	// * `u`     micro   (10^-6)
	// * `n`     nano    (10^-9)
	// * `p`     pico    (10^-12)
	// * `f`     femto   (10^-15)
	// * `a`     atto    (10^-18)
	// * `z`     zepto   (10^-21)
	// * `y`     yocto   (10^-24)
	//
	// * `Ki`    kibi    (2^10)
	// * `Mi`    mebi    (2^20)
	// * `Gi`    gibi    (2^30)
	// * `Ti`    tebi    (2^40)
	// * `Pi`    pebi    (2^50)
	//
	// **Grammar**
	//
	// The grammar also includes these connectors:
	//
	//   - `/`    division or ratio (as an infix operator). For examples,
	//     `kBy/{email}` or `MiBy/10ms` (although you should almost never
	//     have `/s` in a metric `unit`; rates should always be computed at
	//     query time from the underlying cumulative or delta value).
	//   - `.`    multiplication or composition (as an infix operator). For
	//     examples, `GBy.d` or `k{watt}.h`.
	//
	// The grammar for a unit is as follows:
	//
	//	Expression = Component { "." Component } { "/" Component } ;
	//
	//	Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ]
	//	          | Annotation
	//	          | "1"
	//	          ;
	//
	//	Annotation = "{" NAME "}" ;
	//
	// Notes:
	//
	//   - `Annotation` is just a comment if it follows a `UNIT`. If the annotation
	//     is used alone, then the unit is equivalent to `1`. For examples,
	//     `{request}/s == 1/s`, `By{transmitted}/s == By/s`.
	//   - `NAME` is a sequence of non-blank printable ASCII characters not
	//     containing `{` or `}`.
	//   - `1` represents a unitary [dimensionless
	//     unit](https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such
	//     as in `1/s`. It is typically used when none of the basic units are
	//     appropriate. For example, "new users per day" can be represented as
	//     `1/d` or `{new-users}/d` (and a metric value `5` would mean "5 new
	//     users). Alternatively, "thousands of page views per day" would be
	//     represented as `1000/d` or `k1/d` or `k{page_views}/d` (and a metric
	//     value of `5.3` would mean "5300 page views per day").
	//   - `%` represents dimensionless value of 1/100, and annotates values giving
	//     a percentage (so the metric values are typically in the range of 0..100,
	//     and a metric value `3` means "3 percent").
	//   - `10^2.%` indicates a metric contains a ratio, typically in the range
	//     0..1, that will be multiplied by 100 and displayed as a percentage
	//     (so a metric value `0.03` means "3 percent").
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	// A detailed description of the metric, which can be used in documentation.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// A concise name for the metric, which can be displayed in user interfaces.
	// Use sentence case without an ending period, for example "Request count".
	// This field is optional but it is recommended to be set for any metrics
	// associated with user-visible concepts, such as Quota.
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Metadata which can be used to guide usage of the metric.
	Metadata *MetricDescriptor_MetricDescriptorMetadata `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Optional. The launch stage of the metric definition.
	LaunchStage api.LaunchStage `protobuf:"varint,12,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
	// Read-only. If present, then a [time
	// series][google.monitoring.v3.TimeSeries], which is identified partially by
	// a metric type and a
	// [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor], that
	// is associated with this metric type can only be associated with one of the
	// monitored resource types listed here.
	MonitoredResourceTypes []string `protobuf:"bytes,13,rep,name=monitored_resource_types,json=monitoredResourceTypes,proto3" json:"monitored_resource_types,omitempty"`
}

func (x *MetricDescriptor) Reset() {
	*x = MetricDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDescriptor) ProtoMessage() {}

func (x *MetricDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDescriptor.ProtoReflect.Descriptor instead.
func (*MetricDescriptor) Descriptor() ([]byte, []int) {
	return file_google_api_metric_proto_rawDescGZIP(), []int{0}
}

func (x *MetricDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricDescriptor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MetricDescriptor) GetLabels() []*label.LabelDescriptor {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricDescriptor) GetMetricKind() MetricDescriptor_MetricKind {
	if x != nil {
		return x.MetricKind
	}
	return MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (x *MetricDescriptor) GetValueType() MetricDescriptor_ValueType {
	if x != nil {
		return x.ValueType
	}
	return MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (x *MetricDescriptor) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *MetricDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MetricDescriptor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *MetricDescriptor) GetMetadata() *MetricDescriptor_MetricDescriptorMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MetricDescriptor) GetLaunchStage() api.LaunchStage {
	if x != nil {
		return x.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

func (x *MetricDescriptor) GetMonitoredResourceTypes() []string {
	if x != nil {
		return x.MonitoredResourceTypes
	}
	return nil
}

// A specific metric, identified by specifying values for all of the
// labels of a [`MetricDescriptor`][google.api.MetricDescriptor].
type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An existing metric type, see
	// [google.api.MetricDescriptor][google.api.MetricDescriptor]. For example,
	// `custom.googleapis.com/invoice/paid/amount`.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The set of label values that uniquely identify this metric. All
	// labels listed in the `MetricDescriptor` must be assigned values.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_google_api_metric_proto_rawDescGZIP(), []int{1}
}

func (x *Metric) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Metric) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Additional annotations that can be used to guide the usage of a metric.
type MetricDescriptor_MetricDescriptorMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated. Must use the
	// [MetricDescriptor.launch_stage][google.api.MetricDescriptor.launch_stage]
	// instead.
	//
	// Deprecated: Do not use.
	LaunchStage api.LaunchStage `protobuf:"varint,1,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
	// The sampling period of metric data points. For metrics which are written
	// periodically, consecutive data points are stored at this time interval,
	// excluding data loss due to errors. Metrics with a higher granularity have
	// a smaller sampling period.
	SamplePeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=sample_period,json=samplePeriod,proto3" json:"sample_period,omitempty"`
	// The delay of data points caused by ingestion. Data points older than this
	// age are guaranteed to be ingested and available to be read, excluding
	// data loss due to errors.
	IngestDelay *durationpb.Duration `protobuf:"bytes,3,opt,name=ingest_delay,json=ingestDelay,proto3" json:"ingest_delay,omitempty"`
}

func (x *MetricDescriptor_MetricDescriptorMetadata) Reset() {
	*x = MetricDescriptor_MetricDescriptorMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDescriptor_MetricDescriptorMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDescriptor_MetricDescriptorMetadata) ProtoMessage() {}

func (x *MetricDescriptor_MetricDescriptorMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDescriptor_MetricDescriptorMetadata.ProtoReflect.Descriptor instead.
func (*MetricDescriptor_MetricDescriptorMetadata) Descriptor() ([]byte, []int) {
	return file_google_api_metric_proto_rawDescGZIP(), []int{0, 0}
}

// Deprecated: Do not use.
func (x *MetricDescriptor_MetricDescriptorMetadata) GetLaunchStage() api.LaunchStage {
	if x != nil {
		return x.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

func (x *MetricDescriptor_MetricDescriptorMetadata) GetSamplePeriod() *durationpb.Duration {
	if x != nil {
		return x.SamplePeriod
	}
	return nil
}

func (x *MetricDescriptor_MetricDescriptorMetadata) GetIngestDelay() *durationpb.Duration {
	if x != nil {
		return x.IngestDelay
	}
	return nil
}

var File_google_api_metric_proto protoreflect.FileDescriptor

var file_google_api_metric_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x07, 0x0a,
	0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x48,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x6c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0xd8, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e,
	0x0a, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3c,
	0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x4f, 0x0a, 0x0a,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45,
	0x54, 0x52, 0x49, 0x43, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x55, 0x4d, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x22, 0x71, 0x0a,
	0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x10, 0x06,
	0x22, 0x8f, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x5f, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x42, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0xa2, 0x02, 0x04, 0x47,
	0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_metric_proto_rawDescOnce sync.Once
	file_google_api_metric_proto_rawDescData = file_google_api_metric_proto_rawDesc
)

func file_google_api_metric_proto_rawDescGZIP() []byte {
	file_google_api_metric_proto_rawDescOnce.Do(func() {
		file_google_api_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_metric_proto_rawDescData)
	})
	return file_google_api_metric_proto_rawDescData
}

var file_google_api_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_api_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_api_metric_proto_goTypes = []interface{}{
	(MetricDescriptor_MetricKind)(0),                  // 0: google.api.MetricDescriptor.MetricKind
	(MetricDescriptor_ValueType)(0),                   // 1: google.api.MetricDescriptor.ValueType
	(*MetricDescriptor)(nil),                          // 2: google.api.MetricDescriptor
	(*Metric)(nil),                                    // 3: google.api.Metric
	(*MetricDescriptor_MetricDescriptorMetadata)(nil), // 4: google.api.MetricDescriptor.MetricDescriptorMetadata
	nil,                           // 5: google.api.Metric.LabelsEntry
	(*label.LabelDescriptor)(nil), // 6: google.api.LabelDescriptor
	(api.LaunchStage)(0),          // 7: google.api.LaunchStage
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_google_api_metric_proto_depIdxs = []int32{
	6, // 0: google.api.MetricDescriptor.labels:type_name -> google.api.LabelDescriptor
	0, // 1: google.api.MetricDescriptor.metric_kind:type_name -> google.api.MetricDescriptor.MetricKind
	1, // 2: google.api.MetricDescriptor.value_type:type_name -> google.api.MetricDescriptor.ValueType
	4, // 3: google.api.MetricDescriptor.metadata:type_name -> google.api.MetricDescriptor.MetricDescriptorMetadata
	7, // 4: google.api.MetricDescriptor.launch_stage:type_name -> google.api.LaunchStage
	5, // 5: google.api.Metric.labels:type_name -> google.api.Metric.LabelsEntry
	7, // 6: google.api.MetricDescriptor.MetricDescriptorMetadata.launch_stage:type_name -> google.api.LaunchStage
	8, // 7: google.api.MetricDescriptor.MetricDescriptorMetadata.sample_period:type_name -> google.protobuf.Duration
	8, // 8: google.api.MetricDescriptor.MetricDescriptorMetadata.ingest_delay:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_api_metric_proto_init() }
func file_google_api_metric_proto_init() {
	if File_google_api_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDescriptor_MetricDescriptorMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_metric_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_metric_proto_goTypes,
		DependencyIndexes: file_google_api_metric_proto_depIdxs,
		EnumInfos:         file_google_api_metric_proto_enumTypes,
		MessageInfos:      file_google_api_metric_proto_msgTypes,
	}.Build()
	File_google_api_metric_proto = out.File
	file_google_api_metric_proto_rawDesc = nil
	file_google_api_metric_proto_goTypes = nil
	file_google_api_metric_proto_depIdxs = nil
}
