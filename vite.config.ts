import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const DEV_API_PORT = process.env.DEV_API_PORT || 3000;

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/_api': {
        target: `http://0.0.0.0:${DEV_API_PORT}`,
      },
      '^/_api/rooms/.*/ws': {
        target: `ws://0.0.0.0:${DEV_API_PORT}`,
        ws: true,
      }
    }
  }
});
