// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/human_annotation_config.proto

package datalabeling

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StringAggregationType int32

const (
	StringAggregationType_STRING_AGGREGATION_TYPE_UNSPECIFIED StringAggregationType = 0
	// Majority vote to aggregate answers.
	StringAggregationType_MAJORITY_VOTE StringAggregationType = 1
	// Unanimous answers will be adopted.
	StringAggregationType_UNANIMOUS_VOTE StringAggregationType = 2
	// Preserve all answers by crowd compute.
	StringAggregationType_NO_AGGREGATION StringAggregationType = 3
)

var StringAggregationType_name = map[int32]string{
	0: "STRING_AGGREGATION_TYPE_UNSPECIFIED",
	1: "MAJORITY_VOTE",
	2: "UNANIMOUS_VOTE",
	3: "NO_AGGREGATION",
}

var StringAggregationType_value = map[string]int32{
	"STRING_AGGREGATION_TYPE_UNSPECIFIED": 0,
	"MAJORITY_VOTE":                       1,
	"UNANIMOUS_VOTE":                      2,
	"NO_AGGREGATION":                      3,
}

func (x StringAggregationType) String() string {
	return proto.EnumName(StringAggregationType_name, int32(x))
}

func (StringAggregationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{0}
}

// Configuration for how human labeling task should be done.
type HumanAnnotationConfig struct {
	// Required except for LabelAudio case. Instruction resource name.
	Instruction string `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// Required. A human-readable name for AnnotatedDataset defined by
	// users. Maximum of 64 characters
	// .
	AnnotatedDatasetDisplayName string `protobuf:"bytes,2,opt,name=annotated_dataset_display_name,json=annotatedDatasetDisplayName,proto3" json:"annotated_dataset_display_name,omitempty"`
	// Optional. A human-readable description for AnnotatedDataset.
	// The description can be up to 10000 characters long.
	AnnotatedDatasetDescription string `protobuf:"bytes,3,opt,name=annotated_dataset_description,json=annotatedDatasetDescription,proto3" json:"annotated_dataset_description,omitempty"`
	// Optional. A human-readable label used to logically group labeling tasks.
	// This string must match the regular expression `[a-zA-Z\\d_-]{0,128}`.
	LabelGroup string `protobuf:"bytes,4,opt,name=label_group,json=labelGroup,proto3" json:"label_group,omitempty"`
	// Optional. The Language of this question, as a
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	// Default value is en-US.
	// Only need to set this when task is language related. For example, French
	// text classification or Chinese audio transcription.
	LanguageCode string `protobuf:"bytes,5,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Replication of questions. Each question will be sent to up to
	// this number of contributors to label. Aggregated answers will be returned.
	// Default is set to 1.
	// For image related labeling, valid values are 1, 3, 5.
	ReplicaCount int32 `protobuf:"varint,6,opt,name=replica_count,json=replicaCount,proto3" json:"replica_count,omitempty"`
	// Optional. Maximum duration for contributors to answer a question. Default
	// is 1800 seconds.
	QuestionDuration *duration.Duration `protobuf:"bytes,7,opt,name=question_duration,json=questionDuration,proto3" json:"question_duration,omitempty"`
	// Optional. If you want your own labeling contributors to manage and work on
	// this labeling request, you can set these contributors here. We will give
	// them access to the question types in crowdcompute. Note that these
	// emails must be registered in crowdcompute worker UI:
	// https://crowd-compute.appspot.com/
	ContributorEmails []string `protobuf:"bytes,9,rep,name=contributor_emails,json=contributorEmails,proto3" json:"contributor_emails,omitempty"`
	// Email of the user who started the labeling task and should be notified by
	// email. If empty no notification will be sent.
	UserEmailAddress     string   `protobuf:"bytes,10,opt,name=user_email_address,json=userEmailAddress,proto3" json:"user_email_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HumanAnnotationConfig) Reset()         { *m = HumanAnnotationConfig{} }
func (m *HumanAnnotationConfig) String() string { return proto.CompactTextString(m) }
func (*HumanAnnotationConfig) ProtoMessage()    {}
func (*HumanAnnotationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{0}
}

func (m *HumanAnnotationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HumanAnnotationConfig.Unmarshal(m, b)
}
func (m *HumanAnnotationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HumanAnnotationConfig.Marshal(b, m, deterministic)
}
func (m *HumanAnnotationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HumanAnnotationConfig.Merge(m, src)
}
func (m *HumanAnnotationConfig) XXX_Size() int {
	return xxx_messageInfo_HumanAnnotationConfig.Size(m)
}
func (m *HumanAnnotationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HumanAnnotationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HumanAnnotationConfig proto.InternalMessageInfo

func (m *HumanAnnotationConfig) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *HumanAnnotationConfig) GetAnnotatedDatasetDisplayName() string {
	if m != nil {
		return m.AnnotatedDatasetDisplayName
	}
	return ""
}

func (m *HumanAnnotationConfig) GetAnnotatedDatasetDescription() string {
	if m != nil {
		return m.AnnotatedDatasetDescription
	}
	return ""
}

func (m *HumanAnnotationConfig) GetLabelGroup() string {
	if m != nil {
		return m.LabelGroup
	}
	return ""
}

func (m *HumanAnnotationConfig) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *HumanAnnotationConfig) GetReplicaCount() int32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *HumanAnnotationConfig) GetQuestionDuration() *duration.Duration {
	if m != nil {
		return m.QuestionDuration
	}
	return nil
}

func (m *HumanAnnotationConfig) GetContributorEmails() []string {
	if m != nil {
		return m.ContributorEmails
	}
	return nil
}

func (m *HumanAnnotationConfig) GetUserEmailAddress() string {
	if m != nil {
		return m.UserEmailAddress
	}
	return ""
}

// Config for image classification human labeling task.
type ImageClassificationConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. If allow_multi_label is true, contributors are able to choose
	// multiple labels for one image.
	AllowMultiLabel bool `protobuf:"varint,2,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	// Optional. The type of how to aggregate answers.
	AnswerAggregationType StringAggregationType `protobuf:"varint,3,opt,name=answer_aggregation_type,json=answerAggregationType,proto3,enum=google.cloud.datalabeling.v1beta1.StringAggregationType" json:"answer_aggregation_type,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *ImageClassificationConfig) Reset()         { *m = ImageClassificationConfig{} }
func (m *ImageClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationConfig) ProtoMessage()    {}
func (*ImageClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{1}
}

func (m *ImageClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationConfig.Unmarshal(m, b)
}
func (m *ImageClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationConfig.Marshal(b, m, deterministic)
}
func (m *ImageClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationConfig.Merge(m, src)
}
func (m *ImageClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationConfig.Size(m)
}
func (m *ImageClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationConfig proto.InternalMessageInfo

func (m *ImageClassificationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *ImageClassificationConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

func (m *ImageClassificationConfig) GetAnswerAggregationType() StringAggregationType {
	if m != nil {
		return m.AnswerAggregationType
	}
	return StringAggregationType_STRING_AGGREGATION_TYPE_UNSPECIFIED
}

// Config for image bounding poly (and bounding box) human labeling task.
type BoundingPolyConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Instruction message showed on contributors UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoundingPolyConfig) Reset()         { *m = BoundingPolyConfig{} }
func (m *BoundingPolyConfig) String() string { return proto.CompactTextString(m) }
func (*BoundingPolyConfig) ProtoMessage()    {}
func (*BoundingPolyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{2}
}

func (m *BoundingPolyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPolyConfig.Unmarshal(m, b)
}
func (m *BoundingPolyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPolyConfig.Marshal(b, m, deterministic)
}
func (m *BoundingPolyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPolyConfig.Merge(m, src)
}
func (m *BoundingPolyConfig) XXX_Size() int {
	return xxx_messageInfo_BoundingPolyConfig.Size(m)
}
func (m *BoundingPolyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPolyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPolyConfig proto.InternalMessageInfo

func (m *BoundingPolyConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *BoundingPolyConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for image polyline human labeling task.
type PolylineConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Instruction message showed on contributors UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolylineConfig) Reset()         { *m = PolylineConfig{} }
func (m *PolylineConfig) String() string { return proto.CompactTextString(m) }
func (*PolylineConfig) ProtoMessage()    {}
func (*PolylineConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{3}
}

func (m *PolylineConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolylineConfig.Unmarshal(m, b)
}
func (m *PolylineConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolylineConfig.Marshal(b, m, deterministic)
}
func (m *PolylineConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolylineConfig.Merge(m, src)
}
func (m *PolylineConfig) XXX_Size() int {
	return xxx_messageInfo_PolylineConfig.Size(m)
}
func (m *PolylineConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PolylineConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PolylineConfig proto.InternalMessageInfo

func (m *PolylineConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *PolylineConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for image segmentation
type SegmentationConfig struct {
	// Required. Annotation spec set resource name. format:
	// projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Instruction message showed on labelers UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentationConfig) Reset()         { *m = SegmentationConfig{} }
func (m *SegmentationConfig) String() string { return proto.CompactTextString(m) }
func (*SegmentationConfig) ProtoMessage()    {}
func (*SegmentationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{4}
}

func (m *SegmentationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentationConfig.Unmarshal(m, b)
}
func (m *SegmentationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentationConfig.Marshal(b, m, deterministic)
}
func (m *SegmentationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentationConfig.Merge(m, src)
}
func (m *SegmentationConfig) XXX_Size() int {
	return xxx_messageInfo_SegmentationConfig.Size(m)
}
func (m *SegmentationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentationConfig proto.InternalMessageInfo

func (m *SegmentationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *SegmentationConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for video classification human labeling task.
// Currently two types of video classification are supported:
// 1. Assign labels on the entire video.
// 2. Split the video into multiple video clips based on camera shot, and
// assign labels on each video clip.
type VideoClassificationConfig struct {
	// Required. The list of annotation spec set configs.
	// Since watching a video clip takes much longer time than an image, we
	// support label with multiple AnnotationSpecSet at the same time. Labels
	// in each AnnotationSpecSet will be shown in a group to contributors.
	// Contributors can select one or more (depending on whether to allow multi
	// label) from each group.
	AnnotationSpecSetConfigs []*VideoClassificationConfig_AnnotationSpecSetConfig `protobuf:"bytes,1,rep,name=annotation_spec_set_configs,json=annotationSpecSetConfigs,proto3" json:"annotation_spec_set_configs,omitempty"`
	// Optional. Option to apply shot detection on the video.
	ApplyShotDetection   bool     `protobuf:"varint,2,opt,name=apply_shot_detection,json=applyShotDetection,proto3" json:"apply_shot_detection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationConfig) Reset()         { *m = VideoClassificationConfig{} }
func (m *VideoClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*VideoClassificationConfig) ProtoMessage()    {}
func (*VideoClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{5}
}

func (m *VideoClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationConfig.Unmarshal(m, b)
}
func (m *VideoClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationConfig.Marshal(b, m, deterministic)
}
func (m *VideoClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationConfig.Merge(m, src)
}
func (m *VideoClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationConfig.Size(m)
}
func (m *VideoClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationConfig proto.InternalMessageInfo

func (m *VideoClassificationConfig) GetAnnotationSpecSetConfigs() []*VideoClassificationConfig_AnnotationSpecSetConfig {
	if m != nil {
		return m.AnnotationSpecSetConfigs
	}
	return nil
}

func (m *VideoClassificationConfig) GetApplyShotDetection() bool {
	if m != nil {
		return m.ApplyShotDetection
	}
	return false
}

// Annotation spec set with the setting of allowing multi labels or not.
type VideoClassificationConfig_AnnotationSpecSetConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. If allow_multi_label is true, contributors are able to
	// choose multiple labels from one annotation spec set.
	AllowMultiLabel      bool     `protobuf:"varint,2,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) Reset() {
	*m = VideoClassificationConfig_AnnotationSpecSetConfig{}
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) String() string {
	return proto.CompactTextString(m)
}
func (*VideoClassificationConfig_AnnotationSpecSetConfig) ProtoMessage() {}
func (*VideoClassificationConfig_AnnotationSpecSetConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{5, 0}
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Unmarshal(m, b)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Marshal(b, m, deterministic)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Merge(m, src)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Size(m)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig proto.InternalMessageInfo

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

// Config for video object detection human labeling task.
// Object detection will be conducted on the images extracted from the video,
// and those objects will be labeled with bounding boxes.
// User need to specify the number of images to be extracted per second as the
// extraction frame rate.
type ObjectDetectionConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Required. Number of frames per second to be extracted from the video.
	ExtractionFrameRate  float64  `protobuf:"fixed64,3,opt,name=extraction_frame_rate,json=extractionFrameRate,proto3" json:"extraction_frame_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectDetectionConfig) Reset()         { *m = ObjectDetectionConfig{} }
func (m *ObjectDetectionConfig) String() string { return proto.CompactTextString(m) }
func (*ObjectDetectionConfig) ProtoMessage()    {}
func (*ObjectDetectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{6}
}

func (m *ObjectDetectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDetectionConfig.Unmarshal(m, b)
}
func (m *ObjectDetectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDetectionConfig.Marshal(b, m, deterministic)
}
func (m *ObjectDetectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDetectionConfig.Merge(m, src)
}
func (m *ObjectDetectionConfig) XXX_Size() int {
	return xxx_messageInfo_ObjectDetectionConfig.Size(m)
}
func (m *ObjectDetectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDetectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDetectionConfig proto.InternalMessageInfo

func (m *ObjectDetectionConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *ObjectDetectionConfig) GetExtractionFrameRate() float64 {
	if m != nil {
		return m.ExtractionFrameRate
	}
	return 0
}

// Config for video object tracking human labeling task.
type ObjectTrackingConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet    string   `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectTrackingConfig) Reset()         { *m = ObjectTrackingConfig{} }
func (m *ObjectTrackingConfig) String() string { return proto.CompactTextString(m) }
func (*ObjectTrackingConfig) ProtoMessage()    {}
func (*ObjectTrackingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{7}
}

func (m *ObjectTrackingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectTrackingConfig.Unmarshal(m, b)
}
func (m *ObjectTrackingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectTrackingConfig.Marshal(b, m, deterministic)
}
func (m *ObjectTrackingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectTrackingConfig.Merge(m, src)
}
func (m *ObjectTrackingConfig) XXX_Size() int {
	return xxx_messageInfo_ObjectTrackingConfig.Size(m)
}
func (m *ObjectTrackingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectTrackingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectTrackingConfig proto.InternalMessageInfo

func (m *ObjectTrackingConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

// Config for video event human labeling task.
type EventConfig struct {
	// Required. The list of annotation spec set resource name. Similar to video
	// classification, we support selecting event from multiple AnnotationSpecSet
	// at the same time.
	AnnotationSpecSets   []string `protobuf:"bytes,1,rep,name=annotation_spec_sets,json=annotationSpecSets,proto3" json:"annotation_spec_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventConfig) Reset()         { *m = EventConfig{} }
func (m *EventConfig) String() string { return proto.CompactTextString(m) }
func (*EventConfig) ProtoMessage()    {}
func (*EventConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{8}
}

func (m *EventConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventConfig.Unmarshal(m, b)
}
func (m *EventConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventConfig.Marshal(b, m, deterministic)
}
func (m *EventConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventConfig.Merge(m, src)
}
func (m *EventConfig) XXX_Size() int {
	return xxx_messageInfo_EventConfig.Size(m)
}
func (m *EventConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EventConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EventConfig proto.InternalMessageInfo

func (m *EventConfig) GetAnnotationSpecSets() []string {
	if m != nil {
		return m.AnnotationSpecSets
	}
	return nil
}

// Config for text classification human labeling task.
type TextClassificationConfig struct {
	// Optional. If allow_multi_label is true, contributors are able to choose
	// multiple labels for one text segment.
	AllowMultiLabel bool `protobuf:"varint,1,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,2,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Configs for sentiment selection.
	SentimentConfig      *SentimentConfig `protobuf:"bytes,3,opt,name=sentiment_config,json=sentimentConfig,proto3" json:"sentiment_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TextClassificationConfig) Reset()         { *m = TextClassificationConfig{} }
func (m *TextClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*TextClassificationConfig) ProtoMessage()    {}
func (*TextClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{9}
}

func (m *TextClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationConfig.Unmarshal(m, b)
}
func (m *TextClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationConfig.Marshal(b, m, deterministic)
}
func (m *TextClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationConfig.Merge(m, src)
}
func (m *TextClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_TextClassificationConfig.Size(m)
}
func (m *TextClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationConfig proto.InternalMessageInfo

func (m *TextClassificationConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

func (m *TextClassificationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *TextClassificationConfig) GetSentimentConfig() *SentimentConfig {
	if m != nil {
		return m.SentimentConfig
	}
	return nil
}

// Config for setting up sentiments.
type SentimentConfig struct {
	// If set to true, contributors will have the option to select sentiment of
	// the label they selected, to mark it as negative or positive label. Default
	// is false.
	EnableLabelSentimentSelection bool     `protobuf:"varint,1,opt,name=enable_label_sentiment_selection,json=enableLabelSentimentSelection,proto3" json:"enable_label_sentiment_selection,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *SentimentConfig) Reset()         { *m = SentimentConfig{} }
func (m *SentimentConfig) String() string { return proto.CompactTextString(m) }
func (*SentimentConfig) ProtoMessage()    {}
func (*SentimentConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{10}
}

func (m *SentimentConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SentimentConfig.Unmarshal(m, b)
}
func (m *SentimentConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SentimentConfig.Marshal(b, m, deterministic)
}
func (m *SentimentConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentimentConfig.Merge(m, src)
}
func (m *SentimentConfig) XXX_Size() int {
	return xxx_messageInfo_SentimentConfig.Size(m)
}
func (m *SentimentConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SentimentConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SentimentConfig proto.InternalMessageInfo

func (m *SentimentConfig) GetEnableLabelSentimentSelection() bool {
	if m != nil {
		return m.EnableLabelSentimentSelection
	}
	return false
}

// Config for text entity extraction human labeling task.
type TextEntityExtractionConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet    string   `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextEntityExtractionConfig) Reset()         { *m = TextEntityExtractionConfig{} }
func (m *TextEntityExtractionConfig) String() string { return proto.CompactTextString(m) }
func (*TextEntityExtractionConfig) ProtoMessage()    {}
func (*TextEntityExtractionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{11}
}

func (m *TextEntityExtractionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextEntityExtractionConfig.Unmarshal(m, b)
}
func (m *TextEntityExtractionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextEntityExtractionConfig.Marshal(b, m, deterministic)
}
func (m *TextEntityExtractionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextEntityExtractionConfig.Merge(m, src)
}
func (m *TextEntityExtractionConfig) XXX_Size() int {
	return xxx_messageInfo_TextEntityExtractionConfig.Size(m)
}
func (m *TextEntityExtractionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TextEntityExtractionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TextEntityExtractionConfig proto.InternalMessageInfo

func (m *TextEntityExtractionConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datalabeling.v1beta1.StringAggregationType", StringAggregationType_name, StringAggregationType_value)
	proto.RegisterType((*HumanAnnotationConfig)(nil), "google.cloud.datalabeling.v1beta1.HumanAnnotationConfig")
	proto.RegisterType((*ImageClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.ImageClassificationConfig")
	proto.RegisterType((*BoundingPolyConfig)(nil), "google.cloud.datalabeling.v1beta1.BoundingPolyConfig")
	proto.RegisterType((*PolylineConfig)(nil), "google.cloud.datalabeling.v1beta1.PolylineConfig")
	proto.RegisterType((*SegmentationConfig)(nil), "google.cloud.datalabeling.v1beta1.SegmentationConfig")
	proto.RegisterType((*VideoClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.VideoClassificationConfig")
	proto.RegisterType((*VideoClassificationConfig_AnnotationSpecSetConfig)(nil), "google.cloud.datalabeling.v1beta1.VideoClassificationConfig.AnnotationSpecSetConfig")
	proto.RegisterType((*ObjectDetectionConfig)(nil), "google.cloud.datalabeling.v1beta1.ObjectDetectionConfig")
	proto.RegisterType((*ObjectTrackingConfig)(nil), "google.cloud.datalabeling.v1beta1.ObjectTrackingConfig")
	proto.RegisterType((*EventConfig)(nil), "google.cloud.datalabeling.v1beta1.EventConfig")
	proto.RegisterType((*TextClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.TextClassificationConfig")
	proto.RegisterType((*SentimentConfig)(nil), "google.cloud.datalabeling.v1beta1.SentimentConfig")
	proto.RegisterType((*TextEntityExtractionConfig)(nil), "google.cloud.datalabeling.v1beta1.TextEntityExtractionConfig")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/human_annotation_config.proto", fileDescriptor_331725facbee63fc)
}

var fileDescriptor_331725facbee63fc = []byte{
	// 932 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
	0x17, 0xfe, 0x95, 0xfc, 0xed, 0x1a, 0xba, 0x4d, 0x6c, 0xa6, 0x46, 0x95, 0x74, 0xed, 0x3c, 0x15,
	0xc3, 0x8c, 0x62, 0x93, 0x57, 0xef, 0x66, 0xc0, 0x2e, 0x0a, 0xc7, 0x76, 0x3c, 0x0f, 0xb5, 0x1d,
	0xc8, 0x4e, 0x81, 0x16, 0x18, 0x08, 0x5a, 0x3a, 0x51, 0xb8, 0x49, 0xa4, 0x2a, 0x52, 0x6d, 0x8c,
	0x3e, 0xc6, 0xde, 0x6c, 0xc0, 0x6e, 0xf6, 0x06, 0x7b, 0x8b, 0x41, 0xa4, 0x14, 0x7b, 0xa9, 0xbd,
	0x76, 0x19, 0xb6, 0x4b, 0x7f, 0xdf, 0x77, 0xce, 0x47, 0x9e, 0x73, 0x74, 0x68, 0xf4, 0x34, 0x14,
	0x22, 0x8c, 0xa0, 0xe5, 0x47, 0x22, 0x0b, 0x5a, 0x01, 0x55, 0x34, 0xa2, 0x73, 0x88, 0x18, 0x0f,
	0x5b, 0xaf, 0x9f, 0xcc, 0x41, 0xd1, 0x27, 0xad, 0xf3, 0x2c, 0xa6, 0x9c, 0x50, 0xce, 0x85, 0xa2,
	0x8a, 0x09, 0x4e, 0x7c, 0xc1, 0xcf, 0x58, 0xe8, 0x26, 0xa9, 0x50, 0x02, 0x7f, 0x6a, 0x12, 0xb8,
	0x3a, 0x81, 0xbb, 0x9a, 0xc0, 0x2d, 0x12, 0x1c, 0x7e, 0x5c, 0x78, 0xd0, 0x84, 0xb5, 0x96, 0x69,
	0xa4, 0x49, 0x70, 0xf8, 0xb0, 0x60, 0xf5, 0xaf, 0x79, 0x76, 0xd6, 0x0a, 0xb2, 0x54, 0x0b, 0x0c,
	0xef, 0xfc, 0xb2, 0x8d, 0xea, 0xdf, 0xe5, 0x47, 0xe8, 0x5c, 0x86, 0x76, 0xf5, 0x01, 0x70, 0x03,
	0x55, 0x18, 0x97, 0x2a, 0xcd, 0xfc, 0x1c, 0xb4, 0xad, 0x86, 0xd5, 0xdc, 0xf1, 0x56, 0x21, 0xdc,
	0x45, 0x0f, 0x0b, 0x43, 0x08, 0x48, 0x7e, 0x36, 0x09, 0x8a, 0x04, 0x4c, 0x26, 0x11, 0x5d, 0x10,
	0x4e, 0x63, 0xb0, 0xb7, 0x74, 0xd0, 0xfd, 0x4b, 0x55, 0xcf, 0x88, 0x7a, 0x46, 0x33, 0xa6, 0x31,
	0xe0, 0x23, 0xf4, 0x60, 0x4d, 0x12, 0x90, 0x7e, 0xca, 0x12, 0x6d, 0xbc, 0xbd, 0x21, 0xc7, 0x52,
	0x82, 0x3f, 0x41, 0x15, 0x5d, 0x16, 0x12, 0xa6, 0x22, 0x4b, 0xec, 0xff, 0xeb, 0x08, 0xa4, 0xa1,
	0x41, 0x8e, 0xe0, 0x47, 0xe8, 0x4e, 0x44, 0x79, 0x98, 0xd1, 0x10, 0x88, 0x2f, 0x02, 0xb0, 0x6f,
	0x68, 0xc9, 0xed, 0x12, 0xec, 0x8a, 0x00, 0x72, 0x51, 0x0a, 0x49, 0xc4, 0x7c, 0x4a, 0x7c, 0x91,
	0x71, 0x65, 0xdf, 0x6c, 0x58, 0xcd, 0x1b, 0xde, 0xed, 0x02, 0xec, 0xe6, 0x18, 0x3e, 0x46, 0xb5,
	0x57, 0x19, 0x48, 0xdd, 0xa9, 0xb2, 0x94, 0xf6, 0x47, 0x0d, 0xab, 0x59, 0x69, 0x1f, 0xb8, 0x45,
	0xb3, 0xca, 0x5a, 0xbb, 0xbd, 0x42, 0xe0, 0x55, 0xcb, 0x98, 0x12, 0xc1, 0x5f, 0x22, 0xec, 0x0b,
	0xae, 0x52, 0x36, 0xcf, 0x94, 0x48, 0x09, 0xc4, 0x94, 0x45, 0xd2, 0xde, 0x69, 0x6c, 0x37, 0x77,
	0xbc, 0xda, 0x0a, 0xd3, 0xd7, 0x04, 0xfe, 0x02, 0xe1, 0x4c, 0x42, 0xa1, 0x23, 0x34, 0x08, 0x52,
	0x90, 0xd2, 0x46, 0xfa, 0x16, 0xd5, 0x9c, 0xd1, 0xba, 0x8e, 0xc1, 0x9d, 0xdf, 0x2d, 0x74, 0x30,
	0x8c, 0xf3, 0x7b, 0x45, 0x54, 0x4a, 0x76, 0xc6, 0xfc, 0xd5, 0xc6, 0xba, 0x68, 0x7f, 0x65, 0xdc,
	0x64, 0x02, 0x3e, 0x91, 0xa0, 0x8a, 0x06, 0xd7, 0x96, 0xd4, 0x34, 0x01, 0x7f, 0x0a, 0x0a, 0x3f,
	0x46, 0x35, 0x1a, 0x45, 0xe2, 0x0d, 0x89, 0xb3, 0x48, 0x31, 0xa2, 0xcb, 0xaa, 0x3b, 0x7b, 0xcb,
	0xdb, 0xd3, 0xc4, 0x28, 0xc7, 0x9f, 0xe5, 0x30, 0x4e, 0xd0, 0x3d, 0xca, 0xe5, 0x1b, 0x48, 0x09,
	0x0d, 0xc3, 0x14, 0x42, 0xe3, 0xa1, 0x16, 0x09, 0xe8, 0x3e, 0xee, 0xb6, 0xbf, 0x71, 0xdf, 0x3b,
	0xd1, 0xee, 0x54, 0xa5, 0x8c, 0x87, 0x9d, 0x65, 0x82, 0xd9, 0x22, 0x01, 0xaf, 0x6e, 0x12, 0x5f,
	0x81, 0x9d, 0x0c, 0xe1, 0x23, 0x91, 0xf1, 0x80, 0xf1, 0xf0, 0x44, 0x44, 0x8b, 0x6b, 0xde, 0xb1,
	0x85, 0xf6, 0x57, 0x26, 0x9b, 0xc4, 0x20, 0x25, 0x0d, 0xcb, 0xf9, 0xc5, 0x2b, 0xd4, 0xc8, 0x30,
	0xce, 0x2b, 0xb4, 0x9b, 0xdb, 0x45, 0x8c, 0xc3, 0x7f, 0x65, 0x99, 0x21, 0x3c, 0x85, 0x30, 0x06,
	0xae, 0xfe, 0x49, 0x37, 0xff, 0xb6, 0xed, 0x6f, 0x5b, 0xe8, 0xe0, 0x39, 0x0b, 0x40, 0xac, 0x1d,
	0xa6, 0x9f, 0x2d, 0x74, 0x7f, 0x8d, 0x7f, 0xb1, 0xc5, 0xa4, 0x6d, 0x35, 0xb6, 0x9b, 0x95, 0xf6,
	0xec, 0x03, 0xba, 0xbe, 0xd1, 0xc3, 0xed, 0x5c, 0xbd, 0x84, 0xc1, 0x3d, 0x9b, 0xae, 0x27, 0x24,
	0xfe, 0x0a, 0xdd, 0xa5, 0x49, 0x12, 0x2d, 0x88, 0x3c, 0x17, 0xf9, 0x36, 0x51, 0x60, 0x96, 0x98,
	0x99, 0x5a, 0xac, 0xb9, 0xe9, 0xb9, 0x50, 0xbd, 0x92, 0x39, 0xcc, 0xd0, 0xbd, 0x0d, 0x36, 0xff,
	0xe6, 0xf7, 0xe2, 0xbc, 0x45, 0xf5, 0xc9, 0xfc, 0x47, 0xf0, 0x97, 0x27, 0xb9, 0xa6, 0x69, 0x1b,
	0xd5, 0xe1, 0x42, 0xa5, 0xd4, 0x74, 0xf5, 0x2c, 0xa5, 0x31, 0x90, 0x94, 0x2a, 0xf3, 0xd9, 0x59,
	0xde, 0xfe, 0x92, 0x3c, 0xce, 0x39, 0x8f, 0x2a, 0x70, 0x8e, 0xd1, 0x5d, 0x63, 0x3e, 0x4b, 0xa9,
	0xff, 0x13, 0xe3, 0xe1, 0xf5, 0xbc, 0x9d, 0xa7, 0xa8, 0xd2, 0x7f, 0x0d, 0xbc, 0xac, 0x57, 0x5e,
	0xfc, 0x77, 0xc3, 0xcd, 0x28, 0xec, 0x78, 0xf8, 0x9d, 0x78, 0xe9, 0xfc, 0x6a, 0x21, 0x7b, 0x06,
	0x17, 0x6a, 0xed, 0x84, 0xad, 0x2d, 0xa7, 0xb5, 0x7e, 0xfd, 0x6c, 0x38, 0xf9, 0xd6, 0xa6, 0xaa,
	0xfd, 0x80, 0xaa, 0x12, 0xb8, 0x62, 0xf9, 0x47, 0x55, 0x8c, 0xac, 0x2e, 0x58, 0xa5, 0xdd, 0xfe,
	0x90, 0x3d, 0x55, 0x86, 0x16, 0xf3, 0xb8, 0x27, 0xff, 0x0c, 0x38, 0x2f, 0xd1, 0xde, 0x15, 0x0d,
	0x1e, 0xa0, 0x06, 0x70, 0x3a, 0x8f, 0xc0, 0x5c, 0x84, 0x2c, 0xed, 0x25, 0x44, 0xb0, 0x7c, 0x6a,
	0x6f, 0x79, 0x0f, 0x8c, 0x4e, 0x5f, 0xec, 0x32, 0xcb, 0xb4, 0x14, 0x39, 0xcf, 0xd0, 0x61, 0x5e,
	0xb2, 0x3e, 0x57, 0x4c, 0x2d, 0xfa, 0x97, 0xdd, 0xbd, 0x5e, 0x0b, 0x1f, 0xbf, 0x45, 0xf5, 0xb5,
	0x5b, 0x17, 0x7f, 0x8e, 0x1e, 0x4d, 0x67, 0xde, 0x70, 0x3c, 0x20, 0x9d, 0xc1, 0xc0, 0xeb, 0x0f,
	0x3a, 0xb3, 0xe1, 0x64, 0x4c, 0x66, 0x2f, 0x4e, 0xfa, 0xe4, 0x74, 0x3c, 0x3d, 0xe9, 0x77, 0x87,
	0xc7, 0xc3, 0x7e, 0xaf, 0xfa, 0x3f, 0x5c, 0x43, 0x77, 0x46, 0x9d, 0xef, 0x27, 0xde, 0x70, 0xf6,
	0x82, 0x3c, 0x9f, 0xcc, 0xfa, 0x55, 0x0b, 0x63, 0xb4, 0x7b, 0x3a, 0xee, 0x8c, 0x87, 0xa3, 0xc9,
	0xe9, 0xd4, 0x60, 0x5b, 0x39, 0x36, 0x9e, 0xac, 0xe6, 0xaa, 0x6e, 0x1f, 0x5d, 0xa0, 0xcf, 0x7c,
	0x11, 0xbf, 0xbf, 0xe0, 0x27, 0xd6, 0xcb, 0x51, 0x21, 0x0a, 0x45, 0xfe, 0x74, 0xbb, 0x22, 0x0d,
	0x5b, 0x21, 0x70, 0xfd, 0xe0, 0xb6, 0x0c, 0x45, 0x13, 0x26, 0xff, 0xe2, 0xff, 0xd6, 0xb7, 0xab,
	0xe0, 0xfc, 0xa6, 0x8e, 0xfc, 0xfa, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xb3, 0x56, 0xdc,
	0xa8, 0x09, 0x00, 0x00,
}
