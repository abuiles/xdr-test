// Automatically generated on 2020-07-29T17:03:38-05:00
// DO NOT EDIT or your changes may be overwritten

/* jshint maxstatements:2147483647  */
/* jshint esnext:true  */

import XDR from 'js-xdr';


var types = XDR.config(xdr => {
// === xdr source ============================================================
//
//   enum PublicKeyType
//   {
//       PUBLIC_KEY_TYPE_ED25519 = 2
//   };
//
// ===========================================================================
xdr.enum("PublicKeyType", {
  publicKeyTypeEd25519: 2,
});

// === xdr source ============================================================
//
//   union PublicKey switch (PublicKeyType type)
//     {
//     case PUBLIC_KEY_TYPE_ED25519:
//       int ed25519;
//     };
//
// ===========================================================================
xdr.union("PublicKey", {
  switchOn: xdr.lookup("PublicKeyType"),
  switchName: "type",
  switches: [
    ["publicKeyTypeEd25519", "ed25519"],
  ],
  arms: {
    ed25519: xdr.int(),
  },
});

// === xdr source ============================================================
//
//   typedef PublicKey AccountID;
//
// ===========================================================================
xdr.typedef("AccountId", xdr.lookup("PublicKey"));

// === xdr source ============================================================
//
//   typedef AccountID* SponsorshipDescriptor;
//
// ===========================================================================
xdr.typedef("SponsorshipDescriptor", xdr.option(xdr.lookup("AccountId")));

// === xdr source ============================================================
//
//   struct AccountEntryExtensionV2
//   {
//     SponsorshipDescriptor signerSponsoringIDs<2>;
//     SponsorshipDescriptor sponsoringID;
//   };
//
// ===========================================================================
xdr.struct("AccountEntryExtensionV2", [
  ["signerSponsoringIDs", xdr.varArray(xdr.lookup("SponsorshipDescriptor"), 2)],
  ["sponsoringId", xdr.lookup("SponsorshipDescriptor")],
]);

});
export default types;
