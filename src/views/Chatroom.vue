<template>
  <div class="flex items-end">
    <router-link class="text-3xl" to="/">Rooms</router-link>
    <h1 class="text-3xl">&#20; > {{ state.roomName }}</h1>
  </div>
  <div class="flex-1 flex flex-col overflow-y-auto space-y-2 p-8">
    <div class="flex-1 overflow-y-auto space-y-2">
      <div v-for="message in posts" :key="message.id" class="flex items-center bg-slate-600 rounded-md px-2 justify-between">
        <span>
          {{ message.content }}
        </span>
        <span><pre>{{ message.username }} @ {{ formatCreatedAt(message.created_at) }}</pre></span>
      </div>
    </div>
    <div class="flex space-x-2 ">
      <input type="text" v-model="state.message" class="flex-1 text-black" @keyup.enter="sendMessage()" />
      <button @click="sendMessage()">Send</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getroom, listposts, createpost } from '../api';

const route = useRoute();
const router = useRouter();

const roomId = route.params.id as string;

type Posts = {
  id: string;
  content: string;
  username: string;
  created_at: string;
};

const posts: Posts[] = reactive([]);

const state = reactive({
  roomName: '',
  message: '',
});

const formatCreatedAt = (datestr: string) => {
  const date = new Date(datestr);
  return date.toTimeString().split(' ')[0];
};

onMounted(async () => {
  try {
    const [room, lposts] = await Promise.all([getroom(roomId), listposts(roomId)]);
    state.roomName = room.name;
    posts.push(...lposts.reverse());
  } catch (e) {
    console.log(e);
    router.push('/');
  }
});

const sendMessage = async () => {
  const { message } = state;
  if (!message) return;

  const post = await createpost(route.params.id as string, message);
  console.log(post);
  posts.push(post);
  if (posts.length > 50) {
    posts.shift();
  }

  state.message = '';
};
</script>

<style scoped></style>
