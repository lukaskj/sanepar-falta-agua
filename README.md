# Sanepar - Falta de Água

English version [here](README-en.md).

### Sanepar
A Sanepar (Companhia de Saneamento do Paraná) é responsável pela prestação de serviços de saneamento básico a 345 cidades paranaenses e a Porto União, em Santa Catarina, além de 297 localidades de menor porte. ([Fonte](https://site.sanepar.com.br/a-sanepar/perfil)).

## Sobre o projeto
Atualmente a Sanepar não envia notificações em caso de possível interrupmento no fornecimendo de água. Essa informação está presente no aplicativo [Sanepar Mobile](https://play.google.com/store/apps/details?id=br.com.sanepar.saneparmobile) e é acessada passivamente entrando no app e escolhendo a opção "Falta d'água".

Por conta disso, resolvi desenvolver um programa simples que verifica essa informação de falta de água e envia uma notificação SMS para os interessados.

Resolvi me desafiar criando esse projeto em uma linguagem que não tenho familiaridade e em somente um dia.
A linguagem escolhida foi a [`Golang`](https://go.dev/) pelo interesse de aprender mais sobre ela.

## Tecnologias
- [Golang](https://go.dev/)
- [AWS SNS](https://aws.amazon.com/sns/)

### Funcionamento básico
Esse software é simplesmente um loop que bate no endpoint de falta de água da Sanepar e, dependendo do retorno, envia uma mensagem para o tópico SNS que então envia a notificação para os inscritos no tópico. <br>
O tipo de notificação no SNS fica a critério do desenvolvedor. O que estou utilizando é envio de SMS.

### Configuração
O arquivo `.env` é responsável pela configuração do sistema.<br>
Também é possível utilizar variáveis de ambiente ao invés do arquivo `.env`.

```bash
# Intervalo com o que será feito a requisição no endpoint da Sanepar. Default: 60 segundos
TIME_LOOP_SECONDS=60

# Ver abaixo
SANEPAR_BASE_URL= 
SANEPAR_CLIENT_ID=

# Nome do arquivo temporário onde serão salvos as requisições já feitas
SENT_NOTIFICATIONS_JSON_FILENAME=notifications_sent_at.json 

# Configuração das chaves AWS
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_SNS_TOPIC_ARN=
```
As configurações `SANEPAR_BASE_URL` e `SANEPAR_CLIENT_ID` consegui utilizando o proxy HTTPS [mitmproxy](https://mitmproxy.org/) ao selecionar a opção "Falta d'água" no applicativo e verificando qual URL é chamada para retornar os dados desejados.<br>
__Motivo:__ Como esse serviço não é disponível publicamente, foi necessário verificar quais requisições o aplicativo faz para buscar a informação de falta de água.


<br>
<br>
<br>

___
_For educational purposes only_
