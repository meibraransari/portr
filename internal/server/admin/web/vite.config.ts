import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    manifest: true,
    outDir: "dist/static",
  },
  base: "/static/",
  resolve: {
    alias: {
      $lib: "/src/lib",
    },
  },
});
