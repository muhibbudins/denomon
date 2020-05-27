# Denomon

## Also read in [English](https://github.com/muhibbudins/denomon/blob/master/README.md) 

Um observador de arquivos para construir aplicações usando Deno

[![asciicast](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy.png)](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy)

> Nota : O indicador é exibido somente no asciinema, mas não no bash / zsh

### Instalação

Instale usando Wget

```bash
$ wget -O - https://raw.githubusercontent.com/muhibbudins/denomon/master/install.sh | sh
```

ou, CURL
```bash
$ curl -s https://raw.githubusercontent.com/muhibbudins/denomon/master/install.sh | sh
```

### Utilização

```bash
$ denomon <opções> <arquivo>
```

Exemplo:

- Exibir mensagem de ajuda

```bash
denomon --help
```

- Único comando para inicializar a pasta atual recursivamente

```bash
$ denomon
```

- Inicializar arquivos com permissões [vide nota abaixo]

```bash
$ denomon --allow net,read server.ts
```

- Inicializar arquivo em pasta específica e permissões [vide nota abaixo]

```bash
$ denomon --dir fixtures --allow net,read server.ts
```

- Inicialiar pasta específica com permissões

```bash
$ denomon --dir fixtures --allow net,read
```

> Nota : Se você executar o denomon para inicializar um único arquivo, todas as mudanças na pasta raíz vão gatilhar o recarregamento no arquivo principal.


### Opções

#### --version

Exibe a versão do denomon

#### --help

Exibe a mensagem de ajuda

#### --dir

Atribui um diretório para observar

#### --allow

Atribui permissões para seu projeto

### Funcionalidades

- Automaticamente compila para um único arquivo
- Observa todos os arquivos na pasta recursivamente
- Compilação automática ao recarregar para processos filhos (ie. net)

### Licença

Este projeto está sob a licença MIT.

### Stargazers ao longo do tempo

[![Stargazers ao longo do tempo](https://starchart.cc/muhibbudins/denomon.svg)](https://starchart.cc/muhibbudins/denomon)
      