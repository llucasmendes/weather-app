Claro! Aqui está um README detalhado para o projeto:

---

# Weather App

Este é um sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O sistema está preparado para ser publicado no Google Cloud Run.

#### Teste no google cloud Run

```sh
curl -i "https://messengage-test-ohymxcqbya-uc.a.run.app/weather/49035655"
```

## Estrutura do Projeto

```
weather-app/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── weather/
    ├── handler.go
    ├── weather.go
    └── weather_test.go
```

- **Dockerfile**: Arquivo de configuração do Docker para construir a imagem do contêiner.
- **go.mod**: Arquivo de dependências do Go.
- **go.sum**: Arquivo de soma de dependências do Go.
- **main.go**: Ponto de entrada da aplicação.
- **weather/handler.go**: Implementação dos handlers HTTP.
- **weather/weather.go**: Funções para buscar informações do CEP e do clima.
- **weather/weather_test.go**: Testes automatizados das funções.

## Requisitos

- Go 1.20 ou superior
- Docker
- Conta no Google Cloud e configuração do Google Cloud SDK

## Configuração

### 1. Clonar o Repositório

```sh
git clone https://github.com/seu-usuario/weather-app.git
cd weather-app
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto e adicione sua chave da API do WeatherAPI:

```
WEATHER_API_KEY=your_weather_api_key
```

### 3. Construir o Projeto

#### Construir com Go

Para construir e executar o projeto localmente com Go, utilize:

```sh
go run main.go
```

#### Construir com Docker

## Executar Docker Compose

No diretório do projeto, execute:

```sh
docker-compose up --build
```

Para construir a imagem Docker, execute:

```sh
docker build -t weather-app .
```

## Executar o Projeto

### Executar com Go

```sh
go run main.go
```

### Executar com Docker

```sh
docker run -p 8080:8080 weather-app
```

## Testar o Projeto

### Testar Localmente

Com a aplicação em execução, você pode testar utilizando `curl` ou `httpie`:

#### Teste com um CEP válido

```sh
curl -i "http://localhost:8080/weather/01001000"
```

#### Teste com um CEP inválido (formato incorreto)

```sh
curl -i "http://localhost:8080/weather/123"
```

#### Teste com um CEP não encontrado

```sh
curl -i "http://localhost:8080/weather/99999999"
```

### Executar Testes Automatizados

Para executar os testes automatizados, utilize:

```sh
go test ./weather
```

## Deploy no Google Cloud Run

1. **Autenticar no Google Cloud**:

    ```sh
    gcloud auth login
    gcloud config set project YOUR_PROJECT_ID
    ```

2. **Criar e enviar a imagem Docker para o Google Container Registry**:

    ```sh
    docker tag weather-app gcr.io/YOUR_PROJECT_ID/weather-app
    docker push gcr.io/YOUR_PROJECT_ID/weather-app
    ```

3. **Deploy no Google Cloud Run**:

    ```sh
    gcloud run deploy weather-app --image gcr.io/YOUR_PROJECT_ID/weather-app --platform managed
    ```

## Observações

- Certifique-se de substituir `YOUR_PROJECT_ID` pelo ID do seu projeto no Google Cloud.
- Certifique-se de substituir `your_weather_api_key` pela sua chave da API do WeatherAPI.
- Siga as melhores práticas de segurança ao lidar com variáveis de ambiente e chaves de API.
