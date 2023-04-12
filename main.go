package main

import(
	"crypto"
	"golang.org/x/crypto/ripemd160"
	"pgpcipher-go/openPGP"
)
func init(){
	crypto.RegisterHash(crypto.RIPEMD160,ripemd160.New)
}

func main(){
	pgp_pub_key = ""
    file_to_enc = ""
    enc_output_file = ""
    openPGP.EncryptFile(file_to_enc,enc_output_file,pgp_pub_key)

    pgp_priv_key = ""
    passphrase = ""
    enc_file = ""
    dec_output_file = ""
    openPGP.DecryptFile(enc_file,pgp_priv_key, passphrase, dec_output_file)
}
