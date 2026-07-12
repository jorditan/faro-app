# Finance AI

> Aplicación Full Stack para la gestión de finanzas personales con análisis asistido por IA.

## Objetivo

Desarrollar una aplicación de finanzas personales con el objetivo de practicar el desarrollo de un producto Full Stack moderno utilizando Go y React.

El proyecto busca equilibrar simplicidad, buenas prácticas de arquitectura y espacio para incorporar funcionalidades de IA en futuras iteraciones.

---

# Objetivos del proyecto

## Técnicos

* Desarrollar un backend en Go.
* Desarrollar un frontend en React + TypeScript.
* Diseñar una API REST.
* Utilizar Docker para el entorno de desarrollo.
* Implementar autenticación mediante JWT.
* Practicar arquitectura por capas y por dominio.
* Aplicar principios SOLID cuando aporten valor.
* Mantener una estructura fácilmente escalable.

## Funcionales

La aplicación permitirá:

* Registrar ingresos.
* Registrar egresos.
* Administrar categorías.
* Visualizar indicadores financieros.
* Filtrar movimientos.
* Obtener estadísticas.
* Incorporar un asistente conversacional en una etapa posterior.

---

# Filosofía del proyecto

Este proyecto **no busca ser un ejercicio académico**.

La idea es desarrollar un producto que pueda utilizarse diariamente y que evolucione mediante iteraciones.

Cada nueva funcionalidad deberá aportar valor al usuario antes que complejidad técnica.

---

# Roadmap

## MVP

### Usuarios

* Registro
* Inicio de sesión
* JWT

### Transacciones

* Crear movimiento
* Editar movimiento
* Eliminar movimiento
* Listar movimientos

Campos:

* título
* descripción
* monto
* fecha
* tipo (Ingreso / Egreso)
* categoría

### Categorías

CRUD de categorías.

Separadas por tipo:

* Ingresos
* Egresos

### Dashboard

Mostrar:

* Balance actual
* Ingresos del mes
* Gastos del mes
* Ahorro mensual
* Últimos movimientos

### Filtros

* Mes
* Categoría
* Tipo

---

## V2

* Gráfico por categorías
* Evolución mensual
* Gastos por período

---

## V3

Presupuestos

Ejemplo:

* Supermercado
* Transporte
* Ocio

Mostrar porcentaje utilizado.

---

## V4

Objetivos de ahorro.

Ejemplo:

* Notebook
* Viaje
* Fondo de emergencia

---

## V5

Asistente financiero con IA.

---

# Arquitectura

Se utilizará un **monolito desacoplado**.

No se implementarán microservicios.

El frontend y backend serán proyectos independientes.

```text
finance-ai/

├── frontend/
├── backend/
├── docker-compose.yml
└── README.md
```

---

# Stack tecnológico

## Frontend

* React
* TypeScript
* Vite
* React Router
* TanStack Query
* React Hook Form
* Zod
* Tailwind CSS
* shadcn/ui
* Recharts

---

## Backend

* Go
* Gin
* PostgreSQL
* JWT

---

## Infraestructura

* Docker
* Docker Compose

---

# ¿Por qué PostgreSQL?

Aunque MongoDB es una excelente base de datos documental, el dominio de una aplicación financiera posee relaciones bien definidas.

Ejemplos:

* Usuario → Transacciones
* Usuario → Categorías
* Usuario → Presupuestos
* Usuario → Objetivos

Además, la aplicación requerirá consultas frecuentes utilizando:

* SUM
* GROUP BY
* ORDER BY
* BETWEEN
* JOIN

Estas operaciones resultan más naturales y eficientes en una base de datos relacional.

PostgreSQL también permite almacenar datos semiestructurados mediante JSONB, lo cual será útil para futuras integraciones con IA sin necesidad de incorporar otra base de datos.

---

# Arquitectura Backend

Se utilizará una arquitectura por dominio.

```
backend/

cmd/
    api/
        main.go

internal/

    auth/
    user/
    transaction/
    category/

    database/
    middleware/
    config/
    routes/

go.mod
```

Cada dominio encapsula toda su lógica.

Ejemplo:

```
transaction/

    handler.go
    service.go
    repository.go
    model.go
    dto.go
```

---

# Capas

## Handler

Responsabilidades:

* Recibir requests HTTP.
* Validar datos de entrada.
* Invocar servicios.
* Devolver respuestas HTTP.

No contiene lógica de negocio.

---

## Service

Responsabilidades:

* Reglas de negocio.
* Validaciones.
* Coordinación entre repositorios.
* Procesamiento de información.

Es el corazón de la aplicación.

---

## Repository

Responsabilidades:

* Consultar PostgreSQL.
* Ejecutar operaciones CRUD.
* No conoce HTTP.

---

## Model

Representa las entidades persistidas.

No contiene lógica de negocio.

---

## DTO

Representa los objetos enviados y recibidos por la API.

Permite desacoplar la base de datos del contrato HTTP.

---

# Flujo de una petición

```text
Cliente

↓

HTTP

↓

Router

↓

Middleware

↓

Handler

↓

Service

↓

Repository

↓

PostgreSQL
```

---

# IA (Post MVP)

La IA no reemplazará la lógica de negocio.

Todos los cálculos financieros seguirán implementados mediante código.

La IA actuará como una interfaz conversacional capaz de interpretar consultas en lenguaje natural y utilizar herramientas para acceder a la información.

Ejemplos:

* ¿En qué gasté más dinero este mes?
* Compará abril con mayo.
* Registrá un gasto de $12.000 en supermercado.
* ¿Cuánto llevo ahorrado para mi notebook?

---

# Arquitectura de IA

```text
Usuario

↓

Chat

↓

Agente

↓

Tools

↓

Servicios del Backend

↓

PostgreSQL
```

El agente no accederá directamente a la base de datos.

Utilizará herramientas bien definidas como:

* getTransactions()
* createTransaction()
* getMonthlyBalance()
* compareMonths()
* getCategoryTotals()

Esto mantiene la lógica de negocio centralizada y evita duplicación.

---

# Principios de desarrollo

* Mantener el código simple.
* Evitar sobreingeniería.
* Priorizar claridad sobre abstracción.
* Diseñar primero el dominio.
* Agregar nuevas capas solo cuando exista una necesidad real.
* Construir funcionalidades pequeñas e incrementales.

---

# Próximos pasos

## Fase 1

* Definir el dominio.
* Diseñar el modelo de datos.
* Crear el esquema relacional.
* Definir entidades.

## Fase 2

* Crear la API REST.
* Implementar autenticación.
* CRUD de categorías.
* CRUD de transacciones.

## Fase 3

* Dashboard.
* Indicadores.
* Gráficos.

## Fase 4

* Presupuestos.
* Objetivos.

## Fase 5

* Asistente financiero basado en IA.
* Herramientas (Tools).
* Contexto de conversaciones.
* Acciones mediante lenguaje natural.
