<script setup>
import AuthenticatedLayout from '@/Layouts/AuthenticatedLayout.vue';
import { Head } from '@inertiajs/vue3';
</script>

<template>
    <Head title="Inicio" />

    <AuthenticatedLayout title="Panel de Control">
        <!-- KPI Cards placeholder -->
        <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
            <div
                v-for="i in 4"
                :key="i"
                class="rounded-xl border p-5 flex flex-col gap-3"
                style="background: #fff; border-color: rgba(44,26,14,0.08); box-shadow: 0 1px 2px 0 rgba(44,26,14,0.06);"
            >
                <div class="h-3 w-24 rounded-full" style="background: rgba(44,26,14,0.07);" />
                <div class="h-7 w-16 rounded-lg" style="background: rgba(44,26,14,0.05);" />
                <div class="h-2 w-20 rounded-full" style="background: rgba(44,26,14,0.04);" />
            </div>
        </div>
    </AuthenticatedLayout>
</template>
