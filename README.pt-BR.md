![Getl Banner](./assets/getl_banner.png)

---

**Getl: Um Gerenciador Eficiente de Sincronização e ETL para Integração de Dados**

Getl é uma ferramenta poderosa e flexível escrita em Go que facilita a extração, transformação e carga (ETL) de dados, além de sincronia contínua entre diversas fontes e destinos. Ele suporta múltiplos bancos de dados, formatos (CSV, JSON, XML, YAML, etc.) e até integra sistemas de mensageria, permitindo a criação de fluxos de dados dinâmicos e customizáveis.

---

## **Table of Contents**
1. [Sobre o Projeto](#sobre-o-projeto)
2. [Funcionalidades](#funcionalidades)
3. [Instalação](#instalação)
4. [Uso](#uso)
    - [CLI](#cli)
    - [Exemplos de Configuração](#exemplos-de-configuração)
5. [Configuração](#configuração)
6. [Roadmap](#roadmap)
7. [Contribuindo](#contribuindo)
8. [Licença](#licença)
9. [Contato](#contato)

---

## **Sobre o Projeto**
Getl é uma solução robusta para integração de dados, projetada para atender desde operações batch até processos de sincronização contínua. Com ele, você pode:
- Conectar-se a diversas fontes (ex.: Oracle, PostgreSQL, MySQL, SQLite, MongoDB, etc.)
- Aplicar transformações detalhadas entre campos
- Sincronizar dados entre sistemas heterogêneos, inclusive com suporte a mensageria via Kafka e Redis
- Exportar dados para múltiplos formatos

Getl proporciona uma abordagem flexível e configurável para que qualquer usuário (de iniciantes a experts) possa gerenciar fluxos de dados com facilidade e eficiência.

---

## **Funcionalidades**
- **Extração de Dados**: Suporte para diversas fontes de dados.
- **Transformação Personalizada**: Permite definir mapeamentos e operações entre campos.
- **Carga e Sincronização**: Insere e atualiza dados de forma automatizada, com verificação de integridade.
- **Suporte a Mensageria**: Integração com Kafka e Redis para fluxos em tempo real.
- **Exportação de Dados**: Gera saídas em CSV, JSON, XML, YAML, PDF, entre outros formatos.
- **Configuração Dinâmica**: Configure o Getl utilizando arquivos JSON ou YAML com comentários (JSONC).
- **Operação Batch e Contínua**: Execute operações únicas ou sincronizações periódicas com intervalos configuráveis.

---

## **Instalação**
### Requisitos
- **Go** (versão 1.19 ou superior)
- Permissões para acessar as fontes e destinos de dados

```shell
# Clone o repositório
git clone https://github.com/rafa-mori/getl.git
cd getl

# Construa o binário usando o Makefile
make build

# Instale o binário (opcional)
make install

# Adicione o binário ao PATH, se necessário:
export PATH=$PATH:$(pwd)
```

---

## **Uso**

### CLI
Exemplos de comandos para utilizar o Getl:

```shell
# Sincronização simples: extrai dados de uma fonte e os carrega para um destino
getl sync -f examples/configFiles/exp_config_a.json

# Extração de dados específicos via SQL
getl extract --source "oracle_db" --query "SELECT * FROM produtos"

# Transformação e carga de dados com configurações customizadas
getl transform -f examples/configFiles/exp_config_b.json
```

### Exemplos de Configuração

#### **Exemplo 1: Sincronização Básica**
```json
{
  "sourceType": "godror",
  "sourceConnectionString": "username/password@127.0.0.1:1521/orcl",
  "destinationType": "sqlite3",
  "destinationConnectionString": "/home/user/.kubex/web/gorm.db",
  "destinationTable": "erp_products",
  "destinationTablePrimaryKey": "CODPARC",
  "sqlQuery": "SELECT P.CODPARC, P.NOMEPARC FROM TABLE P",
  "outputFormat": "csv",
  "outputPath": "/home/user/Documents/erp_products.csv",
  "needCheck": true,
  "checkMethod": "SELECT * FROM erp_products WHERE CODPARC = ? AND NOMEPARC = ?",
  "kafkaURL": "",
  "kafkaTopic": "",
  "kafkaGroupID": ""
}
```

#### **Exemplo 2: Sincronização Contínua**
```json
{
  "sourceType": "godror",
  "sourceConnectionString": "username/password@127.0.0.1:1521/orcl",
  "destinationType": "sqlite3",
  "destinationConnectionString": "/home/user/.kubex/web/gorm.db",
  "destinationTable": "erp_products",
  "destinationTablePrimaryKey": "CODPARC",
  "sqlQuery": "SELECT P.CODPARC, P.NOMEPARC FROM TABLE P",
  "syncInterval": "30 * * * * *",
  "kafkaURL": "",
  "kafkaTopic": "",
  "kafkaGroupID": ""
}
```

#### **Exemplo 3: Transformações Detalhadas**
```json
{
  "sourceType": "sqlite3",
  "sourceConnectionString": "/home/user/.kubex/web/gorm.db",
  "sourceTable": "erp_products",
  "sourceTablePrimaryKey": "CODPROD",
  "sqlQuery": "",
  "destinationType": "sqlServer",
  "destinationConnectionString": "sqlserver://username:password@localhost:1433?database=my_db_test&encrypt=disable&trustservercertificate=true",
  "destinationTable": "erp_products_test",
  "destinationTablePrimaryKey": "id_v",
  "kafkaURL": "",
  "kafkaTopic": "",
  "kafkaGroupID": "",
  "transformations": [
    {
      "sourceField": "CODPROD",
      "destinationField": "id_v",
      "operation": "none",
      "sPath": "erp_products",
      "dPath": "erp_products_test"
    },
    {
      "sourceField": "PRODDESCR",
      "destinationField": "name_v",
      "operation": "none",
      "sPath": "erp_products",
      "dPath": "erp_products_test"
    }
    // Outras transformações...
  ]
}
```

---

## **Configuração**
O Getl utiliza arquivos de configuração em JSON ou YAML para definir parâmetros de conexão, transformações e sincronização. Esses arquivos podem ser personalizados conforme as necessidades do usuário. Para mais detalhes, consulte a [documentação de configuração](https://github.com/rafa-mori/getl/README.md#configuration-file).

---

## **Roadmap**
🔜 **Recursos Futuros Planejados**:
- Novos conectores para fontes/destinos adicionais (ex.: MongoDB, Redis).
- Mais operações de transformação customizadas.
- Integração expandida com sistemas de mensageria.
- Dashboard e métricas em tempo real para visualizar a sincronização e status dos jobs.

---

## **Contribuindo**
Contribuições são sempre bem-vindas!  
Sinta-se à vontade para abrir issues ou enviar pull requests.  
Confira o [Guia de Contribuição](CONTRIBUTING.md) para mais detalhes.

---

## **Licença**
Este projeto é licenciado sob a [MIT License](LICENSE).

---

## **Contato**
💌 **Desenvolvedor**: [Rafael Mori](mailto:faelmori@gmail.com)  
💼 [GitHub - Rafael Mori](https://github.com/rafa-mori)

Se você achar o projeto interessante ou quiser colaborar, entre em contato!
💼 Estou aberto a novas oportunidades de trabalho e colaborações.

---