/// Type: Vote payloads send across the network

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: common/vote.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VoteEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`             // Unique number per vote attempt, so that replay attacks can't reuse this payload
	ProcessId   []byte `protobuf:"bytes,2,opt,name=processId,proto3" json:"processId,omitempty"`     // The process for which the vote is casted
	Proof       *Proof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`             // Merkle proof, ZK Proof or Ethereum storage proof
	VotePackage []byte `protobuf:"bytes,4,opt,name=votePackage,proto3" json:"votePackage,omitempty"` // base64 encoded bytes[] of the VotePackage
	// optional bytes nullifier = 5;
	Nullifier            []byte   `protobuf:"bytes,5,opt,name=nullifier,proto3" json:"nullifier,omitempty"`                               // Hash of the private key + processId
	EncryptionKeyIndexes []uint32 `protobuf:"varint,6,rep,packed,name=encryptionKeyIndexes,proto3" json:"encryptionKeyIndexes,omitempty"` // On encrypted votes, contains the (sorted) indexes of the keys used to encrypt
}

func (x *VoteEnvelope) Reset() {
	*x = VoteEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteEnvelope) ProtoMessage() {}

func (x *VoteEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteEnvelope.ProtoReflect.Descriptor instead.
func (*VoteEnvelope) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{0}
}

func (x *VoteEnvelope) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *VoteEnvelope) GetProcessId() []byte {
	if x != nil {
		return x.ProcessId
	}
	return nil
}

func (x *VoteEnvelope) GetProof() *Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *VoteEnvelope) GetVotePackage() []byte {
	if x != nil {
		return x.VotePackage
	}
	return nil
}

func (x *VoteEnvelope) GetNullifier() []byte {
	if x != nil {
		return x.Nullifier
	}
	return nil
}

func (x *VoteEnvelope) GetEncryptionKeyIndexes() []uint32 {
	if x != nil {
		return x.EncryptionKeyIndexes
	}
	return nil
}

type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*Proof_Graviton
	//	*Proof_Iden3
	//	*Proof_EthereumStorage
	//	*Proof_EthAccount
	Payload isProof_Payload `protobuf_oneof:"payload"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{1}
}

func (m *Proof) GetPayload() isProof_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Proof) GetGraviton() *ProofGraviton {
	if x, ok := x.GetPayload().(*Proof_Graviton); ok {
		return x.Graviton
	}
	return nil
}

func (x *Proof) GetIden3() *ProofIden3 {
	if x, ok := x.GetPayload().(*Proof_Iden3); ok {
		return x.Iden3
	}
	return nil
}

func (x *Proof) GetEthereumStorage() *ProofEthereumStorage {
	if x, ok := x.GetPayload().(*Proof_EthereumStorage); ok {
		return x.EthereumStorage
	}
	return nil
}

func (x *Proof) GetEthAccount() *ProofEthereumAccount {
	if x, ok := x.GetPayload().(*Proof_EthAccount); ok {
		return x.EthAccount
	}
	return nil
}

type isProof_Payload interface {
	isProof_Payload()
}

type Proof_Graviton struct {
	Graviton *ProofGraviton `protobuf:"bytes,1,opt,name=graviton,proto3,oneof"`
}

type Proof_Iden3 struct {
	Iden3 *ProofIden3 `protobuf:"bytes,2,opt,name=iden3,proto3,oneof"`
}

type Proof_EthereumStorage struct {
	EthereumStorage *ProofEthereumStorage `protobuf:"bytes,3,opt,name=ethereumStorage,proto3,oneof"`
}

type Proof_EthAccount struct {
	EthAccount *ProofEthereumAccount `protobuf:"bytes,4,opt,name=ethAccount,proto3,oneof"`
}

func (*Proof_Graviton) isProof_Payload() {}

func (*Proof_Iden3) isProof_Payload() {}

func (*Proof_EthereumStorage) isProof_Payload() {}

func (*Proof_EthAccount) isProof_Payload() {}

type ProofGraviton struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Siblings []byte `protobuf:"bytes,1,opt,name=siblings,proto3" json:"siblings,omitempty"`
}

func (x *ProofGraviton) Reset() {
	*x = ProofGraviton{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofGraviton) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofGraviton) ProtoMessage() {}

func (x *ProofGraviton) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofGraviton.ProtoReflect.Descriptor instead.
func (*ProofGraviton) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{2}
}

func (x *ProofGraviton) GetSiblings() []byte {
	if x != nil {
		return x.Siblings
	}
	return nil
}

type ProofIden3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Siblings []byte `protobuf:"bytes,1,opt,name=siblings,proto3" json:"siblings,omitempty"`
}

func (x *ProofIden3) Reset() {
	*x = ProofIden3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofIden3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofIden3) ProtoMessage() {}

func (x *ProofIden3) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofIden3.ProtoReflect.Descriptor instead.
func (*ProofIden3) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{3}
}

func (x *ProofIden3) GetSiblings() []byte {
	if x != nil {
		return x.Siblings
	}
	return nil
}

type ProofEthereumStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Siblings [][]byte `protobuf:"bytes,3,rep,name=siblings,proto3" json:"siblings,omitempty"`
}

func (x *ProofEthereumStorage) Reset() {
	*x = ProofEthereumStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofEthereumStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofEthereumStorage) ProtoMessage() {}

func (x *ProofEthereumStorage) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofEthereumStorage.ProtoReflect.Descriptor instead.
func (*ProofEthereumStorage) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{4}
}

func (x *ProofEthereumStorage) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ProofEthereumStorage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ProofEthereumStorage) GetSiblings() [][]byte {
	if x != nil {
		return x.Siblings
	}
	return nil
}

type ProofEthereumAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance     []byte   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"` // Big Int encoded as bytes
	StorageHash []byte   `protobuf:"bytes,3,opt,name=storageHash,proto3" json:"storageHash,omitempty"`
	CodeHash    []byte   `protobuf:"bytes,4,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	Siblings    [][]byte `protobuf:"bytes,5,rep,name=siblings,proto3" json:"siblings,omitempty"`
}

func (x *ProofEthereumAccount) Reset() {
	*x = ProofEthereumAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_vote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofEthereumAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofEthereumAccount) ProtoMessage() {}

func (x *ProofEthereumAccount) ProtoReflect() protoreflect.Message {
	mi := &file_common_vote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofEthereumAccount.ProtoReflect.Descriptor instead.
func (*ProofEthereumAccount) Descriptor() ([]byte, []int) {
	return file_common_vote_proto_rawDescGZIP(), []int{5}
}

func (x *ProofEthereumAccount) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *ProofEthereumAccount) GetBalance() []byte {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *ProofEthereumAccount) GetStorageHash() []byte {
	if x != nil {
		return x.StorageHash
	}
	return nil
}

func (x *ProofEthereumAccount) GetCodeHash() []byte {
	if x != nil {
		return x.CodeHash
	}
	return nil
}

func (x *ProofEthereumAccount) GetSiblings() [][]byte {
	if x != nil {
		return x.Siblings
	}
	return nil
}

var File_common_vote_proto protoreflect.FileDescriptor

var file_common_vote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0xe3, 0x01, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x05, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x3b, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x47, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x6e,
	0x12, 0x32, 0x0a, 0x05, 0x69, 0x64, 0x65, 0x6e, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x64, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x49, 0x64, 0x65, 0x6e, 0x33, 0x48, 0x00, 0x52, 0x05, 0x69,
	0x64, 0x65, 0x6e, 0x33, 0x12, 0x50, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x64, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x76, 0x6f,
	0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2b, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69,
	0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69,
	0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x28, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x49,
	0x64, 0x65, 0x6e, 0x33, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x5a, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xa0, 0x01, 0x0a,
	0x14, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f,
	0x63, 0x64, 0x6f, 0x6e, 0x69, 0x2f, 0x64, 0x76, 0x6f, 0x74, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_vote_proto_rawDescOnce sync.Once
	file_common_vote_proto_rawDescData = file_common_vote_proto_rawDesc
)

func file_common_vote_proto_rawDescGZIP() []byte {
	file_common_vote_proto_rawDescOnce.Do(func() {
		file_common_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_vote_proto_rawDescData)
	})
	return file_common_vote_proto_rawDescData
}

var file_common_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_vote_proto_goTypes = []interface{}{
	(*VoteEnvelope)(nil),         // 0: dvote.types.v1.VoteEnvelope
	(*Proof)(nil),                // 1: dvote.types.v1.Proof
	(*ProofGraviton)(nil),        // 2: dvote.types.v1.ProofGraviton
	(*ProofIden3)(nil),           // 3: dvote.types.v1.ProofIden3
	(*ProofEthereumStorage)(nil), // 4: dvote.types.v1.ProofEthereumStorage
	(*ProofEthereumAccount)(nil), // 5: dvote.types.v1.ProofEthereumAccount
}
var file_common_vote_proto_depIdxs = []int32{
	1, // 0: dvote.types.v1.VoteEnvelope.proof:type_name -> dvote.types.v1.Proof
	2, // 1: dvote.types.v1.Proof.graviton:type_name -> dvote.types.v1.ProofGraviton
	3, // 2: dvote.types.v1.Proof.iden3:type_name -> dvote.types.v1.ProofIden3
	4, // 3: dvote.types.v1.Proof.ethereumStorage:type_name -> dvote.types.v1.ProofEthereumStorage
	5, // 4: dvote.types.v1.Proof.ethAccount:type_name -> dvote.types.v1.ProofEthereumAccount
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_vote_proto_init() }
func file_common_vote_proto_init() {
	if File_common_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteEnvelope); i {
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
		file_common_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
		file_common_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofGraviton); i {
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
		file_common_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofIden3); i {
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
		file_common_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofEthereumStorage); i {
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
		file_common_vote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofEthereumAccount); i {
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
	file_common_vote_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Proof_Graviton)(nil),
		(*Proof_Iden3)(nil),
		(*Proof_EthereumStorage)(nil),
		(*Proof_EthAccount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_vote_proto_goTypes,
		DependencyIndexes: file_common_vote_proto_depIdxs,
		MessageInfos:      file_common_vote_proto_msgTypes,
	}.Build()
	File_common_vote_proto = out.File
	file_common_vote_proto_rawDesc = nil
	file_common_vote_proto_goTypes = nil
	file_common_vote_proto_depIdxs = nil
}
