<template>
  <div class="flex justify-between">
    <h1 class="text-3xl">Rooms</h1>
    <button @click="logout">Logout</button>
  </div>
  <div>
    <router-link :to="{ name: 'new-room' }">
      New Room
    </router-link>
  </div>
  <div class="flex-1 flex flex-col space-y-2 p-8">
    <router-link  v-for="room in rooms" :key="room.id" :to="{ name: 'room', params: { id: room.id } }">
      {{ room.name }}
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { onMounted, reactive } from 'vue';
import { cookielogout, listrooms } from '../api';

const router = useRouter();

const logout = async () => {
  await cookielogout();
  router.push({ name: 'auth' });
};

type Room = {
  id: string;
  name: string;
};

const rooms: Room[] = reactive([]);

onMounted(async () => {
  rooms.push(...await listrooms());
});

</script>

<style scoped></style>
../api