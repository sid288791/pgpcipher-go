package openPGP

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"

	"io/ioutil"
	"os"
)

func DecryptFile(encryptedFilePath, privateKeyPath, privateKeyPassphrase, outputFilePath string) error {

	encryptedFile, err := os.Open(encryptedFilePath)
	if err != nil{
		return err
	}
	defer encryptedFile.Close()

	privateKeyFile,err := os.Open(privateKeyPath)
	if err != nil{
		return err
	}
	defer privateKeyFile.Close()

	privateKeyRing, err := openpgp.ReadArmoredKeyRing(privateKeyFile)
	if err != nil{
		return err
	}

	armorReader,err := armor.Decode(encryptedFile)
	if err != nil{
		return err
	}

	_= privateKeyRing[0].privateKeyFile

	decryptedMessage, err := openpgp.ReadMessage(armorReader.Body,privateKeyRing,func(keys []openpgp.Key, symmetric bool)([]byte,error){
		for _,key := range keys {
			if key.PrivateKey !=nil && key.PrivateKey.Encrypted{
				err := key.PrivateKey.Decrypt ([]byte(privateKeyPassphrase))
				if err != nil{
					return err
				}
			}
		}
		return nil,nil
	},nil
)
if err != nil{
	return err
}

reader := bufio.NewReader(decryptedMessage.UnverifiedBody)
decryptedData, err := ioutil.ReadAll(reader)
if err != nil{
	return err
}

err = ioutil.WriteFile(outputFilePath, decryptedData,0644)
if err != nil{
	return err
}
return nil
}