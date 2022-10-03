<template>
  <h1>UserHome</h1>
  <h1 v-show="username.length !== 0">用户名：{{ username }}</h1>
  <el-button @click="logout" type="danger">注销</el-button>
</template>

<script lang="ts">
import axios from 'axios'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
export default {
    name: "UserHome",
    setup() {
        const router = useRouter()
        let username = ref("")
        if (localStorage.getItem("userInfo") === null) {
            alert("未登录！跳转到登录页面")
            router.push("login")
        } else {
            const usesrInfo = JSON.parse(localStorage.getItem("userInfo")!)    // 非空断言，叹号
            let formData = new FormData()
            formData.append('userID', usesrInfo.userID)
            formData.append('token', usesrInfo.token)
            const url = 'http://localhost:8081/token'
            axios.post(url, formData).then(
                function(response) {
                    if (response.data.message !== "ok") {
                        console.log(response.data)
                    } else {
                        username.value = response.data.userID
                    }
                    
                }
            ).catch(
                function(err) {
                    console.log(err.message)
                }
            )
        }

        function logout() {
            alert("注销！跳转登录页面")
            localStorage.removeItem("userInfo")
            router.push("login")
        }

        return {
            username,
            logout
        }
    }
}
</script>

<style>

</style>