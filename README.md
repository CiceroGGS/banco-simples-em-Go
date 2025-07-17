# Projeto Banco em Go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Este é um projeto de estudo desenvolvido em Go para simular as operações básicas de um sistema bancário. O objetivo foi aplicar e solidificar conceitos fundamentais e avançados da linguagem Go, como orientação a objetos, interfaces, pacotes e gerenciamento de módulos.

> 🚀 Projeto concluído e funcional, servindo como um portfólio prático dos conceitos aprendidos.

## ✨ Funcionalidades Principais

* Criação de diferentes tipos de contas: `Conta Corrente` e `Conta Poupança`.
* Operações de `Saque`, `Depósito` e `Transferência` entre contas.
* Consulta de saldo individual.
* Pagamento de boletos utilizando qualquer tipo de conta que satisfaça a interface de pagamento.

## 🧠 Conceitos de Go Aplicados

Este projeto não é apenas sobre o resultado final, mas sobre a jornada de aprendizado. Os seguintes conceitos foram explorados e implementados:

* **Estrutura de Pacotes**: O projeto é organizado em pacotes (`contas`, `clientes`) para separar as responsabilidades, melhorando a organização e a manutenibilidade do código.

* **Structs e Métodos (Modelagem Orientada a Objetos)**: Uso de `structs` (`ContaCorrente`, `ContaPoupanca`, `Titular`) para modelar entidades do mundo real e `métodos` com `receivers` de ponteiro (`*`) para associar comportamentos a essas entidades, permitindo a modificação de seus estados.

* **Encapsulamento**: O campo `saldo` das contas é privado (inicia com letra minúscula), o que força o uso de métodos públicos como `Sacar()`, `Depositar()` e `ObterSaldo()` para interagir com ele. Isso garante que as regras de negócio sejam sempre respeitadas.

* **Interfaces e Polimorfismo**: A interface `verificarConta` define um "contrato" que permite à função `PagarBoleto` trabalhar com qualquer tipo de conta. Isso demonstra o polimorfismo em Go, tornando o código mais flexível e desacoplado.

* **Gerenciamento de Módulos**: O projeto é inicializado como um módulo Go (`go mod init`), permitindo um gerenciamento de dependências e importações de pacotes locais de forma clara e moderna.

## 🚀 Como Executar

Para executar o projeto, siga os passos abaixo:

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/banco.git](https://github.com/seu-usuario/banco.git)
    ```

2.  **Navegue até a pasta raiz do projeto:**
    ```bash
    cd banco
    ```

3.  **Execute o arquivo principal:**
    ```bash
    go run main.go
    ```

## 📁 Estrutura de Arquivos

A estrutura do projeto foi organizada da seguinte forma:

```
banco/
├── go.mod
├── main.go
├── clientes/
│   └── clientes.go       # Define a struct Titular
└── contas/
    └── contas.go         # Define as structs ContaCorrente e ContaPoupanca e seus métodos
```

## 👨‍💻 Autor

Feito com ❤️ por **[Cicero]**.

[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/cicero-guilherme-a9473a260/)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/CiceroGGS/)
