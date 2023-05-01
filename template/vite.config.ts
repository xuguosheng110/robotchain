import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      "/service": {
        target: "https://api.geekros.com",
        secure: false,
        changeOrigin: true,
        headers: {
          Referer: 'https://api.geekros.com'
        },
        rewrite: (path) => path.replace(/^\/service/, '')
      }
    }
  },
  build: {
    outDir: "./release",
  }
})