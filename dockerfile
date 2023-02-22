# Define a imagem base
FROM golang:1.17-alpine3.14

# Define o diretório de trabalho dentro da imagem
WORKDIR /app

# Copia o código fonte para dentro da imagem
COPY . .

# Baixa as dependências do projeto
RUN go mod download

# Compila o código fonte
RUN go build -o main .

# Expõe a porta em que o servidor irá rodar
EXPOSE 8018

# Define o comando que irá rodar a aplicação
CMD [ "./main" ]