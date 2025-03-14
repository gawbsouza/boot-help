<p align="center" width="100%">
    <img width="15%" src="boot-help.png"> 
</p>

# Boot-Help

Tornando a vida do bootcamper mais simples :)

---

<br>

⚠️ Boot-Help está em construção, a maneira de utilizá-lo pode sofrer alterações.

<br>

Boot-Help tem como objetivo facilitar o trabalho Web com Go quando utilizamos a biblioteca [Chi](https://github.com/go-chi/chi), provê uma maneira mais simples de trabalhar com erros e respostas HTTP.

## Comece instalando:

```sh
go get github.com/gawbsouza/boot-help
```

## Trabalhando com erros

Para facilitar o trabalho com erros o pacote `httperr` possui diversas funcionalidades.

### Formato padronizado

Por padrão o Boot-Help formata os seus erros na seguinte estrutura:

```go
type HttpError struct {
	StatusCode int    // Status code na resposta HTTP
	Message    string // Mensagem principal do erro
	Details    string // Adicionar mais detalhes ao erro
}
```

### Criando erros

O pacote `httperr` possui construtores para erros de vários tipos, assim é possível criar erros de maneira padronizada e com pouco esforço.

```go
// Retorna um HttpError com StatusCode 400 e uma Message customizada
httperr.Bad("dados inválidos")

// Retorna um HttpError com StatusCode 404 e uma Message customizada
httperr.NotFound("usuário não existe")

// Retorna um HttpError com StatusCode 409 e uma Message customizada
httperr.Conflict("email já cadastrado")

// Retorna um HttpError com StatusCode 412 e uma Message customizada
httperr.Condition("usuario precisa estar habilitado")

// Retorna um HttpError com StatusCode 500 e uma Message customizada
httperr.Internal("erro interno")
```

### Adicionando mais detalhes

Com o erro criado é possível adicionar mais detalhes ao erro caso seja necessário.

```go
// Retorna um HttpError com StatusCode, Message e Details peenchidos
httperr.Bad("nome inválido").WithDetails("precisa ter ao menos 3 caracteres")

// HttpError {
//    StatusCode: 400,
//    Message: "nome inválido"
//    Details: "precisa ter ao menos 3 caractes"
// }

// Pode ser utilizado com qualquer tipo de erro
httperr.NotFound("").WithDetails("")
httperr.Internal("").WithDetails("")
httperr.Conflict("").WithDetails("")
```

## Resposta HTTP

Para facilitar o retorno de resposta HTTP o pacote `response` possui diversas funcionalidades.

### Criando novas respostas HTTP

```go
func main() {
	rt := chi.NewRouter()

	rt.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.To(w).Content("Hello from Boot-Help!").SendText()
	})

	log.Println("Server running at port 80")
	log.Fatal(http.ListenAndServe(":80", rt))
}
```

Resposta

```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Thu, 13 Mar 2025 10:27:40 GMT
Content-Length: 21
Connection: close

Hello from Boot-Help!
```

Aqui temos algumas dicas de como utilizar os métodos disponíveis:

```go
// .To() para criar um Response e configurar o ResponseWriter onde o a resposta HTTP será escrita, tudo começa a partir daqui!
response.To(w)

// .Content() para configurar qual será o conteúdo (payload) utilizado na resposta, pode ser de qualquer tipo.
response.To(w).Content("my content")

// .Status() para configurar qual será o status code da resposta, o padrão é 200.
response.To(w).Status(404)

// .Type() para configurar qual será o Content-Type, deve ser utilizando juntamente com o .Send() para conservar o tipo informado, o padrão é "text/plain"
response.To(w).Type("text/csv")

// .Header() para adicionar uma nova entrada no header de resposta HTTP
response.To(w).Header("key", "value")

// .Headers() para adicionar várias entradas no header da resposta HTTP, deve ser informado em formato map[string]string
response.To(w).Headers(map[string]string{"key":"value"})
```

### Enviando respostas HTTP

Após criado a Response como coms métodos anteriores é hora de enviar a resposta, utilize os seguintes métodos para realizar e finalizar o envio.

```go
// .Send() para enviar a resposta HTTP, considera do Content-Type configurado por .Type(), por padrão é "text/plain"
response.To(w).Send()

// .SendText() para enviar a resposta HTTP com o conteúdo formatado para texto, Content-Type "text/plain"
response.To(w).SendText()

// .SendJSON() para enviar a resposta HTTP com o conteúdo formatado para JSON
response.To(w).SendJSON()

// É possível combinar os métodos para um resposta mais personalizada
response.To(w).Status(201).Content("criado com sucesso!").SendText()
response.To(w).Type("text/csv").Content(myReport).Send()
response.To(w).Content(usersList).SendJSON()
```

### Respostas HTTP com erros

O pacote `response` possui integração com os erros de `httperr` e facilita a criação de respostas com erros.

```go
// Criando um erro e adicionando na resposta
notFoundErr := httperr.NotFound("usuário não encontrado")
response.To(w).Err(notFoundErr).SendJSON()
```

Resposta

```
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Thu, 13 Mar 2025 09:49:34 GMT
Content-Length: 57
Connection: close

{
  "status_code": 404,
  "message": "usuário não encontrado"
}
```

#### Crie uma response com o erro já embutido

O pacote `response` possui métodos que já criam e adicionam o erro na resposta, facilitando ainda mais.

```go
// Equivalente a
// myErr := httperr.Bad("nome inválido")
// response.To(w).Err(myErr)
response.To(w).BadErr("mensagem")

// Equivalente a criar o erro com httperr.NotFound()
response.To(w).NotFoundErr("mensagem")

// Equivalente a criar o erro com httperr.Conflict()
response.To(w).ConflictErr("mensagem")

// Equivalente a criar o erro com httperr.Condition()
response.To(w).ConditionErr("mensagem")

// Equivalente a criar o erro com httperr.Internal()
response.To(w).InternalErr("mensagem")
```

## Requisições HTTP

Para facilitar a recepção da requisição HTTP o pacote `request` possui diversas funcionalidades.

### Criando novas requisições HTTP

```go
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	rt := chi.NewRouter()

	rt.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var input Product

		if err := request.From(r).ParseJSON(&input); err != nil {
			response.To(w).BadErr(err.Error()).Send()
			return
		}

		response.To(w).Content(input).SendJSON()
	})

	if err := http.ListenAndServe(":3030", rt); err != nil {
		panic(err)
	}
}
```

Resposta

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 13 Mar 2025 19:02:27 GMT
Content-Length: 25
Connection: close

{
  "name": "Gopher",
  "price": 12.2
}
```

Aqui temos algumas dicas de como utilizar os métodos disponíveis:

```go
// .From() para criar um Request e configurar o http.Request onde a requisição HTTP será recebida, tudo começa a partir daqui!
request.From(r)

// .ParseJSON() para converter o valor recebido no corpo da requisição em uma struct
request.From(w).ParseJSON(&object)

// .ParseValidJSON() para converter o valor recebido no corpo da requisição em uma struct e verificar se os valores das propriedades condizem com as validações requisitadas
request.From(w).ParseValidJSON(&object)
```

### Requisição HTTP com erros

O pacote `request` possui retorno genérico para os erros.

```go
// Implementando request com ParseJSON
// err será uma saida do tipo error
err := request.From(r).ParseJSON(&input)

// Implementando request com ParseValidJSON
// val será uma saida do tipo []string com as mensagens de validação
// err será um saida do tipo error
val, err := request.From(r).ParseValidJSON(&input)
//Your logic here...

```
