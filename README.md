# 🪖 Challenge

Crear una interfaz para buscar información de bases de datos correos.

1. Indexar base de datos de correo electrónico
Primero descargar la base de datos de correos de [Enron Corp (423MB)](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz).
Después escribe un programa que indexe sus contenidos en la herramienta ZincSearch. Por ejemplo:
``` bash
$ ./indexer enron_mail_20110402
```

2. Proling
Hazle [proling](https://go.dev/doc/diagnostics#proling) a tu indexer. Genera el gráco para analizarlo durante la sustentación.

3. Visualizador
Crea una interfaz simple para buscar los contenidos. Por ejemplo:
``` bash
$ ./app --port 3000

App is running in http://localhost:3000
```

4. Optimización (Opcional)
Usa el proling de la parte 2 para optimizar tu código. Documentar que mejoras de optimización encontrasdas.

5. Despliegue (Opcional)
Desplegar todo a AWS o Local Stack usando terraform.

Tecnologías:
- Lenguaje Backend: Go
- Base de Datos: ZincSearch
- API Router: chi (pronunciado kai)
- Interfaz: Vue 3
- CSS: Tailwind
- OS: Linux

# 🚀 Start
``` bash
./indexer enron_mail_20110402
./app --port 3000
```
