import xdr from './generated/stellar-xdr_generated.js'

const sponsor = xdr.PublicKey.publicKeyTypeEd25519(9);
const entry = new xdr.AccountEntryExtensionV2({
  signerSponsoringIDs: [null, sponsor],
  sponsoringId: xdr.PublicKey.publicKeyTypeEd25519(3),
});

process.stdout.write(entry.toXDR('base64'));
// const read = xdr.AccountEntryExtensionV2.fromXDR(entry.toXDR('base64'), 'base64');
// console.log(read.signerSponsoringIDs()[1])
