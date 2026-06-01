<script setup>
import { computed } from 'vue';
import { Link, usePage } from '@inertiajs/vue3';
import {
    LayoutDashboard,
    Building2,
    Banknote,
    TrendingUp,
    Package,
    ShoppingCart,
    Users,
    HeartHandshake,
    Settings,
    ChevronRight,
    LogOut,
} from 'lucide-vue-next';

const props = defineProps({
    isOpen: {
        type: Boolean,
        default: true,
    },
});

const page = usePage();

const navItems = [
    {
        label: 'Inicio',
        route: 'dashboard',
        icon: LayoutDashboard,
        color: 'text-erp-text-inverse',
    },
    {
        label: 'Obras',
        route: 'dashboard',
        icon: Building2,
        color: 'text-amber-300',
        badge: null,
    },
    {
        label: 'Finanzas',
        route: 'dashboard',
        icon: Banknote,
        color: 'text-blue-300',
    },
    {
        label: 'Ventas',
        route: 'dashboard',
        icon: TrendingUp,
        color: 'text-sky-300',
    },
    {
        label: 'Inventario',
        route: 'dashboard',
        icon: Package,
        color: 'text-emerald-300',
    },
    {
        label: 'Compras',
        route: 'dashboard',
        icon: ShoppingCart,
        color: 'text-yellow-300',
    },
    {
        label: 'RRHH',
        route: 'dashboard',
        icon: Users,
        color: 'text-purple-300',
    },
    {
        label: 'CRM',
        route: 'dashboard',
        icon: HeartHandshake,
        color: 'text-rose-300',
    },
];

const isActive = (routeName) => {
    try {
        return route().current(routeName);
    } catch {
        return false;
    }
};
</script>

<template>
    <!-- Sidebar overlay for mobile -->
    <Transition name="overlay">
        <div
            v-if="isOpen"
            class="fixed inset-0 z-20 bg-black/50 backdrop-blur-sm lg:hidden"
            @click="$emit('close')"
        />
    </Transition>

    <!-- Sidebar panel -->
    <Transition name="sidebar">
        <aside
            v-show="isOpen"
            class="fixed top-0 left-0 z-30 flex h-screen w-[260px] flex-col overflow-hidden lg:static lg:z-auto lg:translate-x-0"
            style="background-color: #2C1A0E;"
        >
            <!-- Logo / Brand -->
            <div class="flex items-center gap-3 px-6 py-5 border-b" style="border-color: rgba(250,245,240,0.08);">
                <div class="flex h-9 w-9 items-center justify-center rounded-lg flex-shrink-0"
                     style="background: linear-gradient(135deg, #ea580c 0%, #B5432E 100%);">
                    <span class="text-white font-bold text-sm" style="font-family: 'Literata', serif;">S</span>
                </div>
                <div class="flex flex-col leading-none">
                    <span class="font-semibold text-sm tracking-wide" style="color: #FAF5F0; font-family: 'Literata', serif; font-style: italic;">
                        Sunset
                    </span>
                    <span class="text-[10px] uppercase tracking-[0.2em] font-medium" style="color: rgba(250,245,240,0.45);">
                        ERP
                    </span>
                </div>
            </div>

            <!-- Navigation -->
            <nav class="flex-1 overflow-y-auto py-4 px-3 space-y-0.5">
                <!-- Section label -->
                <p class="px-3 mb-2 text-[10px] uppercase tracking-[0.2em] font-semibold" style="color: rgba(250,245,240,0.3);">
                    Módulos
                </p>

                <Link
                    v-for="item in navItems"
                    :key="item.label"
                    :href="route(item.route)"
                    class="group flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-all duration-150 relative"
                    :class="[
                        isActive(item.route)
                            ? 'bg-white/10 text-white shadow-inner'
                            : 'hover:bg-white/6 text-white/70 hover:text-white'
                    ]"
                    style="text-decoration: none;"
                    @click="$emit('close')"
                >
                    <!-- Active indicator -->
                    <span
                        v-if="isActive(item.route)"
                        class="absolute left-0 top-1/2 -translate-y-1/2 h-5 w-0.5 rounded-full bg-erp-accent"
                    />

                    <!-- Icon -->
                    <component
                        :is="item.icon"
                        class="h-[18px] w-[18px] flex-shrink-0 transition-colors duration-150"
                        :class="[isActive(item.route) ? item.color : 'text-white/40 group-hover:text-white/70']"
                        :stroke-width="1.75"
                    />

                    <!-- Label -->
                    <span class="flex-1 font-montserrat">{{ item.label }}</span>

                    <!-- Chevron -->
                    <ChevronRight
                        class="h-3.5 w-3.5 opacity-0 -translate-x-1 transition-all duration-150 group-hover:opacity-40 group-hover:translate-x-0"
                        style="color: rgba(250,245,240,0.5);"
                        :stroke-width="2"
                    />
                </Link>

                <!-- Divider -->
                <div class="my-3 mx-2" style="border-top: 1px solid rgba(250,245,240,0.08);" />

                <!-- Settings -->
                <Link
                    :href="route('profile.edit')"
                    class="group flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-all duration-150"
                    :class="[
                        route().current('profile.edit')
                            ? 'bg-white/10 text-white'
                            : 'text-white/50 hover:bg-white/6 hover:text-white/80'
                    ]"
                    style="text-decoration: none;"
                    @click="$emit('close')"
                >
                    <Settings
                        class="h-[18px] w-[18px] flex-shrink-0 transition-colors duration-150"
                        :class="route().current('profile.edit') ? 'text-white' : 'text-white/30 group-hover:text-white/60'"
                        :stroke-width="1.75"
                    />
                    <span class="flex-1">Configuración</span>
                </Link>
            </nav>

            <!-- User Footer -->
            <div class="px-3 py-4 border-t" style="border-color: rgba(250,245,240,0.08);">
                <div class="flex items-center gap-3 rounded-lg px-3 py-2.5">
                    <!-- Avatar placeholder -->
                    <div class="flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-full text-xs font-semibold uppercase"
                         style="background: rgba(234,88,12,0.2); color: #ea580c;">
                        {{ ($page.props.auth?.user?.name || 'U').charAt(0) }}
                    </div>
                    <div class="flex-1 min-w-0">
                        <p class="text-xs font-semibold truncate" style="color: rgba(250,245,240,0.9);">
                            {{ $page.props.auth?.user?.name || 'Usuario' }}
                        </p>
                        <p class="text-[10px] truncate" style="color: rgba(250,245,240,0.35);">
                            {{ $page.props.auth?.user?.email || '' }}
                        </p>
                    </div>
                    <Link
                        :href="route('logout')"
                        method="post"
                        as="button"
                        class="flex h-7 w-7 flex-shrink-0 items-center justify-center rounded-md transition-colors duration-150"
                        style="color: rgba(250,245,240,0.3);"
                        @mouseover="$event.currentTarget.style.color = 'rgba(250,245,240,0.8)'; $event.currentTarget.style.background = 'rgba(250,245,240,0.06)'"
                        @mouseleave="$event.currentTarget.style.color = 'rgba(250,245,240,0.3)'; $event.currentTarget.style.background = 'transparent'"
                    >
                        <LogOut class="h-4 w-4" :stroke-width="1.75" />
                    </Link>
                </div>
            </div>
        </aside>
    </Transition>
</template>

<style scoped>
.sidebar-enter-active,
.sidebar-leave-active {
    transition: transform 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}
.sidebar-enter-from,
.sidebar-leave-to {
    transform: translateX(-100%);
}

.overlay-enter-active,
.overlay-leave-active {
    transition: opacity 0.2s ease;
}
.overlay-enter-from,
.overlay-leave-to {
    opacity: 0;
}

@media (min-width: 1024px) {
    .sidebar-enter-from,
    .sidebar-leave-to {
        transform: none;
    }
}
</style>
