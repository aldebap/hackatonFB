# TIME EBCDIC #

O time é composto por Aldebaran, Rafael, Eric e Marcelo.

A aplicação foi implementada em go com partes de persistencia utilizando GOLang e Java.

## Arquitetura ##

A solução foi implementada baseando-se no conceito de eventos. Como message broker foi utilizado o [Apache Kafka](http://kafka.apache.org/).

```
                                                                                     __________
 ______                        ________                          [Persister A] ---> |__________|  <---- [External App]
|      |        topic         |        |        topic          /      ...           |          |
|loader| ---> [][][][][] ---> |enricher| ---> [][][][][] --->  --[Persister B] ---> | DATABASE |  <---- [External App]
|______|                      |________|                       \      ....          |__________|
                                                                 [Persister C] ---> |__________|  <---- [External App]
```

### Submodulos ###

#### Loader ####

[fonte](https://github.com/aldebap/hackatonFB/tree/master/requestLoader)

Componente implentado em GoLang responsável por varrer o diretorio onde o arquivo é disponibilizado e transforma-lo em um stream de mensagens json(canonino do stream). Em seguida, aplicação envia o canonico para um topico kafka.
