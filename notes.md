# Notes

## Modules

### jeiko
- Consultas de stock del mercado
get			/checkStock	JeikoController.check() -> Trae un json con los resultados del mercado

### jogua
- Compras del mercado
get			/checkBuy	JoguaController.checkBuy()->Consulta de compras
post			/buy		JoguaController.buy(fad)-> Compas

### nhemu
- Ventas del mercado
get			/checkSell	NhemuController.checkBuy()->
post			/sell

### Observations

- Why 3 modules?
Because the 3 modules have diferent configuration of scalability requirements 
