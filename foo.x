namespace stellar
{

  union PublicKey switch (int type)
  {
  case 1:
      int ed25519;
  };


  typedef PublicKey AccountID;
  typedef AccountID* SponsorshipDescriptor;

  struct AccountEntryExtensionV2
  {
    SponsorshipDescriptor signerSponsoringIDs<1>;

  };

  struct LedgerEntryExtensionV1
  {
    SponsorshipDescriptor sponsoringID;
  };
}
