// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/classification.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type of the classification problem.
type ClassificationType int32

const (
	// Should not be used, an un-set enum has this value by default.
	ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED ClassificationType = 0
	// At most one label is allowed per example.
	ClassificationType_MULTICLASS ClassificationType = 1
	// Multiple labels are allowed for one example.
	ClassificationType_MULTILABEL ClassificationType = 2
)

var ClassificationType_name = map[int32]string{
	0: "CLASSIFICATION_TYPE_UNSPECIFIED",
	1: "MULTICLASS",
	2: "MULTILABEL",
}

var ClassificationType_value = map[string]int32{
	"CLASSIFICATION_TYPE_UNSPECIFIED": 0,
	"MULTICLASS":                      1,
	"MULTILABEL":                      2,
}

func (x ClassificationType) String() string {
	return proto.EnumName(ClassificationType_name, int32(x))
}

func (ClassificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{0}
}

// Contains annotation details specific to classification.
type ClassificationAnnotation struct {
	// Output only. A confidence estimate between 0.0 and 1.0. A higher value
	// means greater confidence that the annotation is positive. If a user
	// approves an annotation as negative or positive, the score value remains
	// unchanged. If a user creates an annotation, the score is 0 for negative or
	// 1 for positive.
	Score                float32  `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationAnnotation) Reset()         { *m = ClassificationAnnotation{} }
func (m *ClassificationAnnotation) String() string { return proto.CompactTextString(m) }
func (*ClassificationAnnotation) ProtoMessage()    {}
func (*ClassificationAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{0}
}
func (m *ClassificationAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationAnnotation.Unmarshal(m, b)
}
func (m *ClassificationAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationAnnotation.Marshal(b, m, deterministic)
}
func (m *ClassificationAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationAnnotation.Merge(m, src)
}
func (m *ClassificationAnnotation) XXX_Size() int {
	return xxx_messageInfo_ClassificationAnnotation.Size(m)
}
func (m *ClassificationAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationAnnotation proto.InternalMessageInfo

func (m *ClassificationAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Model evaluation metrics for classification problems.
// Visible only to v1beta1
type ClassificationEvaluationMetrics struct {
	// Output only. The Area under precision recall curve metric.
	AuPrc float32 `protobuf:"fixed32,1,opt,name=au_prc,json=auPrc,proto3" json:"au_prc,omitempty"`
	// Output only. The Area under precision recall curve metric based on priors.
	BaseAuPrc float32 `protobuf:"fixed32,2,opt,name=base_au_prc,json=baseAuPrc,proto3" json:"base_au_prc,omitempty"`
	// Output only. Metrics that have confidence thresholds.
	// Precision-recall curve can be derived from it.
	ConfidenceMetricsEntry []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entry,json=confidenceMetricsEntry,proto3" json:"confidence_metrics_entry,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for MULTICLASS classification problems where number
	// of labels is no more than 10.
	// Only set for model level evaluation, not for evaluation per label.
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,4,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	// Output only. The annotation spec ids used for this evaluation.
	AnnotationSpecId     []string `protobuf:"bytes,5,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics) Reset()         { *m = ClassificationEvaluationMetrics{} }
func (m *ClassificationEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*ClassificationEvaluationMetrics) ProtoMessage()    {}
func (*ClassificationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{1}
}
func (m *ClassificationEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Size(m)
}
func (m *ClassificationEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics) GetAuPrc() float32 {
	if m != nil {
		return m.AuPrc
	}
	return 0
}

func (m *ClassificationEvaluationMetrics) GetBaseAuPrc() float32 {
	if m != nil {
		return m.BaseAuPrc
	}
	return 0
}

func (m *ClassificationEvaluationMetrics) GetConfidenceMetricsEntry() []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry {
	if m != nil {
		return m.ConfidenceMetricsEntry
	}
	return nil
}

func (m *ClassificationEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if m != nil {
		return m.ConfusionMatrix
	}
	return nil
}

func (m *ClassificationEvaluationMetrics) GetAnnotationSpecId() []string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return nil
}

// Metrics for a single confidence threshold.
type ClassificationEvaluationMetrics_ConfidenceMetricsEntry struct {
	// Output only. The confidence threshold value used to compute the metrics.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Recall under the given confidence threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision under the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. The recall when only considering the label that has the
	// highest prediction score and not below the confidence threshold for each
	// example.
	RecallAt1 float32 `protobuf:"fixed32,5,opt,name=recall_at1,json=recallAt1,proto3" json:"recall_at1,omitempty"`
	// Output only. The precision when only considering the label that has the
	// highest predictionscore and not below the confidence threshold for each
	// example.
	PrecisionAt1 float32 `protobuf:"fixed32,6,opt,name=precision_at1,json=precisionAt1,proto3" json:"precision_at1,omitempty"`
	// Output only. The harmonic mean of [recall_at1][google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision_at1].
	F1ScoreAt1           float32  `protobuf:"fixed32,7,opt,name=f1_score_at1,json=f1ScoreAt1,proto3" json:"f1_score_at1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Reset() {
	*m = ClassificationEvaluationMetrics_ConfidenceMetricsEntry{}
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{1, 0}
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if m != nil {
		return m.ConfidenceThreshold
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecallAt1() float32 {
	if m != nil {
		return m.RecallAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecisionAt1() float32 {
	if m != nil {
		return m.PrecisionAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1ScoreAt1() float32 {
	if m != nil {
		return m.F1ScoreAt1
	}
	return 0
}

// Confusion matrix of the model running the classification.
type ClassificationEvaluationMetrics_ConfusionMatrix struct {
	// Output only. IDs of the annotation specs used in the confusion matrix.
	AnnotationSpecId []string `protobuf:"bytes,1,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. Rows in the confusion matrix. The number of rows is equal to
	// the size of `annotation_spec_id`.
	// `row[i].value[j]` is the number of examples that have ground truth of the
	// `annotation_spec_id[i]` and are predicted as `annotation_spec_id[j]` by
	// the model being evaluated.
	Row                  []*ClassificationEvaluationMetrics_ConfusionMatrix_Row `protobuf:"bytes,2,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) Reset() {
	*m = ClassificationEvaluationMetrics_ConfusionMatrix{}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfusionMatrix) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfusionMatrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{1, 1}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) GetAnnotationSpecId() []string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return nil
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) GetRow() []*ClassificationEvaluationMetrics_ConfusionMatrix_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

// Output only. A row in the confusion matrix.
type ClassificationEvaluationMetrics_ConfusionMatrix_Row struct {
	// Output only. Value of the specific cell in the confusion matrix.
	// The number of values each row is equal to the size of
	// annotatin_spec_id.
	ExampleCount         []int32  `protobuf:"varint,1,rep,packed,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) Reset() {
	*m = ClassificationEvaluationMetrics_ConfusionMatrix_Row{}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{1, 1, 0}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) GetExampleCount() []int32 {
	if m != nil {
		return m.ExampleCount
	}
	return nil
}

func init() {
	proto.RegisterType((*ClassificationAnnotation)(nil), "google.cloud.automl.v1beta1.ClassificationAnnotation")
	proto.RegisterType((*ClassificationEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfidenceMetricsEntry)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfusionMatrix)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfusionMatrix_Row)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix.Row")
	proto.RegisterEnum("google.cloud.automl.v1beta1.ClassificationType", ClassificationType_name, ClassificationType_value)
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/classification.proto", fileDescriptor_7b436fefe6ae5367)
}

var fileDescriptor_7b436fefe6ae5367 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x4f, 0xdb, 0x3e,
	0x14, 0xff, 0xa6, 0xa5, 0xe5, 0xcb, 0x83, 0x8d, 0xca, 0x30, 0x94, 0x75, 0x6c, 0x54, 0x70, 0xa9,
	0xd0, 0x94, 0x10, 0x76, 0xdc, 0x29, 0x64, 0x45, 0x8a, 0x54, 0x58, 0x95, 0x94, 0x03, 0xbb, 0x58,
	0xae, 0xeb, 0x86, 0x48, 0x69, 0x1c, 0x39, 0x0e, 0x85, 0xeb, 0xee, 0xfb, 0xdb, 0xf6, 0x5f, 0xec,
	0xef, 0x98, 0x62, 0x87, 0x74, 0x95, 0x2a, 0x36, 0x69, 0xbb, 0xe5, 0x7d, 0x7e, 0xf9, 0x3d, 0xbf,
	0xc8, 0x70, 0x16, 0x71, 0x1e, 0x25, 0xcc, 0xa6, 0x09, 0x2f, 0xa6, 0x36, 0x29, 0x24, 0x9f, 0x27,
	0xf6, 0xbd, 0x33, 0x61, 0x92, 0x38, 0x36, 0x4d, 0x48, 0x9e, 0xc7, 0xb3, 0x98, 0x12, 0x19, 0xf3,
	0xd4, 0xca, 0x04, 0x97, 0x1c, 0xbd, 0xd1, 0x0e, 0x4b, 0x39, 0x2c, 0xed, 0xb0, 0x2a, 0x47, 0xf7,
	0xb0, 0x8a, 0x23, 0x59, 0x6c, 0x93, 0x34, 0xe5, 0x52, 0x39, 0x73, 0x6d, 0x3d, 0x3e, 0x03, 0xd3,
	0x5b, 0x89, 0x74, 0x6b, 0x09, 0xda, 0x87, 0x56, 0x4e, 0xb9, 0x60, 0xa6, 0xd1, 0x33, 0xfa, 0x8d,
	0x40, 0x17, 0xc7, 0x3f, 0xda, 0x70, 0xb4, 0x6a, 0x19, 0xdc, 0x93, 0xa4, 0x50, 0x5f, 0x57, 0x4c,
	0x8a, 0x98, 0xe6, 0xe8, 0x15, 0xb4, 0x49, 0x81, 0x33, 0x41, 0x9f, 0xac, 0xa4, 0x18, 0x09, 0x8a,
	0xde, 0xc1, 0xf6, 0x84, 0xe4, 0x0c, 0x57, 0x5c, 0x43, 0x71, 0x5b, 0x25, 0xe4, 0x2a, 0xfe, 0x9b,
	0x01, 0x26, 0xe5, 0xe9, 0x2c, 0x9e, 0xb2, 0x94, 0x32, 0x3c, 0xd7, 0x69, 0x98, 0xa5, 0x52, 0x3c,
	0x9a, 0xcd, 0x5e, 0xb3, 0xbf, 0x7d, 0x1e, 0x5a, 0xcf, 0xcc, 0x6a, 0xfd, 0xa6, 0x2f, 0xcb, 0xab,
	0xc3, 0x2b, 0x64, 0x50, 0x46, 0x07, 0x07, 0x74, 0x2d, 0x8e, 0x16, 0xd0, 0x29, 0x99, 0x22, 0x8f,
	0x79, 0x8a, 0xe7, 0x44, 0x8a, 0xf8, 0xc1, 0xdc, 0xe8, 0x19, 0xfd, 0xed, 0xf3, 0xe1, 0x5f, 0xb7,
	0xa1, 0x42, 0xaf, 0x54, 0x66, 0xb0, 0x4b, 0x57, 0x01, 0xf4, 0x1e, 0xd0, 0x72, 0x55, 0x38, 0xcf,
	0x18, 0xc5, 0xf1, 0xd4, 0x6c, 0xf5, 0x9a, 0xfd, 0xad, 0xa0, 0xb3, 0x64, 0xc2, 0x8c, 0x51, 0x7f,
	0xda, 0xfd, 0xda, 0x80, 0x83, 0xf5, 0x93, 0x21, 0x07, 0xf6, 0x7f, 0xb9, 0x50, 0x79, 0x27, 0x58,
	0x7e, 0xc7, 0x93, 0x69, 0xb5, 0x96, 0xbd, 0x25, 0x37, 0x7e, 0xa2, 0xd0, 0x01, 0xb4, 0x05, 0xa3,
	0x24, 0x49, 0xaa, 0xfd, 0x54, 0x15, 0x3a, 0x84, 0xad, 0x4c, 0x30, 0x1a, 0x97, 0x6d, 0x9a, 0x4d,
	0xbd, 0xba, 0x1a, 0x40, 0xaf, 0xe1, 0xff, 0x99, 0x83, 0xf5, 0xef, 0xb2, 0xa1, 0xc8, 0xcd, 0x99,
	0x13, 0x96, 0x25, 0x7a, 0x0b, 0xa0, 0x23, 0x30, 0x91, 0x8e, 0xd9, 0xd2, 0x4e, 0x8d, 0xb8, 0xd2,
	0x41, 0x27, 0xf0, 0xa2, 0x8e, 0x51, 0x8a, 0xb6, 0x52, 0xec, 0xd4, 0x60, 0x29, 0xea, 0xc1, 0xce,
	0x53, 0xbc, 0xd2, 0x6c, 0x2a, 0x0d, 0x54, 0x47, 0xb8, 0xd2, 0xe9, 0x7e, 0x37, 0x60, 0xd7, 0xfb,
	0xa3, 0x6b, 0x34, 0xd6, 0x5f, 0x23, 0x9a, 0x40, 0x53, 0xf0, 0x85, 0xd9, 0x50, 0xff, 0xd9, 0xe8,
	0x5f, 0x2e, 0xd8, 0x0a, 0xf8, 0x22, 0x28, 0xc3, 0xbb, 0xa7, 0xd0, 0x0c, 0xf8, 0xa2, 0x9c, 0x99,
	0x3d, 0x90, 0x79, 0x96, 0x30, 0x4c, 0x79, 0x91, 0x4a, 0xd5, 0x53, 0x2b, 0xd8, 0xa9, 0x40, 0xaf,
	0xc4, 0x4e, 0x6f, 0x01, 0xad, 0x9e, 0x33, 0x7e, 0xcc, 0x18, 0x3a, 0x81, 0x23, 0x6f, 0xe8, 0x86,
	0xa1, 0x7f, 0xe9, 0x7b, 0xee, 0xd8, 0xff, 0x7c, 0x8d, 0xc7, 0xb7, 0xa3, 0x01, 0xbe, 0xb9, 0x0e,
	0x47, 0x03, 0xcf, 0xbf, 0xf4, 0x07, 0x9f, 0x3a, 0xff, 0xa1, 0x97, 0x00, 0x57, 0x37, 0xc3, 0xb1,
	0xaf, 0x94, 0x1d, 0xa3, 0xae, 0x87, 0xee, 0xc5, 0x60, 0xd8, 0x69, 0x5c, 0x3c, 0xc2, 0x11, 0xe5,
	0xf3, 0xe7, 0x46, 0xbc, 0xd8, 0x5b, 0x3d, 0x7b, 0x54, 0xbe, 0x16, 0x5f, 0xdc, 0xca, 0x11, 0xf1,
	0x84, 0xa4, 0x91, 0xc5, 0x45, 0x64, 0x47, 0x2c, 0x55, 0x2f, 0x89, 0xad, 0x29, 0x92, 0xc5, 0xf9,
	0xda, 0x97, 0xeb, 0xa3, 0x2e, 0x27, 0x6d, 0xa5, 0xfe, 0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0x22,
	0xd4, 0x8f, 0x68, 0xe6, 0x04, 0x00, 0x00,
}
