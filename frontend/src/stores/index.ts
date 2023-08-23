import { ref } from "vue";
import { defineStore } from "pinia";

export const useStore = defineStore("auth", () => {
  const isAuthenticated = ref(false);
  return { isAuthenticated };
});
