package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const walletFile = "./tmp/wallets.data"

// Stores the wallets inside a map of type: address -> Wallet.
type Wallets struct {
	Wallets map[string]*Wallet
}

// Populates the Wallets struct.
func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()

	return &wallets, err
}

// Adds a wallet to the Wallets struct.
func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet

	return address
}

// Gets all the wallet adresses inside the Wallets struct.
func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// Gets a wallet with its corresponding string address.
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

// Loads all the wallets indide the local file.
func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets

	return nil
}

// Saves the wallets in a local file.
func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		s := strings.Split(walletFile, "/")
		pathToWalletFile := strings.Join(s[:len(s)-1], "/")
		os.MkdirAll(pathToWalletFile, os.ModePerm)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
