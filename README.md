[<img src="https://www.freecodecamp.org/news/content/images/2021/10/golang.png" width="150"/>](Golang.org)

# Imersão Full Stack & FullCycle 5.0 - Fincycle - Backend de processamento das transações feito com Golang
## Descrição

Repositório do back-end de processamento das transações feito com Golang

**Importante**: A aplicação do Apache Kafka deve estar rodando primeiro.

## Rodar a aplicação

### Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

Execute os comandos:

```
docker-compose up
```

Use o Apache Kafka para produzir e consumir mensagens para testar a aplicação.

Quer configurar um ambiente de desenvolvimento produtivo? Veja o vídeo: [https://www.youtube.com/watch?v=nTlM4sd-W3E](https://www.youtube.com/watch?v=nTlM4sd-W3E) 

### Para Windows 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart) 

## Mockgen Golang

* https://github.com/golang/mock
* go install github.com/golang/mock/mockgen@v1.6.0

><p>mockgen -destination=domain/repository/mock/mock.go -source=domain/repository/repository.go</p>

## Test - GO, Kafka and Sqlite3

![test_image](/test_go_kafka_sqlite_ms.png)