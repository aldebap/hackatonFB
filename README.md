```
████████╗██╗███╗   ███╗███████╗        ███████╗██████╗  ██████╗██████╗ ██╗ ██████╗
╚══██╔══╝██║████╗ ████║██╔════╝        ██╔════╝██╔══██╗██╔════╝██╔══██╗██║██╔════╝
   ██║   ██║██╔████╔██║█████╗          █████╗  ██████╔╝██║     ██║  ██║██║██║     
   ██║   ██║██║╚██╔╝██║██╔══╝          ██╔══╝  ██╔══██╗██║     ██║  ██║██║██║     
   ██║   ██║██║ ╚═╝ ██║███████╗        ███████╗██████╔╝╚██████╗██████╔╝██║╚██████╗
   ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝        ╚══════╝╚═════╝  ╚═════╝╚═════╝ ╚═╝ ╚═════╝
                                                                                  
           
```
# A FUGA DO BATCH #


O time é composto por [Aldebaran](https://github.com/aldebap), [Rafael Camilo](https://github.com/rcamilo), [Eric Rodrigo Ramos](https://github.com/eric-r-ramos), Paulo Eduardo Domingues e [Marcelo Sousa Lima](https://github.com/marceluxvk).

A aplicação foi implementada em go com partes de persistencia utilizando GOLang e Java.

## Arquitetura ##

A solução foi implementada baseando-se no conceito de eventos. Como message broker foi utilizado o [Apache Kafka](http://kafka.apache.org/).

Vale citar que a aplicação em seus módulos de enriquecimento e persistencia, está preparada para escala horizontal utilizando o consumer groups do Kafka.

```
                                                                                __________
 ______                      ________                        [Persister A] --> |__________| <-- [External App]
|      |        topic       |        |        topic        /      ...          |          |
|Loader| --> [][][][][] --> |Enricher| --> [][][][][] -->  --[Persister B] --> | DATABASE | <-- [External App]
|______|                    |________|                     \      ....         |__________|
                                                             [Persister C] --> |__________| <-- [External App]
```

O time entende que para um melhor desempenho das aplicações Fim(External App no desenho), as mesmas deveriam consumir as mensagens diretamente do kafka, executando assim o efetivo "fim do batch" e eliminando uma camada desnecessária de persistencia, que seria mantida somente em casos específicos.



### Submodulos ###

#### Loader ####

[fonte](https://github.com/aldebap/hackatonFB/tree/master/requestLoader)

Componente implentado em GoLang responsável por varrer o diretorio onde o arquivo é disponibilizado e transforma-lo em um stream de mensagens json(canonino do stream). Em seguida, aplicação envia o canonico para um topico kafka.

#### Enricher ####

[fonte](https://github.com/aldebap/hackatonFB/tree/master/project)

Componente responsável por simular as consultas a serviçoes externos e, adicionar dados da consulta no objeto canonico, que é enviado para um novo tópico, de onde será persistido pelas aplicações de persistencia.

#### Applicacoes de Persistencia ####

Na ausência de um banco de dados oracle, foi utilizado uma unica instancia de banco de dados MySQL que preserva as mesmas caracteristicas de escala vertical de um oracle, porém, executado dentro do mesmo host da aplicação.

##### Caracteristicas da persistencia #####

Para persistencia foi definido que o melhor modelo seria o de inserções em batch, para evitar o excessivo tráfego de rede e multiplas transações em paralelo o que degradaria naturalmente o desempenho das operações no banco de dados.

[fonte 1](https://github.com/aldebap/hackatonFB/tree/master/ajuste-persistence/) [fonte 2](https://github.com/aldebap/hackatonFB/tree/master/ebcdic-spring-boot)


### Arquitetura Escalável ###

O resultado obtido foi alcançado utilizando um ```node``` para cada instância de módulo. No entanto, a arquitetura da aplicação foi elaborada pensando na escalabilidade horizontal de cada um dos módulos. Desta forma, ao identificar os gargalos do processo de carga, podemos subir novas instâncias sob demanda para processar a parte que está necessitando melhoria no processamento. 

```
                                 _______
                               _|_____  |                    [Persister A1]      __________
 ______                      _|______ | |                    [Persister A2] --> |__________| <-- [External App]
|      |        topic       |        ||_|      topic       /      ...           |          |
|Loader| --> [][][][][] --> |Enricher||--> [][][][][] -->  --[Persister B1] --> | DATABASE | <-- [External App]
|______|                    |________|                     \      ....          |__________|
                                                             [Persister C1] --> |__________| <-- [External App]
                                                             [Persister C2]
```

No exemplo acima o modulo ```Enricher``` foi escalado com mais 2 instancias para dar vazão ao consumo das mensagens produzidas pelo ```Loader```. No caso dos módulos de persistência, o ```Persister A``` e o ```Persister C``` demanadou a criação de novas instâncias, enquanto o ```Persister B``` não teve esta nessdidade.


## O Teste ##

Para o teste foi utilizado Notebook com a seguinte caracteristica:

* CPU: Intel Core I3-3110M 2.4GHz 
* Memoria RAM: 6GB de Memória RAM 
* Systema Operacional: Windows 10

O ambiente executava:
1. Zookeeper
2. Kafka
3. MySQL Server
4. Loader
5. Enricher
6. Persistence 1

Foi observado que o ambiente executava com 100% de uso de cpu desde o inicio do processamento até o final do teste, o que sinaliza um gargalo de hardware e não de aplicação. Isso nos permite imaginar que o throughput final pode ser escalado de forma interessante em um ambiente apropriado.

---

### Resultados ###

O teste de carga do arquivo provido pelo hackaton com o seguinte resultado:

* Tempo Total de Teste: 3:53 min
* Throughput medio: 4291 tps
* Latência média por operacao: 2,33 miliseconds
* Utilização de CPU: 100% durante todo o teste
* Utilização de Memória: 20 mb por módulo de aplicação

### Evidencias ####

#### CPU ####

<img src="https://github.com/aldebap/hackatonFB/blob/master/CPUEvidence.png" alt="CPU" style="width: 80px; height: 50px;" />


#### Tempo de Teste ####

<img src="https://github.com/aldebap/hackatonFB/blob/master/cronometro.jpg" alt="Tempo de Teste" style="width: 80px; height: 50px;" />


#### DB Load ####

![DB Load](https://github.com/aldebap/hackatonFB/blob/master/dbevidence.PNG)

#### Loader Start ####

![Loader Start](https://github.com/aldebap/hackatonFB/blob/master/loader-start-evidence.PNG)

#### Loader End ####

![Loader End](https://github.com/aldebap/hackatonFB/blob/master/loader-end-evidence.PNG)

#### Persistence End ####

![Persistence End](https://github.com/aldebap/hackatonFB/blob/master/db-end-evidence.PNG)


