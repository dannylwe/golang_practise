## How to generate cert and key .pem files  
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out cert.pem -keyout key.pem  
After which, follow the prompts  

## References  
-https://ericchiang.github.io/post/go-tls/
-https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go
-https://posener.github.io/http2/