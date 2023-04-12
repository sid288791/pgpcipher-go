package openPGP

import (
	"io"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func EncryptFile(inputFile string,outputFile string,publicKey string){

	input,err := os.Open(inputFile)
	if err != nil{
		return err
	}
	defer input.Close()

	output,err := os.Create(outputFile)
	if err != nil{
		return nil
	}
	defer output.Close()

	keyringFile, err :=os.Open(publicKey)
	if err != nil{
		return nil
	}
	defer keyringFile.Close()

	keyring, err := openpgp.ReadArmoredRing(keyringFile)
	if err != nil{
		return nil
	}

	armoredwriter,err := armor.Encode(output,"PGP MESSAGE",make(map[string]string))
	if err != nil{
		return nil
	}

	w,err := openpgp.Encrypt(armoredwriter,keyring,nil,nil.nil)
	if err != nil{
		return nil
	}

	_,err = io.Copy(w,input)
	if err != nil{
		return nil
	}

	err = w.Close()
	if err != nil{
		return nil
	}

	return nil

}