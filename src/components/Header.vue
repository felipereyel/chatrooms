<template>
    <div class="flex justify-between">
        <div class="flex flex-row">
            <router-link class="text-3xl" to="/">Rooms</router-link>
            <slot name="left"></slot>
        </div>
        <div class="flex flex-row">
            <slot name="right"></slot>
            <button @click="logout" class="border-2 rounded-md border-white px-2">{{state.username}} Logout</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { onMounted, reactive } from 'vue';
import { cookielogout, getUsername } from '../api';

const router = useRouter();

const logout = async () => {
    await cookielogout();
    router.push({ name: 'auth' });
};

const state = reactive({
    username: '',
});


onMounted(async () => {
    state.username = await getUsername();
});

</script>

<style scoped></style>