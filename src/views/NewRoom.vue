<template>
  <Header>
    <template #left>
      <h1 class="text-3xl">: New Room</h1>
    </template>
  </Header>
  <div class="flex-1 flex align-center justify-center space-y-2 p-8">
    <div class="flex flex-col align-center w-80">
      <input type="text" placeholder="room name" v-model="state.name" class="mb-2 text-black" />
      <div class="flex justify-center space-x-3">
        <button @click="create()" class="border-2 rounded-md border-white px-2" >Create Room</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { createroom } from '../api';
import { useRouter } from 'vue-router';
import Header from '../components/Header.vue';

const router = useRouter();

const state = reactive({
  name: '',
});

const create = async () => {
  const { name } = state;

  if (!name) {
    alert('name is required');
    return;
  }

  try {
    const room = await createroom(name);
    router.push({ name: 'room', params: { id: room.id } });
  } catch (e) {
    console.log(e);
    alert('Creation failed');
  }
}
</script>

<style scoped></style>
../api