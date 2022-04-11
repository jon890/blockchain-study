package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jon890/foscoin/blockchain"
	"github.com/jon890/foscoin/p2p"
	"github.com/jon890/foscoin/utils"
	"github.com/jon890/foscoin/wallet"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type addTxPayload struct {
	To     string
	Amount int
}

type myWalletResponse struct {
	Address string `json:"address"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addPeerPayload struct {
	Address, Port string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of block chain",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "Add a Block",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
		{
			URL:         url("/mempool"),
			Method:      "GET",
			Description: "See a mempool",
		},
		{
			URL:         url("/wallet"),
			Method:      "GET",
			Description: "get my wallet pubkey",
		},
		{
			URL:         url("/ws"),
			Method:      "GET",
			Description: "Upgrade to WebSockets",
		},
	}
	json.NewEncoder(rw).Encode(data)

	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain()))
	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		next.ServeHTTP(rw, r)
	})
}

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")

	switch total {
	case "true":
		amount := blockchain.BalanceByAddress(blockchain.Blockchain(), address)
		json.NewEncoder(rw).Encode(balanceResponse{address, amount})
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(blockchain.Blockchain(), address)))

	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))

}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errorResponse{err.Error()})
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func myWallet(rw http.ResponseWriter, r *http.Request) {
	address := wallet.Wallet().Address
	json.NewEncoder(rw).Encode(myWalletResponse{Address: address})
}

func peers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var payload addPeerPayload
		json.NewDecoder(r.Body).Decode(&payload)
		p2p.AddPeer(payload.Address, payload.Port, port)
		rw.WriteHeader(http.StatusOK)
	case "GET":
		json.NewEncoder(rw).Encode(p2p.Peers)
	}
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance)
	router.HandleFunc("/mempool", mempool).Methods("GET")
	router.HandleFunc("/wallet", myWallet).Methods("GET")
	router.HandleFunc("/transactions", transactions).Methods("POST")
	router.HandleFunc("/ws", p2p.Upgrade).Methods("GET")
	router.HandleFunc("/peers", peers).Methods("GET", "POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
