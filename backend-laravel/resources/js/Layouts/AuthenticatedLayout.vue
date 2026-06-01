<script setup>
import { ref } from 'vue';
import Sidebar from '@/Components/Sidebar.vue';
import Navbar from '@/Components/Navbar.vue';

defineProps({
    title: {
        type: String,
        default: '',
    },
});

// Sidebar is open by default on desktop, closed on mobile
const isSidebarOpen = ref(window.innerWidth >= 1024);

const toggleSidebar = () => {
    isSidebarOpen.value = !isSidebarOpen.value;
};

const closeSidebar = () => {
    // Only close on mobile
    if (window.innerWidth < 1024) {
        isSidebarOpen.value = false;
    }
};
</script>

<template>
    <div class="flex h-screen overflow-hidden" style="background-color: #FAF5F0;">

        <!-- Sidebar -->
        <Sidebar
            :is-open="isSidebarOpen"
            @close="closeSidebar"
        />

        <!-- Main content area -->
        <div
            class="flex flex-1 flex-col overflow-hidden transition-all duration-250"
        >
            <!-- Navbar -->
            <Navbar
                :title="title"
                @toggle-sidebar="toggleSidebar"
            />

            <!-- Page content -->
            <main
                class="flex-1 overflow-y-auto"
                style="background-color: #FAF5F0;"
            >
                <!-- Optional page header slot -->
                <div
                    v-if="$slots.header"
                    class="px-6 pt-6 pb-2"
                >
                    <slot name="header" />
                </div>

                <!-- Main slot -->
                <div class="px-6 py-6">
                    <slot />
                </div>
            </main>
        </div>
    </div>
</template>
