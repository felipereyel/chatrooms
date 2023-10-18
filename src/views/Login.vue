<template>
  <h1 class="text-3xl">Login</h1>
  <div class="flex-1 flex flex-col space-y-2 p-8">
    <div>
      <label for="username">Username</label>
      <input type="text" id="username" v-model="state.username" class="flex-1 text-black" />
    </div>
    <div>
      <label for="passwd">Password</label>
      <input type="password" id="passwd" v-model="state.passwd" class="flex-1 text-black" />
    </div>
    <div>
      <button @click="login()">
        Login
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const state = reactive({
  username: '',
  passwd: '',
});

const login = async () => {
  const { username, passwd } = state;

  if (!username || !passwd) {
    alert('Username and password are required');
    return;
  }

  try {
    const r = await fetch('http://localhost:3000/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, passwd }),
    });
  
    if (!r.ok) {
      throw new Error('Login failed');
    }
      
    router.push({ name: 'home' })
  } catch (e) {
    alert('Login failed');
  }
}
</script>

<style scoped></style>
