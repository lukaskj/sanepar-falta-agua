# Sanepar - Falta de Água

pt-BR version [here](README.md).

### Sanepar
Sanepar (Companhia de Saneamento do Paraná) is responsible for providing basic sanitation services to 345 cities in Paraná and Porto União, in Santa Catarina, in addition to 297 smaller locations. ([Source](https://site.sanepar.com.br/a-sanepar/perfil)).

## About the project
Currently, Sanepar does not send notifications in case of a possible interruption in the supply of water. This information is present in the [Sanepar Mobile](https://play.google.com/store/apps/details?id=br.com.sanepar.saneparmobile) mobile app and is accessed passively by entering the app and choosing the "Lack of water" (Falta d'água) option.

Because of this, I decided to develop a simple program that checks this water shortage information and sends an SMS notification to those interested.

I decided to challenge myself by creating this project in a language I'm not familiar with and in just one day.<br>
The language chosen was [`Golang`](https://go.dev/) due to the interest of learning more about it.

## Technologies
- [Golang](https://go.dev/)
- [AWS SNS](https://aws.amazon.com/sns/)

### Basic functioning
This software is simply a loop that keeps making a request to Sanepar's water shortage endpoint and, depending on the response, sends a message to the SNS topic which then sends the notification to subscribers to the topic. <br>
The type of notification on SNS is up to the developer. What I am using is SMS sending.

### Configuration
The `.env` file is responsible for the configuration.<br>
It is also possible to use environment variables instead of the `.env` file.

```bash
# Interval with which the request will be made on the Sanepar endpoint. Default: 60 seconds
TIME_LOOP_SECONDS=60

# See below
SANEPAR_BASE_URL= 
SANEPAR_CLIENT_ID=

# Name of the temporary file where requests already made will be saved
SENT_NOTIFICATIONS_JSON_FILENAME=notifications_sent_at.json 

# AWS access keys and SNS topic ARN
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_SNS_TOPIC_ARN=
```
The `SANEPAR_BASE_URL` and `SANEPAR_CLIENT_ID` settings were achieved using the HTTPS proxy [mitmproxy](https://mitmproxy.org/) by selecting the "Lack of water" option in the application and checking which URL is called to return the desired data.<br>
__Reason:__ As this service is not publicly available, it was necessary to check which requests the application makes to fetch water shortage information.


<br>
<br>
<br>

___
_For educational purposes only_
