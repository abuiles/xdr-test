enum PublicKeyType
{
    PUBLIC_KEY_TYPE_ED25519 = 2
};

union PublicKey switch (PublicKeyType type)
  {
  case PUBLIC_KEY_TYPE_ED25519:
    int ed25519;
  };


typedef PublicKey AccountID;
typedef AccountID* SponsorshipDescriptor;

struct AccountEntryExtensionV2
{
  SponsorshipDescriptor signerSponsoringIDs<2>;
  SponsorshipDescriptor sponsoringID;
};
