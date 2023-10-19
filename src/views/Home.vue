<template>
  <div class="flex justify-between">
    <h1 class="text-3xl">Rooms</h1>
    <button @click="logout">Logout</button>
  </div>
  <div class="flex-1 flex flex-col space-y-2 p-8">
    <router-link  v-for="room in rooms" :key="room.name" :to="{ name: 'room', params: { name: room.name } }">
      {{ room.name }}
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { cookielogout } from '../auth';

const router = useRouter();

const logout = async () => {
  await cookielogout();
  router.push({ name: 'auth' });
};

const rooms = [
  {
    'name': 'room-1',
  },
  {
    'name': 'room-2',
  },
  {
    'name': 'room-2',
  },
]

</script>

<style scoped></style>
