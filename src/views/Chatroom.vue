<template>
  <div class="flex items-end">
    <router-link class="text-3xl" to="/">Rooms</router-link>
    <h1 class="text-3xl">&#20; > {{ route.params.name }}</h1>
  </div>
  <div class="flex-1 flex flex-col space-y-2 p-8">
    <div class="flex-1 overflow-y-auto space-y-2">
      <div v-for="message in messages" :key="message.id">
        {{ message.content }}
      </div>
    </div>
    <div class="flex space-x-2 ">
      <input type="text" v-model="state.message" class="flex-1 text-black" @keyup.enter="sendMessage()" />
      <button @click="sendMessage()">Send</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
const route = useRoute();
const router = useRouter();

const messages = reactive([
  {
    id: 1,
    content: 'Hello, world!',
  },
  {
    id: 2,
    content: 'Hello, Vue 3!',
  },
]);

const state = reactive({
  message: '',
});

const sendMessage = () => {
  const { message } = state;

  if (!message) return;

  messages.push({
    id: messages.length + 1,
    content: message,
  });

  state.message = '';
};

const goHome = () => {
  router.push({ name: 'home' });
};
</script>

<style scoped></style>
