
# Sunset Sur - ERP & CRM Inmobiliario

Sistema integral de Planificación de Recursos Empresariales (ERP) y Gestión de Relaciones con Clientes (CRM) diseñado específicamente para la operación y administración de desarrollos inmobiliarios.

El sistema centraliza la gestión de prospectos, control de obras, finanzas, cálculo de comisiones multinivel y automatización de documentos en una plataforma de alta densidad visual, rápida y escalable.

## Arquitectura del Proyecto

El proyecto utiliza un enfoque de Monorepo, combinando un monolito robusto para la lógica de negocio y seguridad, respaldado por microservicios de alto rendimiento para tareas intensivas. Toda la infraestructura está orquestada ocn Docker.

### Stack Tecnológico

* **Backend (API Gateway & Core):** Laravel 11 (PHP 8.4)
* **Frontend:** Vue.js 3 (Composition API) + Inertia.js
* **Base de Datos:** PostgreSQL 15
* **Estilos y UI:** Tailwind CSS v4 (Sistema de diseño mediante Tokens personalizados)
* **Microservicios:** Go (Golang)
* **Infraestructura:** Docker & Docker Compose (Imágenes Alpine y compilación Multi-stage)

## Módulos Principales

1.  **Panel de Control (Dashboard):** Métricas de ventas, alertas de obras, citas próximas y comisiones.
2.  **Usuarios y Roles:** Control de acceso granular para directivos, vendedores (individuales o líderes de equipo) y personal de obra.
3.  **CRM (Ventas):** Embudo de conversión visual, directorio de contactos, expedientes digitales y agenda de citas.
4.  **Propiedades e Inventario:** Catálogo de desarrollos, lotes y control de disponibilidad en tiempo real.
5.  **Obras y Proyectos:** Fichas técnicas, presupuesto, registro de gastos de material y bitácora fotográfica de avances.
6.  **Finanzas y Comisiones:** Registro de ingresos, enganches y motor de cálculo de comisiones multinivel.
7.  **Documentos:** Generador automatizado de contratos y cotizaciones en PDF.

## Sistema de Diseño (UI/UX)

La interfaz de usuario abandona las plantillas genéricas para implementar la identidad de marca premium de "Sunset Condominios". 

* **Esquema:** Modo Claro (Light Mode) obligatorio. Fondos cálidos (Arena) y textos oscuros (Café).
* **Tokens:** Uso estricto de variables semánticas (`erp-tokens.css`) para estados, finanzas y acciones (ej. `bg-erp-primary`, `text-erp-status-pending`).
* **Tipografía:** * *Montserrat:* Tipografía principal para la UI, tablas y formularios, optimizada para alta legibilidad de datos.
    * *Literata:* Uso exclusivo para títulos principales (Display).
* **Geometría:** Uso de bordes sutiles y esquinas moderadas (`rounded-lg`, `rounded-xl`), evitando el uso de elementos excesivamente redondeados o con estilo "glassmorphism".

## Estructura del Monorepo

* `/backend-laravel`: Monolito principal (Rutas, Controladores, Vistas de Vue, Base de Datos).
* `/go-imagenes`: Microservicio para la compresión y optimización de fotografías de obra.
* `/go-pdfs`: Microservicio para el ensamblaje rápido de contratos y recibos.
* `/go-websockets`: Servidor para notificaciones en tiempo real al panel de administración.
* `/go-worker`: Motor en segundo plano para procesar tareas y enviar recordatorios.

## Instalación y Entorno Local

Requisitos previos: Docker y Docker Desktop instalados (con integración WSL 2 habilitada en Windows).

1.  **Clonar el repositorio:**
    ```bash
    git clone https://github.com/tu-usuario/erp-sunset.git
    cd erp-sunset
    ```

2.  **Levantar la infraestructura con Docker Compose:**
    ```bash
    docker-compose up -d --build
    ```

3.  **Instalar dependencias del Backend (Laravel):**
    ```bash
    docker exec -it sunset-laravel composer install
    ```

4.  **Configurar variables de entorno:**
    Duplicar el archivo `.env.example` a `.env` dentro de `/backend-laravel` y asegurar la conexión a la base de datos interna:
    ```env
    DB_CONNECTION=pgsql
    DB_HOST=postgres-db
    DB_PORT=5432
    DB_DATABASE=erp_sunset
    DB_USERNAME=root
    DB_PASSWORD=secretpassword
    ```

5.  **Generar la clave de la aplicación y ejecutar migraciones:**
    ```bash
    docker exec -it sunset-laravel php artisan key:generate
    docker exec -it sunset-laravel php artisan migrate
    ```

6.  **Compilar los assets del Frontend (Vue/Tailwind):**
    Desde la terminal local (fuera de Docker), navegar a la carpeta de Laravel e iniciar Vite:
    ```bash
    cd backend-laravel
    npm install
    npm run dev
    ```

El sistema estará disponible en `http://localhost:8000`.

ENG 

# Sunset Sur - Real Estate ERP & CRM

Comprehensive Enterprise Resource Planning (ERP) and Customer Relationship Management (CRM) system explicitly designed for the operation and administration of real estate developments.

The system centralizes lead management, construction tracking, finances, multi-level commission calculations, and document automation within a highly responsive, scalable, and high-density visual platform.

## Project Architecture

The project utilizes a Monorepo approach, combining a robust monolith for business logic and security, supported by high-performance microservices for resource-intensive tasks. The entire infrastructure is orchestrated using Docker.

### Tech Stack

* **Backend (API Gateway & Core):** Laravel 11 (PHP 8.4)
* **Frontend:** Vue.js 3 (Composition API) + Inertia.js
* **Database:** PostgreSQL 15
* **Styling & UI:** Tailwind CSS v4 (Custom Token-based Design System)
* **Microservices:** Go (Golang)
* **Infrastructure:** Docker & Docker Compose (Alpine images and Multi-stage builds)

## Core Modules

1. **Dashboard:** Key metrics, sales of the month, upcoming appointments, and construction alerts.
2. **Users & Roles:** Granular access control for executives, sales representatives (individuals or team leaders), and construction staff.
3. **CRM (Sales):** Visual conversion funnel, contact directory, digital dossiers, and appointment agenda.
4. **Properties & Inventory:** Real estate development catalog, lot tracking, and real-time availability status.
5. **Construction & Projects:** Technical specs, initial budgets, material expense tracking, and photographic progress logs.
6. **Finances & Commissions:** Income registration, down payments, and a multi-level commission calculation engine.
7. **Documents:** Automated generator for PDF contracts and quotes.

## Design System (UI/UX)

The user interface discards generic templates to implement the premium brand identity of "Sunset Condominios".

* **Theme:** Strict Light Mode. Warm backgrounds (Sand) and dark text (Coffee).
* **Tokens:** Mandatory use of semantic variables (`erp-tokens.css`) for UI states, finance indicators, and actions (e.g., `bg-erp-primary`, `text-erp-status-pending`).
* **Typography:** * *Montserrat:* Primary font for UI, tables, and forms, optimized for high data density readability.
    * *Literata:* Exclusive use for main display headers (H1/H2).
* **Geometry:** Subtle borders and moderate border-radius (`rounded-lg`, `rounded-xl`), avoiding fully rounded elements or glassmorphism effects for better clarity.

## Monorepo Structure

* `/backend-laravel`: Main monolith (Routes, Controllers, Vue Views, Database).
* `/go-imagenes`: Microservice for compressing and optimizing construction progress photographs.
* `/go-pdfs`: Microservice for the rapid assembly of PDF contracts and receipts.
* `/go-websockets`: Server for real-time notifications sent to the administrative panel.
* `/go-worker`: Background engine to process asynchronous tasks and send reminders.

## Local Environment & Installation

Prerequisites: Docker and Docker Desktop installed (with WSL 2 integration enabled on Windows).

1. **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/erp-sunset.git](https://github.com/your-username/erp-sunset.git)
    cd erp-sunset
    ```

2. **Spin up the infrastructure with Docker Compose:**
    ```bash
    docker-compose up -d --build
    ```

3. **Install Backend (Laravel) dependencies:**
    ```bash
    docker exec -it sunset-laravel composer install
    ```

4. **Configure environment variables:**
    Duplicate the `.env.example` file to `.env` inside `/backend-laravel` and ensure the internal database connection is set:
    ```env
    DB_CONNECTION=pgsql
    DB_HOST=postgres-db
    DB_PORT=5432
    DB_DATABASE=erp_sunset
    DB_USERNAME=root
    DB_PASSWORD=secretpassword
    ```

5. **Generate the application key and run migrations:**
    ```bash
    docker exec -it sunset-laravel php artisan key:generate
    docker exec -it sunset-laravel php artisan migrate
    ```

6. **Compile Frontend assets (Vue/Tailwind):**
    From your local terminal (outside Docker), navigate to the Laravel folder and start Vite:
    ```bash
    cd backend-laravel
    npm install
    npm run dev
    ```

The system will be available at `http://localhost:8000`.
