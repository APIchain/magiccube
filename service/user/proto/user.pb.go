// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	PushTxRequest
	PushTxResponse
	RegisterRequest
	AccountInfo
	UserInfo
	RegisterResponse
	LoginRequest
	LoginResponse
	GetBlockHeaderRequest
	GetBlockHeaderResponse
	BlockHeader
	GetAccountInfoRequest
	GetAccountInfoResponse
	AccountInfoData
	FavoriteRequest
	FavoriteResponse
	GetFavoriteRequest
	GetFavoriteResponse
	FavoriteArr
	FavoriteData
	GetBalanceRequest
	GetBalanceResponse
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PushTxRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *PushTxRequest) Reset()                    { *m = PushTxRequest{} }
func (m *PushTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PushTxRequest) ProtoMessage()               {}
func (*PushTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PushTxRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushTxRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *PushTxRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *PushTxRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *PushTxRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PushTxRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *PushTxRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PushTxRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *PushTxRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *PushTxRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type PushTxResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *PushTxResponse) Reset()                    { *m = PushTxResponse{} }
func (m *PushTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PushTxResponse) ProtoMessage()               {}
func (*PushTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PushTxResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PushTxResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RegisterRequest struct {
	Account     *AccountInfo `protobuf:"bytes,1,opt,name=account" json:"account"`
	User        *UserInfo    `protobuf:"bytes,2,opt,name=user" json:"user"`
	VerifyId    string       `protobuf:"bytes,3,opt,name=verify_id,json=verifyId" json:"verify_id"`
	VerifyValue string       `protobuf:"bytes,4,opt,name=verify_value,json=verifyValue" json:"verify_value"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterRequest) GetAccount() *AccountInfo {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *RegisterRequest) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RegisterRequest) GetVerifyId() string {
	if m != nil {
		return m.VerifyId
	}
	return ""
}

func (m *RegisterRequest) GetVerifyValue() string {
	if m != nil {
		return m.VerifyValue
	}
	return ""
}

type AccountInfo struct {
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name"`
	Pubkey string `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey"`
}

func (m *AccountInfo) Reset()                    { *m = AccountInfo{} }
func (m *AccountInfo) String() string            { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()               {}
func (*AccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type UserInfo struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfo) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UserInfo) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *UserInfo) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *UserInfo) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *UserInfo) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *UserInfo) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *UserInfo) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UserInfo) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *UserInfo) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *UserInfo) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type RegisterResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegisterResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginRequest struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username"`
	Random      string `protobuf:"bytes,2,opt,name=random" json:"random"`
	VerifyId    string `protobuf:"bytes,3,opt,name=verify_id,json=verifyId" json:"verify_id"`
	VerifyValue string `protobuf:"bytes,4,opt,name=verify_value,json=verifyValue" json:"verify_value"`
	Signture    string `protobuf:"bytes,5,opt,name=signture" json:"signture"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *LoginRequest) GetVerifyId() string {
	if m != nil {
		return m.VerifyId
	}
	return ""
}

func (m *LoginRequest) GetVerifyValue() string {
	if m != nil {
		return m.VerifyValue
	}
	return ""
}

func (m *LoginRequest) GetSignture() string {
	if m != nil {
		return m.Signture
	}
	return ""
}

type LoginResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoginResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetBlockHeaderRequest struct {
}

func (m *GetBlockHeaderRequest) Reset()                    { *m = GetBlockHeaderRequest{} }
func (m *GetBlockHeaderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockHeaderRequest) ProtoMessage()               {}
func (*GetBlockHeaderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetBlockHeaderResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *BlockHeader `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetBlockHeaderResponse) Reset()                    { *m = GetBlockHeaderResponse{} }
func (m *GetBlockHeaderResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockHeaderResponse) ProtoMessage()               {}
func (*GetBlockHeaderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetBlockHeaderResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBlockHeaderResponse) GetData() *BlockHeader {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlockHeaderResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type BlockHeader struct {
	HeadBlockNum          uint32 `protobuf:"varint,1,opt,name=head_block_num,json=headBlockNum" json:"head_block_num"`
	HeadBlockHash         string `protobuf:"bytes,2,opt,name=head_block_hash,json=headBlockHash" json:"head_block_hash"`
	HeadBlockTime         uint64 `protobuf:"varint,3,opt,name=head_block_time,json=headBlockTime" json:"head_block_time"`
	HeadBlockDelegate     string `protobuf:"bytes,4,opt,name=head_block_delegate,json=headBlockDelegate" json:"head_block_delegate"`
	CursorLabel           uint32 `protobuf:"varint,5,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	LastConsensusBlockNum uint32 `protobuf:"varint,6,opt,name=last_consensus_block_num,json=lastConsensusBlockNum" json:"last_consensus_block_num"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BlockHeader) GetHeadBlockNum() uint32 {
	if m != nil {
		return m.HeadBlockNum
	}
	return 0
}

func (m *BlockHeader) GetHeadBlockHash() string {
	if m != nil {
		return m.HeadBlockHash
	}
	return ""
}

func (m *BlockHeader) GetHeadBlockTime() uint64 {
	if m != nil {
		return m.HeadBlockTime
	}
	return 0
}

func (m *BlockHeader) GetHeadBlockDelegate() string {
	if m != nil {
		return m.HeadBlockDelegate
	}
	return ""
}

func (m *BlockHeader) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *BlockHeader) GetLastConsensusBlockNum() uint32 {
	if m != nil {
		return m.LastConsensusBlockNum
	}
	return 0
}

type GetAccountInfoRequest struct {
	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
}

func (m *GetAccountInfoRequest) Reset()                    { *m = GetAccountInfoRequest{} }
func (m *GetAccountInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoRequest) ProtoMessage()               {}
func (*GetAccountInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAccountInfoRequest) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

type GetAccountInfoResponse struct {
	Code uint32           `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *AccountInfoData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string           `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetAccountInfoResponse) Reset()                    { *m = GetAccountInfoResponse{} }
func (m *GetAccountInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoResponse) ProtoMessage()               {}
func (*GetAccountInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetAccountInfoResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAccountInfoResponse) GetData() *AccountInfoData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetAccountInfoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AccountInfoData struct {
	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
	Pubkey      string `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey"`
}

func (m *AccountInfoData) Reset()                    { *m = AccountInfoData{} }
func (m *AccountInfoData) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoData) ProtoMessage()               {}
func (*AccountInfoData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AccountInfoData) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *AccountInfoData) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type FavoriteRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *FavoriteRequest) Reset()                    { *m = FavoriteRequest{} }
func (m *FavoriteRequest) String() string            { return proto.CompactTextString(m) }
func (*FavoriteRequest) ProtoMessage()               {}
func (*FavoriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *FavoriteRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FavoriteRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *FavoriteRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *FavoriteRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *FavoriteRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FavoriteRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *FavoriteRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *FavoriteRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *FavoriteRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *FavoriteRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type FavoriteResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *FavoriteResponse) Reset()                    { *m = FavoriteResponse{} }
func (m *FavoriteResponse) String() string            { return proto.CompactTextString(m) }
func (*FavoriteResponse) ProtoMessage()               {}
func (*FavoriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *FavoriteResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FavoriteResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FavoriteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetFavoriteRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	GoodsType string `protobuf:"bytes,2,opt,name=goods_type,json=goodsType" json:"goods_type"`
	PageSize  uint32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   uint32 `protobuf:"varint,4,opt,name=page_num,json=pageNum" json:"page_num"`
}

func (m *GetFavoriteRequest) Reset()                    { *m = GetFavoriteRequest{} }
func (m *GetFavoriteRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFavoriteRequest) ProtoMessage()               {}
func (*GetFavoriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetFavoriteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetFavoriteRequest) GetGoodsType() string {
	if m != nil {
		return m.GoodsType
	}
	return ""
}

func (m *GetFavoriteRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetFavoriteRequest) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetFavoriteResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *FavoriteArr `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetFavoriteResponse) Reset()                    { *m = GetFavoriteResponse{} }
func (m *GetFavoriteResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFavoriteResponse) ProtoMessage()               {}
func (*GetFavoriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetFavoriteResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetFavoriteResponse) GetData() *FavoriteArr {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetFavoriteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type FavoriteArr struct {
	PageNum  uint64          `protobuf:"varint,1,opt,name=page_num,json=pageNum" json:"page_num"`
	RowCount uint64          `protobuf:"varint,2,opt,name=row_count,json=rowCount" json:"row_count"`
	Row      []*FavoriteData `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *FavoriteArr) Reset()                    { *m = FavoriteArr{} }
func (m *FavoriteArr) String() string            { return proto.CompactTextString(m) }
func (*FavoriteArr) ProtoMessage()               {}
func (*FavoriteArr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *FavoriteArr) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *FavoriteArr) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *FavoriteArr) GetRow() []*FavoriteData {
	if m != nil {
		return m.Row
	}
	return nil
}

type FavoriteData struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	GoodsId   string `protobuf:"bytes,2,opt,name=goods_id,json=goodsId" json:"goods_id"`
	GoodsName string `protobuf:"bytes,3,opt,name=goods_name,json=goodsName" json:"goods_name"`
	Price     uint64 `protobuf:"varint,4,opt,name=price" json:"price"`
	Time      uint64 `protobuf:"varint,5,opt,name=time" json:"time"`
}

func (m *FavoriteData) Reset()                    { *m = FavoriteData{} }
func (m *FavoriteData) String() string            { return proto.CompactTextString(m) }
func (*FavoriteData) ProtoMessage()               {}
func (*FavoriteData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *FavoriteData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FavoriteData) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *FavoriteData) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *FavoriteData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *FavoriteData) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type GetBalanceRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username"`
	Random   string `protobuf:"bytes,2,opt,name=random" json:"random"`
}

func (m *GetBalanceRequest) Reset()                    { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()               {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *GetBalanceRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetBalanceRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

type GetBalanceResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetBalanceResponse) Reset()                    { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()               {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetBalanceResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBalanceResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetBalanceResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PushTxRequest)(nil), "PushTxRequest")
	proto.RegisterType((*PushTxResponse)(nil), "PushTxResponse")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*AccountInfo)(nil), "AccountInfo")
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*GetBlockHeaderRequest)(nil), "GetBlockHeaderRequest")
	proto.RegisterType((*GetBlockHeaderResponse)(nil), "GetBlockHeaderResponse")
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
	proto.RegisterType((*GetAccountInfoRequest)(nil), "GetAccountInfoRequest")
	proto.RegisterType((*GetAccountInfoResponse)(nil), "GetAccountInfoResponse")
	proto.RegisterType((*AccountInfoData)(nil), "AccountInfoData")
	proto.RegisterType((*FavoriteRequest)(nil), "FavoriteRequest")
	proto.RegisterType((*FavoriteResponse)(nil), "FavoriteResponse")
	proto.RegisterType((*GetFavoriteRequest)(nil), "GetFavoriteRequest")
	proto.RegisterType((*GetFavoriteResponse)(nil), "GetFavoriteResponse")
	proto.RegisterType((*FavoriteArr)(nil), "FavoriteArr")
	proto.RegisterType((*FavoriteData)(nil), "FavoriteData")
	proto.RegisterType((*GetBalanceRequest)(nil), "GetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "GetBalanceResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1024 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0x73, 0xf7, 0x71, 0x6e, 0x9d, 0xee, 0xb6, 0xde, 0x40, 0x45, 0xd6, 0x5a, 0xad, 0xf2,
	0xc2, 0x48, 0x14, 0xc1, 0xc2, 0xbe, 0x95, 0xae, 0xe8, 0x56, 0xaa, 0x2a, 0x14, 0x0a, 0x6f, 0x28,
	0x4c, 0xed, 0x89, 0x63, 0xad, 0xe3, 0x09, 0x33, 0x76, 0x4b, 0xf7, 0x0f, 0xf0, 0x06, 0xaf, 0x48,
	0x48, 0x08, 0xf1, 0x5b, 0xf8, 0x61, 0x68, 0x8e, 0xed, 0xc4, 0x49, 0xda, 0x68, 0xab, 0xbe, 0xf6,
	0x6d, 0xce, 0x65, 0x2e, 0xe7, 0x3b, 0xdf, 0x39, 0xc7, 0x06, 0x48, 0x14, 0x97, 0x74, 0x2e, 0x45,
	0x2c, 0x9c, 0x7f, 0xca, 0xd0, 0xfe, 0x2e, 0x51, 0xd3, 0x8b, 0x5f, 0x47, 0xfc, 0x97, 0x84, 0xab,
	0x98, 0xd8, 0xd0, 0xb8, 0xe2, 0x52, 0x05, 0x22, 0xb2, 0x8d, 0x81, 0x31, 0x6c, 0x8f, 0x72, 0x91,
	0x1c, 0x00, 0xb8, 0x89, 0x54, 0x42, 0x8e, 0xa3, 0x64, 0x66, 0x97, 0xd1, 0x68, 0xa6, 0x9a, 0xf3,
	0x64, 0x46, 0x9e, 0x43, 0x2b, 0x33, 0x87, 0xec, 0x92, 0x87, 0x76, 0x05, 0x1d, 0xac, 0x54, 0x77,
	0xa6, 0x55, 0xa4, 0x0f, 0xcd, 0x30, 0x98, 0xf0, 0x38, 0x98, 0x71, 0xbb, 0x3a, 0x30, 0x86, 0xd5,
	0xd1, 0x42, 0x26, 0x7b, 0x50, 0x57, 0x3c, 0xf2, 0xb8, 0xb4, 0x6b, 0x03, 0x63, 0x68, 0x8e, 0x32,
	0x49, 0xef, 0x71, 0x45, 0x14, 0x4b, 0xe6, 0xc6, 0x76, 0x1d, 0x2d, 0x0b, 0x59, 0xef, 0x99, 0xf1,
	0x78, 0x2a, 0x3c, 0xbb, 0x91, 0xee, 0x49, 0x25, 0xf2, 0x04, 0x6a, 0x73, 0x26, 0xd9, 0xcc, 0x6e,
	0xa2, 0x3a, 0x15, 0xc8, 0x3e, 0x34, 0x54, 0xe0, 0x8f, 0x59, 0xe8, 0xdb, 0x26, 0xbe, 0xad, 0xae,
	0x02, 0xff, 0x28, 0xf4, 0xc9, 0xc7, 0x60, 0xaa, 0xc0, 0x8f, 0x58, 0x9c, 0x48, 0x6e, 0x03, 0x6e,
	0x59, 0x2a, 0x9c, 0x2f, 0xa1, 0x93, 0x23, 0xa4, 0xe6, 0x22, 0x52, 0x9c, 0x10, 0xa8, 0xba, 0xc2,
	0xe3, 0x19, 0x3e, 0xb8, 0x26, 0x3d, 0xa8, 0xcc, 0x94, 0x8f, 0xa8, 0x98, 0x23, 0xbd, 0x74, 0xfe,
	0x34, 0xa0, 0x3b, 0xe2, 0x7e, 0xa0, 0x62, 0x2e, 0x73, 0x70, 0x5f, 0x42, 0x83, 0xb9, 0xae, 0x48,
	0xa2, 0x18, 0x37, 0x5b, 0x87, 0x2d, 0x7a, 0x94, 0xca, 0xa7, 0xd1, 0x44, 0x8c, 0x72, 0x23, 0x39,
	0x80, 0xaa, 0x4e, 0x12, 0x1e, 0x67, 0x1d, 0x9a, 0xf4, 0x07, 0xc5, 0x25, 0x7a, 0xa0, 0x9a, 0x7c,
	0x04, 0xe6, 0x15, 0x97, 0xc1, 0xe4, 0x66, 0x1c, 0x78, 0x88, 0xb3, 0x39, 0x6a, 0xa6, 0x8a, 0x53,
	0x4f, 0xe7, 0x21, 0x33, 0x5e, 0xb1, 0x30, 0x49, 0x81, 0x36, 0x47, 0x56, 0xaa, 0xfb, 0x51, 0xab,
	0x9c, 0xaf, 0xc1, 0x2a, 0x5c, 0xab, 0xe3, 0x89, 0xd8, 0x8c, 0x67, 0x8f, 0xc7, 0xb5, 0x86, 0x76,
	0x9e, 0x5c, 0xbe, 0xe3, 0x37, 0xd9, 0xf9, 0x99, 0xe4, 0xfc, 0x55, 0x86, 0x66, 0xfe, 0x9a, 0x47,
	0xae, 0xac, 0x73, 0xe5, 0x2b, 0xe8, 0x2d, 0x53, 0x7e, 0x2f, 0xb6, 0xfc, 0x6d, 0x40, 0xeb, 0x4c,
	0xf8, 0x41, 0x94, 0x53, 0xa5, 0x0f, 0x4d, 0x9d, 0x6b, 0x4c, 0x8c, 0x91, 0xc6, 0x92, 0xcb, 0x3a,
	0x16, 0xc9, 0x22, 0x4f, 0xcc, 0xb2, 0x13, 0x32, 0xe9, 0xa1, 0xbc, 0xd0, 0x77, 0xea, 0x58, 0x30,
	0xb6, 0x14, 0xd9, 0x85, 0xec, 0x7c, 0x01, 0xed, 0xec, 0x7d, 0xf7, 0x8a, 0x6b, 0x1f, 0x9e, 0x9e,
	0xf0, 0xf8, 0x9b, 0x50, 0xb8, 0xef, 0xde, 0x72, 0xe6, 0x2d, 0x4a, 0xc1, 0xf9, 0x19, 0xf6, 0xd6,
	0x0d, 0x5b, 0x0e, 0x1e, 0x40, 0xd5, 0x63, 0x31, 0xcb, 0x0a, 0xa2, 0x45, 0x8b, 0xfb, 0xd0, 0x92,
	0x5f, 0x5d, 0x59, 0x5e, 0xfd, 0x47, 0x19, 0xac, 0x82, 0x1f, 0x79, 0x01, 0x9d, 0x29, 0x67, 0xde,
	0xf8, 0x52, 0xeb, 0x90, 0x97, 0xe9, 0x0d, 0x2d, 0xad, 0x45, 0x47, 0x4d, 0xcd, 0x97, 0xd0, 0x2d,
	0x78, 0x4d, 0x99, 0x9a, 0x66, 0xe1, 0xb4, 0x17, 0x6e, 0x6f, 0x99, 0x9a, 0xae, 0xf9, 0x21, 0x4d,
	0x2b, 0x48, 0xd3, 0xa5, 0xdf, 0x85, 0xe6, 0x2a, 0x85, 0xdd, 0x82, 0x9f, 0xc7, 0x43, 0xee, 0xb3,
	0x38, 0x47, 0x7f, 0x67, 0xe1, 0xfb, 0x26, 0x33, 0x6c, 0x94, 0x46, 0x6d, 0xb3, 0x34, 0x5e, 0x81,
	0x1d, 0x32, 0x15, 0x8f, 0x5d, 0x0d, 0x57, 0xa4, 0x12, 0x55, 0x08, 0xa9, 0x8e, 0xee, 0x4f, 0xb5,
	0xfd, 0x38, 0x37, 0xe7, 0xb1, 0x39, 0xaf, 0x31, 0x19, 0xc5, 0x8e, 0x93, 0x91, 0xed, 0x39, 0xb4,
	0xb2, 0xd6, 0x33, 0x2e, 0x10, 0xce, 0xca, 0x74, 0xe7, 0x6c, 0xc6, 0x1d, 0x0f, 0xf3, 0xb5, 0xb2,
	0x77, 0x4b, 0xbe, 0x5e, 0xac, 0xe4, 0xab, 0x57, 0xec, 0x72, 0x6f, 0x58, 0xcc, 0xee, 0xcc, 0xd9,
	0x19, 0x74, 0xd7, 0x5c, 0x3f, 0xe0, 0x6d, 0x85, 0x66, 0x55, 0x5e, 0x69, 0x56, 0xff, 0x96, 0xa1,
	0xfb, 0x2d, 0xbb, 0x12, 0x32, 0x88, 0xf9, 0xe3, 0x7c, 0xbb, 0xa3, 0x67, 0x9d, 0x41, 0x6f, 0x89,
	0xd1, 0x96, 0x94, 0x92, 0x42, 0x4a, 0xcd, 0x3b, 0x13, 0xf8, 0x9b, 0x01, 0xe4, 0x84, 0xc7, 0xeb,
	0xa8, 0x6f, 0xeb, 0x66, 0x07, 0x00, 0xbe, 0x10, 0x9e, 0x1a, 0xc7, 0x37, 0xf3, 0x7c, 0x08, 0x99,
	0xa8, 0xb9, 0xb8, 0x99, 0x73, 0xdd, 0xd4, 0xe6, 0xcc, 0xe7, 0x63, 0x15, 0xbc, 0xe7, 0x19, 0xe8,
	0x4d, 0xad, 0xf8, 0x3e, 0x78, 0xcf, 0xc9, 0x33, 0xc0, 0x35, 0x66, 0xac, 0x9a, 0xa6, 0x53, 0xcb,
	0x9a, 0xec, 0x3f, 0xc1, 0xee, 0xca, 0x43, 0xee, 0xd1, 0x5d, 0xf2, 0x4d, 0x47, 0xf2, 0xee, 0xee,
	0x32, 0x01, 0xab, 0xe0, 0xb6, 0xf2, 0x10, 0x03, 0x53, 0x9f, 0x3f, 0x44, 0x07, 0x20, 0xc5, 0xf5,
	0x38, 0x1d, 0xfb, 0xe5, 0x94, 0x16, 0x52, 0x5c, 0x1f, 0xe3, 0xa4, 0xff, 0x04, 0x2a, 0x52, 0x5c,
	0xdb, 0x95, 0x41, 0x65, 0x68, 0x1d, 0xb6, 0x17, 0x37, 0x63, 0x91, 0x68, 0x8b, 0xf3, 0xbb, 0x01,
	0xad, 0xa2, 0x76, 0x2b, 0x94, 0xcf, 0xa0, 0x99, 0x42, 0x19, 0x78, 0x19, 0x90, 0x0d, 0x94, 0x4f,
	0xbd, 0x25, 0xca, 0xb8, 0xb1, 0x52, 0x40, 0x19, 0x4b, 0x48, 0x53, 0x4a, 0x06, 0x6e, 0xce, 0xdb,
	0x54, 0xd0, 0x60, 0x21, 0x99, 0x6b, 0xa8, 0xc4, 0xb5, 0x73, 0x02, 0x3b, 0xba, 0x71, 0xb3, 0x90,
	0x45, 0x2e, 0x7f, 0xc0, 0xb4, 0x72, 0xce, 0x91, 0x29, 0x8b, 0x83, 0x1e, 0x4a, 0xbd, 0xc3, 0xff,
	0x2a, 0x50, 0xd5, 0x9f, 0x26, 0xe4, 0x33, 0x68, 0xe6, 0x53, 0x98, 0xf4, 0xe8, 0xda, 0x37, 0x58,
	0x7f, 0x87, 0xae, 0x8f, 0x68, 0xa7, 0x44, 0x86, 0x50, 0xc3, 0xe9, 0x46, 0xda, 0xb4, 0x38, 0x85,
	0xfb, 0x1d, 0xba, 0x32, 0xf4, 0x9c, 0x12, 0x39, 0x86, 0xce, 0xea, 0xdc, 0x22, 0x7b, 0xf4, 0xd6,
	0x09, 0xd7, 0xdf, 0xa7, 0xb7, 0x0f, 0xb8, 0xc5, 0x21, 0xc5, 0x6f, 0x30, 0x3c, 0x64, 0xb3, 0x33,
	0xa7, 0x87, 0xdc, 0xd2, 0x75, 0x9d, 0x92, 0x0e, 0x33, 0x27, 0x06, 0xe9, 0xd1, 0xb5, 0x8a, 0xeb,
	0xef, 0xd0, 0x75, 0xea, 0x3b, 0x25, 0xf2, 0x1a, 0xac, 0x42, 0x4d, 0x90, 0x5d, 0xba, 0x59, 0xaa,
	0xfd, 0x27, 0xf4, 0x96, 0xb2, 0x71, 0x4a, 0xe4, 0x53, 0x68, 0x5e, 0x48, 0x16, 0xa9, 0x09, 0x97,
	0xa4, 0x43, 0x57, 0x7e, 0x1a, 0xfa, 0x5d, 0xba, 0xfa, 0x89, 0xec, 0x94, 0xc8, 0x2b, 0x80, 0x65,
	0x76, 0x09, 0xa1, 0x1b, 0x9c, 0xe9, 0xef, 0xd2, 0xcd, 0xf4, 0x3b, 0xa5, 0xcb, 0x3a, 0xfe, 0x99,
	0x7c, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x0d, 0xb9, 0x4b, 0xa7, 0x0c, 0x00, 0x00,
}
