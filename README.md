# Currency Converter 💱

Um conversor de moedas via linha de comando (CLI), desenvolvido em Go, que consulta taxas de câmbio em tempo real através de uma API pública.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-conclu%C3%ADdo-brightgreen)

## 📖 Sobre o projeto

O **Currency Converter** é uma ferramenta de linha de comando que converte valores entre moedas usando taxas de câmbio atualizadas, consultadas em tempo real na [Frankfurter API](https://frankfurter.dev/) — uma API pública, gratuita e sem necessidade de chave de acesso, que fornece taxas de referência do Banco Central Europeu.

Foi construído com foco em praticar o consumo de APIs externas em Go: fazer requisições HTTP como cliente, decodificar respostas JSON dinâmicas (usando maps) e tratar erros de rede e de entrada do usuário.

## ✨ Funcionalidades

- 💵 Conversão entre qualquer par de moedas suportado pela API
- 🔤 Aceita códigos de moeda em maiúsculas ou minúsculas
- ⚠️ Tratamento de erros: valores inválidos, moedas inexistentes, falhas de rede

## 🚀 Como usar

### Pré-requisitos

- [Go](https://go.dev/dl/) 1.22 ou superior
- Conexão com a internet (a ferramenta consulta uma API externa)

### Instalação

```bash
git clone https://github.com/fshenrique6/currency-converter.git
cd currency-converter
go build -o converter
```

### Uso

```bash
./converter <valor> <moeda_origem> <moeda_destino>
```

### Exemplos

```bash
$ ./converter 100 USD BRL
100.00 USD = 511.53 BRL

$ ./converter 50 usd eur
50.00 USD = 43.28 EUR

$ ./converter 100 USD XYZ
erro: API retornou status 404 (verifique os códigos de moeda)

$ ./converter 100 USD
uso: converter <valor> <moeda_origem> <moeda_destino>
exemplo: converter 100 USD BRL
```

## 🏗️ Estrutura do projeto

```
currency-converter/
├── go.mod
├── main.go                  # Ponto de entrada, leitura e validação dos argumentos
└── converter/
    └── converter.go         # Consumo da API externa e cálculo da conversão
```

## 🛠️ Tecnologias e conceitos aplicados

- [Go](https://go.dev/)
- Consumo de API REST externa (`net/http` como cliente)
- Decodificação de JSON dinâmico com `map[string]float64`
- Tratamento de erros com `fmt.Errorf` e `%w` (error wrapping)
- [Frankfurter API](https://frankfurter.dev/) — taxas de câmbio públicas e gratuitas

## 🔮 Próximos passos

- [ ] Suporte a taxas históricas (conversão em datas específicas)
- [ ] Modo interativo (loop de conversões sem reiniciar o programa)
- [ ] Cache local das taxas para reduzir chamadas à API


## 👤 Autor

**Henrique Souza**
Projeto desenvolvido para fins de aprendizado e portfólio.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)

---

# Currency Converter 💱 *(English version)*

A command-line currency converter built in Go, which fetches real-time exchange rates from a public API.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-completed-brightgreen)

## 📖 About the project

**Currency Converter** is a command-line tool that converts amounts between currencies using up-to-date exchange rates, fetched in real time from the [Frankfurter API](https://frankfurter.dev/) — a free, public API that requires no access key and provides reference rates from the European Central Bank.

It was built to practice consuming external APIs in Go: making HTTP requests as a client, decoding dynamic JSON responses (using maps), and handling both network errors and invalid user input.

## ✨ Features

- 💵 Conversion between any currency pair supported by the API
- 🔤 Accepts currency codes in uppercase or lowercase
- ⚠️ Error handling for invalid amounts, unknown currencies, and network failures

## 🚀 Getting started

### Prerequisites

- [Go](https://go.dev/dl/) 1.22 or higher
- Internet connection (the tool queries an external API)

### Installation

```bash
git clone https://github.com/fshenrique6/currency-converter.git
cd currency-converter
go build -o converter
```

### Usage

```bash
./converter <amount> <from_currency> <to_currency>
```

### Examples

```bash
$ ./converter 100 USD BRL
100.00 USD = 511.53 BRL

$ ./converter 50 usd eur
50.00 USD = 43.28 EUR

$ ./converter 100 USD XYZ
error: API returned status 404 (check the currency codes)

$ ./converter 100 USD
usage: converter <amount> <from_currency> <to_currency>
example: converter 100 USD BRL
```

## 🏗️ Project structure

```
currency-converter/
├── go.mod
├── main.go                  # Entry point, argument parsing and validation
└── converter/
    └── converter.go         # External API consumption and conversion logic
```

## 🛠️ Technologies and concepts applied

- [Go](https://go.dev/)
- Consuming an external REST API (`net/http` as a client)
- Decoding dynamic JSON with `map[string]float64`
- Error handling with `fmt.Errorf` and `%w` (error wrapping)
- [Frankfurter API](https://frankfurter.dev/) — free, public exchange rate data

## 🔮 Roadmap

- [ ] Support for historical rates (conversion on specific dates)
- [ ] Interactive mode (repeated conversions without restarting the program)
- [ ] Local caching of rates to reduce API calls


## 👤 Author

**Henrique Souza**
Project built for learning and portfolio purposes.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)
