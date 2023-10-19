<template>
  <h1 class="text-3xl">Login / Register</h1>
  <div class="flex-1 flex align-center justify-center space-y-2 p-8">
    <div class="flex flex-col align-center justify-center  w-80">
      <input type="text" placeholder="username" v-model="state.username" class="mb-2 text-black" />
      <input type="password" placeholder="passwd" v-model="state.passwd" class="mb-2 text-black" />
      <div class="flex justify-center space-x-3">
        <button @click="login()" class="border-2 rounded-md border-white px-2" >Login</button>
        <button @click="register()" class="border-2 rounded-md border-white px-2">Register</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { apilogin, apiregister } from '../auth';

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
    await apilogin(username, passwd);
    router.push({ name: 'home' })
  } catch (e) {
    console.log(e);
    alert('Login failed');
  }
}

const register = async () => {
  const { username, passwd } = state;

  if (!username || !passwd) {
    alert('Username and password are required');
    return;
  }

  try {
    await apiregister(username, passwd);
    router.push({ name: 'home' })
  } catch (e) {
    console.log(e);
    alert('Register failed');
  }
}
</script>

<style scoped></style>
