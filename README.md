# mutual-tls

- Generating CA file <br>
`openssl req -newkey rsa:2048 -nodes -x509 -days 365 -out ca.crt -keyout
  ca.key -subj "/C=CN/ST=GD/L=SZ/O=AppsCode, Inc./CN=example.test Root CA" `
- Generating Server Certificate <br>
`openssl genrsa -out server.key 2048`
- `openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=GD/L=S
Z/O=AppsCode, Inc./CN=server.test"`
- `echo "subjectAltName=DNS:server.test,DNS:server.test" > altsubjserver.ext`
- `openssl x509 -req -extfile altsubjserver.ext -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt`
- Generating The Client Certificate
- same way.
