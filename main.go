package main

import (
	"fmt"
	"sync"
	"os"
	"encoding/json"
	"io/ioutil"
)

var wg sync.WaitGroup

type Caso struct{
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State        string `json:"state"`
}

type CasosContenedor struct {
	Casos []Caso `json:"Casos"`
}

func (t Caso) toString() string {
    bytes, err := json.Marshal(t)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return string(bytes)
}

func getCasos(path string) CasosContenedor {
    var casoContenedor CasosContenedor
    raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
	}
	json.Unmarshal(raw, &casoContenedor)
    return casoContenedor
}

func main(){
	var url string
	var gorutinas int
	var solicitudes int
	var path string

	fmt.Println("Ingrese la url:")
	fmt.Scanf("%s",&url)
	fmt.Println("Cantidad de gorutinas a utilizar:")
	fmt.Scanf("%d",&gorutinas)
	fmt.Println("Cantidad de solicitudes:")
	fmt.Scanf("%d",&solicitudes)
	fmt.Println("Ruta del archivo:")
	fmt.Scanf("%s",&path)

	if gorutinas > solicitudes{
		fmt.Println("Las gorutinas no pueden ser mayor a las solicitudes")
		return
	}

	casos := getCasos(path)

	if(solicitudes > len(casos.Casos)){
		fmt.Println("Las solicitudes son mayores a las contenidas en el archivo")
		return
	}

	indice:=0
	rango := solicitudes / gorutinas
	faltante := (solicitudes%gorutinas)
	
	wg.Add(gorutinas)

	for gorutinas > 0 {
		if gorutinas == 1{
			go enviarCasos(casos.Casos,indice,indice+rango+faltante)	
		}else{
			go enviarCasos(casos.Casos,indice,indice+rango)	
		}
		indice += rango
		gorutinas--
	}
	wg.Wait()
	fmt.Println("===================Terminando programa===================")
}

func enviarCasos(casos []Caso, indiceInicial int, indiceFinal int){
	defer wg.Done()
	for indiceInicial < indiceFinal{
		fmt.Println(casos[indiceInicial].toString(),indiceInicial)		
		indiceInicial +=1
	}
}