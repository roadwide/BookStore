import { defineStore } from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore('main', {
  // a function that returns a fresh state
  state: () => ({
    isLogin: false,
    token: 'default token',
    username: 'default username',
  }),
  // optional actions
  actions: {
    getUserState() {
        if (localStorage.getItem("userInfo") !== null) {
            this.isLogin = true
            const usesrInfo = JSON.parse(localStorage.getItem("userInfo")!)    // 非空断言，叹号
            this.username = usesrInfo.username
            this.token = usesrInfo.token
        }
    },
    logout() {
        this.isLogin = false
        this.token = 'default token'
        this.username = 'default username'
        localStorage.removeItem("userInfo")
    },
  },
})