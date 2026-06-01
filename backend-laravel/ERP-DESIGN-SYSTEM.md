# Sistema de diseño Sunset → ERP

Lineamientos extraídos de `src/app/globals.css`, componentes UI y páginas internas. Pensados para reutilizar la identidad Sunset en una aplicación operativa (tablas, formularios, estados, permisos).

**Archivo de tokens listo para copiar:** [`erp-tokens.css`](./erp-tokens.css)

---

## 1. Principios de marca (landing)

| Concepto | Valor / patrón |
|----------|----------------|
| Tono | Cálido, editorial, premium — fondos arena (`#FAF5F0`), texto café (`#2C1A0E`) |
| Contraste | Azul corporativo + naranja energético + verde confianza |
| Motion | Framer Motion en marketing; en ERP preferir transiciones CSS cortas |
| Densidad visual | Landing: mucho aire, `rounded-3xl`, glassmorphism. **ERP: más compacto** |

---

## 2. Tipografía

| Rol | Familia | Pesos usados | Uso en landing | Uso recomendado en ERP |
|-----|---------|--------------|----------------|------------------------|
| UI / labels / tablas | **Montserrat** (`font-montserrat`) | 300 light, 400, 500 medium, 600 semibold, 700 bold | Formularios, badges, navegación | **Principal** — todo el producto |
| Display / títulos | **Literata** (`font-literata`) | 300 light, 400, italic | H1–H3, citas | Solo títulos de módulo o dashboard (1 por vista) |
| Sistema (layout root) | Geist / Arial fallback | — | Variables cargadas en `layout.tsx` | Opcional en ERP; Montserrat basta |

### Escalas observadas en componentes

| Elemento | Clases típicas |
|----------|----------------|
| Eyebrow / sección | `font-montserrat text-[10px] sm:text-xs tracking-[0.2em] uppercase font-medium` |
| Label de formulario | `font-montserrat text-[10px] tracking-[0.2em] uppercase font-semibold` |
| Cuerpo | `font-montserrat font-light text-sm` / `text-base leading-relaxed` |
| Título hero | `font-literata font-light tracking-tight` — `text-4xl` → `text-6xl` |
| Título sección | `font-literata font-light italic text-3xl sm:text-5xl` |
| Badge / pill | `font-montserrat text-[10px] tracking-[0.2em] uppercase` |

### Escala ERP sugerida

```text
text-xs   → metadatos, celdas secundarias
text-sm   → cuerpo, inputs, botones
text-base → contenido principal
text-lg   → subtítulos de panel
text-xl+  → Literata solo en encabezado de módulo
```

### Escala fluida de `html` (heredada)

En `globals.css`, el `font-size` del root baja en viewports grandes (93.75% → 87.5%) y vuelve a 100% en ≥1920px. En ERP conviene **fijar 16px** (`html { font-size: 100% }`) para legibilidad en datos densos.

---

## 3. Paleta de color (extraída)

### Marca

| Token Tailwind | Hex | Uso landing |
|----------------|-----|-------------|
| `brand-blue` | `#1e3a8a` | Confianza, Casas Sur, links hover |
| `brand-green` | `#15803d` | Éxito, Casas Sur nav |
| `brand-orange` | `#ea580c` | CTA, acentos fuertes |
| `brand-sand` | `#e3cba8` | Detalles cálidos |
| `brand-purple` | `#2e1065` | Gradientes contacto |

### Páginas internas (`page-*`)

| Token | Hex |
|-------|-----|
| `page-bg` | `#FAF5F0` |
| `page-bg-alt` | `#F3EBE0` |
| `page-text` | `#2C1A0E` |
| `page-text-hover` | `#3d2820` |

### Sub-marca Sunset Condominios (`sc-*`)

| Token | Hex | Notas |
|-------|-----|-------|
| `sc-primary` | `#e6846a` | Coral — focus rings en formularios |
| `sc-primary-dark` | `#B5432E` | Énfasis |
| `sc-contrast` | `#53b076` | Verde suave |
| `sc-accent` | `#D4728C` | Rosa — selección texto |

### Opacidades recurrentes

- Texto secundario: `text-page-text/80`, `/50`, `/30`
- Bordes: `border-page-text/5`, `/10`, `/20`
- Fondos glass: `bg-white/40`, `bg-white/50`, `bg-white/70` + `backdrop-blur-xl`
- Hover en cards: `hover:border-{color}/20`, `hover:shadow-{color}/5`

---

## 4. Paleta ERP (complemento)

Tokens en [`erp-tokens.css`](./erp-tokens.css). Mapeo por **intención**:

### Acciones de interfaz

| Intención | Token | Color | Ejemplos ERP |
|-----------|-------|-------|--------------|
| Primaria | `erp-primary` | `#1e3a8a` | Guardar, buscar, navegación activa |
| Acento / crear | `erp-accent` | `#ea580c` | Nuevo registro, exportar destacado |
| Secundaria | `erp-bg-muted` + borde | — | Cancelar, volver |
| Peligro destructiva | `erp-danger` | `#b91c1c` | Eliminar, anular factura |
| Éxito | `erp-success` | `#15803d` | Confirmar pago, sincronizado |
| Advertencia | `erp-warning` | `#b45309` | Stock bajo, vence pronto |
| Información | `erp-info` | `#1e40af` | Ayuda, tips, en revisión |

### Estados de documento

| Estado | Fondo badge | Texto |
|--------|-------------|-------|
| Borrador | `erp-status-draft-bg` | `erp-status-draft` |
| Pendiente autorización | `erp-status-pending-bg` | `erp-status-pending` |
| En revisión | `erp-status-in-review-bg` | `erp-status-in-review` |
| Aprobado / pagado | `erp-status-approved-bg` | `erp-status-approved` |
| Rechazado | `erp-status-rejected-bg` | `erp-status-rejected` |
| Cancelado | `erp-status-cancelled-bg` | `erp-status-cancelled` |
| Archivado | `erp-status-archived-bg` | `erp-status-archived` |

### Inventario

| Nivel | Token |
|-------|-------|
| OK | `erp-stock-ok` |
| Bajo | `erp-stock-low` |
| Crítico | `erp-stock-critical` |
| Sobrestock | `erp-stock-overstock` |

### Finanzas

| Significado | Token |
|-------------|-------|
| Ingreso / saldo a favor | `erp-money-positive` |
| Egreso / adeudo | `erp-money-negative` |
| Neutro / conciliado | `erp-money-neutral` |
| Por cobrar / pendiente | `erp-money-pending` |

### Módulos (sidebar / chips)

Ventas → azul · Inventario → verde SC · Compras → ámbar · RRHH → púrpura · Finanzas → azul claro · CRM → rosa SC.

---

## 5. Componentes (patrones extraídos)

### Botón (`Button.tsx`)

| Variant | Estilo |
|---------|--------|
| `default` | `bg-white text-black rounded-full` |
| `brand` | `bg-brand-orange text-white` |
| `outline` | borde blanco/20, glass |
| `ghost` | texto blanco, hover `bg-white/10` |
| Tamaños | `sm`: h-9 px-4 · `lg`: px-8 py-6 |
| Extra | Shine hover con gradiente skew |

**ERP:** usar `rounded-lg` o `rounded-xl` en lugar de `rounded-full`; variantes `primary`, `accent`, `outline`, `ghost`, `danger`.

### Tarjetas (landing)

```text
rounded-3xl | rounded-[2rem]
border border-page-text/10
bg-page-bg o bg-white/40 + backdrop-blur
p-6 sm:p-10
transition-card duration-500
hover:shadow-sm hover:border-{accent}/20
```

**ERP:**

```text
rounded-xl border border-erp-border bg-erp-bg-elevated
p-4 shadow-erp-sm
hover:shadow-erp-md (sin transform grande)
```

### Inputs (contacto, profesionales)

```text
h-12 rounded-2xl px-5
bg-white/50 border border-page-text/10
font-montserrat font-light text-sm
focus:ring-1 focus:ring-sc-accent/30
label: text-[10px] uppercase tracking-[0.2em] font-semibold
```

**ERP:** `h-10 rounded-lg`, focus `ring-erp-primary/30` o `ring-erp-border-focus`.

### Badges / pills

```text
px-4 py-1.5 rounded-full
font-montserrat text-[10px] tracking-[0.2em] uppercase
bg-{semantic}/10 text-{semantic} border border-{semantic}/20
```

### Alertas (formulario error)

```text
bg-red-50/50 border border-red-200/50 rounded-xl p-4
```

**ERP:** alinear con `erp-danger-bg` + `erp-danger-border`.

### Iconos en tarjetas

Contenedor: `w-10 h-10 rounded-xl bg-{color}/10` · icono `w-5 h-5 text-{color}`.

---

## 6. Espaciado y layout

| Patrón landing | Valor típico | ERP sugerido |
|----------------|--------------|--------------|
| Sección vertical | `py-24 lg:py-40` | `py-6` – `py-8` |
| Contenedor | `max-w-5xl` / `max-w-7xl` `px-6` | full width con sidebar 240–280px |
| Grid formulario | `grid-cols-1 sm:grid-cols-2 gap-6` | igual, `gap-4` en modales |
| Breakpoint extra | `3xl: 120rem` | útil para tablas anchas |

---

## 7. Sombras y elevación

Landing unifica todo en `--shadow-sm` (sutil, editorial).

ERP (`erp-tokens.css`):

- `shadow-erp-sm` — filas, inputs
- `shadow-erp-md` — cards, dropdowns
- `shadow-erp-lg` — modales, drawer

---

## 8. Bordes y radios

| Contexto | Landing | ERP |
|----------|---------|-----|
| Cards marketing | `rounded-3xl`, `rounded-[2.5rem]` | `rounded-xl` |
| Inputs | `rounded-2xl` | `rounded-lg` |
| Botones CTA | `rounded-full` | `rounded-lg` |
| Tablas | — | `rounded-lg` contenedor, celdas sin radio |

---

## 9. Motion y transiciones

| Utilidad | Definición |
|----------|------------|
| `.transition-card` | color, background, border, shadow, transform — `cubic-bezier(0.4, 0, 0.2, 1)` |
| `standardFadeUp` | opacity 0→1, 0.8s easeOut |
| `standardViewport` | `once: true, amount: 0.1` |
| CinematicHeading | stagger por carácter/palabra — **no usar en ERP** |

**ERP:** `transition-colors duration-150` en filas; modales 200–300ms.

---

## 10. Accesibilidad y contraste

- Texto principal `#2C1A0E` sobre `#FAF5F0`: ratio alto ✓
- `brand-orange` sobre blanco: verificar en botones pequeños (usar texto blanco en filled)
- Estados no depender solo del color: icono + texto (ej. “Aprobado”, “Stock bajo”)
- Focus visible: heredar `focus-visible:ring-2` del Button; en ERP `ring-erp-primary/40`

---

## 11. Integración en proyecto ERP (Tailwind 4)

```css
@import "tailwindcss";
@import "../design/erp-tokens.css";
```

Ejemplos de clases:

```html
<button class="bg-erp-primary text-erp-text-on-primary hover:bg-erp-primary-hover rounded-lg px-4 h-10 font-montserrat text-sm font-medium">
  Guardar
</button>

<span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-montserrat font-medium bg-erp-status-pending-bg text-erp-status-pending">
  Pendiente
</span>

<tr class="border-b border-erp-border hover:bg-erp-bg-muted/50 transition-colors">
```

---

## 12. Qué no trasladar del marketing al ERP

- Hero a pantalla completa, blur orbs, textura noise
- `CinematicHeading` y animaciones por carácter
- `rounded-full` en todos los botones
- Glassmorphism excesivo en tablas (legibilidad)
- `color-scheme: dark` del root layout del landing (ERP: `light`)

---

## 13. Checklist al implementar pantallas ERP

1. Fondo `erp-bg`, texto `erp-text`
2. Sidebar `erp-bg-sidebar` + texto `erp-text-inverse`
3. Una sola familia display (Literata) por vista si hace falta
4. Badges de estado con par fondo/texto de la tabla §4
5. Botón primario = azul; acción destacada “crear” = naranja
6. Destructivo siempre rojo + confirmación modal
7. Tablas: Montserrat `text-sm`, zebra opcional con `erp-bg-alt`
8. Importar tokens y fijar `font-size: 100%` en `html`

---

*Generado a partir del repositorio sunsetlanding — Mayo 2026*
