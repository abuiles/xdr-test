// Package stellar-xdr is generated from:
//
//  src/foo.x
//
// DO NOT EDIT or your changes may be overwritten
package main

import (
	"bytes"
	"encoding"
	"fmt"
	"io"

	xdr "github.com/stellar/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// PublicKeyType is an XDR Enum defines as:
//
//   enum PublicKeyType
//    {
//        PUBLIC_KEY_TYPE_ED25519 = 2
//    };
//
type PublicKeyType int32

const (
	PublicKeyTypePublicKeyTypeEd25519 PublicKeyType = 2
)

var publicKeyTypeMap = map[int32]string{
	2: "PublicKeyTypePublicKeyTypeEd25519",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PublicKeyType
func (e PublicKeyType) ValidEnum(v int32) bool {
	_, ok := publicKeyTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e PublicKeyType) String() string {
	name, _ := publicKeyTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PublicKeyType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PublicKeyType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PublicKeyType)(nil)
	_ encoding.BinaryUnmarshaler = (*PublicKeyType)(nil)
)

// PublicKey is an XDR Union defines as:
//
//   union PublicKey switch (PublicKeyType type)
//      {
//      case PUBLIC_KEY_TYPE_ED25519:
//        int ed25519;
//      };
//
type PublicKey struct {
	Type    PublicKeyType
	Ed25519 *int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PublicKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u PublicKey) ArmForSwitch(sw int32) (string, bool) {
	switch PublicKeyType(sw) {
	case PublicKeyTypePublicKeyTypeEd25519:
		return "Ed25519", true
	}
	return "-", false
}

// NewPublicKey creates a new  PublicKey.
func NewPublicKey(aType PublicKeyType, value interface{}) (result PublicKey, err error) {
	result.Type = aType
	switch PublicKeyType(aType) {
	case PublicKeyTypePublicKeyTypeEd25519:
		tv, ok := value.(int32)
		if !ok {
			err = fmt.Errorf("invalid value, must be int32")
			return
		}
		result.Ed25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u PublicKey) MustEd25519() int32 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PublicKey) GetEd25519() (result int32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PublicKey) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PublicKey) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PublicKey)(nil)
	_ encoding.BinaryUnmarshaler = (*PublicKey)(nil)
)

// AccountId is an XDR Typedef defines as:
//
//   typedef PublicKey AccountID;
//
type AccountId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u AccountId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewAccountId creates a new  AccountId.
func NewAccountId(aType PublicKeyType, value interface{}) (result AccountId, err error) {
	u, err := NewPublicKey(aType, value)
	result = AccountId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u AccountId) MustEd25519() int32 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountId) GetEd25519() (result int32, ok bool) {
	return PublicKey(u).GetEd25519()
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountId)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountId)(nil)
)

// SponsorshipDescriptor is an XDR Typedef defines as:
//
//   typedef AccountID* SponsorshipDescriptor;
//
type SponsorshipDescriptor AccountId

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SponsorshipDescriptor) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SponsorshipDescriptor) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SponsorshipDescriptor)(nil)
	_ encoding.BinaryUnmarshaler = (*SponsorshipDescriptor)(nil)
)

// AccountEntryExtensionV2 is an XDR Struct defines as:
//
//   struct AccountEntryExtensionV2
//    {
//      SponsorshipDescriptor signerSponsoringIDs<2>;
//      SponsorshipDescriptor sponsoringID;
//    };
//
type AccountEntryExtensionV2 struct {
	SignerSponsoringIDs []*SponsorshipDescriptor `xdrmaxsize:"2"`
	SponsoringId        *SponsorshipDescriptor
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExtensionV2) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExtensionV2) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExtensionV2)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExtensionV2)(nil)
)

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
