<script setup>
import { Link } from '@inertiajs/vue3';
import { Menu, Bell, Search, ChevronDown } from 'lucide-vue-next';
import Dropdown from '@/Components/Dropdown.vue';
import DropdownLink from '@/Components/DropdownLink.vue';

defineProps({
    title: {
        type: String,
        default: '',
    },
});

defineEmits(['toggleSidebar']);
</script>

<template>
    <header
        class="sticky top-0 z-10 flex h-14 items-center gap-3 border-b px-4 sm:px-6"
        style="
            background: rgba(250, 245, 240, 0.85);
            backdrop-filter: blur(12px);
            -webkit-backdrop-filter: blur(12px);
            border-color: rgba(44, 26, 14, 0.08);
        "
    >
        <!-- Hamburger (mobile + desktop toggle) -->
        <button
            id="sidebar-toggle-btn"
            class="flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg transition-colors duration-150"
            style="color: rgba(44,26,14,0.5);"
            @click="$emit('toggleSidebar')"
            @mouseover="$event.currentTarget.style.background = 'rgba(44,26,14,0.06)'; $event.currentTarget.style.color = '#2C1A0E'"
            @mouseleave="$event.currentTarget.style.background = 'transparent'; $event.currentTarget.style.color = 'rgba(44,26,14,0.5)'"
            aria-label="Toggle sidebar"
        >
            <Menu class="h-5 w-5" :stroke-width="1.75" />
        </button>

        <!-- Divider -->
        <div class="h-5 w-px" style="background: rgba(44,26,14,0.1);" />

        <!-- Module title -->
        <div class="flex-1 min-w-0">
            <h1
                v-if="title"
                class="truncate text-base font-light italic leading-tight"
                style="font-family: 'Literata', serif; color: #2C1A0E;"
            >
                {{ title }}
            </h1>
        </div>

        <!-- Right actions -->
        <div class="flex items-center gap-1.5">

            <!-- Search button -->
            <button
                id="global-search-btn"
                class="flex h-8 w-8 items-center justify-center rounded-lg transition-colors duration-150"
                style="color: rgba(44,26,14,0.45);"
                @mouseover="$event.currentTarget.style.background = 'rgba(44,26,14,0.06)'; $event.currentTarget.style.color = '#2C1A0E'"
                @mouseleave="$event.currentTarget.style.background = 'transparent'; $event.currentTarget.style.color = 'rgba(44,26,14,0.45)'"
                aria-label="Buscar"
            >
                <Search class="h-4 w-4" :stroke-width="1.75" />
            </button>

            <!-- Notifications button -->
            <button
                id="notifications-btn"
                class="relative flex h-8 w-8 items-center justify-center rounded-lg transition-colors duration-150"
                style="color: rgba(44,26,14,0.45);"
                @mouseover="$event.currentTarget.style.background = 'rgba(44,26,14,0.06)'; $event.currentTarget.style.color = '#2C1A0E'"
                @mouseleave="$event.currentTarget.style.background = 'transparent'; $event.currentTarget.style.color = 'rgba(44,26,14,0.45)'"
                aria-label="Notificaciones"
            >
                <Bell class="h-4 w-4" :stroke-width="1.75" />
                <!-- Notification dot -->
                <span
                    class="absolute top-1.5 right-1.5 h-1.5 w-1.5 rounded-full"
                    style="background: #ea580c;"
                />
            </button>

            <!-- Divider -->
            <div class="h-5 w-px mx-1" style="background: rgba(44,26,14,0.1);" />

            <!-- User Dropdown -->
            <Dropdown align="right" width="48">
                <template #trigger>
                    <button
                        id="user-menu-btn"
                        class="flex h-8 items-center gap-2 rounded-lg px-2 transition-colors duration-150"
                        style="color: rgba(44,26,14,0.7);"
                        @mouseover="$event.currentTarget.style.background = 'rgba(44,26,14,0.06)'; $event.currentTarget.style.color = '#2C1A0E'"
                        @mouseleave="$event.currentTarget.style.background = 'transparent'; $event.currentTarget.style.color = 'rgba(44,26,14,0.7)'"
                    >
                        <!-- User avatar -->
                        <div
                            class="flex h-6 w-6 flex-shrink-0 items-center justify-center rounded-full text-[10px] font-semibold uppercase"
                            style="background: rgba(234,88,12,0.15); color: #ea580c;"
                        >
                            {{ ($page.props.auth?.user?.name || 'U').charAt(0) }}
                        </div>
                        <span class="hidden text-xs font-medium sm:block">
                            {{ ($page.props.auth?.user?.name || 'Usuario').split(' ')[0] }}
                        </span>
                        <ChevronDown class="h-3 w-3 opacity-50" :stroke-width="2.5" />
                    </button>
                </template>

                <template #content>
                    <div class="px-4 py-3 border-b" style="border-color: rgba(44,26,14,0.08);">
                        <p class="text-xs font-semibold" style="color: #2C1A0E;">
                            {{ $page.props.auth?.user?.name }}
                        </p>
                        <p class="text-[10px] mt-0.5 truncate" style="color: rgba(44,26,14,0.45);">
                            {{ $page.props.auth?.user?.email }}
                        </p>
                    </div>
                    <DropdownLink :href="route('profile.edit')">
                        Perfil
                    </DropdownLink>
                    <DropdownLink
                        :href="route('logout')"
                        method="post"
                        as="button"
                    >
                        Cerrar sesión
                    </DropdownLink>
                </template>
            </Dropdown>
        </div>
    </header>
</template>
