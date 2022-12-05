import { defineStore } from "pinia"; // 引入pinia

export const UserStore = defineStore("user", {
  state: () => {
    return {
      token: "",
    };
  },
  actions: {},
  getters: {},
});