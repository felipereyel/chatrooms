<template>
  <Header class="flex justify-between">
    <template #right>
      <button @click="newRoom" class="border-2 rounded-md border-white px-2">
        New Room
      </button>
    </template>
  </Header>
  <div class="flex-1 flex flex-col space-y-2 p-8">
    <router-link  v-for="room in rooms" :key="room.id" :to="{ name: 'room', params: { id: room.id } }">
      {{ room.name }}
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { listrooms } from '../api';
import { useRouter } from 'vue-router';
import { onMounted, reactive } from 'vue';
import Header from '../components/Header.vue';

const router = useRouter();

type Room = {
  id: string;
  name: string;
};

const rooms: Room[] = reactive([]);

onMounted(async () => {
  rooms.push(...await listrooms());
});

const newRoom = () => {
  router.push({ name: 'new-room' });
};

</script>

<style scoped></style>
../api