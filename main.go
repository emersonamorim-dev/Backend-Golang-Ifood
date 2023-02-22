package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"googlemaps.github.io/maps"
)

var (
	client *maps.Client
	)
	
	func main() {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
	fmt.Println("Chave da API do Google Maps não definida")
	return
	}

// Cria um cliente para a API do Google Maps
var err error
client, err = maps.NewClient(maps.WithAPIKey(apiKey))
if err != nil {
	fmt.Printf("Erro ao criar cliente para a API do Google Maps: %v\n", err)
	return
}

	router := gin.Default()
	router.POST("/solicitar-pedido", solicitarPedidoHandler)
	router.GET("/pedidos", listarPedidosHandler)	
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

// Definição da struct de Pedido
type Pedido struct {
	ID int json:"id"
	Cliente string json:"cliente" validate:"required"
	Endereco string json:"endereco" validate:"required"
	Items []string json:"items" validate:"required"
	}
	
// Handler para a rota de solicitação de pedido
func solicitarPedidoHandler(c *gin.Context) {
    // Decodifica o corpo da requisição JSON em uma struct Pedido
    var pedido Pedido
    if err := c.ShouldBindJSON(&pedido); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    // Verifica se o cliente já existe
    if cliente == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cliente não encontrado"})
        return
    }

    // Verifica se o endereço já existe
    if pedido.Endereco == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Endereço não encontrado"})
        return
    }

    // Calcula o valor total do pedido
    var valorTotal float64
    for _, item := range pedido.Items {
        valorTotal += buscarValorItemNoBancoDeDados(item)
    }
    valorTotal = valorTotal / float64(len(pedido.Items))

    // Gera o payload Pix para o pagamento do pedido
    payload, err := pixpayload.PagamentoPayload{
        PixKey:       "sua_chave_pix",
        Description:  "Pagamento do pedido",
        MerchantName: "Nome do seu estabelecimento",
        Amount:       valorTotal,
        TxID:         "id_da_transacao",
    }.Payload()

    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    // Cria a struct Pedidos com os dados do pedido
    restaurante := Restaurante{
        ID:       1,
        Nome:     "Restaurante 1",
        Endereco: "Rua A, 123",
        Cozinha:  "Pizza",
        Distancia: 500,
    }
    pagamento := Pagamento{
        Method: "PIX",
        Amount: int(valorTotal * 100),
    }
    pedidoNovo := Pedidos{
        ID:          1,
        Restaurante: restaurante,
        Items:       pedido.Items,
        Total:       int(valorTotal * 100),
        Pagamento:   pagamento,
    }

    // Adiciona o pedido à lista de pedidos
    pedidos = append(pedidos, pedidoNovo)

    // Retorna a resposta para o cliente
    c.JSON(http.StatusOK, gin.H{
        "message": "Pedido realizado com sucesso!",
        "payload": payload,
    })
}

 // Verifica se o cliente já existe
cliente, err := buscarClienteNoBancoDeDados(pedido.Cliente)
if err != nil || cliente == nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Cliente não encontrado"})
	return
}

// Verifica se o endereço já existe
if pedido.Endereco == "" {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Endereço não encontrado"})
	return
}

// Calcula o valor total do pedido
var valorTotal float64
for _, item := range pedido.Items {
	valorTotal += buscarValorItemNoBancoDeDados(item)
}

valorTotal = valorTotal / float64(len(pedido.Items))

// Gera o payload Pix para o pagamento do pedido
payload, err := pixpayload.PagamentoPayload{
	PixKey:       "sua_chave_pix",
	Description:  "Pagamento do pedido",
	MerchantName: "Nome do seu estabelecimento",
	Amount:       valorTotal,
	TxID:         "id_da_transacao",
}.Payload()

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

type Restaurante struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Cozinha  string `json:"cozinha"`
	Distancia int    `json:"distancia"`
}

type Pagamento struct {
	Method string `json:"method"`
	Amount int    `json:"amount"`
}

type PedidoResponse struct {
	ID          int         `json:"id"`
	Restaurante Restaurante `json:"restaurante"`
	Items       []string    `json:"items"`
	Total       float64     `json:"total"`
	Pagamento   Pagamento   `json:"pagamento"`
}
var pedidos []Pedidos

// Simula a consulta a uma API de mapas para obter restaurantes próximos
func getRestaurante(w http.ResponseWriter, r *http.Request) {
	restaurantes := []Restaurante{
		{
			ID:       1,
			Nome:     "Restaurante 1",
			Endereco:  "Rua A, 123",
			Cozinha:  "Pizza",
			Distancia: 500,
		},
		{
			ID:       2,
			Nome:     "Restaurante 2",
			Endereco:  "Rua B, 456",
			Cozinha:  "Hambúrguer",
			Distancia: 750,
		},
		{
			ID:       3,
			Nome:     "Restaurante 3",
			Endereco:  "Rua C, 789",
			Cozinha:  "Churrasco",
			Distancia: 1000,
		},
	}

	// Converte os dados em JSON e envia a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
func solicitarPedidoHandler(c *gin.Context) {
	// Decodifica o corpo da requisição JSON em uma struct Pedido
	var pedido Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Calcula o valor total do pedido
	var valorTotal float64
	for _, item := range pedido.Items {
		valorTotal += buscarValorItemNoBancoDeDados(item)
	}

	// Gera o payload Pix para o pagamento do pedido
	payload, err := pixpayload.PagamentoPayload{
		PixKey:       "sua_chave_pix",
		Description:  "Pagamento do pedido",
		MerchantName: "Nome do seu estabelecimento",
		Amount:       valorTotal,
		TxID:         "id_da_transacao",
	}.Payload()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Envie a resposta com o payload do Pix
	c.JSON(http.StatusOK, gin.H{
		"payload": payload,
	})
}


func createpedido(w http.ResponseWriter, r *http.Request) {
	// Simula a criação de um pedido com Pix como forma de pagamento
	var pedido pedido
	json.NewDecoder(r.Body).Decode(&pedido)
	pedido.ID = len(pedidos) + 1
	pedido.Payment.Method = "Pix"
	pedido.Payment.Amount = pedido.Total
	pedidos = append(pedidos, pedido)

	// Converte os dados em JSON e envia a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

// Obtém o ID do pedido a partir dos parâmetros da URL
func getpedido(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Procura o pedido correspondente na lista de pedidos
	var pedido Pedido
	for _, item := range pedidos {
		if fmt.Sprintf("%v", item.ID) == id {
			pedido = item
			break
		}
	}

	// Se o pedido não foi encontrado, retorna um erro
	if pedido.ID == 0 {
		http.NotFound(w, r)
		return
	}
}
