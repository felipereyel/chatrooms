<template>
  <Header class="flex justify-between">
    <template #left>
      <h1 class="text-3xl">&#20; > {{ state.roomName }}</h1>
    </template>
  </Header>
  <div class="flex-1 flex flex-col overflow-y-auto space-y-2 p-8">
    <div class="flex-1 overflow-y-auto space-y-2" id="msg-container">
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
import Header from '../components/Header.vue';
import { nextTick, onMounted, reactive } from 'vue';
import { onBeforeRouteLeave, useRoute } from 'vue-router';
import { getroom, listposts, createpost, getRoomWs } from '../api';


const route = useRoute();
const roomId = route.params.id as string;
let ws: WebSocket;

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
    const elem = document.getElementById('msg-container')!;
    const [room, lposts] = await Promise.all([getroom(roomId), listposts(roomId)]);
    state.roomName = room.name;
    posts.push(...lposts.reverse());
    elem.scrollTop = elem.scrollHeight;

    ws = getRoomWs(roomId);
    ws.addEventListener('message', (event) => {
      posts.push(JSON.parse(event.data));
      if (posts.length > 50) {
        posts.shift();
      }

      nextTick(() => {
        elem.scrollTop = elem.scrollHeight;
      });
    });

  } catch (e) {
    console.log(e);
    // router.push('/');
  }
});

onBeforeRouteLeave(() => ws.close());

const sendMessage = async () => {
  const { message } = state;
  if (!message) return;

  await createpost(route.params.id as string, message);
  state.message = '';
};
</script>

<style scoped></style>
